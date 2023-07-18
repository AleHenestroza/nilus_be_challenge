package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alehenestroza/fifa_review/internal/file_utils"
)

func main() {
	var (
		rulesPath string
		matchPath string
	)

	flag.StringVar(&rulesPath, "rules", "", "Path to rules file")
	flag.StringVar(&matchPath, "match", "", "Path to match file")
	flag.Parse()

	rules, err := file_utils.ReadFile(rulesPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	match, err := file_utils.ReadFile(matchPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(rules))
	fmt.Println(string(match))
}
