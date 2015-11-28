// A game which will generate a date and ask you to enter the day of the week.
// This is a training game to help improve the speed of calculating the day of
// the week using the doomsday method
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Ganners/doomsday"
)

const (
	minYear = 1900
	maxYear = 2100

	title = `
================================================================
    ______  _____  _____ ___  ___ _____ ______   ___ __   __
    |  _  \|  _  ||  _  ||  \/  |/  ___||  _  \ / _ \\ \ / /
    | | | || | | || | | || .  . |\ '--. | | | |/ /_\ \\ V /
    | | | || | | || | | || |\/| | '--. \| | | ||  _  | \ /
    | |/ / \ \_/ /\ \_/ /| |  | |/\__/ /| |/ / | | | | | |
    |___/   \___/  \___/ \_|  |_/\____/ |___/  \_| |_/ \_/

================================================================

Speed training for Conway's doomsday method to work out the day
of the week for a given day.

Answer keyboard mapping:

1. Sunday
2. Monday
3. Tuesday
4. Wednesday
5. Thursday
6. Friday
7. Saturday

Press CTRL+C to close :-)
`

	AnswerSelected  = "You've gone for %s\n"
	AskQuestion     = "\033[0;32m" + "Calculate the day for: -- %d %s, %d --: " + "\033[1;32m"
	ValidationError = "Please only input a number between 1 - 7\n"
	//IncorrectAnswer = "Sorry, that's incorrect. The answer is: %s\n"
	IncorrectAnswer = "\033[0;31m" + `
☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠

                  Wrong! The answer was %s

☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠ ☠

` + "\033[1;32m"

	Congratulations = "\033[0;33m" + `
                  *
      ★                        ★           *           ★
               *                                   ★
            Congratulations! That is correct!
    ★                                     *                 ★
 ★                              ★                       ★


` + "\033[1;32m"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	// Print the title
	fmt.Println(title)

	// Start the game
	for {
		startGame()
	}
}

func startGame() {

	randYear := genRandYear()
	randMonth := genRandMonth()
	randDay := genRandDay(randYear, randMonth)
	dayOfWeek := doomsday.DayOfWeek(randYear, randMonth, randDay)

	// Wait a moment
	time.Sleep(1)

	// Prompts input and reads in the number
	number := readNumber(randYear, randMonth, randDay)
	dayGuess := doomsday.Day(number - 1)

	fmt.Printf(AnswerSelected, dayGuess)

	if dayGuess == dayOfWeek {
		fmt.Printf(Congratulations)
	} else {
		fmt.Printf(IncorrectAnswer, dayOfWeek)
	}
}

func readNumber(year, month, day int) int {

	fmt.Printf(AskQuestion, day, doomsday.Month(month), year)

	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	if len(str) < 1 || len(str) > 1 {
		fmt.Printf(ValidationError)
		return readNumber(year, month, day)
	}

	firstCharacter := []rune(str)[0]

	if firstCharacter < 49 || firstCharacter > 55 {
		fmt.Printf(ValidationError)
		return readNumber(year, month, day)
	}

	return int(firstCharacter - 48)
}

func genRandYear() int {

	yearRange := maxYear - minYear
	year := rand.Intn(yearRange) + minYear
	return year
}

func genRandMonth() int {

	monthRange := 12
	return rand.Intn(monthRange) + 1
}

func genRandDay(year, month int) int {

	daysInMonth := []int{
		0, // Start keys at 1
		31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31}

	// Set february to 29
	if doomsday.IsLeapYear(year) {
		daysInMonth[2] += 1
	}

	dayRange := daysInMonth[month]
	return rand.Intn(dayRange) + 1
}
