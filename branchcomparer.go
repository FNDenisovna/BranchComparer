package main

import (
	branchapi "branchcomparer/internal/branchApi"
	comparer "branchcomparer/internal/comparer"
	"flag"
	"fmt"
)

var (
	branch1 string
	branch2 string
	url     string
)

type ExtComparerType struct{}

var ExtComparer ExtComparerType

func init() {
	flag.StringVar(&branch1, "branch1", "sisyphus", "First branch for compare")
	flag.StringVar(&branch2, "branch2", "p10", "First branch for compare")
	flag.Parse()

	url = "https://rdb.altlinux.org/api/export/branch_binary_packages/"
}

func main() {
	if branch1 == "" {
		fmt.Printf("Name of -branch1 is empty. Try again with value.\n")
		return
	}
	if branch2 == "" {
		fmt.Printf("Name of -branch2 is empty. Try again with value.\n")
		return
	}

	fmt.Printf("branch1 = %v\n", branch1)
	fmt.Printf("branch2 = %v\n", branch2)

	api := branchapi.New(url)
	comp := comparer.New(branch1, branch2, api)
	result, err := comp.Compare()
	if err != nil {
		fmt.Printf("Comparing is failed. Error: %v.\n", err)
		return
	}

	fmt.Println(string(result[:]))
}

func (ec ExtComparerType) Compare(b1 string, b2 string) (string, error) {
	branch1 = b1
	branch2 = b2
	url = "https://rdb.altlinux.org/api/export/branch_binary_packages/"

	if branch1 == "" {
		return "", fmt.Errorf("Name of -branch1 is empty. Try again with value.\n")
	}
	if branch2 == "" {
		return "", fmt.Errorf("Name of -branch2 is empty. Try again with value.\n")
	}

	fmt.Printf("branch1 = %v\n", branch1)
	fmt.Printf("branch2 = %v\n", branch2)

	api := branchapi.New(url)
	comp := comparer.New(branch1, branch2, api)
	result, err := comp.Compare()
	if err != nil {
		return "", fmt.Errorf("Comparing is failed. Error: %v.\n", err)

	}

	return string(result[:]), nil
}
