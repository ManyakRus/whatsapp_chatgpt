package main

import (
	"github.com/ManyakRus/whatsapp_chatgpt/internal/v0/app/programdir"
	"github.com/ManyakRus/whatsapp_chatgpt/internal/v0/app/whatsapp"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/config"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/contextmain"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/logger"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/stopapp"
)

// log - глобальный логгер
var log = logger.GetLog()

// main - старт приложения
func main() {
	StartApp()
}

// StartApp - выполнение всех операций для старта приложения
func StartApp() {
	ProgramDir := programdir.ProgramDir()
	config.LoadEnv(ProgramDir)
	contextmain.GetContext()

	stopapp.StartWaitStop()

	contextmain.GetContext()

	whatsapp.FillSettings()

	go func() {
		err := whatsapp.CreateClient()
		if err != nil {
			log.Panicln("whatsapp.CreateClient() error: ", err)
		}
	}()

	stopapp.GetWaitGroup_Main().Wait()

}
