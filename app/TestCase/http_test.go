package testcase

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"testing"
)

type Info struct {
	Name  string `json:"name"`
	Age   int64  `json:"age"`
	Store Store  `json:"store"`
}

type Store struct {
	StoreName string `json:"store_name"`
	StoreId   int64  `json:"store_id"`
}

func TestHttp1(t *testing.T) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	/*	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/static/a.html")
	})*/

	http.ListenAndServe(":80", nil)
}


func TestHttp2(t *testing.T) {
	r := mux.NewRouter()

	brooder := r.PathPrefix("/books").Subrouter()
	brooder.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "books-test---")
	})

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)
}


func TestString1(t *testing.T) {
	str := "name,game,price"
	fmt.Println(len(strings.Split(str, ",")))
}
func TestJson1(t *testing.T) {
	str := `{"name":"buffer","age":30,"store":{"store_name":"STORE100","store_id":20}}`
	var InfoData Info
	json.Unmarshal([]byte(str), &InfoData)
	fmt.Println("info--"+InfoData.Name+"----", InfoData.Age, InfoData.Store.StoreName, InfoData.Store.StoreId)
}

func TestJson2(t *testing.T) {
	str := `[{"name":"buffer","age":30,"store":{"store_name":"STORE100","store_id":20}},{"name":"buffer","age":30,"store":{"store_name":"STORE100","store_id":20}},{"name":"buffer","age":30,"store":{"store_name":"STORE100","store_id":20}},{"name":"buffer","age":30,"store":{"store_name":"STORE100","store_id":20}},{"name":"buffer","age":30,"store":{"store_name":"STORE100","store_id":20}},{"name":"buffer","age":30,"store":{"store_name":"STORE100","store_id":20}},{"name":"buffer","age":30,"store":{"store_name":"STORE100","store_id":20}},{"name":"buffer","age":30,"store":{"store_name":"STORE100","store_id":20}},{"name":"buffer","age":30,"store":{"store_name":"STORE100","store_id":20}}]`
	var InfoData []Info
	json.Unmarshal([]byte(str), &InfoData)

	//json 字符串转结构体
	fmt.Println(InfoData[0].Name)

	data, err := json.Marshal(InfoData)

	if err != nil {
		panic(err)
	}
	//结构转JSON字符串
	fmt.Println(string(data))
}
