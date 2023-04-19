package repository

import (
    "Project/domain"

    "github.com/stretchr/testify/mock"
)

// mockStickerRepository is the mock implementation of StickerRepository
type mockStickerRepository struct {
    mock.Mock
}

// NewMockStickerRepository creates a new instance of mockStickerRepository
func NewMockStickerRepository() *mockStickerRepository {
    return &mockStickerRepository{}
}

// GetTrendingStickers mocks the GetTrendingStickers method of StickerRepository
func (r *mockStickerRepository) GetTrendingStickers(currentTime string, numStickers int) ([]domain.Sticker, error) {
    args := r.Called(currentTime, numStickers)
    return args.Get(0).([]domain.Sticker), args.Error(1)
}
