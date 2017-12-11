package main

import (
	"encoding/csv"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	input := flag.String("i", "", "Input XML filename")
	output := flag.String("o", "", "Output CSV filename; default STDOUT")
	maxattr := flag.Int("maxattr", 5, "Maximum number of attributes for an element; default 5")
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

	// open input file and read into memory
	bytes, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal("ioutil.ReadFile() error:" + err.Error())
	}
	// convert byte slice to a string
	r := strings.NewReader(string(bytes))

	// Output the headers
	headers := []string{"Depth", "Type", "Name", "Text"}
	for i := 0; i < *maxattr; i++ {
		headers = append(headers, fmt.Sprintf("Attribute %v", i+1))
		headers = append(headers, fmt.Sprintf("Value %v", i+1))
	}
	err = w.Write(headers)
	if err != nil {
		log.Fatal("w.Write() headers error:" + err.Error())
	}

	// start decode loop and track depth for future
	// pretty printing purposes
	lastStartName := ""
	decoder := xml.NewDecoder(r)
	depth := 0
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("decoder.Token():\n%v\n", err)
		}
		switch t := token.(type) {
		case xml.StartElement:
			lastStartName = processStart(w, xml.StartElement(t), depth)
			depth++
		case xml.EndElement:
			depth--
			processEnd(w, xml.EndElement(t), depth)
		case xml.CharData:
			processCharData(w, lastStartName, xml.CharData(t), depth)
		case xml.Comment:
			processComment(w, xml.Comment(t), depth)
		case xml.ProcInst:
			processProcInst(w, xml.ProcInst(t), depth)
		case xml.Directive:
			processDirective(w, xml.Directive(t), depth)
		default:
			fmt.Println("Unknown!?")
		}
	}
	w.Flush()
}

func processStart(w *csv.Writer, e xml.StartElement, depth int) string {
	var outcells []string
	outcells = append(outcells, fmt.Sprintf("%v", depth))
	outcells = append(outcells, "Start")
	outcells = append(outcells, e.Name.Local)
	outcells = append(outcells, "") // text val column

	for _, a := range e.Attr {
		outcells = append(outcells, a.Name.Local)
		outcells = append(outcells, a.Value)
	}

	err := w.Write(outcells)
	if err != nil {
		log.Fatal("w.Write() start row error:" + err.Error())
	}
	return e.Name.Local
}

func processEnd(w *csv.Writer, e xml.EndElement, depth int) {
	var outcells []string
	outcells = append(outcells, fmt.Sprintf("%v", depth))
	outcells = append(outcells, "End")
	outcells = append(outcells, e.Name.Local)

	err := w.Write(outcells)
	if err != nil {
		log.Fatal("w.Write() end row error:" + err.Error())
	}
}
func processCharData(w *csv.Writer, ename string, e xml.CharData, depth int) {
	txt := strings.TrimSpace(string(e))
	if txt == "" {
		return
	}
	var outcells []string
	outcells = append(outcells, fmt.Sprintf("%v", depth))
	outcells = append(outcells, "CharData")
	outcells = append(outcells, ename)
	outcells = append(outcells, txt)

	err := w.Write(outcells)
	if err != nil {
		log.Fatal("w.Write() chardata row error:" + err.Error())
	}
}
func processComment(w *csv.Writer, e xml.Comment, depth int) {
	var outcells []string
	outcells = append(outcells, fmt.Sprintf("%v", depth))
	outcells = append(outcells, "Comment")
	outcells = append(outcells, strings.TrimSpace(string(e)))

	err := w.Write(outcells)
	if err != nil {
		log.Fatal("w.Write() comment row error:" + err.Error())
	}
}
func processProcInst(w *csv.Writer, e xml.ProcInst, depth int) {
	var outcells []string
	outcells = append(outcells, fmt.Sprintf("%v", depth))
	outcells = append(outcells, "ProcInst")
	outcells = append(outcells, e.Target)
	outcells = append(outcells, string(e.Inst))

	err := w.Write(outcells)
	if err != nil {
		log.Fatal("w.Write() procinst row error:" + err.Error())
	}
}
func processDirective(w *csv.Writer, e xml.Directive, depth int) {
	var outcells []string
	outcells = append(outcells, fmt.Sprintf("%v", depth))
	outcells = append(outcells, "Directive")
	outcells = append(outcells, strings.TrimSpace(string(e)))

	err := w.Write(outcells)
	if err != nil {
		log.Fatal("w.Write() directive row error:" + err.Error())
	}
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: parseAny [options]\n")
	flag.PrintDefaults()
}
