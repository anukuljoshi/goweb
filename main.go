package main

import (
	"context"
	"fmt"
	"log"

	"github.com/anukuljoshi/goweb/cmd/api"
	"github.com/anukuljoshi/goweb/config"
	"github.com/anukuljoshi/goweb/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	dbPool, err := connectToDB(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer dbPool.Close()

	service, err := services.NewService(dbPool)
	if err != nil {
		log.Fatalln(err)
	}
	api, err := api.NewApi(cfg, service)
	if err != nil {
		log.Fatalln(err)
	}

	api.Server.Logger.Fatal(
		api.Server.Start(fmt.Sprintf(":%d", cfg.Port)),
	)
}

func connectToDB(cfg *config.Config) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), cfg.DSN)
	if err != nil {
		return nil, err
	}
	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return dbPool, nil
}
