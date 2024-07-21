package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

const (
	secondsInMinute = 60
	secondsInHour   = 60 * secondsInMinute
	secondsInDay    = 24 * secondsInHour
	secondsInWeek   = 7 * secondsInDay
	secondsInMonth  = 30 * secondsInDay  // Approximate month length
	secondsInYear   = 365 * secondsInDay // Non-leap year
)

type TimeDuration struct {
	Years   int
	Months  int
	Weeks   int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func convertSeconds(seconds int) TimeDuration {

	years := seconds / secondsInYear
	seconds %= secondsInYear

	months := seconds / secondsInMonth
	seconds %= secondsInMonth

	weeks := seconds / secondsInWeek
	seconds %= secondsInWeek

	days := seconds / secondsInDay
	seconds %= secondsInDay

	hours := seconds / secondsInHour
	seconds %= secondsInMinute

	minutes := seconds / secondsInMinute
	seconds %= secondsInMinute

	return TimeDuration{
		Years:   years,
		Months:  months,
		Weeks:   weeks,
		Days:    days,
		Hours:   hours,
		Seconds: seconds,
		Minutes: minutes,
	}
}

func getPassword() (int, int, int, int, int, int, int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter password: ")
	password, _ := reader.ReadString('\n')
	passLength := len(strings.TrimSpace(password))
	passNumsPattern := (`[0-9]+`)
	passLowerPattern := (`[a-z]+`)
	passUpperPattern := (`[A-Z]+`)
	passSpecialPattern := (`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?\]\*\$ ]+`)
	passNums, _ := regexp.Compile(passNumsPattern)
	passLower, _ := regexp.Compile(passLowerPattern)
	passUpper, _ := regexp.Compile(passUpperPattern)
	passSpecial, _ := regexp.Compile(passSpecialPattern)
	var timeToCrack float64
	var duration TimeDuration
	lowerCount := 26
	specialCount := 32
	numsCount := 10
	fmt.Println(passLength)
	hashesPerSecond := 1102200000
	if passNums.MatchString(password) && !passLower.MatchString(password) && !passUpper.MatchString(password) && !passSpecial.MatchString(password) {
		numPossibilities := math.Pow(float64(numsCount), float64(passLength))
		fmt.Println(numPossibilities)
		timeToCrack = float64(numPossibilities) / float64(hashesPerSecond)
		duration = convertSeconds(int(timeToCrack))
	}
	if !passNums.MatchString(password) && passLower.MatchString(password) && !passUpper.MatchString(password) && !passSpecial.MatchString(password) {
		lowerPossibilities := math.Pow(float64(lowerCount), float64(passLength))
		timeToCrack = float64(lowerPossibilities) / float64(hashesPerSecond)
		duration = convertSeconds(int(timeToCrack))

	}
	if !passNums.MatchString(password) && !passLower.MatchString(password) && passUpper.MatchString(password) && !passSpecial.MatchString(password) {
		lowerPossibilities := math.Pow(float64(lowerCount), float64(passLength))
		timeToCrack = float64(lowerPossibilities) / float64(hashesPerSecond)
		duration = convertSeconds(int(timeToCrack))

	}
	if !passNums.MatchString(password) && !passLower.MatchString(password) && !passUpper.MatchString(password) && passSpecial.MatchString(password) {
		lowerPossibilities := math.Pow(float64(specialCount), float64(passLength))
		timeToCrack = float64(lowerPossibilities) / float64(hashesPerSecond)
		duration = convertSeconds(int(timeToCrack))

	}
	if passNums.MatchString(password) && passLower.MatchString(password) && passUpper.MatchString(password) && passSpecial.MatchString(password) {
		lowerPossibilities := math.Pow((float64(lowerCount)*2)+float64(specialCount)+float64(numsCount), float64(passLength))
		fmt.Println(lowerPossibilities)
		timeToCrack = float64(lowerPossibilities) / float64(hashesPerSecond)
		duration = convertSeconds(int(timeToCrack))

	}
	return duration.Days, duration.Hours, duration.Months, duration.Weeks, duration.Years, duration.Seconds, duration.Minutes
}
func main() {
	days, hours, months, weeks, years, seconds, minutes := getPassword()
	fmt.Printf("Time to crack: %d seconds, %d minutes, %d hours, %d days, %d weeks, %d months, %d years to crack! \n", seconds, minutes, hours, days, weeks, months, years)

}
