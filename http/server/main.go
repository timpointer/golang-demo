package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	http.HandleFunc("/bar/good", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.String())
		fmt.Println(r.RequestURI)

		//u := url.PathEscape("http://www.baidu.com/" + r.RequestURI)
		u := r.RequestURI
		fmt.Println("u", u)
		fmt.Println(strings.Index(u, "?"))
		fmt.Println("url------", u[:strings.Index(u, "?")])
		u, _ = url.PathUnescape(u)
		fmt.Println("u", u)
		if r.Method != "POST" {
			log.Println(r.Method)
			http.Error(w, "Not Supported", http.StatusNotAcceptable)
			return
		}
		//读取body,然后通过base64编码转为string
		body, _ := ioutil.ReadAll(r.Body)
		log.Println("body", string(body))
		//签名正确后，把post过来的json转为struct对象
		apidate := APIData{}
		//err := json.NewDecoder(body).Decode(&apidate)
		err := json.Unmarshal(body, &apidate)
		if err != nil {
			log.Println("json解析错误：", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("server error", err)
		http.Error(w, "server error", http.StatusBadRequest)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
