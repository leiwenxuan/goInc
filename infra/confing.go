package infra

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql     MysqlUrl     `json:"mysql"`
	Mongodb   MongodbUrl   `json:"mongodb"`
	Inception IncServerUrl `json:"inception"`
}
type MongodbUrl struct {
	Url string
}
type IncServerUrl struct {
	Url string
}

type MysqlUrl struct {
	User     string
	Password string
	Host     string
	Port     string
	Url      string
}

var Conf Config

func InitConfig() {
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unable to decode into struct：  %s \n", err))
	}
	fmt.Println(Conf)
}
