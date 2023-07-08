package main

import (
	"fmt"
	"io/ioutil"
	// "log"
	// "net/http"
)

type Page struct {
	Title string
	Body []byte
}

// 0600のパーミッションはウェブサーバーを立ち上げた人が読み書きできる

func (p *Page) save() error {
filename := p.Title + ".txt"
return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil,err
	}
	return &Page{Title: title, Body: body},nil
}

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// title:= r.URL.Path

// }

func main() {
	p1 := &Page{Title: "test", Body: []byte("this is a sample page")}
	p1.save()

	p2,_ :=loadPage(p1.Title)
	fmt.Println(string(p2.Body))
	// http.HandleFunc("/view",viewHandler)
	// log.Fatal(http.ListenAndServe(":8080",nil))
}
