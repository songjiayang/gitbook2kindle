package kindle

import (
	"fmt"
	"io"
	"net/smtp"
	"strings"

	"github.com/gitbook2kindle/app/cfg"
	"github.com/jordan-wright/email"
)

var smptAuth smtp.Auth

func InitSmtp() {
	smptAuth = smtp.PlainAuth("", cfg.Cfg.SmtpAccount, cfg.Cfg.SmtpPassword, cfg.Cfg.SmtpHost)
}

func Send(books map[string]io.ReadCloser) {
	defer func() {
		for _, r := range books {
			if r != nil {
				r.Close()
			}
		}
	}()

	e := email.NewEmail()
	e.From = "gitbook2kindle <" + cfg.Cfg.SmtpAccount + ">"
	e.To = []string{cfg.Cfg.KindleAccount}
	e.Subject = "sync gitbook to kindle"

	bookNames := make([]string, 0)
	for bookName, r := range books {
		bookNames = append(bookNames, bookName)
		e.Attach(r, bookName+".mobi", "")
	}

	fmt.Println(fmt.Sprintf("--> Syncing: %s", strings.Join(bookNames, ", ")))
	err := e.Send(cfg.Cfg.SmtpServer, smptAuth)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("--> Synced: %s", strings.Join(bookNames, ", ")))
}
