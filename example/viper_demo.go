package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql   MysqlUrl   `json:"mysql"`
	Mongodb MongodbUrl `json:"mongodb"`
}
type MongodbUrl struct {
	Url string
}

type MysqlUrl struct {
	User     string
	Password string
	Host     string
	Port     string
	Url      string
}

func main() {
	var Conf Config
	viper.SetConfigType("toml")
	//viper.AddConfigPath("F:\\code\\002Golang\\006Goinception\\conf")
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unable to decode into struct：  %s \n", err))
	}
	fmt.Println(Conf)
	url := viper.GetString("mysql.url")
	fmt.Println(url)
}
