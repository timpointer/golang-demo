// Package main provides ...
package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/objx"
	"github.com/timpointer/golang-demo/web/trace"
)

// set the active Avatar implementation
var avatars Avatar = TryAvatars{
	UseFileSystemAvatar,
	UseAuthAvatar,
	UseGravatar,
}

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
	config   *configration
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//t.once.Do(func() {
	t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	//})
	data := map[string]interface{}{
		"Host": t.config.Site.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {

		log.Printf("%v", authCookie.Value)
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	log.Printf("%v", data)
	t.templ.Execute(w, data)
}

type handlerBuilder struct {
	config *configration
}

func (b *handlerBuilder) New(filename string) http.Handler {
	return &templateHandler{filename: filename, config: b.config}
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()                            // parse the flags
	config := NewConfig("config/conf.json") // get config from config file

	// init social Oauth client
	gomniauth.SetSecurityKey(config.Site.SecurityKey)
	cgithub := config.Auth.Github
	gomniauth.WithProviders(github.New(cgithub.Key, cgithub.Secret, "http://"+config.Site.Host+"/auth/callback/github"))

	// make chat room
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	builder := &handlerBuilder{config: config}

	http.Handle("/", builder.New("index.html"))
	http.Handle("/login", builder.New("login.html"))
	http.HandleFunc("/logout", logOutHandler)

	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/chat", MustAuth(builder.New("chat.html")))

	// upload user avatar
	http.Handle("/upload", MustAuth(builder.New("upload.html")))
	http.HandleFunc("/uploader", uploaderHandler)

	http.Handle("/room", r)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/avatars/", http.StripPrefix("/avatars/", http.FileServer(http.Dir("./avatars/"))))
	// get the room going
	go r.run()
	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
