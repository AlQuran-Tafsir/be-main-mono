package domain

import (
	"context"
	"github.com/alqurantafsir/be-main-monolith/dto"
)

type SurahService interface {
	Initialize() error
	GetSurah(ctx context.Context, number, part int) (dto.Surah, error)
}
