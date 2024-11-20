package handlers

import (
	"example.com/m/v2/internal/pkg"
	"example.com/m/v2/internal/service/requests"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net"
	"net/http"
)

func SendRariMeNotification(w http.ResponseWriter, r *http.Request) {
	chi.URLParam(r, "project")
	request, err := requests.NewSendNotification(r)
	if err != nil {
		Log(r).WithError(err).Error("NewRequestSendNotification")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	notificationManger := GetNotificationManager(r)

	err = notificationManger.SendNotification(pkg.ProjectRarime, request)

	if err != nil {
		Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, net.Interface{})
}

func SendUnitedSpaceNotification(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewSendNotification(r)
	if err != nil {
		Log(r).Error(err)
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	notificationManger := GetNotificationManager(r)

	err = notificationManger.SendNotification(pkg.UnitedSpace, request)

	if err != nil {
		Log(r).Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, net.Interface{})
}
