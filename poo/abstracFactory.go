package main

import "fmt"

// SMS, EMAIL

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}
type ISender interface {
   GetSenderMethod() string
   GetSenderChannel() string
}

type SmsNotification struct {
}

type EmailNotification struct {
}

func (SmsNotification) SendNotification()  {
	fmt.Println("Mensaje enviado por SMS")
}

func (SmsNotification) GetSender() ISender   {
	return SmsNotificationSender{}
}

type SmsNotificationSender struct {
}

type EmailNotificationSender struct {
}

func (EmailNotification) SendNotification()  {
	fmt.Println("Mesaje enviado por email")
}

func (EmailNotification) GetSender() ISender  {
	return EmailNotificationSender{}
}
func (SmsNotificationSender) GetSenderMethod() string  {
	return "SMS"
}

func (SmsNotificationSender) GetSenderChannel() string  {
	return "Twillio"
}
func (EmailNotificationSender) GetSenderMethod() string  {
	return "EMAIL"
}

func (EmailNotificationSender) GetSenderChannel() string  {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {

	if notificationType == "SMS" {
		return SmsNotification{}, nil
	}

	if notificationType == "EMAIL" {
		return EmailNotification{}, nil
	}

	return nil, fmt.Errorf("Tipo de notificacion invalido")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getSenderMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
  smsNotification, _ := getNotificationFactory("SMS")
  emailNotification, _ := getNotificationFactory("EMAIL")

  sendNotification(smsNotification)
  sendNotification(emailNotification)

  getSenderMethod(smsNotification)
  getSenderMethod(emailNotification)
}
