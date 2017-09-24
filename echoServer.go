package main

import (
  "fmt"
  "log"
  "net/http"
  "net/http/httputil"
  "io/ioutil"
  "github.com/k0kubun/pp"
)

func main() {
  var httpServer http.Server
  http.HandleFunc("/", handler2)
  http.HandleFunc("/digest", handlerDigest)
  log.Println("start http listening :1888")
  httpServer.Addr = ":1888"
  log.Println(httpServer.ListenAndServe())
}

func handler2(w http.ResponseWriter, r *http.Request)  {
  dump, err := httputil.DumpRequest(r, true)
  if err != nil {
    http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
    return
  }
  fmt.Println(string(dump))
  fmt.Fprintf(w, "<html><body>hello</body></html>¥n")
}


func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Set-Cookie", "VISIT=TRUE")
  if _, ok := r.Header["Cookie"]; ok {
    // クッキーがあるということは一度来たことがある人
    fmt.Fprintf(w, "<html><body>2回目以降</body></html>")
  } else {
    fmt.Fprintf(w, "<html><body>初訪問</body></html>")
  }
}

func handlerDigest(w http.ResponseWriter, r *http.Request) {
  pp.Printf("URL: %s\n", r.URL.String())
  pp.Printf("Query: %v\n", r.URL.Query())
  pp.Printf("Proto: %s\n", r.Proto)
  pp.Print("Method: %s\n", r.Method)
  pp.Print("Header: %s\n", r.Header)
  defer r.Body.Close()
  body,_ := ioutil.ReadAll(r.Body)
  fmt.Printf("\n--body--\n%s\n", string(body))
  if _, ok := r.Header["Authorization"]; !ok {
    w.Header().Add("WWW-Authenticate", `Digest realm="Secret Zone", nonce="TgLc25U2BQA=f510a2780473e18e6587be702c2e67fe2b04afd", algorithm=MD5, qop="auth"`)
    w.WriteHeader(http.StatusUnauthorized)
  } else {
    fmt.Fprint(w, "<head><body>secret page</body></html>")
  }

}
