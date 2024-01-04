package main

import (
	branchapi "branchcomparer/internal/branchApi"
	comparer "branchcomparer/internal/comparer"
	"flag"
	"fmt"
	"reflect"
)

var Branches = map[string]struct{}{
	"sisyphus": {},
	"p10":      {},
	"p9":       {},
	"p8":       {},
	"p7":       {},
	"c10f2":    {},
	"c10f1":    {},
	"c9f2":     {},
	"c9f1":     {},
	"c8.1":     {},
	"c8":       {},
	"c7.1":     {},
	"c7":       {},
}

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
	if err := checkBranch(branch1, "branch1"); err != nil {
		fmt.Println(err)
		return
	}

	if err := checkBranch(branch2, "branch2"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("branch1 = %v\n", branch1)
	fmt.Printf("branch2 = %v\n", branch2)

	api := branchapi.New(url)
	comp := comparer.New(branch1, branch2, api)
	fmt.Println("Start comparing")
	result, err := comp.Compare()
	if err != nil {
		fmt.Printf("Comparing is failed. Error: %v.\n", err)
		return
	}
	fmt.Println("Finish comparing")
	fmt.Println(string(result[:]))
}

func checkBranch(b string, name string) (err error) {
	if b == "" {
		return fmt.Errorf("Name of -%s is empty. Try again with value.\n", name)
	}

	if _, ok := Branches[b]; !ok {
		return fmt.Errorf("Name of -%s is not allowed. Try again with value in: %v.\n", name, reflect.ValueOf(Branches).MapKeys())
	}

	return nil
}

func (ec ExtComparerType) Compare(b1 string, b2 string) (string, error) {
	branch1 = b1
	branch2 = b2
	url = "https://rdb.altlinux.org/api/export/branch_binary_packages/"

	if err := checkBranch(branch1, "branch1"); err != nil {
		return "", err
	}

	if err := checkBranch(branch2, "branch2"); err != nil {
		return "", err
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
