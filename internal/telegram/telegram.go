package telegram

import (
	"context"
	"fmt"
	"github.com/ManyakRus/starter/chatgpt_connect"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/telegram_client"
	"github.com/ManyakRus/whatsapp_chatgpt/internal/constants"
	"github.com/gotd/td/tg"
	_ "github.com/gotd/td/tgerr"
	"reflect"
	"strings"
	"time"
)

// MessageTelegram - сообщение из WhatsApp сокращённо
type MessageTelegram struct {
	Text      string
	FromID    int64
	ChatID    int64
	IsFromMe  bool
	MediaType string
	//NameTo    string
	IsGroup  bool
	ID       int
	TimeSent time.Time
}

// OnNewMessage - функция для получения новых сообщений
func OnNewMessage(ctx context.Context, entities tg.Entities, u *tg.UpdateNewMessage) error {
	var err error

	m, ok := u.Message.(*tg.Message)
	if !ok || m.Out {
		// Outgoing message, not interesting.
		return nil
	}

	mess := FillMessageTelegram(m)
	log.Printf("new telegram message:\n%v\n", mess.String())
	ReceiveMessage(mess)

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
func ReceiveMessage(mess MessageTelegram) {
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

	if mess.IsFromMe == true {
		//ON
		if strings.ToLower(mess.Text) == "on" {
			constants.Whatsapp_AutoAnswer_Enabled = true
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
			constants.Whatsapp_AutoAnswer_Enabled = false
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

	////Отправка в ChatGPT
	//TextRequest := mess.Text
	//OtvetGPT, err := chatgpt_connect.SendMessage(TextRequest, mess.NameFrom)
	//if err != nil {
	//	log.Error("chatgpt_connect.SendMessage() error: ", err)
	//	return
	//}
	//if OtvetGPT == "" {
	//	log.Debug("error: response text gpt=''")
	//	return
	//}
	//
	////Отправка в WhatsApp
	//SendMessage_WithChatGPTName(mess.FromID, OtvetGPT)

}

// FillMessageTelegram - заполнение струткру MessageTelegram из сообщения от Telegram
func FillMessageTelegram(m *tg.Message) MessageTelegram {
	Otvet := MessageTelegram{}

	ctxMain := contextmain.GetContext()
	IsGroup := false

	Otvet.Text = m.Message
	Otvet.ID = m.ID
	Otvet.MediaType = m.TypeName()
	TimeInt := m.GetDate()
	Otvet.TimeSent = time.UnixMilli(int64(TimeInt * 1000))
	var ChatID int64

	if m.PeerID != nil && CheckNilInterface(m.PeerID) == false {
		switch v := m.PeerID.(type) {
		case *tg.PeerUser:
			ChatID = v.UserID
		case *tg.PeerChat:
			{
				ChatID = v.ChatID
				IsGroup = true
			}
		case *tg.PeerChannel:
			{
				ChatID = v.ChannelID
				IsGroup = true
			}
		default:
			{
				IsGroup = true
			}
		}
	}
	Otvet.ChatID = ChatID

	MyUser, err := telegram_client.Client.Self(ctxMain)
	if err != nil {
		return Otvet
	}
	MyID := MyUser.ID
	var SenderID int64

	IsFromMe := false
	if m.FromID != nil && CheckNilInterface(m.FromID) == false {
		switch v := m.FromID.(type) {
		case *tg.PeerUser:
			{
				SenderID = v.UserID
			}
		//case *tg.PeerChat: // peerChat#36c6019a
		//case *tg.PeerChannel: // peerChannel#a2a5371e
		default:
		}
	} else {
		IsFromMe = true
	}
	Otvet.IsGroup = IsGroup //m.GroupedID != 0

	if MyID == SenderID {
		IsFromMe = true
	}
	Otvet.IsFromMe = IsFromMe
	Otvet.FromID = SenderID

	return Otvet
}

// SendMessage_WithChatGPTName - отправка сообщения в WhatsApp с добавлением имени CHATGPT_NAME
func SendMessage_WithChatGPTName(PhoneFrom, TextMess string) {

	TextMess = strings.Trim(TextMess, "\n")

	if chatgpt_connect.Settings.CHATGPT_NAME != "" {
		TextMess = chatgpt_connect.Settings.CHATGPT_NAME + "\n" + TextMess
	}

	_, err := telegram_client.SendMessage(PhoneFrom, TextMess)
	if err != nil {
		log.Error("telegram_client.SendMessage() error: ", err)
	}

}

// CheckNilInterface - проверка интерфейса на nil
func CheckNilInterface(i any) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}

	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}

func (m MessageTelegram) String() string {
	Otvet := ""

	Otvet = Otvet + fmt.Sprint("Text: ", m.Text, "\n")
	Otvet = Otvet + fmt.Sprint("MediaType: ", m.MediaType, "\n")
	Otvet = Otvet + fmt.Sprint("FromID: ", m.FromID, "\n")
	Otvet = Otvet + fmt.Sprint("IsFromMe: ", m.IsFromMe, "\n")
	Otvet = Otvet + fmt.Sprint("IsGroup: ", m.IsGroup, "\n")
	Otvet = Otvet + fmt.Sprint("ID: ", m.ID, "\n")
	Otvet = Otvet + fmt.Sprint("TimeSent: ", m.TimeSent, "\n")

	return Otvet
}
