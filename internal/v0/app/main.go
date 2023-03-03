package main

import (
	"github.com/ManyakRus/whatsapp_chatgpt/internal/v0/app/whatsapp"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/chatgpt_connect"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/config"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/contextmain"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/logger"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/stopapp"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/whatsapp_connect"
)

// log - глобальный логгер
var log = logger.GetLog()

// main - старт приложения
func main() {
	StartApp()
}

// StartApp - выполнение всех операций для старта приложения
func StartApp() {
	//ProgramDir := programdir.ProgramDir()
	config.LoadEnv()
	contextmain.GetContext()

	stopapp.StartWaitStop()

	contextmain.GetContext()

	chatgpt_connect.Start()

	whatsapp_connect.Connect(whatsapp.EventHandler)
	//whatsapp_connect.Start(whatsapp.EventHandler)

	stopapp.GetWaitGroup_Main().Wait()

}
