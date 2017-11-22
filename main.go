package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Database root node
type Database struct {
	Object *Object `xml:"object,omitempty"`
}

// Object encapsulates a single note
type Object struct {
	AttrID    string       `xml:"id,attr"`
	AttrType  string       `xml:"type,attr"`
	Attribute []*Attribute `xml:"attribute,omitempty"`
}

// Attribute individual note data
type Attribute struct {
	AttrName string `xml:"name,attr"`
	AttrType string `xml:"type,attr"`
	Text     string `xml:",chardata"`
}

const usage = `no file name supplied

usage: snexport /path/to/Simplenote.storedata path/to/backup_dir
`

func main() {
	if len(os.Args) != 3 {
		log.Fatalln(usage)
	}

	fileName := os.Args[1]
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Fatalln("file '" + fileName + "'' does not exist")
	}

	exportDir := os.Args[2]
	if err := os.MkdirAll(exportDir, os.ModePerm); err != nil {
		log.Fatalln(err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	var simperiumKey string
	var content string
	var nExported int
	var db Database
	xml.Unmarshal(data, &db)
	for _, o := range db.Object.Attribute {
		if o.AttrName == "simperiumkey" {
			simperiumKey = o.Text
		}
		if o.AttrName == "content" {
			content = o.Text
		}
		if simperiumKey != "" && content != "" {
			path := exportDir + "/" + simperiumKey + ".txt"
			if err := ioutil.WriteFile(path, []byte(content), os.ModePerm); err != nil {
				log.Println(err)
			} else {
				nExported++
			}
			simperiumKey = ""
			content = ""
		}
	}
	fmt.Println("Exported", nExported, "files")
	absPath, err := filepath.Abs(exportDir)
	if err != nil {
		fmt.Println("Location", exportDir)
	} else {
		fmt.Println("Location", absPath)
	}
}
