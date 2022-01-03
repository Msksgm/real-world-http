package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")
	part := make(textproto.MIMEHeader)
	part.Set("Content-TYpe", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumnail"; filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
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
	resp, err := http.Post("http://localhost:18888", "text/plain", &buffer)
	if err != nil {
		// 送信失敗
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
