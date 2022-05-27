package handlers

import (
	"fmt"
	"github.com/tulir/whatsmeow/types/events"
)

type WaHandler struct {
}

func (u *WaHandler)eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println(
			"Received a message!", v.Message.GetConversation())
	}
}