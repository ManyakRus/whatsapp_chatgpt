package telegram

import (
	"context"
	"github.com/ManyakRus/starter/chatgpt_connect"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/telegram_client"
	"github.com/ManyakRus/whatsapp_chatgpt/internal/constants"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/tg"
	_ "github.com/gotd/td/tgerr"
	"strings"
	"time"
)

// Contacts - список контактов в Telegram
var Contacts *tg.ContactsContacts

// OnNewMessage - функция для получения новых сообщений
func OnNewMessage(ctx context.Context, entities tg.Entities, u *tg.UpdateNewMessage) error {
	var err error

	m, ok := u.Message.(*tg.Message)
	if !ok || m.Out {
		// Outgoing message, not interesting.
		return nil
	}

	mess := telegram_client.FillMessageTelegramFromMessage(m)
	log.Debugf("new telegram message:\n%v\n", mess.String())

	ReceiveMessage(mess, entities, u)

	//// тестовый пример эхо
	//// Helper for sending messages.
	//api := telegram_client.Client.API()
	//sender := message.NewSender(api)
	//
	//// Sending reply.
	//_, err = sender.Reply(entities, u).Text(ctx, m.Message)

	return err
}

// ReceiveMessage - обработка полученного сообщения от WhatsApp
func ReceiveMessage(mess telegram_client.MessageTelegram, entities tg.Entities, u *tg.UpdateNewMessage) {
	var err error

	if mess.Text == "" {
		return
	}

	if mess.IsGroup == true {
		return
	}

	if mess.MediaType != "message" {
		return
	}

	//сообщение от меня и мне
	if mess.IsFromMe == true && mess.FromID == telegram_client.UserSelf.ID {
		//ON
		if strings.ToLower(mess.Text) == "on" {
			constants.Telegram_AutoAnswer_Enabled = true
			Text := constants.TextEnabled
			log.Info(Text)
			_, err = telegram_client.SendMessage(telegram_client.Settings.TELEGRAM_PHONE_FROM, Text)
			if err != nil {
				log.Error("telegram_client.SendMessage() error: ", err)
			}
			return
		}

		//OFF
		if strings.ToLower(mess.Text) == "off" {
			constants.Telegram_AutoAnswer_Enabled = false
			Text := constants.TextDisabled
			log.Info(Text)
			_, err = telegram_client.SendMessage(telegram_client.Settings.TELEGRAM_PHONE_FROM, Text)
			if err != nil {
				log.Error("telegram_client.SendMessage() error: ", err)
			}
			return
		}
	}

	//Проверка включено
	if constants.Telegram_AutoAnswer_Enabled == false {
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
	//if mess.IsFromMe != mess.FromID {
	//	return
	//}

	// прошёл 1 час // потом вернуть
	if time.Now().Sub(mess.TimeSent) > time.Minute*60 {
		return
	}

	OtvetGPT := "test"
	////Отправка в ChatGPT
	//TextRequest := mess.Text
	//Name := micro.StringFromInt64(mess.FromID)
	//OtvetGPT, err = chatgpt_connect.SendMessage(TextRequest, Name)
	//if err != nil {
	//	log.Error("chatgpt_connect.SendMessage() error: ", err)
	//	return
	//}
	//if OtvetGPT == "" {
	//	log.Debug("error: response text gpt=''")
	//	return
	//}

	//Отправка в Telegram
	//PhoneTo := micro.StringFromInt64(mess.FromID)
	SendMessage_WithChatGPTName(entities, u, OtvetGPT)

}

// SendMessage_WithChatGPTName - отправка сообщения в Telegram с добавлением имени CHATGPT_NAME
func SendMessage_WithChatGPTName(entities tg.Entities, u *tg.UpdateNewMessage, TextMess string) {

	TextMess = strings.Trim(TextMess, "\n")

	if chatgpt_connect.Settings.CHATGPT_NAME != "" {
		TextMess = chatgpt_connect.Settings.CHATGPT_NAME + "\n" + TextMess
	}

	ctxMain := contextmain.GetContext()
	ctx, cancel := context.WithTimeout(ctxMain, time.Second*60)
	defer cancel()

	api := telegram_client.Client.API()
	sender := message.NewSender(api)
	rb := sender.Answer(entities, u)
	_, err := rb.Text(ctx, TextMess)
	if err != nil {
		log.Error("Answer() error: ", err)
	}

}

// ContactsGetContacts - обновляет список контактов
func ContactsGetContacts() {
	ctxMain := contextmain.GetContext()
	ctx, cancel_func := context.WithTimeout(ctxMain, time.Second*60)
	defer cancel_func()
	IContacts, err := telegram_client.Client.API().ContactsGetContacts(ctx, 0)
	if err != nil {
		log.Error("ContactsGetContacts() error: ", err)
		return
	}
	ok := false
	Contacts, ok = IContacts.(*tg.ContactsContacts)
	if ok == false {
		log.Error("IContacts() error: ", err)
		return
	}
	log.Debug("Contacts len: ", len(Contacts.Contacts))
}
