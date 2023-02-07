package whatsapp

import (
	"fmt"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/logger"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/whatsapp_connect"
	"go.mau.fi/whatsmeow/types/events"
)

// log - глобальный логгер приложения
var log = logger.GetLog()

// EventHandler - получение событий из сервера whatsapp
func EventHandler(evt interface{}) {
	if evt == nil {
		log.Error("evt is null !")
	}
	switch v := evt.(type) {
	case *events.Message:
		mess := evt.(*events.Message)
		messW := whatsapp_connect.FillMessageWhatsapp(mess)
		fmt.Println("Received a message from: ", messW.NameFrom, " phone: ", messW.PhoneFrom, "text: ", messW.Text)
		//fmt.Println("Received a message: ", mess.Message, " from: ", v.Message.GetContactMessage(), "text: ", v.Info.MediaType)
	default:
		fmt.Printf("Received: %#v \n", v)
	}
}
