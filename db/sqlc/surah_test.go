package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gafar-code/quran-server/util"

	"github.com/stretchr/testify/require"
)

func createRandomSurah(t *testing.T) Surah {
	arg := CreateSurahParams{
		Page:      1,
		Ar:        util.RandomString(5),
		Tr:        util.RandomString(5),
		Number:    util.RandomInt(7, 200),
		Idn:       util.RandomString(10),
		Type:      util.RandomType(),
		TotalAyah: util.RandomInt(7, 200),
	}

	surah, err := testQueries.CreateSurah(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, surah)

	require.Equal(t, arg.Ar, surah.Ar)
	require.Equal(t, arg.Number, surah.Number)
	require.Equal(t, arg.Tr, surah.Tr)
	require.Equal(t, arg.Idn, surah.Idn)
	require.Equal(t, arg.Type, surah.Type)
	require.Equal(t, arg.TotalAyah, surah.TotalAyah)

	require.NotZero(t, surah.ID)
	require.NotZero(t, surah.CreatedAt)

	return surah
}

func TestCreateSurah(t *testing.T) {
	createRandomSurah(t)
}

func TestGetSurah(t *testing.T) {
	surah1 := createRandomSurah(t)
	surah2, err := testQueries.GetSurah(context.Background(), surah1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, surah2)

	require.Equal(t, surah1.ID, surah2.ID)
	require.Equal(t, surah1.Ar, surah2.Ar)
	require.Equal(t, surah1.Number, surah2.Number)
	require.Equal(t, surah1.Tr, surah2.Tr)
	require.Equal(t, surah1.Idn, surah2.Idn)
	require.Equal(t, surah1.Type, surah2.Type)
	require.Equal(t, surah1.TotalAyah, surah2.TotalAyah)
	require.WithinDuration(t, surah1.CreatedAt, surah2.CreatedAt, time.Second)
}

func TestUpdateSurah(t *testing.T) {
	surah1 := createRandomSurah(t)

	arg := UpdateSurahParams{
		ID:        surah1.ID,
		Ar:        util.RandomString(5),
		Number:    util.RandomInt(1, 200),
		Tr:        util.RandomString(10),
		Idn:       util.RandomString(10),
		Type:      util.RandomType(),
		TotalAyah: util.RandomInt(7, 200),
	}

	surah2, err := testQueries.UpdateSurah(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, surah2)

	require.Equal(t, arg.Ar, surah2.Ar)
	require.Equal(t, arg.Number, surah2.Number)
	require.Equal(t, arg.Tr, surah2.Tr)
	require.Equal(t, arg.Idn, surah2.Idn)
	require.Equal(t, arg.Type, surah2.Type)
	require.Equal(t, arg.TotalAyah, surah2.TotalAyah)

	require.WithinDuration(t, surah1.CreatedAt, surah2.CreatedAt, time.Second)

}

func TestDeleteSurah(t *testing.T) {
	surah1 := createRandomSurah(t)

	err := testQueries.DeleteSurah(context.Background(), surah1.ID)
	require.NoError(t, err)

	surah2, err := testQueries.GetSurah(context.Background(), surah1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, surah2)
}

func TestListSurah(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomSurah(t)
	}

	arg := ListSurahParams{
		Limit:  5,
		Offset: 5,
	}

	surah1, err := testQueries.ListSurah(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, surah1, 5)

	for _, surah := range surah1 {
		require.NotEmpty(t, surah)
	}
}

func TestListSurahByPage(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomSurah(t)
	}

	surah, err := testQueries.ListSurahByPage(context.Background(), 1)

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(surah), 10)

	for _, entry := range surah {
		require.NotEmpty(t, entry)
	}
}
