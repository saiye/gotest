package engine

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type HttpServer struct {
}

func (server HttpServer) Start() {
	/*	//首页
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
			if err != nil {
				return
			}
		})

		//login
		http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
			user:=r.URL.Query().Get("user")
			password:=r.URL.Query().Get("password")
			_, err := fmt.Fprintf(w, "user: %s--password:%s", user,password)
			if err != nil {
				return
			}
		})

		//静态资源目录
		fs := http.FileServer(http.Dir("static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
		err := http.ListenAndServe(":81", nil)*/

	/**
		gorilla/mux是一个适配 Go 的默认 HTTP 路由器的包。它具有许多功能，可在编写 Web 应用程序时提高生产力。
		它还符合 Go 的默认请求处理程序签名func (w http.ResponseWriter, r *http.Request)，因此该包可以与其他 HTTP 库（如中间件或现有应用程序）混合和匹配。
		使用go get命令从 GitHub 安装包，如下所示：
	   go get -u github.com/gorilla/mux
	*/
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		_, err := fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
		if err != nil {
			return
		}
	})
	//将请求处理程序限制为特定的 HTTP 方法。
	/*	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
		r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
		r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
		r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

		//将请求处理程序限制为特定的主机名或子域。
		r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")


		//将请求处理程序限制为 http/https。
		r.HandleFunc("/secure", SecureHandler).Schemes("https")
		r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

		//将请求处理程序限制为特定路径前缀。
		bookrouter := r.PathPrefix("/books").Subrouter()
		bookrouter.HandleFunc("/", AllBooks)
		bookrouter.HandleFunc("/{title}", GetBook)*/

	err := http.ListenAndServe(":81", r)
	if err != nil {
		return
	}
}

func (server HttpServer) Exit() {
	fmt.Println("http Exit")
}
