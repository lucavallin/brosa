package main

import (
	"github.com/lucavallin/mau/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("mau")
	viper.SetConfigType("yml")
	viper.SetConfigFile("./mau.yml")
	viper.ReadInConfig()

	cmd.Execute()
}
