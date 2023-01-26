package main

import (
	"testing"

	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/contextmain"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/micro"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/common/pkg/v0/stopapp"
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
