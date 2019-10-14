package main

import (
	"flag"
	"github.com/RyotaNakajimakun/Training/Golang/web_application/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

const (
	//@TODO 外部ファイルに定義　
	SecurityKey       = ""
	FacebookClientID  = ""
	FacebookSecretKey = ""
	GitHubClientID    = ""
	GitHubSecretKey   = ""
	GoogleClientID    = ""
	GoogleSecretKey   = ""
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
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.temp1.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションアドレス")
	flag.Parse()
	gomniauth.SetSecurityKey(SecurityKey)
	gomniauth.WithProviders(
		facebook.New(FacebookClientID, FacebookSecretKey, "http://127.0.0.1:8080/auth/callback/facebook"),
		github.New(GitHubClientID, GitHubSecretKey, "http://127.0.0.1:8080/auth/callback/github"),
		google.New(GoogleClientID, GoogleSecretKey, "http://127.0.0.1:8080/auth/callback/google"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	// ルート
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	go r.run()

	// WEBサーバー開始
	log.Println("Webサーバーを起動します。　ポート" + *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
