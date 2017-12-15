package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// note that this struct does not include either
// ProcInst or Directive since tags are not
// provided for them (probably for reasonable cause)
type xmlany struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Content string     `xml:",chardata"`
	Comment string     `xml:",comment"`
	Nested  []*xmlany  `xml:",any"`
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: identityXform [options]\n")
	flag.PrintDefaults()
}

func main() {
	input := flag.String("i", "", "Input XML filename")
	output := flag.String("o", "", "Output XML filename; default STDOUT")
	indent := flag.Bool("indent", false, "Use indented format (pretty print); default is false")
	help := flag.Bool("help", false, "Show usage message")
	flag.Parse()

	if *help {
		usage("Help Message")
		os.Exit(0)
	}

	// open input file
	fi, fierr := os.Open(*input)
	if fierr != nil {
		log.Fatal("os.Open() Error:" + fierr.Error())
	}
	defer fi.Close()
	// ... and read into memory
	data, err := ioutil.ReadAll(fi)

	v := &xmlany{}
	err = xml.Unmarshal([]byte(data), v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	var doc []byte
	if *indent {
		doc, err = xml.MarshalIndent(*v, "", "    ")
		if err != nil {
			fmt.Printf("xml.MarshalIndent() error: %v", err)
			return
		}
	} else {
		doc, err = xml.Marshal(*v)
		if err != nil {
			fmt.Printf("xml.Marshal() error: %v", err)
			return
		}

	}

	if *output == "" {
		fmt.Println(string(doc))
	} else {
		ioutil.WriteFile(*output, doc, 0600)
	}

}
