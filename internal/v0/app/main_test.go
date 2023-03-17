package main

import (
	"testing"

	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/stopapp"
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
