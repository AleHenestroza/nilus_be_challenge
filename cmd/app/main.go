package main

import (
	"flag"
	"fmt"
	"os"
)

func readFile(filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	var (
		rulesPath string
		matchPath string
	)

	flag.StringVar(&rulesPath, "rules", "", "Path to rules file")
	flag.StringVar(&matchPath, "match", "", "Path to match file")
	flag.Parse()

	rules, err := readFile(rulesPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	match, err := readFile(matchPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(rules))
	fmt.Println(string(match))
}
