// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"time"
)

type Ayah struct {
	ID            int64     `json:"id"`
	Page          int64     `json:"page"`
	SurahNumber   int64     `json:"surah_number"`
	Number        int64     `json:"number"`
	NumberInSurah int64     `json:"number_in_surah"`
	Ar            string    `json:"ar"`
	Tr            string    `json:"tr"`
	Idn           string    `json:"idn"`
	AudioUrl      string    `json:"audio_url"`
	CreatedAt     time.Time `json:"created_at"`
}

type Surah struct {
	ID        int64     `json:"id"`
	Page      int64     `json:"page"`
	Ar        string    `json:"ar"`
	Number    int64     `json:"number"`
	Tr        string    `json:"tr"`
	Idn       string    `json:"idn"`
	Type      string    `json:"type"`
	TotalAyah int64     `json:"total_ayah"`
	CreatedAt time.Time `json:"created_at"`
}
