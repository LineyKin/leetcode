package easy

// https://leetcode.com/problems/distribute-money-to-maximum-children/

func distMoney(money int, children int) int {

	// монет должно быть больше, чем детей
	if money < children {
		return -1
	}

	// каждому ребёнку по монете
	if money == children {
		return 0
	}

	// частный случай
	if money == 8 && children == 1 {
		return 1
	}

	// никто не должен получать 4 монеты
	if money == 4 && children == 1 {
		return -1
	}

	// все деньги должны быть распределены
	if children == 0 && money > 0 {
		return -1
	}

	return 1 + distMoney(money-8, children-1)
}
