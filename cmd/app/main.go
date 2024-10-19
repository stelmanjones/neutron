package main

import (
	"os"

	"github.com/charmbracelet/log"
	neutron "github.com/stelmanjones/neutron"
	_ "github.com/stelmanjones/neutron/internal/mswebview2"
	_ "github.com/stelmanjones/neutron/internal/mswebview2/include"
	_ "github.com/stelmanjones/neutron/internal/webview"
	_ "github.com/stelmanjones/neutron/internal/webview/include"
)

var logger = log.NewWithOptions(os.Stderr, log.Options{
	Level:           log.DebugLevel,
	TimeFormat:      "15:04:05",
	ReportTimestamp: true,
})

type Result struct {
	Value uint `json:"value"`
}

func main() {
	var count uint = 0
	d := neutron.New(true)
	err := d.Bind("increment", func() Result {
		count++
		logger.Info("Called increment.", "count", count)
		return Result{Value: count}
	})
	if err != nil {
		logger.Error(err.Error())
	}

	d.Run()
}
