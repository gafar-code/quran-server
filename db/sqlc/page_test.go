package db

import (
	"context"
	"testing"

	"github.com/gafar-code/quran-server/util"
	"github.com/stretchr/testify/require"
)

func TestGetPage(t *testing.T) {
	getPage := NewPage(testDB)
	pageID := int64(1)

	result, err := getPage.GetPage(context.Background(), pageID)
	require.NoError(t, err)

	require.Equal(t, result.PageID, result.PageID)

}

func TestCreatePage(t *testing.T) {
	getPage := NewPage(testDB)
	pageID := int64(1)

	listAyah := []CreatePageAyah{
		{
			SurahNumber:   pageID,
			Number:        util.RandomInt(1, 100),
			NumberInSurah: util.RandomInt(1, 100),
			AudioUrl:      util.RandomString(10),
			Ar:            util.RandomString(10),
			Tr:            util.RandomString(10),
			Idn:           util.RandomString(10),
		},
		{
			SurahNumber:   pageID,
			Number:        util.RandomInt(1, 100),
			NumberInSurah: util.RandomInt(1, 100),
			AudioUrl:      util.RandomString(10),
			Ar:            util.RandomString(10),
			Tr:            util.RandomString(10),
			Idn:           util.RandomString(10),
		},
		{
			SurahNumber:   pageID,
			Number:        util.RandomInt(1, 100),
			NumberInSurah: util.RandomInt(1, 100),
			AudioUrl:      util.RandomString(10),
			Ar:            util.RandomString(10),
			Tr:            util.RandomString(10),
			Idn:           util.RandomString(10),
		},
	}

	listSurah := []CreatePageSurah{
		{
			Ar:        util.RandomString(5),
			Number:    util.RandomInt(7, 200),
			Tr:        util.RandomString(10),
			Idn:       util.RandomString(10),
			Type:      util.RandomType(),
			TotalAyah: util.RandomInt(7, 200),
		},
	}

	arg := CreatePageParams{
		PageID:    pageID,
		ListAyah:  listAyah,
		ListSurah: listSurah,
	}

	result, err := getPage.CreatePage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

}
