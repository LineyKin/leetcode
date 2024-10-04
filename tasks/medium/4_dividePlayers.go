package medium

import (
	"sort"
)

// https://leetcode.com/problems/divide-players-into-teams-of-equal-skill/?envType=daily-question&envId=2024-10-04

type teamList []int64

func DividePlayers(skill []int) int64 {
	totalPlayers := len(skill)
	totalTeams := totalPlayers / 2
	sort.Ints(skill)
	miTeamSkill := skill[0] + skill[1]                            // 3
	maxTeamSkill := skill[totalPlayers-1] + skill[totalPlayers-2] // 9

	for teamSkill := miTeamSkill; teamSkill <= maxTeamSkill; teamSkill++ {

		skillCopy := make([]int, len(skill))
		copy(skillCopy, skill)
		teamList := make(teamList, 0)

		for j := 0; j < totalPlayers; j++ {

			for k := j + 1; k < totalPlayers; k++ {

				if skillCopy[j]+skillCopy[k] == teamSkill {

					if skillCopy[j] == 0 || skillCopy[k] == 0 {
						continue
					}

					teamList = append(teamList, int64(skillCopy[j]*skillCopy[k]))
					skillCopy = deletePlayers(skillCopy, j, k)

				}

				if len(teamList) == totalTeams {
					return CalcChemistry(teamList)
				}

			}
		}
	}

	return -1
}

func deletePlayers(skill []int, j, k int) []int {
	skill[j] = 0
	skill[k] = 0
	return skill
}

func CalcChemistry(list teamList) int64 {
	var sum int64
	for _, v := range list {
		sum += v
	}

	return sum
}
