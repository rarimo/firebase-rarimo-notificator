package service

import (
	"example.com/m/v2/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxNotificationManager(s.notificationManager),
		))

	r.Route("/notifications", func(r chi.Router) {
		r.Route("/{project}", func(r chi.Router) {
			r.Post("/", handlers.SendRariMeNotification)
		})

		r.Route("/united-space", func(r chi.Router) {
			r.Post("/", handlers.SendUnitedSpaceNotification)
		})
	})

	return r
}
