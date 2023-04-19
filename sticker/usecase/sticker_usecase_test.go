package usecase_test

import (
	"Project/domain"
	"Project/sticker/usecase"
	"errors"
	"reflect"
	"testing"
)

type mockStickerRepository struct {
	mockGetTrendingStickers func(currentTime string, numStickers int) ([]domain.Sticker, error)
}

func (m *mockStickerRepository) GetTrendingStickers(currentTime string, numStickers int) ([]domain.Sticker, error) {
	return m.mockGetTrendingStickers(currentTime, numStickers)
}

func TestGetTrendingStickers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		mockRepo      *mockStickerRepository
		currentTime   string
		numStickers   int
		expected      []domain.Sticker
		expectedError error
	}{
		{
			name: "Success",
			mockRepo: &mockStickerRepository{
				mockGetTrendingStickers: func(currentTime string, numStickers int) ([]domain.Sticker, error) {
					return []domain.Sticker{
						{ID: 1, Name: "Sticker 1", ImageURL: "http://example.com/sticker1.png"},
						{ID: 2, Name: "Sticker 2", ImageURL: "http://example.com/sticker2.png"},
					}, nil
				},
			},
			currentTime: "2023-04-19 13:37:00",
			numStickers: 2,
			expected: []domain.Sticker{
				{ID: 1, Name: "Sticker 1", ImageURL: "http://example.com/sticker1.png"},
				{ID: 2, Name: "Sticker 2", ImageURL: "http://example.com/sticker2.png"},
			},
			expectedError: nil,
		},
		{
			name: "Error in retrieval",
			mockRepo: &mockStickerRepository{
				mockGetTrendingStickers: func(currentTime string, numStickers int) ([]domain.Sticker, error) {
					return nil, errors.New("failed to fetch stickers from database")
				},
			},
			currentTime:   "2023-04-19 13:37:00",
			numStickers:   2,
			expected:      nil,
			expectedError: errors.New("failed to fetch stickers from database"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := usecase.NewStickerUsecase(test.mockRepo)

			test.mockRepo.mockGetTrendingStickers = func(currentTime string, numStickers int) ([]domain.Sticker, error) {
				return test.expected, test.expectedError
			}

			stickers, err := uc.GetTrendingStickers(test.currentTime, test.numStickers)

			if !reflect.DeepEqual(stickers, test.expected) {
				t.Errorf("Expected stickers to be %v but got %v", test.expected, stickers)
			}

			if !errors.Is(err, test.expectedError) {
				t.Errorf("Expected error to be %v but got %v", test.expectedError, err)
			}
		})
	}
}
