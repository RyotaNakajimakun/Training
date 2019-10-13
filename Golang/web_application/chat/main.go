package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"github.com/RyotaNakajimakun/Training/Golang/web_application/trace"
)

// tem1は１つのテンプレートを表わす
type templateHandler struct {
	once     sync.Once
	filename string
	temp1    *template.Template
}

// ServerHTTPはHTTPリクエストを処理する
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.temp1 = template.Must(template.ParseFiles(filepath.Join("chat/templates", t.filename)))
	})
	t.temp1.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションアドレス")
	flag.Parse()
	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	// ルート
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.run()

	// WEBサーバー開始
	log.Println("Webサーバーを起動します。　ポート" + *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
