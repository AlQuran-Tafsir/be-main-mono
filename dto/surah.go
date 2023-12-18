package dto

type Quran struct {
	Data []Surah `json:"data"`
}

type Surah struct {
	Number         int `json:"number"`
	Sequence       int `json:"sequence"`
	NumberOfVerses int `json:"numberOfVerses"`
	Name           struct {
		Short           string `json:"short"`
		Long            string `json:"long"`
		Transliteration struct {
			En string `json:"en"`
			Id string `json:"id"`
		} `json:"transliteration"`
		Translation struct {
			En string `json:"en"`
			Id string `json:"id"`
		} `json:"translation"`
	} `json:"name"`
	Revelation struct {
		Arab string `json:"arab"`
		En   string `json:"en"`
		Id   string `json:"id"`
	} `json:"revelation"`
	Verses []struct {
		Number struct {
			InQuran int `json:"inQuran"`
			InSurah int `json:"inSurah"`
		} `json:"number"`
		Meta struct {
			Juz         int `json:"juz"`
			Page        int `json:"page"`
			Manzil      int `json:"manzil"`
			Ruku        int `json:"ruku"`
			HizbQuarter int `json:"hizbQuarter"`
			Sajda       struct {
				Recommended bool `json:"recommended"`
				Obligatory  bool `json:"obligatory"`
			} `json:"sajda"`
		} `json:"meta"`
		Text struct {
			Arab            string `json:"arab"`
			Transliteration struct {
				En string `json:"en"`
			} `json:"transliteration"`
		} `json:"text"`
		Translation struct {
			En string `json:"en"`
			Id string `json:"id"`
		} `json:"translation"`
		Audio struct {
			Primary   string   `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"audio"`
		Tafsir struct {
			Id struct {
				Short string `json:"short"`
				Long  string `json:"long"`
			} `json:"id"`
		} `json:"tafsir"`
	} `json:"verses"`
}
