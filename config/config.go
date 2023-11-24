package config

import (
	"log"

	"github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	DataConfig  DataConfig
	AppConfig   AppConfig
	RedisConfig RedisConfig
	AuthConfig  AuthConfig
}
type DataConfig struct {
	DbHost     string `env:"DB_Host,required=true"`
	DbUser     string `env:"DB_User,required=true"`
	DbPassword string `env:"DB_Password,required=true"`
	DbName     string `env:"DB_Name,required=true"`
	DbPort     string `env:"DB_Port,required=true"`
	Dbsslmode  string `env:"DB_Sslmode,required=true"`
	DbTimeZone string `env:"DB_TimeZone,required=true"`
}

type AppConfig struct {
	AppHost         string `env:"APP_HOST"`
	AppPort         string `env:"APP_PORT,required=true"`
	AppReadTimeout  uint32 `env:"APP_READ_TIMEOUT,required=true"`
	AppWriteTimeout uint32 `env:"APP_WRITE_TIMEOUT,required=true"`
	AppIdleTimeout  uint32 `env:"APP_IDLE_TIMEOUT,required=true"`
}

type RedisConfig struct {
	RedisPort     string `env:"REDiS_PORT,required=true"`
	RedisPassword string `env:"REDIS_PASSWORD,required=true"`
	RedisDb       int    `env:"REDIS_DB,required=true"`
}

type AuthConfig struct {
	PrivateKey string `env:"PRIVATE_KEY,required=true"`
	PublicKey  string `env:"PUBLIC_KEY,required=true"`
}

func init() {
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Panic(err)
	}
}

func GetConfig() Config {
	return cfg
}
