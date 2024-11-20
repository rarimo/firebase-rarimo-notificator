package service

import (
	"example.com/m/v2/internal/service/handlers"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	corsOptions := cors.Options{
		// Allow specific origins
		AllowedOrigins:   []string{"https://*", "http://*"}, // Frontend origin
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true, // Allow cookies, authorization headers, or TLS client certificates
		MaxAge:           300,  // Maximum value for Access-Control-Max-Age in seconds
	}

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxNotificationManager(s.notificationManager),
		),
		cors.New(corsOptions).Handler,
	)

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
