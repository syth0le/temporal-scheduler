package application

import (
	"temporal-docs/internal/handler/publicapi"

	xservers "github.com/syth0le/gopnik/servers"

	"github.com/go-chi/chi/v5"
)

func (a *App) newHTTPServer(env *env) *xservers.HTTPServerWrapper {
	return xservers.NewHTTPServerWrapper(
		a.Logger,
		xservers.WithPublicServer(a.Config.PublicServer, a.publicMux(env)),
	)
}

func (a *App) publicMux(env *env) *chi.Mux {
	mux := chi.NewMux()

	handler := publicapi.NewHandler(a.Logger, env.scheduleService)

	mux.Route("/schedule", func(r chi.Router) {
		//r.Get("/{scheduleID}", handler.GetSchedule)
		//r.Get("/", handler.ListSchedule)
		r.Post("/", handler.CreateSchedule)
		//r.Delete("/{scheduleID}", handler.DeleteSchedule)
	})

	return mux
}
