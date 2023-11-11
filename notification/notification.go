package notification

import (
	"context"
	"log"
	requestdomain "signature-app/domain/request_domain"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type notification struct {
	Client *messaging.Client
	Ctx    context.Context
}

type NotificationInterface interface {
	SendNotification(data requestdomain.NotificationData)
}

func NewNotification(ctx context.Context) (NotificationInterface, error) {
	opt := option.WithCredentialsFile("firebase_credential.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	cl, err := app.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	return &notification{
		Client: cl,
		Ctx:    ctx,
	}, nil
}

func (n *notification) SendNotification(data requestdomain.NotificationData) {
	message := &messaging.Message{
		Data:  data.Data,
		Token: data.Token,
		Notification: &messaging.Notification{
			Title: data.Data["title"],
			Body:  data.Data["body"],
		},
	}

	log.Printf("notification message %+v\n", message)

	resp, err := n.Client.Send(n.Ctx, message)
	if err != nil {
		log.Println("error push notification: ", err)
	}

	log.Println("send notif response : ", resp)
}
