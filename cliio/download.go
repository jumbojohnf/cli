package cliio

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func Download(url string) (*os.File, error) {
	fmt.Println("🌏 Downloading", url)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	downloadDir, err := ioutil.TempDir("", "download")
	if err != nil {
		return nil, err
	}
	dst, err := os.Create(filepath.Join(downloadDir, filepath.Base(url)))
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	_, err = io.Copy(dst, response.Body)
	return dst, err
}
