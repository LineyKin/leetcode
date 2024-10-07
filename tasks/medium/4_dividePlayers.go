package medium

import (
	"sort"
)

// https://leetcode.com/problems/divide-players-into-teams-of-equal-skill/?envType=daily-question&envId=2024-10-04

func dividePlayers(skill []int) int64 {
	totalPlayers := len(skill)
	totalTeams := totalPlayers / 2
	sort.Ints(skill)
	miTeamSkill := skill[0] + skill[1]
	maxTeamSkill := skill[totalPlayers-1] + skill[totalPlayers-2]

	for teamSkill := miTeamSkill; teamSkill <= maxTeamSkill; teamSkill++ {
		skillCopy := make([]int, len(skill))
		copy(skillCopy, skill)
		teamsCount := 0
		j := 0
		k := totalPlayers - 1
		var chemistry int64
		for j < k {
			if skillCopy[j]+skillCopy[k] == teamSkill {

				if skillCopy[j] == 0 || skillCopy[k] == 0 {
					continue
				}

				chemistry += int64(skillCopy[j] * skillCopy[k])
				teamsCount++
				skillCopy = deletePlayers(skillCopy, j, k)

			}

			if skillCopy[j] < skillCopy[k] {
				j++
			} else {
				k--
			}

			if teamsCount == totalTeams {
				return chemistry
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
