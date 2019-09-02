package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templ1は1つのテンプレートを表す
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTPはHTTPリクエストを処理する
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()

	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルームを開始する
	go r.run()

	// Webサーバを起動する
	log.Println("Webサーバを開始します。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
