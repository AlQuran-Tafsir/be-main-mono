package api

import (
	"context"
	"github.com/alqurantafsir/be-main-monolith/domain"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"time"
)

type surahApi struct {
	surahService domain.SurahService
}

func NewSurah(app *fiber.App, surahService domain.SurahService) {
	sa := surahApi{
		surahService: surahService,
	}

	app.Get("/surah/:id", sa.GetSurah)
}

func (s surahApi) GetSurah(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	number, _ := strconv.Atoi(ctx.Params("id"))
	part, _ := strconv.Atoi(ctx.Query("part"))

	surah, err := s.surahService.GetSurah(c, number, part)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(surah)
}
