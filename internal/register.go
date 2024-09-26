package internal

import (
	"github.com/ankorstore/yokai/fxhealthcheck"
	"github.com/ankorstore/yokai/fxmetrics"
	"github.com/ankorstore/yokai/orm/healthcheck"
	"go.uber.org/fx"

	"github.com/Dudeiebot/http-level/internal/repository"
	"github.com/Dudeiebot/http-level/internal/service"
)

// Register is used to register the application dependencies.
func Register() fx.Option {
	return fx.Options(
		// orm probe
		fxhealthcheck.AsCheckerProbe(healthcheck.NewOrmProbe),
		// services
		fx.Provide(
			// gophers repository
			repository.NewGopherRepository,
			// gophers service
			service.NewGopherService,
		),
		fxmetrics.AsMetricsCollector(service.GopherListCounter),
	)
}
