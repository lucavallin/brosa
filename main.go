package main

import (
	"github.com/lucavallin/brosa/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("brosa")
	viper.SetConfigType("yml")
	viper.SetConfigFile("./brosa.yml")
	viper.ReadInConfig()

	cmd.Execute()
}
