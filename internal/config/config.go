package config

import (
	"github.com/ManyakRus/starter/micro"
	"os"
)

type SettingsINI struct {
	WHATSAPP_ENABLED bool
	TELEGRAM_ENABLED bool
}

var Settings = SettingsINI{}

// FillSettings - заполняем настройки из переменных окружения
func FillSettings() {
	sWHATSAPP_ENABLED := os.Getenv("WHATSAPP_ENABLED")
	sTELEGRAM_ENABLED := os.Getenv("TELEGRAM_ENABLED")

	WHATSAPP_ENABLED := micro.BoolFromString(sWHATSAPP_ENABLED)
	TELEGRAM_ENABLED := micro.BoolFromString(sTELEGRAM_ENABLED)

	Settings.WHATSAPP_ENABLED = WHATSAPP_ENABLED
	Settings.TELEGRAM_ENABLED = TELEGRAM_ENABLED

}
