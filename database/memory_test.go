package database_test

import (
	"testing"

	"whatev.com/interface/assert"
	"whatev.com/interface/database"
)

// NOTE: too lazy to test many cases

func Test_MemoryStorage_GetTokenCount(t *testing.T) {
	ms := database.NewMemoryStorage()
	ms.AddToken("kalau")

	count := ms.GetTokenCount("kalau")
	assert.Equal(t, count, 1)
}

func Test_MemoryStorage_AddToken(t *testing.T) {
	ms := database.NewMemoryStorage()
	ms.AddToken("jika")

	// NOTE: feels redundant, but in practice this only checks if no error

	count := ms.GetTokenCount("jika")
	assert.Equal(t, count, 1)
}

func Test_MemoryStorage_UpdateTokenName(t *testing.T) {
	ms := database.NewMemoryStorage()
	ms.AddToken("kapan")

	count := ms.GetTokenCount("kapan")
	assert.Equal(t, count, 1)

	ms.UpdateTokenName("kapan", "dimana")

	oldCount := ms.GetTokenCount("kapan")
	assert.Equal(t, oldCount, 0)

	newCount := ms.GetTokenCount("dimana")
	assert.Equal(t, newCount, 1)
}

func Test_MemoryStorage_UpdateTokenCount(t *testing.T) {
	ms := database.NewMemoryStorage()
	ms.AddToken("bagaimana")

	beforeCount := ms.GetTokenCount("bagaimana")
	assert.Equal(t, beforeCount, 1)

	ms.UpdateTokenCount("bagaimana", 10)

	afterCount := ms.GetTokenCount("bagaimana")
	assert.Equal(t, afterCount, 10)
}

func Test_MemoryStorage_RemoveToken(t *testing.T) {
	ms := database.NewMemoryStorage()
	ms.AddToken("apa")

	beforeCount := ms.GetTokenCount("apa")
	assert.Equal(t, beforeCount, 1)

	ms.RemoveToken("apa")

	afterCount := ms.GetTokenCount("apa")
	assert.Equal(t, afterCount, 0)
}

func Test_MemoryStorage_ResetToken(t *testing.T) {
	ms := database.NewMemoryStorage()
	ms.AddToken("ini")

	beforeCount := ms.GetTokenCount("ini")
	assert.Equal(t, beforeCount, 1)

	ms.ResetToken("ini")

	afterCount := ms.GetTokenCount("ini")
	assert.Equal(t, afterCount, 0)
}

func Test_MemoryStorage_Size(t *testing.T) {
	ms := database.NewMemoryStorage()
	ms.AddToken("kenapa")
	ms.AddToken("begitu")

	assert.Equal(t, ms.Size(), 2)

	ms.RemoveToken("kenapa")
	assert.Equal(t, ms.Size(), 1)

	ms.ResetToken("begitu")
	assert.Equal(t, ms.Size(), 1)
}
