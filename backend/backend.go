package backend

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var Inputf, Outputf string
var fileCount = 0
var fileWritten = 0

type File struct {
	Name, Body string
}

func getfiles(status func(msg string), progressPercent func(p float64)) []*File {
	var files []*File
	progressPercent(-1.00)
	fmt.Println(Outputf)
	walkfunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}
		if !info.IsDir() {
			files = append(files, &File{
				Name: strings.Replace(path, Inputf, "", 1),
				Body: path,
			})
			status(info.Name())
			fileCount++
			//log.Println(path) //
		}
		return nil
	}
	status("********Reading files from the directory**************")
	err := filepath.Walk(Inputf, walkfunc)

	if err != nil {
		log.Panic(err)
	}
	return files

}
func readFromFile(filePath string) string {
	info, err := os.Stat(filePath)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	if !info.IsDir() {
		r, er := ioutil.ReadFile(filePath)
		if er != nil {
			log.Printf(er.Error())
			return ""
		}
		return string(r)
	}
	return ""
}
func StartTakingBackup(status func(msg string), progressPercent func(p float64)) {
	Outputf, e := os.Create(Outputf)
	if e != nil {
		panic(e)
	}
	w := zip.NewWriter(Outputf)
	files := getfiles(status, progressPercent)
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		status(file.Body)
		//log.Println(file.Body)
		_, err = f.Write([]byte(readFromFile(file.Body)))
		fileWritten++
		progressPercent(float64((fileWritten * 100.0) / fileCount))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

}
