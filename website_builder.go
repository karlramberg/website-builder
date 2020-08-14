package main

import (
	"github.com/gomarkdown/markdown"
	"os"
	"path/filepath"
	"io/ioutil"
	"fmt"
	"strings"
)

const CONTENTTOKEN string = "#CONTENT#"
func main() {
	// Exit if 2 arguments are not given
	if len(os.Args) < 2 {
		fmt.Println("USAGE: $ website_builder <markdown folder> <template file>")
	}

	root := os.Args[1]

	// Read the HTML template file into a string
	template, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println("ERROR: Template file not found or is protected")
		return	
	}

	// Walk through the root directory file by file
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// Create mirrored folder structure
		if info.IsDir() {
			newFolder := path[len(root):]
			os.MkdirAll(newFolder, 0755)
			fmt.Println("Generating " + newFolder + "...")
			return nil
		}

		// Generate html file from any markdown file
		newPath := path[len(root):len(path)-len(filepath.Ext(path))] + ".html"
		fmt.Println("Generating " + newPath + "...")
		md,_ := ioutil.ReadFile(path)
		html := string(markdown.ToHTML(md, nil, nil))

		// Replace the content token with generated HTML
		output := strings.Replace(string(template), CONTENTTOKEN, html, 1)	
		ioutil.WriteFile(newPath, []byte(output), 0644)
		return nil
	})
}
