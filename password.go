package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"time"
	"math/rand"
	"unicode"
)

type Scope struct {
	min, max int
}

type Password struct {
	Length
	Letters
	Numbers
	Uppercase
	Symbols
	string string
}

type Length struct {
	Scope
	current int
}

type Letters struct {
	Scope
	amount int
}

type Numbers struct {
	Scope
	amount int
}

type Uppercase struct {
	Scope
	amount int
}

type Symbols struct {
	Scope
	amount int
}

type Runes struct {
	lowercaseLetters []rune
	numbers          []rune
	symbols          []rune
}

var p = Password{Length{Scope{6, 20}, 0}, Letters{}, Numbers{}, Uppercase{}, Symbols{}, ""}
var r = Runes{[]rune("abcdefghijklmnopqrstuvwxyz"), []rune("1234567890"), []rune("`~!@#$%^&*()-=_+[]{};':,./<>?")}

var reader = bufio.NewReader(os.Stdin)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	askLetters()

	askUppercaseLetters()

	askNumbers()

	askSymbols()

	buildPassword()

	fmt.Println("Your password: " + p.string)
}

func askLetters() {
	fmt.Println("Letters amount [" + strconv.Itoa(p.Length.min) + "-" + strconv.Itoa(p.Length.max) + "]: ")

	lettersAmount := parseInputToInt(reader.ReadString('\n'))

	if int(lettersAmount) >= p.Length.min && int(lettersAmount) <= p.Length.max {
		p.Letters.amount = int(lettersAmount)
		updatePasswordLength()
	} else {
		badInput()
	}
}

func askUppercaseLetters() {
	p.Uppercase.min = 0

	fmt.Println("Uppercase letters amount [" + strconv.Itoa(p.Uppercase.min) + "-" + strconv.Itoa(p.Letters.amount) + "]: ")

	uppercaseLettersAmount := parseInputToInt(reader.ReadString('\n'))

	if int(uppercaseLettersAmount) >= p.Uppercase.min && int(uppercaseLettersAmount) <= p.Letters.amount {
		p.Uppercase.amount = int(uppercaseLettersAmount)
	} else {
		badInput()
	}
}

func askNumbers() {
	if p.Length.current < p.Length.max {
		p.Numbers.min = 0
		p.Numbers.max = p.Length.max - p.Length.current

		fmt.Println("Numbers amount [" + strconv.Itoa(p.Numbers.min) + "-" + strconv.Itoa(p.Numbers.max) + "]: ")

		numbersAmount := parseInputToInt(reader.ReadString('\n'))

		if int(numbersAmount) >= 0 && int(numbersAmount) <= p.Length.max-p.Length.current {
			p.Numbers.amount = int(numbersAmount)
			updatePasswordLength()
		} else {
			badInput()
		}
	}
}

func askSymbols() {
	if p.Length.current < p.Length.max {
		p.Symbols.min = 0
		p.Symbols.max = p.Length.max - p.Length.current

		fmt.Println("Symbols amount [" + strconv.Itoa(p.Symbols.min) + "-" + strconv.Itoa(p.Symbols.max) + "]: ")

		symbolsAmount := parseInputToInt(reader.ReadString('\n'))

		if int(symbolsAmount) >= 0 && int(symbolsAmount) <= p.Length.max-p.Length.current {
			p.Symbols.amount = int(symbolsAmount)
			updatePasswordLength()
		} else {
			fmt.Println("Bad input, exiting...")
			os.Exit(1)
		}
	}
}

func parseInputToInt(input string, err error) int {
	if err == nil {
		replacer := strings.NewReplacer("\n", "", "\r", "")
		tempInputString := replacer.Replace(input)
		tempInt64, _ := strconv.ParseInt(tempInputString, 10, 0)
		return int(tempInt64)
	}
	return 0
}

func updatePasswordLength() {
	p.Length.current = p.Letters.amount + p.Numbers.amount + p.Symbols.amount
}

func badInput() {
	fmt.Println("Bad input, exiting...")
	time.Sleep(1000 * time.Millisecond)
	os.Exit(1)
}

func buildPassword() {
	letters := make([]rune, p.Letters.amount)
	for i := range letters {
		if i < p.Uppercase.amount {
			letters[i] = unicode.ToUpper(r.lowercaseLetters[rand.Intn(len(r.lowercaseLetters))])
		} else {
			letters[i] = r.lowercaseLetters[rand.Intn(len(r.lowercaseLetters))]
		}
	}
	p.string += string(letters)

	if p.Numbers.amount != 0 {
		numbers := make([]rune, p.Numbers.amount)
		for i := range numbers {
			numbers[i] = r.numbers[rand.Intn(len(r.numbers))]
		}
		p.string += string(numbers)
	}

	if p.Symbols.amount != 0 {
		symbols := make([]rune, p.Symbols.amount)
		for i := range symbols {
			symbols[i] = r.symbols[rand.Intn(len(r.symbols))]
		}
		p.string += string(symbols)
	}

	array := strings.Split(p.string, "")
	rand.Shuffle(len(array), func(i, j int) {
		array[i], array[j] = array[j], array[i]
	})

	p.string = strings.Join(array, "")
}
