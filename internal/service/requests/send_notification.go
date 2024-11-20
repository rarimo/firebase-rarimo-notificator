package requests

import (
	"encoding/json"
	"example.com/m/v2/resources"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
)

func NewSendNotification(r *http.Request) (req resources.NotificationRequest, err error) {

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	errs := validation.Errors{
		"content":     validation.Validate(req.Content, validation.NilOrNotEmpty), // Optional but must not be empty if provided
		"description": validation.Validate(req.Description, validation.Required, validation.Length(1, 500)),
		"target":      validation.Validate(req.Target, validation.Required, validation.In("ios", "android", "ios-and-android")),
		"title":       validation.Validate(req.Title, validation.Required, validation.Length(1, 100)),
		"topic":       validation.Validate(req.Topic, validation.Required, validation.Length(1, 100)),
		"type":        validation.Validate(req.Type, validation.Required),
	}

	return req, errs.Filter()
}

func newDecodeError(what string, err error) error {
	return validation.Errors{
		what: fmt.Errorf("decode request %s: %w", what, err),
	}
}
