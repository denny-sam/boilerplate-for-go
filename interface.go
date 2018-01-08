package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	menuItems := []string{"Basic main function", "Basic server"}

	fmt.Println("\nHola, choose the boilerplate you want to start off with. Also if you have an editor of preference then specify that and we will open the code in it. Or else you can leave it there, in which case a file will be created in the current working directory with a name you specify. Happy Coding!\n")
	for i, item := range menuItems {
		fmt.Printf("%d) %s\n", i+1, item)
	}
	fmt.Println("\nYour choice: ")

	//Taking input about what to build
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	template := input.Text()

	switch template {
	case "1":
		fmt.Println("I will build you a template for ", menuItems[0], "\n")
	case "2":
		fmt.Println("I will build you a template for ", menuItems[1], "\n")
	}

	//Asking user to save the code or save the code and open in their favo editor
	fmt.Println("Do you want to")
	fmt.Println("1) Simply save the file")
	fmt.Println("2) Open it in my favorite code editor")
	input.Scan()
	opt := input.Text()

	switch opt {
	case "1":
		switch template {

		case "1":
			//exec.Command takes consecutive commands and executes
			//basic_func.txt contains the required function which is saved as template.go
			err := exec.Command("sh", "-c", "cat basic_func.txt >template.go").Run()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Saved your file as template.go")
			}

		case "2":
			err := exec.Command("sh", "-c", "cat basic_server.txt >template.go").Run()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Saved your file as template.go")
			}

		default:
			fmt.Println("Work in Progress...")

		}

	case "2":
		switch template {

		case "1":
			fmt.Println("Please enter the command to open your favourite editor ")
			input.Scan()
			editor := input.Text()
			err := exec.Command("sh", "-c", "cat basic_func.txt >template.go").Run()
			if err != nil {
				fmt.Println(err)
			}

			err = exec.Command(editor, "template.go").Run()
			if err != nil {
				fmt.Println(err)
			}

		case "2":
			fmt.Println("Please enter the command to open your favourite editor ")
			input.Scan()
			editor := input.Text()
			err := exec.Command("sh", "-c", "cat basic_server.txt >template.go").Run()
			if err != nil {
				fmt.Println(err)
			}

			err = exec.Command(editor, "template.go").Run()
			if err != nil {
				fmt.Println(err)
			}

		default:
			fmt.Println("Work in Progress...")
		}

	}

}
