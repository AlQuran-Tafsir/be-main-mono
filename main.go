package main

import (
	"flag"
	"github.com/alqurantafsir/be-main-monolith/internal/api"
	"github.com/alqurantafsir/be-main-monolith/internal/repository"
	"github.com/alqurantafsir/be-main-monolith/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	jsonLocation := flag.String("env", "", "Lokasi file env")
	flag.Parse()

	location := *jsonLocation
	if location == "" {
		location = "quran.json"
	}

	cacheRepository := repository.NewBigCache()
	surahService := service.NewSurah(cacheRepository, location)

	err := surahService.Initialize()
	if err != nil {
		log.Fatal("failed when initialize: ", err.Error())
	}

	app := fiber.New()
	app.Use(cors.New())
	api.NewSurah(app, surahService)

	err = app.Listen(":2599")
	if err != nil {
		log.Fatal(err.Error())
	}
}
