package main

import (
	"github.com/ManyakRus/starter/chatgpt_connect"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/logger"
	"github.com/ManyakRus/starter/stopapp"
	"github.com/ManyakRus/starter/whatsapp_connect"
	"github.com/ManyakRus/whatsapp_chatgpt/internal/v0/app/whatsapp"
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
	config_main.LoadEnv()
	contextmain.GetContext()

	stopapp.StartWaitStop()

	contextmain.GetContext()

	chatgpt_connect.Start()

	whatsapp_connect.Connect(whatsapp.EventHandler)
	//whatsapp_connect.Start(whatsapp.EventHandler)

	stopapp.GetWaitGroup_Main().Wait()

}
