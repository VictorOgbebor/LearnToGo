package main

import (
	"fmt"
)

type TeamMember struct {
	Name string
}

func (t *TeamMember) updateTeam(newName string) {
	t.Name = newName
}

func (t *TeamMember) PrintTeamName() {
	fmt.Println(t.Name)
}

func main() {
	var teamMember TeamMember
	teamMember.Name = "BlockchainBic"
	teamMember.updateTeam("BicBlockchainSolutions")
	teamMember.PrintTeamName()
}
