package internal

import (
	"github.com/ankorstore/yokai/fxhttpserver"
	"go.uber.org/fx"

	"github.com/Dudeiebot/http-level/internal/handler"
	"github.com/Dudeiebot/http-level/internal/handler/dude"
)

// Router is used to register the application HTTP middlewares and handlers.
func Router() fx.Option {
	return fx.Options(
		fxhttpserver.AsHandler("GET", "", handler.NewExampleHandler),
		// dude creation
		// we can do a s a group also, check it out in the docs
		fxhttpserver.AsHandler("POST", "/dudepeople", dude.NewCreateGopherHandler),
		fxhttpserver.AsHandler("GET", "/allpeople", dude.NewListGophersHandler),
	)
}
