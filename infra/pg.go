package infra

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"math/rand/v2"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	DatabasePort       int
	DatabaseName       string
	DatabaseHost       string
	DatabaseUsername   string
	DatabasePassword   string
	DatabaseTimezone   string
	DatabaseSslMode    string
	DatabaseLogEnabled bool
}

type TxDB struct {
	db       *bun.DB
	txMap    map[string]*bun.Tx
	mapLock  sync.Mutex
	idLock   sync.Mutex
	instance *snowflake.Node
}

func NewDBWithTX(cfg DatabaseConfig) (*TxDB, func(), error) {
	bunDB, deferFunc, err := connect(cfg)
	txdb := TxDB{
		db:    bunDB,
		txMap: map[string]*bun.Tx{},
	}
	return &txdb, deferFunc, err
}

func connect(cfg DatabaseConfig) (*bun.DB, func(), error) {
	primaryDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.DatabaseHost,
		cfg.DatabaseUsername,
		cfg.DatabasePassword,
		cfg.DatabaseName,
		cfg.DatabasePort,
		cfg.DatabaseSslMode,
		cfg.DatabaseTimezone,
	)

	sqlDB, err := sql.Open("postgres", primaryDSN)
	if err != nil {
		return nil, nil, err
	}
	db := bun.NewDB(sqlDB, pgdialect.New())

	err = db.Ping()
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			fmt.Println(err)
		}
	}

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(cfg.DatabaseLogEnabled)))

	return db, cleanup, nil
}

type transactionID struct{}

type IRunInTransaction func(ctx context.Context, fn func(ctx context.Context) bool) error

func (t *TxDB) Commit(ctx context.Context) error {
	t.mapLock.Lock()
	defer t.mapLock.Unlock()
	tID, err := t.getTransactionID(ctx)
	if err != nil {
		return err
	}
	if tx, ok := t.txMap[tID]; ok {
		err := tx.Commit()
		delete(t.txMap, tID)
		return err
	}
	return nil
}

func (t *TxDB) Rollback(ctx context.Context) error {
	t.mapLock.Lock()
	defer t.mapLock.Unlock()
	tID, err := t.getTransactionID(ctx)
	if err != nil {
		return err
	}
	if tx, ok := t.txMap[tID]; ok {
		err := tx.Rollback()
		delete(t.txMap, tID)
		return err
	}
	return nil
}

func (t *TxDB) GetTX(ctx context.Context, opts *sql.TxOptions) (bun.IDB, error) {
	t.mapLock.Lock()
	defer t.mapLock.Unlock()
	tID, err := t.getTransactionID(ctx)
	if err != nil {
		return t.db, nil
	}

	if tx, ok := t.txMap[tID]; ok {
		return tx, nil
	}

	tx, err := t.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "error in getting transaction")
	}
	t.txMap[tID] = &tx
	return &tx, nil
}

func (t *TxDB) getTransactionID(ctx context.Context) (string, error) {
	cID := ctx.Value(transactionID{})
	if cID == nil {
		return "", errors.New("correlation id is not set")
	}
	return cID.(string), nil
}

func (t *TxDB) setTransactionID(ctx context.Context) context.Context {
	id, err := t.newID()
	if err != nil {
		fmt.Println("error in generating id")
		id = fmt.Sprint(strconv.Itoa(rand.Int()), time.Now().Unix())
	}
	return context.WithValue(ctx, transactionID{}, id)
}

func (t *TxDB) RunInTransaction(ctx context.Context, fn func(ctx context.Context) bool) error {
	ctx = t.setTransactionID(ctx)
	ok := fn(ctx)
	if ok {
		return t.Commit(ctx)
	}
	return t.Rollback(ctx)
}

func (t *TxDB) RegisterModel(models ...interface{}) {
	t.db.RegisterModel(models...)
}

func (t *TxDB) newID() (string, error) {
	if t.instance == nil {
		t.idLock.Lock()
		defer t.idLock.Unlock()

		node, err := snowflake.NewNode(1)
		if err != nil {
			return "", err
		}
		t.instance = node
	}

	return t.instance.Generate().String(), nil
}
