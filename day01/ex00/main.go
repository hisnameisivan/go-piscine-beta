package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var p = fmt.Println

var (
	filename string
)

func init() {
	flag.StringVar(&filename, "f", "", "Filename")
	flag.Parse()
}

func main() {
	var (
		xmldoc  XMLDoc
		jsondoc JSONDoc
		out     []byte
	)

	if filename == "" {
		gracefulExit("No input file")
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		gracefulExit(err.Error())
	}

	if strings.HasSuffix(filename, ".xml") {
		err = xmldoc.Read(data)
		if err != nil {
			gracefulExit(err.Error())
		}
		out, err = xmldoc.Write()
		if err != nil {
			gracefulExit(err.Error())
		}
		p(string(out))
	} else if strings.HasSuffix(filename, ".json") {
		err = jsondoc.Read(data)
		if err != nil {
			gracefulExit(err.Error())
		}
		out, err = jsondoc.Write()
		if err != nil {
			gracefulExit(err.Error())
		}
		p(string(out))
	} else {
		gracefulExit("Invalid file format")
	}
}

type DBReader interface {
	Read(data []byte) error
	Write() ([]byte, error)
}

type recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cakes   []struct {
		Name        string `xml:"name" json:"name"`
		Time        string `xml:"stovetime" json:"time"`
		Ingredients []struct {
			Name  string `xml:"itemname" json:"ingredient_name"`
			Count string `xml:"itemcount" json:"ingredient_count"`
			Unit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
		} `xml:"ingredients>item" json:"ingredients"`
	} `xml:"cake" json:"cake"`
}

type XMLDoc recipes
type JSONDoc recipes

func (doc *XMLDoc) Read(data []byte) error {
	err := xml.Unmarshal(data, doc)
	return err
}

func (doc *XMLDoc) Write() ([]byte, error) {
	out, err := json.MarshalIndent(*doc, "", "    ")
	return out, err
}

func (doc *JSONDoc) Read(data []byte) error {
	err := json.Unmarshal(data, doc)
	return err
}

func (doc *JSONDoc) Write() ([]byte, error) {
	out, err := xml.MarshalIndent(*doc, "", "    ")
	return out, err
}

func gracefulExit(msg string) {
	p(msg)
	os.Exit(1)
}
