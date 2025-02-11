package main

import (
	"github.com/ManyakRus/starter/chatgpt_connect"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/logger"
	"github.com/ManyakRus/starter/stopapp"
	"github.com/ManyakRus/starter/telegram_client"
	"github.com/ManyakRus/starter/whatsapp_connect"
	"github.com/ManyakRus/whatsapp_chatgpt/internal/config"
	"github.com/ManyakRus/whatsapp_chatgpt/internal/telegram"
	"github.com/ManyakRus/whatsapp_chatgpt/internal/whatsapp"
)

// log - глобальный логгер
var log = logger.GetLog()

// main - старт приложения
func main() {
	StartApp()
}

// StartApp - выполнение всех операций для старта приложения
func StartApp() {
	config_main.LoadENV_or_SettingsTXT()

	config.FillSettings()

	stopapp.StartWaitStop()

	chatgpt_connect.Start()

	if config.Settings.WHATSAPP_ENABLED == true {
		whatsapp_connect.Connect(whatsapp.EventHandler)
	}

	if config.Settings.TELEGRAM_ENABLED == true {
		telegram_client.Connect(telegram.OnNewMessage)
		//telegram.GetContactsAll()
	}

	stopapp.GetWaitGroup_Main().Wait()

}
