package easy

// https://leetcode.com/problems/roman-to-integer/description/

type romanSign struct {
	value          int
	possiblePrefix string
	moreThenOne    bool
}

func romanToInt(s string) int {
	romanTable := map[string]romanSign{
		"I": {value: 1, possiblePrefix: "", moreThenOne: true},
		"V": {value: 5, possiblePrefix: "I", moreThenOne: false},
		"X": {value: 10, possiblePrefix: "I", moreThenOne: true},
		"L": {value: 50, possiblePrefix: "X", moreThenOne: false},
		"C": {value: 100, possiblePrefix: "X", moreThenOne: true},
		"D": {value: 500, possiblePrefix: "C", moreThenOne: false},
		"M": {value: 1000, possiblePrefix: "C", moreThenOne: true},
	}
	sRune := []rune(s)
	sum := 0
	var sign romanSign

	for i := len(sRune) - 1; i >= 0; i-- {
		romanLetter := string(sRune[i])

		if sign.value == 0 {
			sign = romanTable[romanLetter]
			sum += sign.value
			continue
		}

		if romanLetter == sign.possiblePrefix {
			nextSign := romanTable[romanLetter]
			sum -= nextSign.value
			sign = nextSign
			continue
		}

		nextSign := romanTable[romanLetter]
		if nextSign.value == sign.value {
			if nextSign.moreThenOne {
				sum += nextSign.value
			}
			continue
		}

		sum += nextSign.value
		sign = nextSign
	}

	return sum
}
