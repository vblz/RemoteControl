package utils

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ReadFile read the file with relative path and return content or error
func ReadFile(relPath string) ([]byte, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}

	f, err := ioutil.ReadFile(dir + relPath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// ReadHTML read the file with relative path and return content or empty HTML
// if any error
func ReadHTML(relPath string) template.HTML {
	c, err := ReadFile(relPath)
	if err != nil {
		log.Println("Can't load html file "+relPath+": ", err)
		return template.HTML("")
	}
	return template.HTML(c)
}
