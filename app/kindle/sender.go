package kindle

import (
	"fmt"
	"io"
	"net/smtp"

	"github.com/gitbook2kindle/app/cfg"
	"github.com/jordan-wright/email"
)

var smptAuth smtp.Auth

func InitSmtp() {
	smptAuth = smtp.PlainAuth("", cfg.Cfg.SmtpAccount, cfg.Cfg.SmtpPassword, cfg.Cfg.SmtpHost)
}

func Send(books map[string]io.ReadCloser) {
	fmt.Println(fmt.Sprintf("start send %d books to kindle: %s", len(books), cfg.Cfg.KindleAccount))
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

	for bookName, r := range books {
		e.Attach(r, bookName, "")
	}

	err := e.Send(cfg.Cfg.SmtpServer, smptAuth)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("end send %d books to kindle: %s", len(books), cfg.Cfg.KindleAccount))
}
