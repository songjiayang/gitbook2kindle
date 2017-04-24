package gitbook

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type book struct {
	ID string `json:"id"`
}

type bookList struct {
	Items []*book `json:"starred"`
}

type starBookAPI struct {
	StarBooks *bookList `json:"props"`
}

func (this *client) ListStarBooks() (books []*book) {
	req := this.newRequest("GET", "https://www.gitbook.com/@songjiayang/starred")
	req.Header.Set("x-pjax", "true")

	resp, err := this.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data starBookAPI
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	books = data.StarBooks.Items
	return
}

func (this *client) DownloadBook(id string) (r io.ReadCloser) {
	fmt.Println("Start download book: ", id)
	defer fmt.Println("End download book: ", id)

	req := this.newRequest("GET", downloadUrl(id))

	resp, err := this.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	return resp.Body
}

func downloadUrl(id string) string {
	return "https://www.gitbook.com/download/mobi/book/" + id
}
