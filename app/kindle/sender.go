package kindle

import (
	"fmt"
	"io"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

var smptAuth smtp.Auth

func init() {
	for _, k := range []string{"SMTP_SERVER", "SMTP_ACCOUNT", "SMTP_PASSWORD", "SMTP_HOST", "KINDLE_ACCOUNT"} {
		if os.Getenv(k) == "" {
			panic("No " + k + " Setting")
		}
	}

	smptAuth = smtp.PlainAuth("", os.Getenv("SMTP_ACCOUNT"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
}

func SendBook(bookName string, r io.ReadCloser) {
	fmt.Println("Start send book: ", bookName)
	defer fmt.Println(fmt.Println("End send book: ", bookName))

	if r == nil {
		return
	}
	defer r.Close()

	e := email.NewEmail()
	e.From = "gitbook2kindle <" + os.Getenv("SMTP_ACCOUNT") + ">"
	e.To = []string{os.Getenv("KINDLE_ACCOUNT")}
	e.Subject = bookName
	e.Attach(r, bookName, "")

	err := e.Send(os.Getenv("SMTP_SERVER"), smptAuth)

	if err != nil {
		fmt.Println(err)
	}
}
