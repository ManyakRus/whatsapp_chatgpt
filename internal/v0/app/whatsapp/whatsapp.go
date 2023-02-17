package whatsapp

import (
	"fmt"
	"github.com/ManyakRus/whatsapp_chatgpt/internal/v0/app/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/chatgpt_connect"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/logger"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/whatsapp_connect"
	"go.mau.fi/whatsmeow/types/events"
	"strings"
)

// log - глобальный логгер приложения
var log = logger.GetLog()

// EventHandler - получение событий из сервера whatsapp
func EventHandler(evt interface{}) {
	if evt == nil {
		log.Error("evt is null !")
	}
	switch evt.(type) {
	case *events.Message:
		mess := evt.(*events.Message)
		messW := whatsapp_connect.FillMessageWhatsapp(mess)
		ReceiveMessage(messW)
		fmt.Println("Received a message from: ", messW.NameFrom, " phone: ", messW.PhoneFrom, "text: ", messW.Text)
		//fmt.Println("Received a message: ", mess.Message, " from: ", v.Message.GetContactMessage(), "text: ", v.Info.MediaType)
	default:
		//fmt.Printf("Received: %#v \n", v)
	}
}

func ReceiveMessage(mess whatsapp_connect.MessageWhatsapp) {
	var err error

	if mess.Text == "" {
		return
	}

	if mess.IsGroup == true {
		return
	}

	if mess.MediaType != "" {
		return
	}

	if mess.IsFromMe == true {
		if strings.ToLower(mess.Text) == "on" {
			constants.AutoAnswer_Enabled = true
			Text := "AutoAnswer_Enabled = true"
			log.Info(Text)
			_, err = whatsapp_connect.SendMessage(whatsapp_connect.Settings.WHATSAPP_PHONE_FROM, Text)
			if err != nil {
				log.Error("whatsapp_connect.SendMessage() error: ", err)
			}
			return
		}
		if strings.ToLower(mess.Text) == "off" {
			constants.AutoAnswer_Enabled = false
			Text := "AutoAnswer_Enabled = false"
			log.Info(Text)
			_, err = whatsapp_connect.SendMessage(whatsapp_connect.Settings.WHATSAPP_PHONE_FROM, Text)
			if err != nil {
				log.Error("whatsapp_connect.SendMessage() error: ", err)
			}
			return
		}

		//потом убрать
		if constants.AutoAnswer_Enabled == true {
			TextRequest := mess.Text
			OtvetGPT, err := chatgpt_connect.SendMessage(TextRequest)
			if err != nil {
				log.Error("chatgpt_connect.SendMessage() error: ", err)
				return
			}
			if OtvetGPT == "" {
				log.Debug("error: response text gpt=''")
				return
			}

			TextMess := OtvetGPT
			if chatgpt_connect.Settings.CHATGPT_NAME != "" {
				TextMess = chatgpt_connect.Settings.CHATGPT_NAME + ":" + "\n" + OtvetGPT
			}
			_, err = whatsapp_connect.SendMessage(whatsapp_connect.Settings.WHATSAPP_PHONE_FROM, TextMess)
			if err != nil {
				log.Error("whatsapp_connect.SendMessage() error: ", err)
			}

		}
	}
}
