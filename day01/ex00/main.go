package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

var p = fmt.Println

func main() {
	var (
		// xmldoc recipes
		// xmldoc XMLDoc

		jsondoc JSONDoc
	)

	// xmlData, err := ioutil.ReadFile("res.xml")
	// if err != nil {
	// 	p(err)
	// }

	// p(string(xmlData))
	// err = xml.Unmarshal(xmlData, &xmldoc)
	// if err != nil {
	// 	p(err)
	// }
	// p(xmldoc)

	// data, err := xml.Marshal(xmldoc)
	// if err != nil {
	// 	p(err)
	// }
	// p(string(data))

	/*
	** json
	 */

	jsonData, err := ioutil.ReadFile("res.json")
	if err != nil {
		p(err)
	}

	p(string(jsonData))
	err = json.Unmarshal(jsonData, &jsondoc)
	if err != nil {
		p(err)
	}
	p(jsondoc)

	data, err := json.Marshal(jsondoc)
	if err != nil {
		p(err)
	}
	p(string(data))
}

type BDReader interface {
	Read(p []byte) (n int, err error)
}

// type recipes struct {
// 	Cakes []struct {
// 		Name        string `xml:"name"`
// 		Stovetime   string `xml:"stovetime"`
// 		Ingredients struct {
// 			Items []struct {
// 				Itemname  string `xml:"itemname"`
// 				Itemcount string `xml:"itemcount"`
// 				Itemunit  string `xml:"itemunit"`
// 			} `xml:"item"`
// 		} `xml:"ingredients"`
// 	} `xml:"cake"`
// }

// type XMLDoc struct {
// 	XMLName xml.Name `xml:"recipes"`
// 	Cakes   []struct {
// 		Name        string `xml:"name"`
// 		Stovetime   string `xml:"stovetime"`
// 		Ingredients []struct {
// 			// Items []struct {
// 			Itemname  string `xml:"itemname"`
// 			Itemcount string `xml:"itemcount"`
// 			Itemunit  string `xml:"itemunit"`
// 			// } `xml:"item"`
// 		} `xml:"ingredients>item"`
// 	} `xml:"cake"`
// }

// type JSONDoc struct {
// 	// JSONName json.Name `json:"recipes"`
// 	Cakes []struct {
// 		Name        string `json:"name"`
// 		Stovetime   string `json:"time"`
// 		Ingredients []struct {
// 			Itemname  string `json:"ingredient_name"`
// 			Itemcount string `json:"ingredient_count"`
// 			Itemunit  string `json:"ingredient_unit"`
// 		} `json:"ingredients"`
// 	} `json:"cake"`
// }

// type XMLDoc struct {
// 	XMLName xml.Name `xml:"recipes"`
// 	Cakes   []Cake   `xml:"cake"`
// }

// type JSONDoc struct {
// 	Cakes []Cake `json:"cake"`
// }

// type Cake struct {
// 	Name        string `xml:"name" json:"name"`
// 	Time        string `xml:"stovetime" json:"time"`
// 	Ingredients []struct {
// 		Name  string `xml:"itemname" json:"ingredient_name"`
// 		Count string `xml:"itemcount" json:"ingredient_count"`
// 		Unit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
// 	} `xml:"ingredients>item" json:"ingredients"`
// }

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

// func (doc *XMLDoc) Read(p byte) (n int, err error) {

// }
