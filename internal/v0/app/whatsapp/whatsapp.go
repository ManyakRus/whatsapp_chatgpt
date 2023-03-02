package whatsapp

import (
	"fmt"
	"github.com/ManyakRus/whatsapp_chatgpt/internal/v0/app/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/chatgpt_connect"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/logger"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/whatsapp_connect"
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
		//ON
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

		//OFF
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
	}

	//Проверка включено
	if constants.AutoAnswer_Enabled == false {
		return
	}

	//сообщение от автоответчика (старое)
	len_name := len(chatgpt_connect.Settings.CHATGPT_NAME)
	if chatgpt_connect.Settings.CHATGPT_NAME != "" && len(mess.Text) >= len_name && mess.Text[0:len_name] == chatgpt_connect.Settings.CHATGPT_NAME {
		return
	}

	//Отправка в ChatGPT
	TextRequest := mess.Text
	OtvetGPT, err := chatgpt_connect.SendMessage(TextRequest, mess.NameFrom)
	if err != nil {
		log.Error("chatgpt_connect.SendMessage() error: ", err)
		return
	}
	if OtvetGPT == "" {
		log.Debug("error: response text gpt=''")
		return
	}

	//Отправка в WhatsApp
	SendMessage_WithChatGPTName(mess.PhoneFrom, OtvetGPT)

	//}
}

func SendMessage_WithChatGPTName(PhoneFrom, TextMess string) {

	//if len(TextMess) > 0 && TextMess[0:1] == "\n" {
	//	TextMess = TextMess[1:]
	//}

	TextMess = strings.Trim(TextMess, "\n")

	if chatgpt_connect.Settings.CHATGPT_NAME != "" {
		TextMess = chatgpt_connect.Settings.CHATGPT_NAME + "\n" + TextMess
	}

	_, err := whatsapp_connect.SendMessage(PhoneFrom, TextMess)
	if err != nil {
		log.Error("whatsapp_connect.SendMessage() error: ", err)
	}

}
