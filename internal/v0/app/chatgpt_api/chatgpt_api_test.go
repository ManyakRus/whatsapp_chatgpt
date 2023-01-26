package chatgpt_api

import (
	"github.com/ManyakRus/whatsapp_chatgpt/internal/v0/app/programdir"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/config"
	"testing"
)

func TestConnect(t *testing.T) {
	ProgramDir := programdir.ProgramDir()
	config.LoadEnv(ProgramDir)
	FillSettings()

	Connect()

}

func TestSendMessage(t *testing.T) {
	ProgramDir := programdir.ProgramDir()
	config.LoadEnv(ProgramDir)
	FillSettings()

	Connect()

	Text := "Test message. Ответить на русском языке"
	Otvet := SendMessage(Text, "", "")
	log.Debugf("chatgpt_api_test.TestSendMessage() Otvet: %#v", Otvet)
}
