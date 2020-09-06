package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
	"image/png"
	"github.com/boombuler/barcode"
 	"github.com/boombuler/barcode/qr"
 	"github.com/boombuler/barcode/code128"
)

func main() {
  fmt.Println("TextToQR Service Up")
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "shoo shoo")
  })

  http.HandleFunc("/barcode", func(w http.ResponseWriter, r *http.Request) {
    keys, ok := r.URL.Query()["text"]
    
    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'text' is missing")
        return
    }

    key := keys[0]

   	bcImg, _ := code128.Encode(key)
   	scaled, _ := barcode.Scale(bcImg, bcImg.Bounds().Dx(), 80)
    png.Encode(w, scaled)
  })


  http.HandleFunc("/qr", func(w http.ResponseWriter, r *http.Request) {
    keys, ok := r.URL.Query()["text"]
    
    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'text' is missing")
        return
    }

    key := keys[0]

    log.Println("Url Param 'text' is: " + string(key))    


    qrCode, _ := qr.Encode(key, qr.M, qr.Auto)
    qrCode, _ = barcode.Scale(qrCode, 200, 200)
    png.Encode(w, qrCode)
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
