package cmd

import (
	"context"
	"github.com/Anarr/entain/internal/config"
	"github.com/Anarr/entain/internal/database"
	"github.com/Anarr/entain/internal/handler"
	"github.com/Anarr/entain/internal/manager"
	"github.com/Anarr/entain/internal/repository"
	"github.com/Anarr/entain/internal/task"
	"github.com/Anarr/entain/internal/util/echoutil"
	"log"
	"time"
)

func Execute() {
	cfg, err := config.New()
	failOnErr(err)

	//int db
	db := database.New(cfg)
	err = db.Connect()
	failOnErr(err)
	//start migration
	err = database.Migrate(db.DB, cfg.Postgres.Name, "./internal/database/migrations")
	failOnErr(err)

	r := repository.New(db.DB) //init repo
	m := manager.New(r)        //init manager

	//start task processor
	processor := task.NewProcessor(m)
	go processor.Start(context.Background(), time.Second*time.Duration(cfg.TaskInterval))

	h := handler.New(m)
	//run http server
	echoServer := getHTTPServer()
	echoServer.HTTPErrorHandler = echoutil.CustomHTTPErrorHandler
	h.RegisterRoutes(echoServer.Group("api"))
	initSwagger(echoServer)

	echoServer.Logger.Fatal(echoServer.Start(":" + cfg.Port))
}

func failOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
