package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func getPassword() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter password: ")
	password, _ := reader.ReadString('\n')
	passLength := len(strings.TrimSpace(password))
	passNumsPattern := (`[0-9]+`)
	passLowerPattern := (`[a-z]+`)
	passUpperPattern := (`[A-Z]+`)
	passSpecialPattern := (`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?\]\*\$]+`)
	passNums, _ := regexp.Compile(passNumsPattern)
	passLower, _ := regexp.Compile(passLowerPattern)
	passUpper, _ := regexp.Compile(passUpperPattern)
	passSpecial, _ := regexp.Compile(passSpecialPattern)
	var crackTimeString string
	var timeToCrack float64
	lowerCount := 26
	// upperCount := 26
	// specialCount := 40
	numsCount := 10
	fmt.Println(passLength)
	hashesPerSecond := 1102200000
	if passNums.MatchString(password) && !passLower.MatchString(password) && !passUpper.MatchString(password) && !passSpecial.MatchString(password) {
		numPossibilities := math.Pow(float64(numsCount), float64(passLength))
		fmt.Println(numPossibilities)
		timeToCrack = float64(numPossibilities) / float64(hashesPerSecond)

	}
	if !passNums.MatchString(password) && passLower.MatchString(password) && !passUpper.MatchString(password) && !passSpecial.MatchString(password) {
		lowerPossibilities := math.Pow(float64(lowerCount), float64(passLength))
		timeToCrack = float64(lowerPossibilities) / float64(hashesPerSecond)
	}

	if float64(timeToCrack) < 1 {
		crackTimeString = "Instantly"
	}
	if float64(timeToCrack) >= 1000 {
		crackTimeString = fmt.Sprintf("%.2f", timeToCrack) + " seconds"
	}
	if float64(timeToCrack) >= 60 {
		crackTimeString = fmt.Sprintf("%.2f", timeToCrack/60) + " minutes"
	}
	if float64(timeToCrack) >= 3600 {
		crackTimeString = fmt.Sprintf("%.2f", timeToCrack/3600) + " hours"
	}
	if float64(timeToCrack) >= 86400 {
		crackTimeString = fmt.Sprintf("%.2f", timeToCrack/3600/24) + " days"
	}
	if float64(timeToCrack) >= 86400 {
		crackTimeString = fmt.Sprintf("%.2f", (timeToCrack/3600/24)/7) + " weeks"
	}
	return crackTimeString
}
func main() {
	fmt.Printf("Time to crack: %s! \n", getPassword())

}
