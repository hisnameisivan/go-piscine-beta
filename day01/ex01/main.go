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
var pf = fmt.Printf

var (
	originalFilename string
	stolenFilename   string
)

func init() {
	flag.StringVar(&originalFilename, "old", "", "Filename of original DB")
	flag.StringVar(&stolenFilename, "new", "", "Filename of stolen DB")
	flag.Parse()
}

func main() {
	var (
		xmldoc  XMLDoc
		jsondoc JSONDoc
		oldData []byte
		newData []byte
		err     error
		cakes   map[string]bool
	)

	if originalFilename == "" && stolenFilename == "" {
		gracefulExit("Missing database names")
		// } else if originalFilename == "" {
		// 	gracefulExit("Missing original database name")
		// } else if stolenFilename == "" {
		// 	gracefulExit("Missing stolen database name")
	} else if !strings.HasSuffix(originalFilename, ".xml") {
		gracefulExit("Incorrect format of the original database")
	} else if !strings.HasSuffix(stolenFilename, ".json") {
		gracefulExit("Incorrect format of the stolen database")
	}

	p(originalFilename)
	p(stolenFilename)

	oldData, err = ioutil.ReadFile(originalFilename)
	if err != nil {
		gracefulExit(err.Error())
	}
	newData, err = ioutil.ReadFile(stolenFilename)
	if err != nil {
		gracefulExit(err.Error())
	}

	err = xmldoc.Read(oldData)
	if err != nil {
		gracefulExit(err.Error())
	}
	err = jsondoc.Read(newData)
	if err != nil {
		gracefulExit(err.Error())
	}

	cakes = make(map[string]bool)
	for _, oldCake := range xmldoc.Cakes {
		cakes[oldCake.Name] = true
		for _, newCake := range jsondoc.Cakes {
			if oldCake.Name == newCake.Name {
				if oldCake.Time != newCake.Time {
					pf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldCake.Name, newCake.Time, oldCake.Time)
				}
				ingredients := make(map[string]bool)
				for _, oldIngredient := range oldCake.Ingredients {
					ingredients[oldIngredient.Name] = true
					for _, newIngredient := range newCake.Ingredients {
						if oldIngredient.Name == newIngredient.Name {
							if oldIngredient.Unit == "" && newIngredient.Unit != "" {
								pf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", newIngredient.Unit, oldIngredient.Name, oldCake.Name)
							} else if oldIngredient.Unit != "" && newIngredient.Unit == "" {
								pf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIngredient.Unit, oldIngredient.Name, oldCake.Name)
							} else if oldIngredient.Unit != newIngredient.Unit {
								pf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIngredient.Name, oldCake.Name, newIngredient.Unit, oldIngredient.Unit)
							}
							if oldIngredient.Count != newIngredient.Count {
								pf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIngredient.Name, oldCake.Name, newIngredient.Count, oldIngredient.Count)
							}
						}
					}
				}
			}
		}
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
