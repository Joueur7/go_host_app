package usecase

import (
	"Project/domain"
	"Project/sticker/repository"
)

// StickerUsecase is the usecase for sticker related actions
type StickerUsecase struct {
    stickerRepository repository.StickerRepository
}

// NewStickerUsecase creates a new instance of StickerUsecase
func NewStickerUsecase(sr repository.StickerRepository) *StickerUsecase {
    return &StickerUsecase{sr}
}

// GetTrendingStickers returns the top numStickers trending stickers
func (su *StickerUsecase) GetTrendingStickers(currentTime string, numStickers int) ([]domain.Sticker, error) {
    return su.stickerRepository.GetTrendingStickers(currentTime, numStickers)
}