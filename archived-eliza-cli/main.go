package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')

		//userInput = strings.Replace(userInput, "\n", "", -1)   //Alternative method

		if userInput == "quit\n" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}

}
