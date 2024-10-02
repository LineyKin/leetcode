package medium

// https://leetcode.com/problems/integer-to-roman/description/

type orderSigns [3]string
type romanOrder [3]orderSigns

func intToRoman(num int) string {
	var ro romanOrder
	ro[0] = orderSigns{"I", "V", "X"}
	ro[1] = orderSigns{"X", "L", "C"}
	ro[2] = orderSigns{"C", "D", "M"}

	var numOrder int
	var romanNumber string
	for num > 0 {

		if numOrder == 3 {
			var str string
			for i := 0; i < num; i++ {
				str += "M"
			}

			return str + romanNumber
		}

		lastDigit := num % 10
		romanNumber = convert(lastDigit, ro[numOrder]) + romanNumber
		numOrder++
		num = (num - lastDigit) / 10
	}

	return romanNumber
}

func convert(n int, signs orderSigns) string {
	if n == 0 {
		return ""
	}

	if n < 4 {
		var str string
		for i := 0; i < n; i++ {
			str += signs[0]
		}

		return str
	}

	if n == 4 {
		return signs[0] + signs[1]
	}

	if n == 5 {
		return signs[1]
	}

	if n < 9 {
		str := signs[1]
		for i := 0; i < n-5; i++ {
			str += signs[0]
		}
		return str
	}

	return signs[0] + signs[2]
}
