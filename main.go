package branchcomparer

import (
	"flag"
	"fmt"
)

var (
	branch1 string
	branch2 string
)

func init() {
	flag.StringVar(&branch1, "branch1", "", "First branch for compare")
	flag.StringVar(&branch2, "branch2", "", "First branch for compare")
	flag.Parse()
}

func main() {
	if branch1 == "" {
		fmt.Printf("Name of -branch1 is empty. Try again with value.")
		return
	}
	if branch2 == "" {
		fmt.Printf("Name of -branch2 is empty. Try again with value.")
		return
	}

	fmt.Printf("branch1 = %v\n", branch1)
	fmt.Printf("branch2 = %v\n", branch2)

}
