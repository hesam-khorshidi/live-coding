//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"live-coding/config"
	"live-coding/infra"
	shareddom "live-coding/internal/shared/adapter/inbound/http"

	userhttp "live-coding/internal/user/adapter/inbound/http"
	userfilesrv "live-coding/internal/user/adapter/outbound/file"
	userrepo "live-coding/internal/user/adapter/outbound/sql/user_repo"
	userinprt "live-coding/internal/user/core/port/inbound"
	useroutprt "live-coding/internal/user/core/port/outbound"
	usersrv "live-coding/internal/user/core/service/user_srv"
)

var inboundSet = wire.NewSet(
	userhttp.Init,
)

var infraSet = wire.NewSet(
	infra.NewHttpServer,
	infra.NewDBWithTX,
	infra.NewRedisClient,
	provideHttpDependencies,
	provideRunInTransaction,
)

var configSet = wire.NewSet(
	provideHttpServerConfig,
	provideDatabaseConfig,
	provideWorkerConfig,
	provideFileConfig,
)

var serviceSet = wire.NewSet(
	usersrv.New, wire.Bind(new(userinprt.UserService), new(usersrv.Service)),
)

var outboundSet = wire.NewSet(
	userrepo.New, wire.Bind(new(useroutprt.UserRepository), new(userrepo.Repository)),
	userfilesrv.New, wire.Bind(new(useroutprt.UserReaderService), new(userfilesrv.Service)),
)

func InitHttp(_ config.Config) (Http, func(), error) {
	wire.Build(infraSet, configSet, provideHttp, serviceSet, inboundSet, outboundSet)
	return Http{}, nil, nil
}

func provideHttpDependencies(e *echo.Echo, cfg infra.HTTPServerConfig) shareddom.Dependencies {
	return shareddom.Dependencies{
		Echo:   e,
		Prefix: cfg.ApiPrefix + cfg.ApiVersion,
		Debug:  cfg.Debug,
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

func provideFileConfig(cfg config.Config) config.FileConfig {
	return cfg.FileConfig
}

func provideWorkerConfig(cfg config.Config) config.WorkerConfig {
	return cfg.WorkerConfig
}

func provideRunInTransaction(txDB *infra.TxDB) infra.IRunInTransaction {
	return txDB.RunInTransaction
}

func provideHttp(
	server *echo.Echo,
	_ userhttp.Controller,
) Http {
	return newHttp(server)
}
