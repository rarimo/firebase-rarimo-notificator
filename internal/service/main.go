package service

import (
	"example.com/m/v2/internal/config"
	"example.com/m/v2/internal/pkg"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net"
	"net/http"
)

type service struct {
	log                 *logan.Entry
	copus               types.Copus
	listener            net.Listener
	notificationManager *pkg.NotificationManager
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	adminSDKPaths := cfg.AdminSDKPaths()

	rarimePath := adminSDKPaths["rarime"]
	unitedSpacePath := adminSDKPaths["united-space"]

	keys := map[pkg.Project]string{
		pkg.ProjectRarime: rarimePath,
		pkg.UnitedSpace:   unitedSpacePath,
	}

	return &service{
		log:                 cfg.Log(),
		copus:               cfg.Copus(),
		listener:            cfg.Listener(),
		notificationManager: pkg.NewNotificationManager(keys),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
