package handler

import (
	"github.com/Yusuf1901-cloud/message_sender/daemon"
	"github.com/Yusuf1901-cloud/message_sender/models"
	"github.com/gin-gonic/gin"
)

//@Router /send [post]
//@Summary Send Message with TGBot
//@Description API for sending message
//@Tags tgbot
//@Accept json
//@Produce json
//@Param request body models.Message true "message"
//@Success 200 {object} models.Response
//@Failure 400 {object} models.Response
//@Failure 500 {object} models.Response
func SendMessage(c *gin.Context) {
	var msg models.Message

	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(400, "could not bind json")
		return
	}

	switch msg.Priority {
	case "low":
		select {
		case daemon.LowPriorityQueue <- msg:
		default:
			c.JSON(400, "low priority queue is full")
			return
		}
	case "medium":
		select {
		case daemon.MediumPriorityQueue <- msg:
		default:
			c.JSON(400, "medium priority queue is full")
			return
		}
	case "high":
		select {
		case daemon.HighPriorityQueue <- msg:
		default:
			c.JSON(400, "high priority queue is full")
			return
		}
	default:
		c.JSON(400, "priority should be one of low, medium, high")
		return
	}

	c.JSON(200, "ok")
}
