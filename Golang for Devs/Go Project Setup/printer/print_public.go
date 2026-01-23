package printer // package name

import "fmt"

func PrintPublic(s string) {
	printPrivate(s)
}

func printPrivate(s string) {
	fmt.Println(s)
}
