package chatgpt_connect

import (
	"errors"
	//"github.com/jackc/pgconn"
	"os"
	"sync"
	//"time"

	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/contextmain"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/logger"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/stopapp"

	gogpt "github.com/sashabaranov/go-gpt3"
)

// Conn - соединение к CHAT GPT
// var Conn *gogpt.CompletionStream
var Conn *gogpt.Client

// log - глобальный логгер
var log = logger.GetLog()

// mutexReconnect - защита от многопоточности Reconnect()
var mutexReconnect = &sync.Mutex{}

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	CHATGPT_API_KEY string
	CHATGPT_NAME    string
}

// Connect_err - подключается к базе данных
func Connect() {

	err := Connect_err()
	if err != nil {
		log.Panicln("CHATGPT Connect_err() api_key: ", Settings.CHATGPT_API_KEY, " Error: ", err)
	} else {
		log.Info("CHATGPT connected. api_key: ", Settings.CHATGPT_API_KEY)
	}

}

// Connect_err - подключается к базе данных
func Connect_err() error {
	var err error

	if Settings.CHATGPT_API_KEY == "" {
		FillSettings()
	}

	Conn = gogpt.NewClient(Settings.CHATGPT_API_KEY)

	//req := gogpt.CompletionRequest{
	//	Model:     gogpt.GPT3Ada,
	//	MaxTokens: 5,
	//	Prompt:    Settings.CHATGPT_NAME,
	//	Stream:    true,
	//}

	//ctx := contextmain.GetContext()
	//stream, err := Conn.CreateCompletionStream(ctx, req)
	//if err != nil {
	//	return err
	//}

	return err
}

// CloseConnection - закрытие соединения с базой данных
func CloseConnection() error {
	if Conn == nil {
		return nil
	}

	err := CloseConnection_err()
	if err != nil {
		log.Error("Postgres sqlx CloseConnection() error: ", err)
	} else {
		log.Debug("Postgres sqlx connection closed")
	}

	return err
}

// CloseConnection - закрытие соединения с базой данных
func CloseConnection_err() error {
	var err error
	if Conn == nil {
		return nil
	}

	//ctx := contextmain.GetContext()
	//ctx := context.Background()
	//Conn.Close()

	return err
}

// WaitStop - ожидает отмену глобального контекста или сигнала завершения приложения
func WaitStop() {

	select {
	//case <-stopapp.SignalInterrupt:
	//	log.Warn("Interrupt clean shutdown.")
	//	contextmain.CancelContext()
	case <-contextmain.GetContext().Done():
		log.Warn("Context app is canceled.")
	}

	//
	stopapp.WaitTotalMessagesSendingNow("CHATGPT")

	//
	err := CloseConnection()
	if err != nil {
		log.Error("CloseConnection() error: ", err)
	}
	stopapp.GetWaitGroup_Main().Done()
}

// Start - делает соединение с БД, отключение и др.
func Start() {
	Connect()

	stopapp.GetWaitGroup_Main().Add(1)
	go WaitStop()

}

// FillSettings загружает переменные окружения в структуру из файла или из переменных окружения
func FillSettings() {
	Settings = SettingsINI{}
	Settings.CHATGPT_API_KEY = os.Getenv("CHATGPT_API_KEY")
	Settings.CHATGPT_NAME = os.Getenv("CHATGPT_NAME")
	if Settings.CHATGPT_API_KEY == "" {
		log.Panicln("Need fill CHATGPT_API_KEY ! in os .env ")
	}
	if Settings.CHATGPT_NAME == "" {
		log.Warnln("Need fill CHATGPT_NAME ! in os .env ")
	}

	//
}

func SendMessage(Text string) (string, error) {
	var Otvet = ""
	var err error

	if Conn == nil {
		Connect()
	}

	ctx := contextmain.GetContext()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Ada, //надо gogpt.GPT3TextDavinci003
		MaxTokens: 5,
		Prompt:    Text,
	}
	resp, err := Conn.CreateCompletion(ctx, req)
	if err != nil {
		log.Debug("ChatGPT CreateCompletion() error: ", err)
		return Otvet, err
	}

	if len(resp.Choices) > 0 {
		Otvet = resp.Choices[0].Text
	} else {
		err = errors.New("error: no response")
	}
	//fmt.Println("Otvet: ", resp.Choices[0].Text)

	//req := gogpt.CompletionRequest{
	//	Model:     gogpt.GPT3Ada,
	//	MaxTokens: 5,
	//	Prompt:    Text,
	//	Stream:    true,
	//}
	//stream, err := Conn.CreateCompletionStream(ctx, req)
	//if err != nil {
	//	return Otvet, err
	//}
	//defer stream.Close()
	//
	//for {
	//	response, err := stream.Recv()
	//	Otvet = response
	//	if errors.Is(err, io.EOF) {
	//		fmt.Println("Stream finished")
	//		err = nil
	//		return Otvet, err
	//	}
	//
	//	if err != nil {
	//		fmt.Printf("Stream error: %v\n", err)
	//		return Otvet, err
	//	}
	//
	//	fmt.Printf("Stream response: %v\n", response)
	//}

	return Otvet, err
}
