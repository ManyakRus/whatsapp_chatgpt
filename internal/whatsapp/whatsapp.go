package whatsapp

import (
	"github.com/ManyakRus/starter/chatgpt_connect"
	"github.com/ManyakRus/starter/logger"
	"github.com/ManyakRus/starter/whatsapp_connect"
	"github.com/ManyakRus/whatsapp_chatgpt/internal/constants"
	"go.mau.fi/whatsmeow/types/events"
	"strings"
	"time"
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
		//fmt.Println("Received a message from: ", messW.NameFrom, " phone: ", messW.FromID, "text: ", messW.Text)
		log.Debug("new whatsapp message:\n", messW.String(), "\n")
		ReceiveMessage(messW)
		//fmt.Println("Received a message: ", mess.Message, " from: ", v.Message.GetContactMessage(), "text: ", v.Info.MediaType)
	default:
		//fmt.Printf("Received: %#v \n", v)
	}
}

// ReceiveMessage - обработка полученного сообщения от WhatsApp
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
			constants.Whatsapp_AutoAnswer_Enabled = true
			Text := constants.TextEnabled
			log.Info(Text)
			_, err = whatsapp_connect.SendMessage(whatsapp_connect.Settings.WHATSAPP_PHONE_FROM, Text)
			if err != nil {
				log.Error("whatsapp_connect.SendMessage() error: ", err)
			}
			return
		}

		//OFF
		if strings.ToLower(mess.Text) == "off" {
			constants.Whatsapp_AutoAnswer_Enabled = false
			Text := constants.TextDisabled
			log.Info(Text)
			_, err = whatsapp_connect.SendMessage(whatsapp_connect.Settings.WHATSAPP_PHONE_FROM, Text)
			if err != nil {
				log.Error("whatsapp_connect.SendMessage() error: ", err)
			}
			return
		}
	}

	//Проверка включено
	if constants.Whatsapp_AutoAnswer_Enabled == false {
		return
	}

	//игнорируем своё сообщение
	if (mess.Text == constants.TextEnabled) || (mess.Text == constants.TextDisabled) {
		return
	}

	//сообщение от автоответчика (старое)
	name := chatgpt_connect.Settings.CHATGPT_NAME + "\n"
	len_name := len(name)
	if chatgpt_connect.Settings.CHATGPT_NAME != "" && len(mess.Text) >= len_name && mess.Text[0:len_name] == name {
		return
	}

	// сообщение не мне
	if mess.PhoneChat != mess.PhoneFrom {
		return
	}

	// прошёл 1 час // потом вернуть
	if time.Now().Sub(mess.TimeSent) > time.Minute*60 {
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

// SendMessage_WithChatGPTName - отправка сообщения в WhatsApp с добавлением имени CHATGPT_NAME
func SendMessage_WithChatGPTName(PhoneFrom, TextMess string) {

	TextMess = strings.Trim(TextMess, "\n")

	if chatgpt_connect.Settings.CHATGPT_NAME != "" {
		TextMess = chatgpt_connect.Settings.CHATGPT_NAME + "\n" + TextMess
	}

	_, err := whatsapp_connect.SendMessage(PhoneFrom, TextMess)
	if err != nil {
		log.Error("whatsapp_connect.SendMessage() error: ", err)
	}

}
