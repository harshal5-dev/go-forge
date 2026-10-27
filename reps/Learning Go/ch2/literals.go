package main

import "fmt"

func literals() {
	var base10Val int = 10
	var base2Val int = 0b1010
	var base8Val int = 0o4574
	var base16Val int = 0xA2323
	var intZeroVal int

	var longerIntVal1, longerIntVal2 int = 1_234, 1_2_3_4

	var base10FloatVal float64 = 6.03e55
	var base16FloatVal float64 = 0x55.44p5
	var floatZeroVal float64

	var longerFloatVal1, longerFloatVal2 float64 = 1_234.4e55, 1_2_3_4.e5

	var unicodeChar rune = 'a'
	var newline rune = '\n'
	var tab rune = '\t'
	var quote rune = '\''
	var backslash rune = '\\'
	var runeZeroVal rune

	var msg1 string = "Greetings and\n \"Salutations\" \\"
	var msg2 string = `it's nice
	day for "work" \`
	var stringZeroVal string

	var flag bool
	var isAwesome = true

	fmt.Println("--------- Literals ----------")
	fmt.Println("Integer Literal Decimal value:", base10Val)
	fmt.Println("Integer Literal Binary value:", base2Val)
	fmt.Println("Integer Literal Octal value:", base8Val)
	fmt.Println("Integer Literal Hexadecimal value:", base16Val)
	fmt.Println("Integer Literal Zero value:", intZeroVal)
	fmt.Println("Integer Literal longer integer literals:", longerIntVal1, longerIntVal2)

	fmt.Println("-----------------------------")

	fmt.Println("Floating Literal Decimal value:", base10FloatVal)
	fmt.Println("Floating Literal Hexadecimal value:", base16FloatVal)
	fmt.Println("Floating Literal Zero value:", floatZeroVal)
	fmt.Println("Floating Literal longer integer literals:", longerFloatVal1, longerFloatVal2)

	fmt.Println("-----------------------------")

	fmt.Println("Rune Literal unicode characters:", unicodeChar)
	fmt.Println("Rune Literal newline:", newline)
	fmt.Println("Rune Literal tab:", tab)
	fmt.Println("Rune Literal quote:", quote)
	fmt.Println("Rune Literal backslash:", backslash)
	fmt.Println("Rune Literal Zero value:", runeZeroVal)

	fmt.Println("-----------------------------")

	fmt.Println("String Literal msg1:", msg1)
	fmt.Println("String Literal msg2:", msg2)
	fmt.Println("String Literal Zero value:", stringZeroVal)

	fmt.Println("-----------------------------")
	fmt.Println("Boolean Literal zero value:", flag)
	fmt.Println("Boolean Literal isAwesome::", isAwesome)

	fmt.Println()
}
