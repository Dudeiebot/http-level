package service

import (
	"context"

	"github.com/ankorstore/yokai/log"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/Dudeiebot/http-level/internal/model"
	"github.com/Dudeiebot/http-level/internal/repository"
)

var GopherListCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "gophers_list_total",
	Help: "The number of times gophers were listed",
})

type GopherService struct {
	repository *repository.GopherRepository
}

func NewGopherService(repository *repository.GopherRepository) *GopherService {
	return &GopherService{
		repository: repository,
	}
}

func (s *GopherService) Create(ctx context.Context, gopher *model.Dude) error {
	return s.repository.Create(ctx, gopher)
}

func (s *GopherService) List(ctx context.Context) ([]model.Dude, error) {
	log.CtxLogger(ctx).Info().Msg("called GopherService.List()")

	GopherListCounter.Inc()

	return s.repository.FindAll(ctx)
}
