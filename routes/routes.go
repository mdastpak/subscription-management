package routes

import (
	"context"
	"net/http"
	"subscription-management/handlers"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewRouter() *mux.Router {
	return mux.NewRouter()
}

func RegisterRoutes(lc fx.Lifecycle, router *mux.Router, adminHandler *handlers.AdminHandler, logger *zap.Logger) {
	router.HandleFunc("/admin/service", adminHandler.CreateServiceHandler).Methods("POST")
	router.HandleFunc("/admin/service", adminHandler.UpdateServiceHandler).Methods("PUT")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("Starting HTTP server", zap.String("address", ":8080"))
				if err := http.ListenAndServe(":8080", router); err != nil {
					logger.Fatal("Failed to start HTTP server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down HTTP server")
			return nil
		},
	})
}
