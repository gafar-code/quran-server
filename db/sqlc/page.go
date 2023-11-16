package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Page struct {
	*Queries
	db *sql.DB
}

func NewPage(db *sql.DB) *Page {
	return &Page{
		db:      db,
		Queries: New(db),
	}
}

func (store *Page) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type GetPageResult struct {
	PageID    int64   `json:"page_id"`
	ListAyah  []Ayah  `json:"list_ayah"`
	ListSurah []Surah `json:"list_surah"`
}

type CreatePageParams struct {
	PageID    int64             `json:"page_id"`
	ListAyah  []CreatePageAyah  `json:"list_ayah"`
	ListSurah []CreatePageSurah `json:"list_surah"`
}

type CreatePageAyah struct {
	SurahNumber   int64  `json:"surah_number"`
	Number        int64  `json:"number"`
	NumberInSurah int64  `json:"number_in_surah"`
	Ar            string `json:"ar"`
	Tr            string `json:"tr"`
	Idn           string `json:"idn"`
	AudioUrl      string `json:"audio_url"`
}

type CreatePageSurah struct {
	Number    int64  `json:"number"`
	Ar        string `json:"ar"`
	Tr        string `json:"tr"`
	Idn       string `json:"idn"`
	Type      string `json:"type"`
	TotalAyah int64  `json:"total_ayah"`
}

func (page *Page) CreatePage(ctx context.Context, arg CreatePageParams) (GetPageResult, error) {
	var result GetPageResult

	err := page.execTx(ctx, func(q *Queries) error {

		var listAyah []Ayah
		var listSurrah []Surah

		for _, item := range arg.ListAyah {
			ayah, err := q.CreateAyah(ctx, CreateAyahParams{
				Page:          arg.PageID,
				SurahNumber:   item.SurahNumber,
				Number:        item.Number,
				NumberInSurah: item.NumberInSurah,
				Ar:            item.Ar,
				Tr:            item.Tr,
				Idn:           item.Idn,
				AudioUrl:      item.AudioUrl,
			})
			if err != nil {
				return err
			}

			listAyah = append(listAyah, ayah)
		}

		result.ListAyah = listAyah

		for _, item := range arg.ListSurah {
			surahArg := CreateSurahParams{
				Page:      arg.PageID,
				Number:    item.Number,
				Type:      item.Type,
				TotalAyah: item.TotalAyah,
				Ar:        item.Ar,
				Tr:        item.Tr,
				Idn:       item.Idn,
			}

			check, err := q.GetSurahByNumber(ctx, item.Number)
			if err == nil {
				continue
			}

			if check.ID == 0 {
				surah, err := q.CreateSurah(ctx, surahArg)
				if err != nil {
					return err
				}

				listSurrah = append(listSurrah, surah)
			}

		}

		result.ListSurah = listSurrah

		result.PageID = arg.PageID

		return nil
	})

	return result, err
}

func (page *Page) GetPage(ctx context.Context, pageID int64) (GetPageResult, error) {
	var result GetPageResult

	err := page.execTx(ctx, func(q *Queries) error {
		var err error

		result.ListAyah, err = q.ListAyahByPage(ctx, pageID)
		if err != nil {
			return err
		}

		result.ListSurah, err = q.ListSurahByPage(ctx, pageID)
		if err != nil {
			return err
		}

		result.PageID = pageID

		return nil
	})

	return result, err
}
