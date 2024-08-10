package services

import (
	"strings"
	"unicode"
)

func ConvertToMorseCode(text string) string {

	fH := make(chan string)
	sH := make(chan string)
	fText := ""
	sText := ""

	txtUpper := strings.ToUpper(text)
	words := strings.Fields(txtUpper)
	midPoint := len(words)

	firstHalf := words[:midPoint]
	secondHalf := words[midPoint:]

	go morseCode(firstHalf, fH)
	go morseCode(secondHalf, sH)

	fText += <-fH
	sText += <-sH

	close(fH)
	close(sH)

	morseTxt := fText + " " + sText

	return morseTxt

}

var morseMap = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",
}

func morseCode(text []string, fh chan string) {

	var result strings.Builder

	for _, word := range text {
		for _, r := range word {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				result.WriteString(morseMap[string(r)] + " ")
			}
		}
	}

	fh <- result.String()
}

func MorseToLetter(morse string) string {
	letters := make(map[string]string)
	var result strings.Builder

	for letter, morse := range morseMap {
		letters[morse] = letter
	}

	morseC := strings.Fields(morse)

	for _, m := range morseC {
		result.WriteString(letters[m] + " ")
	}

	return result.String()
}
