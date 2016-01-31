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

	path := filepath.Join(dir, relPath)
	f, err := ioutil.ReadFile(path)
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

// ReadFilesInDir read all files in dir and return it in map[path][]fileBytes
func ReadFilesInDir(relPath string) map[string][]byte {
	files, err := ioutil.ReadDir(relPath)
	if err != nil {
		return nil
	}
	result := make(map[string][]byte, 2)

	for _, f := range files {
		if !f.IsDir() {
			path := filepath.Join(relPath, f.Name())
			bytes, err := ReadFile(path)
			if err == nil {
				result[path] = bytes
			} else {
				log.Println("Can't load file "+path+": ", err)
			}
		}
	}
	return result
}
