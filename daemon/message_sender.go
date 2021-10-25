package daemon

import (
	"context"
	"fmt"
	"time"

	"github.com/Yusuf1901-cloud/message_sender/models"
	"github.com/Yusuf1901-cloud/message_sender/repository"
)

var (
	LowPriorityQueue    chan models.Message
	MediumPriorityQueue chan models.Message
	HighPriorityQueue   chan models.Message
)

func init() {
	LowPriorityQueue = make(chan models.Message, 100)
	MediumPriorityQueue = make(chan models.Message, 100)
	HighPriorityQueue = make(chan models.Message, 100)
}

func SendMessages(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second * 5):
			select {
			case msg := <-HighPriorityQueue:
				err := repository.BotSenderRepo.Send(&msg)
				if err != nil {
					fmt.Println(err)
				}
				continue
			default:
			}

			select {
			case msg := <-MediumPriorityQueue:
				err := repository.BotSenderRepo.Send(&msg)
				if err != nil {
					fmt.Println(err)
				}
				continue
			default:
			}

			select {
			case msg := <-LowPriorityQueue:
				err := repository.BotSenderRepo.Send(&msg)
				if err != nil {
					fmt.Println(err)
				}
			default:
			}
		}

	}
}
