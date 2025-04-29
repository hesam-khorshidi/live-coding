package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
	"time"
)

type AppConfig struct {
	Name        string `mapstructure:"name" validate:"required"`
	Version     string `mapstructure:"version" validate:"required"`
	Environment string `mapstructure:"environment" validate:"oneof=development staging production"`
}

type ServerConfig struct {
	Host       string        `mapstructure:"host" validate:"required"`
	Port       string        `mapstructure:"port" validate:"required"`
	Timeout    time.Duration `mapstructure:"timeout" validate:"required"`
	Protocol   string        `mapstructure:"protocol" validate:"required,oneof=http https"`
	ApiPrefix  string        `mapstructure:"api_prefix" validate:"required"`
	ApiVersion string        `mapstructure:"api_version" validate:"required"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required,min=1,max=65535"`
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Name     string `mapstructure:"name" validate:"required"`
	SSLMode  string `mapstructure:"ssl_mode" validate:"required,oneof=disable"`
	Timezone string `mapstructure:"timezone" validate:"required"`
}

type LoggingConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Level   string `mapstructure:"level" validate:"oneof=debug info warn error"`
	Format  string `mapstructure:"format" validate:"oneof=json text"`
}

type KafkaConfig struct {
	Host      string   `mapstructure:"host" validate:"required"`
	Port      int      `mapstructure:"port" validate:"required"`
	Brokers   []string `mapstructure:"brokers" validate:"required,dive,required"`
	Topic     string   `mapstructure:"topic" validate:"required"`
	Partition int      `mapstructure:"partition" validate:"required"`
	GroupID   string   `mapstructure:"group_id" validate:"required"`
}

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Kafka    KafkaConfig    `mapstructure:"kafka"`
	Logging  LoggingConfig  `mapstructure:"logging"`
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)
	v.AddConfigPath(".")

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, errors.Wrap(err, "config validation failed")
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return &cfg, nil
}
