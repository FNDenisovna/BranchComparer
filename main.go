package branchcomparer

import (
	branchapi "branchComparer/internal/branchApi"
	comparer "branchComparer/internal/comparer"
	"flag"
	"fmt"
)

var (
	branch1 string
	branch2 string
	url     string
)

func init() {
	flag.StringVar(&branch1, "branch1", "", "First branch for compare")
	flag.StringVar(&branch2, "branch2", "", "First branch for compare")
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
