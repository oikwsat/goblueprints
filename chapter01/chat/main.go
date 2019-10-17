package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHander struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP は HTTP リクエストを処理します
func (t *templateHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	// ルート
	http.Handle("/", &templateHander{filename: "chat.html"})

	// Webサーバを開始します
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
