package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// Инициализация конфигурации
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	// Установление префикса для переменных окружения
	viper.SetEnvPrefix("myapp") // viper будет искать переменные окружения с префиксом "MYAPP_"
	viper.AutomaticEnv()

	// Чтение конфигурации
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s", err)
	}

	// Print all keys
	fmt.Println(viper.AllKeys())
	// Print all settings
	fmt.Println(viper.AllSettings())
}
