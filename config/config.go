package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type dbConfig struct {
	Host     string `envconfig:"DB_HOST"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_NAME"`
	Port     int    `envconfig:"DB_PORT"`
}

type kafkaConfig struct {
	Broker  string `envconfig:"KAFKA_BROKER_URL"`
	Topic   string `envconfig:"KAFKA_TOPIC"`
	GroupId string `envconfig:"KAFKA_GROUP_ID"`
}

type appConfig struct {
	Environment string `envconfig:"ENV"`
	Host        string `envconfig:"HOST"`
}

type config struct {
	App   appConfig
	Db    dbConfig
	Kafka kafkaConfig
}

var cfg config

func Load() {
	godotenv.Load()
	envconfig.MustProcess("", &cfg)
}

func Get() config {
	return cfg
}

func GetDb() dbConfig {
	return cfg.Db
}

func GetKafka() kafkaConfig {
	return cfg.Kafka
}
