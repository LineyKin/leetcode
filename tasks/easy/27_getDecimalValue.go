package easy

import (
	"fmt"
	"log"
	"strconv"
)

// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

func GetDecimalValue(head *ListNode) int {
	binString := ""
	for head != nil {
		binString += strconv.Itoa(head.Val)
		head = head.Next
	}

	fmt.Println(binString)

	dec, err := strconv.ParseInt(binString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(dec)
}
