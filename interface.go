package main

import (
	"bufio"
	"fmt"
	"os"
)

func basicMainFunc() {
	template := "package main \nimport 'fmt' \nfunc main(){ \n} "
	fmt.Printf("%s", template)
}

func main() {
	menuItems := []string{"Basic main function", "Hello world", "Basic server"}

	fmt.Println("\nHola, choose the boilerplate you want to start off with. Also if you have an editor of preference then specify that and we will open the code in it. Or else you can leave it there, in which case a file will be created in the current working directory with a name you specify. Happy Coding!\n")
	for i, item := range menuItems {
		fmt.Printf("%d) %s\n", i+1, item)
	}

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	opt := input.Text()

	switch opt {
	case "1":
		fmt.Println("I will build you a template for ", menuItems[0])
		basicMainFunc()
	case "2":
		fmt.Println("I will build you a template for ", menuItems[1])
	case "3":
		fmt.Println("I will build you a template for ", menuItems[2])
	}
}
