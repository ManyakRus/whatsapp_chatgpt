package whatsapp

import (
	"github.com/ManyakRus/whatsapp_chatgpt/internal/v0/app/programdir"
	"testing"
	"time"

	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/config"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/micro"
)

func TestCreateClient(t *testing.T) {
	ProgramDir := programdir.ProgramDir()
	config.LoadEnv(ProgramDir)
	FillSettings()

	err := CreateClient()
	if err != nil {
		t.Error("TestCreateClient() error: ", err)
	}
	StopWhatsApp()
}

func TestSendMessage(t *testing.T) {
	ProgramDir := programdir.ProgramDir()
	config.LoadEnv(ProgramDir)
	FillSettings()

	err := CreateClient()
	if err != nil {
		t.Error("whatsapp_test.TestCreateClient() error: ", err)
	}
	//micro.Pause(500)

	phone_send_to := Settings.WHATSAPP_PHONE_SEND_TEST
	text := "Test www.ya.ru " + time.Now().String()

	id, ok, err := SendMessage(phone_send_to, text)
	if id == "" {
		t.Error("whatsapp_test.TestSendMessage() id=''")
	}
	if ok == false {
		t.Error("whatsapp_test.TestSendMessage() is_sent=false")
	}
	if err != nil {
		t.Error("whatsapp_test.TestSendMessage() error: ", err)
	}

	micro.Pause(100)
	//StopWhatsApp()
}

func Test_eventHandler(t *testing.T) {
	eventHandler("")
}

func TestParseJID(t *testing.T) {
	_, ok := ParseJID("+79055391111")
	if ok != true {
		t.Error("whatsapp_test.TestParseJID() error")
	}
}
