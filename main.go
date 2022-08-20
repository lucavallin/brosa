package main

import (
	"github.com/lucavallin/mau/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("mau")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	cmd.Execute()
}
