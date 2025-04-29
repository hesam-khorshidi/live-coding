//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"nightingale/config"
	"nightingale/infra"
	shareddom "nightingale/internal/shared/adapter/inbound/http"
)

var inboundSet = wire.NewSet()

var infraSet = wire.NewSet(
	infra.NewHttpServer,
	infra.NewDBWithTX,
	infra.NewKafkaClient,
	infra.NewRedisClient,
	provideHttpDependencies,
	provideRunInTransaction,
)

var configSet = wire.NewSet(
	provideHttpServerConfig,
	provideDatabaseConfig,
	provideKafkaConfig,
)

var serviceSet = wire.NewSet()

func InitHttp(_ config.Config) (Http, func(), error) {
	wire.Build(infraSet, configSet, provideHttp)
	return Http{}, nil, nil
}

func provideHttpDependencies(e *echo.Echo, cfg infra.HTTPServerConfig) shareddom.Dependencies {
	//Add middlewares
	return shareddom.Dependencies{
		Echo:        e,
		Middlewares: shareddom.Middlewares{},
		Prefix:      cfg.ApiPrefix + cfg.ApiVersion,
		Debug:       cfg.Debug,
	}
}

func provideDatabaseConfig(cfg config.Config) infra.DatabaseConfig {
	return infra.DatabaseConfig{
		DatabasePort:       cfg.Database.Port,
		DatabaseHost:       cfg.Database.Host,
		DatabaseName:       cfg.Database.Name,
		DatabaseUsername:   cfg.Database.Username,
		DatabasePassword:   cfg.Database.Password,
		DatabaseTimezone:   cfg.Database.Timezone,
		DatabaseSslMode:    cfg.Database.SSLMode,
		DatabaseLogEnabled: cfg.Logging.Enabled,
	}
}

func provideHttpServerConfig(cfg config.Config) infra.HTTPServerConfig {
	return infra.HTTPServerConfig{
		Debug:      cfg.App.Environment == "dev",
		Protocol:   cfg.Server.Protocol,
		Host:       cfg.Server.Host,
		Port:       cfg.Server.Port,
		ApiPrefix:  cfg.Server.ApiPrefix,
		ApiVersion: cfg.Server.ApiVersion,
		LogEnable:  cfg.Logging.Enabled,
	}
}

func provideKafkaConfig(cfg config.Config) config.KafkaConfig {
	return cfg.Kafka
}

func provideRunInTransaction(txDB *infra.TxDB) infra.IRunInTransaction {
	return txDB.RunInTransaction
}

func provideHttp(
	server *echo.Echo,
) Http {
	return newHttp(server)
}
