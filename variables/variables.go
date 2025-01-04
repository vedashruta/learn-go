package variables

import (
	"bufio"
	"fmt"
	"os"
)

func Variables() {
	var varName1 int32
	varName1 = 10
	var varName2 int32 = 20
	var varName3 = 30
	varName4 := 40
	fmt.Println("variable 1", varName1)
	fmt.Println("variable 2", varName2)
	fmt.Println("variable 3", varName3)
	fmt.Println("variable 4", varName4)

	type WeekDay float32
	const (
		Sunday WeekDay = iota
		//some random comment ,iota skips this line & blank line as well

		Monday
		Tuesday
	)
	fmt.Println("Sunday", Sunday)
	fmt.Println("Monday", Monday)
	fmt.Println("Tuesday", Tuesday)

	//Maps
	m := make(map[string]any, 4)
	m["a"] = 10
	m["b"] = "hello"
	fmt.Println(m)
	m["a"] = "world"
	fmt.Println(m)

	//Reading input from console

	//1. Using Scan,Scanln
	var name string
	var age int
	var phone string
	fmt.Print("Enter your name :")
	fmt.Scan(&name)
	fmt.Print("Enter your age :")
	fmt.Scan(&age)
	fmt.Print("Enter your phone number :")
	fmt.Scanln(&phone)
	fmt.Println("Name : ", name, "\nAge : ", age, "\nPhone Number : ", phone)

	//2. Using Scanf
	var model string
	var company string
	fmt.Print("Enter your car model and company : ")
	fmt.Scanf("%s %s", &model, &company)
	fmt.Println("Model :", model)
	fmt.Println("Company :", company)

	//Using bufio & os.Stdin
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Reads the entire line
	text := scanner.Text()
	fmt.Println("Text Read from stdin :", text)

	//Panic, concurrent map iteration and write
	newMap := make(map[string]int)
	go func() {
		for {
			newMap["blog"] = 1
		}
	}()
	go func() {
		for {
			newMap["blog"] = 2
		}
	}()
	select {}
}
