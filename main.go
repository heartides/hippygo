package main

import (
	"golang.org/x/sync/errgroup"
	_ "learning/conf"
	"learning/routers"
	"log"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	apiRouter := &http.Server{
		Addr:         ":3030",
		Handler:      routers.InitApiRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return apiRouter.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
