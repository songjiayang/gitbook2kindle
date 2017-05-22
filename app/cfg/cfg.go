package cfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var Cfg *cfg

var (
	fpath string = ".gitbook2kinle.data"
)

func init() {
	usr, _ := user.Current()
	fpath = filepath.Join(usr.HomeDir, fpath)

	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Println(err)
		Cfg = NewCfg()

		if os.IsNotExist(err) {
			_, err = os.Create(fpath)
			if err != nil {
				fmt.Println("create data cache file with err:", err)
			}
		}

		return
	}

	err = json.Unmarshal(data, &Cfg)
	if err != nil {
		fmt.Println(err)
		Cfg = NewCfg()
		return
	}

	// 初始化 cached_books
	if Cfg != nil && Cfg.CachedBooks == nil {
		Cfg.ResetCache()
	}
}

type cfg struct {
	SmtpAccount  string `json:"smtp_email"`
	SmtpPassword string `json:"smtp_password"`
	SmtpHost     string `json:"smtp_host"`
	SmtpServer   string `json:"smtp_server"`

	KindleAccount string `json:"kindle_account"`
	GitBookCookie string `json:"gitbook_cookie"`

	CachedBooks []string `json:"cached_books"`
}

func NewCfg() *cfg {
	ret := &cfg{}
	ret.ResetCache()
	return ret
}

func (this *cfg) Merge(target *cfg) {
	if target.isEmpty() {
		return
	}

	if target.SmtpAccount != "" {
		this.SmtpAccount = target.SmtpAccount
	}

	if target.SmtpPassword != "" {
		this.SmtpPassword = target.SmtpPassword
	}

	if target.SmtpHost != "" {
		this.SmtpHost = target.SmtpHost
	}

	if target.SmtpServer != "" {
		this.SmtpServer = target.SmtpServer
	}

	if target.KindleAccount != "" {
		this.KindleAccount = target.KindleAccount
	}

	if target.GitBookCookie != "" {
		this.GitBookCookie = target.GitBookCookie
	}

	if strings.Contains(this.SmtpAccount, "gmail.com") {
		this.SmtpHost = "smtp.gmail.com"
		this.SmtpServer = "smtp.gmail.com:587"
	}

	this.Save()
}

func (this *cfg) ToString() string {
	data, err := json.Marshal(this)
	if err != nil {
		return ""
	}

	return string(data)
}

func (this *cfg) Save() {
	d, err := json.Marshal(this)
	if err != nil {
		fmt.Println(err)
		return
	}

	ioutil.WriteFile(fpath, d, os.ModePerm)
}

func (this *cfg) IsValid() bool {
	return this.SmtpAccount != "" &&
		this.SmtpPassword != "" &&
		this.SmtpHost != "" &&
		this.SmtpServer != "" &&
		this.KindleAccount != "" &&
		this.GitBookCookie != ""
}

func (this *cfg) ResetCache() {
	this.CachedBooks = make([]string, 0)
}

func (this *cfg) CacheBook(book string) {
	this.CachedBooks = append(this.CachedBooks, book)
}

func (this *cfg) CachedBook(book string) bool {
	for _, cachedBook := range this.CachedBooks {
		if cachedBook == book {
			return true
		}
	}
	return false
}

func (this *cfg) isEmpty() bool {
	return this.SmtpAccount == "" &&
		this.SmtpPassword == "" &&
		this.SmtpHost == "" &&
		this.SmtpServer == "" &&
		this.KindleAccount == "" &&
		this.GitBookCookie == ""
}
