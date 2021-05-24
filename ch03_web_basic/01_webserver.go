package main

import (
    "fmt"
    "log"
    "net/http"
)

// type HandlerFunc func(ResponseWriter, *Request) 实现了这个类，因此可以被http.HandleFunc注册进去
func sayHelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    log.Printf("request incoming, URL:%s, scheme:%s\n", r.URL, r.URL.Scheme)
    fmt.Fprintf(w, "hello")
}

func main() {
    http.HandleFunc("/", sayHelloName)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("listen error", err)
    }
}
