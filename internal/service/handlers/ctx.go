package handlers

import (
	"context"
	"example.com/m/v2/internal/pkg"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int
type notificationManagerKey int

const (
	logCtxKey ctxKey = iota
	notificationManagerCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxNotificationManager(entry *pkg.NotificationManager) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, notificationManagerCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}
func GetNotificationManager(r *http.Request) *pkg.NotificationManager {
	return r.Context().Value(notificationManagerCtxKey).(*pkg.NotificationManager)
}
