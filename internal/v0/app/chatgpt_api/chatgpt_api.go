package chatgpt_api

import (
	"github.com/chatgp/chatgpt-go"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/logger"
	"net/http"
	"os"
	"time"
)

// log - глобальный логгер приложения
var log = logger.GetLog()

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

var Client *chatgpt.Client

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	SESSION_TOKEN string
	SESSION_CF    string
}

// FillSettings загружает переменные окружения в структуру из файла или из переменных окружения
func FillSettings() {
	Settings = SettingsINI{}
	Settings.SESSION_TOKEN = os.Getenv("SESSION_TOKEN")
	Settings.SESSION_CF = os.Getenv("SESSION_CF")

	if Settings.SESSION_TOKEN == "" {
		log.Panicln("Need fill SESSION_TOKEN ! in os.ENV ")
	}

	if Settings.SESSION_CF == "" {
		log.Panicln("Need fill SESSION_CF ! in os.ENV ")
	}

}

func Connect() {
	// new chatgpt client

	cookies := []*http.Cookie{
		{
			Name:  "__Secure-next-auth.session-token",
			Value: Settings.SESSION_TOKEN,
		},
		{
			Name:  "cf_clearance",
			Value: Settings.SESSION_CF,
		},
	}

	Client = chatgpt.NewClient(
		chatgpt.WithDebug(true),
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithCookies(cookies),
	)

	if Client == nil {
		log.Panic("error: Client =nil")
	}
}

func SendMessage(MessageText, conversationID, parentMessage string) *chatgpt.ChatText {
	// continue conversation with new message

	Otvet, err := Client.GetChatText(MessageText, conversationID, parentMessage)
	if err != nil {
		log.Panicf("get chat text failed: %v", err)
	}

	log.Printf("q: %s, a: %s\n", MessageText, Otvet.Content)

	return Otvet
}
