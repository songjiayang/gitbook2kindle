package main

import (
	"flag"
	"fmt"

	"github.com/gitbook2kindle/app/cfg"
	"github.com/gitbook2kindle/app/gitbook"
	"github.com/gitbook2kindle/app/kindle"
)

var (
	runApp, printConfig bool
	batchSize           int
)

func main() {
	// smtp config
	newCfg := cfg.NewCfg()
	flag.StringVar(&newCfg.SmtpAccount, "config.smtp.email", "", "config send smtp email")
	flag.StringVar(&newCfg.SmtpPassword, "config.smtp.password", "", "config send smtp email")
	flag.StringVar(&newCfg.SmtpHost, "config.smtp.host", "", "config send smtp server host")
	flag.StringVar(&newCfg.SmtpServer, "config.smtp.server", "", "config send smtp server address")

	// gitbook and kindle account config
	flag.StringVar(&newCfg.KindleAccount, "config.kindle", "", "config kindle account")
	flag.StringVar(&newCfg.GitBookCookie, "config.gitbook", "", "config gitbook cookie")

	// app running flag
	flag.BoolVar(&runApp, "run", false, "run app")
	flag.BoolVar(&printConfig, "config", false, "print current config")

	flag.IntVar(&batchSize, "batchSize", 5, "batch sync size with one email")

	flag.Parse()

	cfg.Cfg.Merge(newCfg)

	if printConfig {
		fmt.Println(cfg.Cfg.ToString())
	}

	if runApp {
		run()
	}
}

func run() {
	if !cfg.Cfg.IsValid() {
		panic("Have params not set yet, please run -config and -h to see details")
	}

	gClient := gitbook.NewClient()
	books := gClient.ListStarBooks()

	if len(books) == 0 {
		return
	}

	kindle.InitSmtp()

	sendBooks := make([]string, 0)
	for _, book := range books {
		sendBooks = append(sendBooks, book.ID)
		if len(sendBooks) == batchSize {
			kindle.Send(gClient.DownloadBooks(sendBooks))
			sendBooks = make([]string, 0)
		}
	}

	if len(sendBooks) > 0 {
		kindle.Send(gClient.DownloadBooks(sendBooks))
	}
}
