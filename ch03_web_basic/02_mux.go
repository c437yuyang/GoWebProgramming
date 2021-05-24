package main

import (
    "fmt"
    "log"
    "net/http"
)

type MyMux struct {
}

// 自定义 Mux
func (mux *MyMux) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {
    if req.URL.Path == "/" {
        sayHelloName1(respWriter, req)
        return
    }
    http.NotFound(respWriter, req)
    return
}

func sayHelloName1(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    log.Printf("request incoming, URL:%s, scheme:%s\n", r.URL, r.URL.Scheme)
    fmt.Fprintf(w, "hello")
}

func main() {
    mux := &MyMux{}
    http.ListenAndServe(":9090", mux)
}
