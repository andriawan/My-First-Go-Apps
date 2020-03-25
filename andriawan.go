package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	generateBanner()
	name := getUserInput(reader)
	age, _ := getUserAge(reader)
	var matureStatement string
	if isMature(age) {
		matureStatement = "You are mature, find a job"
	} else {
		matureStatement = "Playing video game is more fun"
	}
	printResult(name, age, matureStatement)
}

func isMature(age int64) bool {
	if age > 25 {
		return true
	}
	return false
}

func printResult(name string, age int64, matrue string) {
	fmt.Println("")
	fmt.Println("------------------------------------------")
	fmt.Println("RESULT")
	fmt.Println("------------------------------------------")
	fmt.Println("Hello", name)
	fmt.Println("You are", age, "Year old now")
	fmt.Println(matrue)
	fmt.Println("------------------------------------------")
	fmt.Println("")
}

func generateBanner() {
	fmt.Println("------------------------------------------")
	fmt.Println("Welcome to Andriawan First Go Simple Apps")
	fmt.Println("------------------------------------------")
}

func getUserAge(reader *bufio.Reader) (int64, error) {
	fmt.Print("Please enter year of birth : ")
	year, _ := reader.ReadString('\n')
	year = strings.Replace(year, "\n", "", -1)
	currentYear := time.Now().Year()
	intYear, err := strconv.ParseInt(year, 10, 64)
	currentYearInt64 := int64(currentYear)
	if err == nil {
		return currentYearInt64 - intYear, nil
	}

	return 0, errors.New("Please provide number")

}

func getUserInput(reader *bufio.Reader) string {
	fmt.Print("Please enter your name : ")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	return name
}
