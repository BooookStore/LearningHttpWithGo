package main

import (
  "bytes"
  "io"
  "log"
  "mime/multipart"
  "net/http"
  "os"
)

func main() {
  var buffer bytes.Buffer
  writer := multipart.NewWriter(&buffer)
  writer.WriteField("name", "Micheal Jackson")
  fileWriter, err := writer.CreateFormFile("thumbail", "photo.jpg")
  if err != nil {
    panic(err)
  }
  readFile, err := os.Open("photo.jpg")
  if err != nil {
    panic(err)
  }
  defer readFile.Close()
  io.Copy(fileWriter, readFile)
  writer.Close()

  resp, err := http.Post("http://localhost:1888", writer.FormDataContentType(), &buffer)
  if err != nil {
    panic(err)
  }
  log.Println("Status:", resp.Status)
}
