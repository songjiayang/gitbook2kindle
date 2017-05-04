package gitbook

import (
	"net/http"
	"time"

	"github.com/gitbook2kindle/app/cfg"
)

var cookie string

type client struct {
	*http.Client
}

func NewClient() *client {
	cookie = cfg.Cfg.GitBookCookie
	if cookie == "" {
		panic("No GITBOOK_COOKIE Setting")
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	return &client{
		&http.Client{Transport: tr},
	}
}

func (this *client) newRequest(method, url string) *http.Request {
	req, _ := http.NewRequest(method, url, nil)

	c := &http.Cookie{
		Name:   "gitbook:sess",
		Value:  cookie,
		Path:   "/",
		Domain: "www.gitbook.com",
	}

	req.AddCookie(c)
	return req
}
