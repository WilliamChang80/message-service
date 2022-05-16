package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	appPort    string
	dbPort     int
	dbUrl      string
	dbUser     string
	dbPassword string
	dbName     string
	dbKeyspace string
}

var appConfig *Config

func readConfig() {
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

func Init() {
	readConfig()

	appConfig = &Config{
		appPort:    strconv.Itoa(viper.GetInt("app.port")),
		dbPort:     viper.GetInt("database.port"),
		dbUrl:      viper.GetString("database.url"),
		dbUser:     viper.GetString("database.user"),
		dbPassword: viper.GetString("database.password"),
		dbName:     viper.GetString("database.name"),
		dbKeyspace: viper.GetString("database.keyspace"),
	}
}

func GetAppPort() string {
	return appConfig.appPort
}

func GetDbPort() int {
	return appConfig.dbPort
}

func GetDbUrl() string {
	return appConfig.dbUrl
}

func GetDbUser() string {
	return appConfig.dbUser
}

func GetDbPassword() string {
	return appConfig.dbPassword
}

func GetDbName() string {
	return appConfig.dbName
}

func GetDbKeyspace() string {
	return appConfig.dbKeyspace
}
