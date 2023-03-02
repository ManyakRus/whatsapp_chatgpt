package main

import (
	"testing"

	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/contextmain"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/micro"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/stopapp"
)

func TestStartApp(t *testing.T) {
	t.SkipNow()

	contextmain.GetContext()
	//stopapp.StartWaitStop()
	go main()

	micro.Pause(500)

	contextmain.CancelContext()

	stopapp.GetWaitGroup_Main().Wait()

	contextmain.GetNewContext()

}
