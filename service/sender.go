package service

import "github.com/Yusuf1901-cloud/message_sender/models"



type SenderI interface {
	Send(msg *models.Message) error
}
