package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`strings can span multiple lines with the \n escape sequence`)
	fmt.Println(`peace be upon you upon you be peace`)
	fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)
	//9.2
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)
	message := "shalom"
	c := message[5]
	fmt.Printf("%c\n", c)

	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star)
	smile := ''
	fmt.Printf("%c %[1]v\n", smile)
	acute := 'é'
	fmt.Printf("%c %[1]v\n", acute)

	cb := 'a'
	cb = cb + 3
	fmt.Printf("%c\n", cb)

	message_2 := "uv vagreangvbany fcnpr fgngvba"
	fmt.Println(len(message_2))

	message_3 := "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(message_3); i++ {
		c_h := message_3[i]
		if c_h >= 'a' && c_h <= 'z' {
			c_h = c_h + 13
			if c_h > 'z' {
				c_h = c_h - 26
			}
		}
		fmt.Printf("%c", c_h)
	}

	question := "¿Cómo estás?"
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")
	c_f, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c_f, size)
}
