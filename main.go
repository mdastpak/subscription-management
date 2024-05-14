package main

import (
	"log"
	"subscription-management/handlers"
	"subscription-management/routes"
	"subscription-management/utils"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
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
	utils.InitLogger()
	logger := utils.GetLogger()

	app := fx.New(
		fx.Provide(
			utils.NewRedisClient,
			routes.NewRouter,
			handlers.NewAdminHandler,
			func() *zap.Logger {
				return logger
			},
		),
		fx.Invoke(
			routes.RegisterRoutes,
		),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)

	app.Run()
	utils.SyncLogger()
}
