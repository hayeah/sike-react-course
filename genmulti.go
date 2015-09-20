package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	startTag   = "<cn>"
	endTag     = "</cn>"
	startTagRE = regexp.MustCompile(startTag)
	endTagRE   = regexp.MustCompile(endTag)
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filepath.Walk(root, walk)
}

func walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Println(err)
	}

	dir, filename := filepath.Split(path)

	if filename == "_index.md" {
		log.Println("generate multi lang:", dir)

		input, err := os.Open(path)
		if err != nil {
			log.Println(err)
			return nil
		}

		enOutput, err := os.Create(filepath.Join(dir, "index.md"))
		if err != nil {
			log.Println(err)
			return err
		}

		strip("en", input, enOutput)

		_, err = input.Seek(0, os.SEEK_SET)
		if err != nil {
			log.Println(err)
			return nil
		}

		cnOutput, err := os.Create(filepath.Join(dir, "index.md.cn"))
		if err != nil {
			log.Println(err)
			return err
		}

		strip("cn", input, cnOutput)

	}

	return nil
}

func strip(mode string, input io.Reader, output io.Writer) {
	content, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Print(string(content))

	if mode == "cn" {
		for {
			startLoc := startTagRE.FindIndex(content)
			if startLoc == nil {
				break
			}

			sectionStart := startLoc[1]
			content = content[sectionStart:]

			endLoc := endTagRE.FindIndex(content)
			if endLoc == nil {
				log.Fatalln("no end tag found:", endTag)
			}

			sectionEnd := endLoc[0]
			// log.Println("sectionStart", startLoc)
			// log.Println("sectionEnd", endLoc)

			chunk := content[0:sectionEnd]
			output.Write(chunk)

			content = content[endLoc[1]:]
		}
	} else {
		for {
			startLoc := startTagRE.FindIndex(content)
			if startLoc == nil {
				output.Write(content)
				break
			}

			output.Write(content[:startLoc[0]])
			content = content[startLoc[1]:]

			// skip to endTag
			endLoc := endTagRE.FindIndex(content)
			if endLoc == nil {
				log.Fatalln("no end tag found:", endTag)
			}

			content = content[endLoc[1]:]
		}
	}
}
