package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ByteFormater(size int64) string {
	if size < 1000 {
		return fmt.Sprintf("%d byte", size)
	} else if size < 1000*1000 {
		return fmt.Sprintf("%d KB", size/1000)
	} else {
		return fmt.Sprintf("%d MB", size/(1000*1000))
	}
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type CustomReader struct {
	io.Reader
	total int64
	size  int64
}

func (cr *CustomReader) Read(p []byte) (int, error) {
	n, err := cr.Reader.Read(p)
	cr.total += int64(n)

	if err == nil {
		fmt.Println("Downloaded ", ByteFormater(cr.total), " from ", ByteFormater(cr.size), " at speed of ", ByteFormater(int64(n)))
	}

	return n, err
}

type ImageInfo struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	FileSize  int64  `json:"file_size"`
	FileUrl   string `json:"file_url"`
	Extension string `json:"file_ext"`
}

func (y *ImageInfo) Download(path string) {
	response, err := http.Get(y.FileUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	file, err := os.Create(fmt.Sprintf("%s.%s", RandStringBytes(5), y.Extension))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	customBody := &CustomReader{Reader: response.Body, size: y.FileSize}
	_, err = io.Copy(file, customBody)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleImages() {
	response, _ := http.Get("https://yande.re/post.json")
	body, _ := ioutil.ReadAll(response.Body)
	images := []ImageInfo{}

	if err := json.Unmarshal(body, &images); err != nil {
		log.Fatal(err)
	}
	n := 0
	for _, image := range images {
		image.Download("")
		n++
	}
}
