package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alqurantafsir/be-main-monolith/domain"
	"github.com/alqurantafsir/be-main-monolith/dto"
	"io"
	"os"
	"time"
)

type surahService struct {
	cacheRepository domain.CacheRepository
	quranJson       string
}

func NewSurah(cacheRepository domain.CacheRepository, quranJson string) domain.SurahService {
	return &surahService{
		cacheRepository: cacheRepository,
		quranJson:       quranJson,
	}
}

func (s surahService) Initialize() error {
	fQuran, err := os.Open(s.quranJson)
	if err != nil {
		return err
	}

	defer func() {
		_ = fQuran.Close()
	}()

	byteValue, _ := io.ReadAll(fQuran)

	var quran dto.Quran
	err = json.Unmarshal(byteValue, &quran)
	if err != nil {
		return err
	}

	for _, v := range quran.Data {
		surah, _ := json.Marshal(v)
		err = s.cacheRepository.Set(fmt.Sprintf("surah-%d", v.Number), surah, 15*time.Minute)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s surahService) GetSurah(ctx context.Context, number, part int) (dto.Surah, error) {
	data, err := s.cacheRepository.Get(fmt.Sprintf("surah-%d", number))
	if err != nil {
		return dto.Surah{}, err
	}

	var surah dto.Surah
	_ = json.Unmarshal(data, &surah)
	offset := part * 10

	if len(surah.Verses) < offset {
		return dto.Surah{}, nil
	}

	if len(surah.Verses) >= offset+10 {
		dte := surah.Verses[offset : offset+10]
		surah.Verses = dte
		return surah, nil
	}

	dte := surah.Verses[offset:]
	surah.Verses = dte
	return surah, nil
}
