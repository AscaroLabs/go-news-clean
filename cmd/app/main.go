package main

import (
	"log"

	"github.com/xlab/closer"

	"go-news-clean/internal/app"
)

func main() {
	closer.Bind(func() {
		log.Print("Stop running...")
	})
	a := app.App{}
	go a.Run()
	closer.Hold()
}
