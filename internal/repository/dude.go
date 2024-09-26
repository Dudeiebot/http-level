package repository

import (
	"context"
	"sync"

	"gorm.io/gorm"

	"github.com/Dudeiebot/http-level/internal/model"
)

type GopherRepository struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewGopherRepository(db *gorm.DB) *GopherRepository {
	return &GopherRepository{
		db: db,
	}
}

func (r *GopherRepository) Create(ctx context.Context, gopher *model.Dude) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	res := r.db.WithContext(ctx).Create(gopher)

	return res.Error
}

func (r *GopherRepository) FindAll(ctx context.Context) ([]model.Dude, error) {
	var gophers []model.Dude

	res := r.db.WithContext(ctx).Find(&gophers)
	if res.Error != nil {
		return nil, res.Error
	}

	return gophers, nil
}
