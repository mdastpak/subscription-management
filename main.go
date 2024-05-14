package main

import (
	"fmt"
	"log"
	"net/http"
	"subscription-management/routes"
	"subscription-management/utils"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func main() {
	initConfig()
	utils.InitRedis()

	port := viper.GetString("server.port")
	router := routes.InitRoutes()

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
