package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"math/rand"
	"unicode"
	"bufio"
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
var r = Runes{[]rune("abcdefghijklmnopqrstuvwxyz"), []rune("1234567890"), []rune("!#$%&'()*+,-./:;<=>?@[]^_`{|}~")}

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	askLetters()

	askUppercaseLetters()

	askNumbers()

	askSymbols()

	buildPassword()

	returnPassword()
}

func askLetters() {
	fmt.Println("Letters amount [" + strconv.Itoa(p.Length.min) + "-" + strconv.Itoa(p.Length.max) + "]: ")

	lettersAmount := parseIntFromStdIn()

	if lettersAmount >= p.Length.min && lettersAmount <= p.Length.max {
		p.Letters.amount = int(lettersAmount)
		updatePasswordLength()
	} else {
		badInput()
	}
}

func askUppercaseLetters() {
	p.Uppercase.min = 0

	fmt.Println("Uppercase letters amount [" + strconv.Itoa(p.Uppercase.min) + "-" + strconv.Itoa(p.Letters.amount) + "]: ")

	uppercaseLettersAmount := parseIntFromStdIn()

	if uppercaseLettersAmount >= p.Uppercase.min && uppercaseLettersAmount <= p.Letters.amount {
		p.Uppercase.amount = uppercaseLettersAmount
	} else {
		badInput()
	}
}

func askNumbers() {
	if p.Length.current < p.Length.max {
		p.Numbers.min = 0
		p.Numbers.max = p.Length.max - p.Length.current

		fmt.Println("Numbers amount [" + strconv.Itoa(p.Numbers.min) + "-" + strconv.Itoa(p.Numbers.max) + "]: ")

		numbersAmount := parseIntFromStdIn()

		if numbersAmount >= 0 && numbersAmount <= p.Length.max-p.Length.current {
			p.Numbers.amount = numbersAmount
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

		symbolsAmount := parseIntFromStdIn()

		if symbolsAmount >= 0 && symbolsAmount <= p.Length.max-p.Length.current {
			p.Symbols.amount = symbolsAmount
			updatePasswordLength()
		} else {
			fmt.Println("Bad input, exiting...")
			os.Exit(1)
		}
	}
}

func parseIntFromStdIn() int {
	var i int
	fmt.Scan()
	_, err := fmt.Scanf("%d\n", &i)
	if err != nil {
		badInput()
	}
	return i
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

func returnPassword() {
	fmt.Println("Your password: " + p.string)

	fmt.Println("Press (Enter) to exit...")
	fmt.Scanln()

	fmt.Println("Success, exiting...")
	time.Sleep(1000 * time.Millisecond)
	os.Exit(1)
}
