package main

import (
	"math/rand"
	"net/http"
	"io"
	"os"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func generateOutputFilename() string {
	return StringWithCharset(64, charset);
}

func (app *application) uploadFromUrl(url string) (string, error) {
	outputFilename := generateOutputFilename()
	out, err := os.Create(outputFilename)
	if err != nil {
		app.errorLog.Println("Failed output file creation.")
		return outputFilename, err
	}
	app.infoLog.Printf("Created output filename %s\n", outputFilename)
	defer out.Close()

	resp, err := http.Get(url)
	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		app.errorLog.Printf("Upload failed. Bytes written: %d\n", n)
		return outputFilename, err
	}
	app.infoLog.Printf("Upload complete. Bytes written: %d\n", n)
	return outputFilename, nil
}
