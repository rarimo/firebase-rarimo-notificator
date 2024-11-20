package pkg

import (
	"context"
	"errors"
	"example.com/m/v2/resources"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type Project string

const (
	ProjectRarime Project = "rarime"
	UnitedSpace   Project = "united-space"
)

type NotificationManager struct {
	keys map[Project]string
}

func NewNotificationManager(keys map[Project]string) *NotificationManager {
	return &NotificationManager{keys: keys}
}

func (nm *NotificationManager) getClient(project Project) (*messaging.Client, error) {
	keyPath, exists := nm.keys[project]
	if !exists {
		return nil, errors.New("Keys not found for: " + string(project))
	}

	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(keyPath))
	if err != nil {
		return nil, errors.New("Cant init Firebase App: " + string(project))
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, errors.New("Cant create client for: " + string(project) + err.Error())
	}

	return client, nil
}

func (nm *NotificationManager) resolveMessage(request resources.NotificationRequest) *messaging.Message {
	// Create the base message
	msg := &messaging.Message{
		Topic: request.Topic,
	}

	// Helper to create Android configuration
	createAndroidConfig := func() *messaging.AndroidConfig {
		return &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Title: request.Title,
				Body:  request.Description,
			},
			Priority: "high",
			Data: map[string]string{
				"type":        request.Type,
				"content":     *request.Content,
				"title":       request.Title,
				"description": request.Description,
			},
		}
	}

	// Helper to create iOS configuration
	createIOSConfig := func() *messaging.APNSConfig {
		return &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					MutableContent: true,
					Alert: &messaging.ApsAlert{
						Title: request.Title,
						Body:  request.Description,
					},
				},
				CustomData: map[string]interface{}{
					"type":    request.Type,
					"content": request.Content,
				},
			},
		}
	}

	// Assign configurations based on the target platform
	switch request.Target {
	case "ios":
		msg.APNS = createIOSConfig()
	case "android":
		msg.Android = createAndroidConfig()
	case "ios-and-android":
		msg.Android = createAndroidConfig()
		msg.APNS = createIOSConfig()
	}

	return msg
}

// SendNotification отправляет уведомление для указанного проекта
func (nm *NotificationManager) SendNotification(project Project, request resources.NotificationRequest) error {
	client, err := nm.getClient(project)
	if err != nil {
		return err
	}

	msg := nm.resolveMessage(request)

	// Send the message
	ctx := context.Background()
	response, err := client.Send(ctx, msg)
	if err != nil {
		return err
	}

	log.Printf("Notification sent: %s\n", response)
	return nil
}
