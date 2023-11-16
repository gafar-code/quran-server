package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gafar-code/quran-server/util"
	"github.com/stretchr/testify/require"
)

func createRandomAyah(t *testing.T) Ayah {
	for i := 0; i < 5; i++ {
		createRandomSurah(t)
	}

	argSurrah := ListSurahParams{
		Limit:  1,
		Offset: 5,
	}

	surah, err := testQueries.ListSurah(context.Background(), argSurrah)
	require.NoError(t, err)

	arg := CreateAyahParams{
		Page:        1,
		SurahNumber: surah[0].ID,
		Number:      util.RandomInt(1, 100),
		Ar:          util.RandomString(10),
		AudioUrl:    util.RandomString(10),
		Tr:          util.RandomString(10),
		Idn:         util.RandomString(10),
	}

	ayah, err := testQueries.CreateAyah(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, ayah)

	require.Equal(t, arg.Page, ayah.Page)
	require.Equal(t, arg.SurahNumber, ayah.SurahNumber)
	require.Equal(t, arg.Number, ayah.Number)
	require.Equal(t, arg.Ar, ayah.Ar)
	require.Equal(t, arg.Tr, ayah.Tr)
	require.Equal(t, arg.Idn, ayah.Idn)
	require.Equal(t, arg.AudioUrl, ayah.AudioUrl)

	require.NotZero(t, ayah.ID)
	require.NotZero(t, ayah.CreatedAt)

	return ayah
}

func TestCreateAyah(t *testing.T) {
	createRandomAyah(t)
}

func TestGetAyah(t *testing.T) {
	ayah1 := createRandomAyah(t)

	ayah2, err := testQueries.GetAyah(context.Background(), ayah1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, ayah2)

	require.Equal(t, ayah1.ID, ayah2.ID)
	require.WithinDuration(t, ayah1.CreatedAt, ayah2.CreatedAt, time.Second)
}

func TestUpdateAyah(t *testing.T) {
	ayah1 := createRandomAyah(t)

	arg := UpdateAyahParams{
		ID:       ayah1.ID,
		Page:     util.RandomPage(),
		Number:   util.RandomInt(1, 100),
		Ar:       util.RandomString(10),
		AudioUrl: util.RandomString(10),
		Tr:       util.RandomString(10),
		Idn:      util.RandomString(10),
	}

	ayah2, err := testQueries.UpdateAyah(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, ayah2)

	require.Equal(t, ayah1.ID, ayah2.ID)

	require.NotZero(t, ayah2.ID)
	require.NotZero(t, ayah2.CreatedAt)
}

func TestDeleteAyah(t *testing.T) {
	ayah1 := createRandomAyah(t)

	err := testQueries.DeleteAyah(context.Background(), ayah1.ID)
	require.NoError(t, err)

	ayah2, err := testQueries.GetAyah(context.Background(), ayah1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ayah2)
}

func TestListAllAyah(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAyah(t)
	}

	arg := ListAllAyahParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListAllAyah(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, ayah := range entries {
		require.NotEmpty(t, ayah)
	}
}

func TestListAyahByPage(t *testing.T) {
	list, err := testQueries.ListAyahByPage(context.Background(), 1)

	require.NoError(t, err)

	for _, ayah := range list {
		require.NotEmpty(t, ayah)
	}
}
