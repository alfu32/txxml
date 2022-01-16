package main

import (
	"encoding/xml"
	"fmt"
)

type Entity struct {
	XMLName  xml.Name `xml:"entity"`
	Text     string   `xml:",chardata"`
	Incident struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Code string `xml:"code,attr"`
	} `xml:"incident"`
	References struct {
		// Text      string `xml:",chardata"`
		ID          string `xml:"id,attr"`
		Code        string `xml:"code,attr"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Author      string `xml:"author"`
		Reference   []struct {
			//Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Code string `xml:"code,attr"`
		} `xml:"reference"`
	} `xml:"references"`
}

func main() {
	fmt.Println("Hello World")
	ent := Entity{}
	if entXml, err := xml.Marshal(ent); err == nil {
		fmt.Println(string(entXml))
	}

}
