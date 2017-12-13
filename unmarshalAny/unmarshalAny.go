package main

import (
	"encoding/csv"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func main() {
	input := flag.String("i", "", "Input XML filename")
	output := flag.String("o", "", "Output CSV filename; default STDOUT")
	help := flag.Bool("help", false, "Show usage message")
	flag.Parse()

	if *help {
		usage("Help Message")
		os.Exit(0)
	}

	var w *csv.Writer
	if *output == "" {
		w = csv.NewWriter(os.Stdout)
	} else {
		fo, foerr := os.Create(*output)
		if foerr != nil {
			log.Fatal("os.Create() error:" + foerr.Error())
		}
		defer fo.Close()
		w = csv.NewWriter(fo)
	}
	headers := []string{"Depth", "Type", "Value"}
	err := w.Write(headers)
	if err != nil {
		log.Fatal("w.Write() error:" + err.Error())
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

	depth := 0
	walk(depth, v, w)
	w.Flush()

}

func walk(depth int, v *xmlany, w *csv.Writer) {

	var cells []string

	// Local Name
	cells = append(cells, fmt.Sprintf("%v", depth))
	cells = append(cells, "XMLName Local")
	cells = append(cells, v.XMLName.Local)
	err := w.Write(cells)
	if err != nil {
		log.Fatal("w.Write() error:" + err.Error())
	}

	// Namespace name if any
	cells[1] = "XMLName Space"
	cells[2] = v.XMLName.Space
	err = w.Write(cells)
	if err != nil {
		log.Fatal("w.Write() error:" + err.Error())
	}

	// Element attributes
	for n := range v.Attrs {
		// attribute name
		cells[1] = "Attr Name"
		cells[2] = v.Attrs[n].Name.Local
		err = w.Write(cells)
		if err != nil {
			log.Fatal("w.Write() error:" + err.Error())
		}

		// attribute value
		cells[1] = "Attr value"
		cells[2] = v.Attrs[n].Value
		err = w.Write(cells)
		if err != nil {
			log.Fatal("w.Write() error:" + err.Error())
		}
	}

	// Character data (whitespace removed)
	content := strings.TrimSpace(v.Content)
	if content != "" {
		cells[1] = "CharData"
		cells[2] = content
		err = w.Write(cells)
		if err != nil {
			log.Fatal("w.Write() error:" + err.Error())
		}
	}

	// Comment data (whitespace removed)
	comment := strings.TrimSpace(v.Comment)
	if comment != "" {
		cells[1] = "Comment"
		cells[2] = comment
		err = w.Write(cells)
		if err != nil {
			log.Fatal("w.Write() error:" + err.Error())
		}
	}
	w.Flush()

	if len(v.Nested) > 0 {
		depth++
		for n := range v.Nested {
			walk(depth, v.Nested[n], w)
		}
		depth--
	}
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: unmarshalAny [options]\n")
	flag.PrintDefaults()
}
