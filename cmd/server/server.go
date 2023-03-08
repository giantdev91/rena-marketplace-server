package main

import (
	"context"
	"fmt"
	"net/http"
	"rena-server-v2/internal/account"
	accountDB "rena-server-v2/internal/account/database"
	"rena-server-v2/internal/article"
	articleDB "rena-server-v2/internal/article/database"
	"rena-server-v2/internal/collection"
	collectionDB "rena-server-v2/internal/collection/database"
	"rena-server-v2/internal/config"
	"rena-server-v2/internal/database"
	"rena-server-v2/internal/event"
	"rena-server-v2/internal/item"
	itemDB "rena-server-v2/internal/item/database"
	"rena-server-v2/internal/marketplace"
	marketplaceDB "rena-server-v2/internal/marketplace/database"
	"rena-server-v2/internal/registry"
	registryDB "rena-server-v2/internal/registry/database"
	"rena-server-v2/internal/token"
	tokenDB "rena-server-v2/internal/token/database"
	"rena-server-v2/pkg/logging"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func newServer(lc fx.Lifecycle, cfg *config.Config) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerConfig.Port),
		Handler: r,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logging.FromContext(ctx).Infof("Start to rest api server :%d", cfg.ServerConfig.Port)
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logging.FromContext(ctx).Infof("Stopped rest api server")
			return srv.Shutdown(ctx)
		},
	})
	return r
}

func printAppInfo(cfg *config.Config) {
	logging.DefaultLogger().Infow("application information", "config", cfg)
}

func loadConfig() (*config.Config, error) {
	return config.Load(configFile)
}

func runApplication() {
	// setup application(di + run server)
	app := fx.New(
		fx.Provide(
			// load config
			loadConfig,
			// setup database
			database.NewDatabase,
			// setup account packages
			accountDB.NewAccountDB,
			account.NewAuthMiddleware,
			account.NewHandler,
			// setup article packages
			articleDB.NewArticleDB,
			article.NewHandler,
			// setup collection packages
			collectionDB.NewCollectionDB,
			collection.NewHandler,
			// setup item packages
			itemDB.NewItemDB,
			item.NewHandler,
			// setup event
			event.NewHandler,
			// setup marketplace
			marketplace.NewHandler,
			marketplaceDB.NewMarketplaceDB,
			// setup registry
			registryDB.NewRegistryDB,
			registry.NewHandler,
			// setup token
			tokenDB.NewTokenDB,
			token.NewHandler,
			// server
			newServer,
		),
		fx.Invoke(
			account.RouteV1,
			article.RouteV1,
			collection.RouteV1,
			item.RouteV1,
			event.RouteV1,
			marketplace.RouteV1,
			registry.RouteV1,
			token.RouteV1,
			printAppInfo,
		),
	)
	app.Run()
}
