package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "github.com/skip2/go-qrcode"
)

func main() {
  fmt.Println("TextToQR Service Up")
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "shoo shoo")
  })

  http.HandleFunc("/qr", func(w http.ResponseWriter, r *http.Request) {
    keys, ok := r.URL.Query()["text"]
    
    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'text' is missing")
        return
    }

    key := keys[0]

    log.Println("Url Param 'text' is: " + string(key))    

    var png []byte
    fmt.Println(r)
    png, err := qrcode.Encode(string(key), qrcode.Medium, 256)

    if err != nil {} 
    w.Write(png)
  })
  
  
  log.Fatal(http.ListenAndServe(getPort(), nil))
}

func getPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":8080"
}

