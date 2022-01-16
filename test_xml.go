package main

import (
	"encoding/xml"
	"fmt"
	"github.com/alfu32/txxml/text"
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
	tt:=text.Text{}
	tt=tt.Init("some text words").
	  Separators(" ")

	fmt.Printf("%+v\n",tt)
	fmt.Printf("ToLowerSnake [%s]\n",tt.ToLowerSnake())
	fmt.Printf("ToUpperSnake [%s]\n",tt.ToUpperSnake())
	fmt.Printf("ToLowerKebap [%s]\n",tt.ToLowerKebap())
	fmt.Printf("ToUpperKebap [%s]\n",tt.ToUpperKebap())
}
