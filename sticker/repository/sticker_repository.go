package repository

import (
    "Project/domain"
    "fmt"

    "gorm.io/gorm"
)

// StickerRepository is the interface that describes the methods
// that a sticker repository should implement.
type StickerRepository interface {
    GetTrendingStickers(currentTime string, numStickers int) ([]domain.Sticker, error)
}

// mysqlStickerRepository is the MySQL implementation of StickerRepository
type mysqlStickerRepository struct {
    db *gorm.DB
}

// NewMySQLStickerRepository creates a new instance of mysqlStickerRepository
func NewMySQLStickerRepository(db *gorm.DB) *mysqlStickerRepository {
    return &mysqlStickerRepository{db}
}

// GetTrendingStickers retrieves the top numStickers trending stickers as of currentTime
func (r *mysqlStickerRepository) GetTrendingStickers(currentTime string, numStickers int) ([]domain.Sticker, error) {
    var stickers []domain.Sticker
    err := r.db.Table("trending_stickers").
        Where("start_time <= ? AND end_time >= ?", currentTime, currentTime).
        Order("priority DESC").
        Limit(numStickers).
        Find(&stickers).
        Error
    if err != nil {
        return nil, fmt.Errorf("failed to fetch stickers: %w", err)
    }
    return stickers, nil
}
