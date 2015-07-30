/************************
** Mission: All Aboard **
*************************
**
** Logging into the agency network, you're quickly into the operation
** manifest code. Sure enough, you're listed with no equipment, a typical
** agency snafu. You're going to need full tactical gear for this mission.
** Let's go ahead and fix that.
**
**/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	totalAgents := 5
	agents := make([]Agent, 0)
	agents = append(agents, Agent{name: "Jason Marshal", equipment: "none"})
	agents = append(agents, Agent{name: "Jake Boswell", equipment: "full"})
	agents = append(agents, Agent{name: "Jane Carter", equipment: "full"})
	agents = append(agents, Agent{name: "Max Miller", equipment: "full"})
	agents = append(agents, Agent{name: "Ty Walker", equipment: "full"})
	gearTable := createGearTable(agents, totalAgents)
	fmt.Println(gearTable)
}

func createGearTable(agents []Agent, totalAgents int) string {
	if len(agents) != totalAgents {
		return "ERROR: Wrong number of agents"
	}
	gearTable := "Operation Go: Agent Manifest\n"
	gearTable += "----------------------------"
	for num, agent := range agents {
		gearTable += "\n" + strconv.Itoa(num+1) + ". " + agent.name
		gearTable += ", Gear: " + agent.equipment
	}
	return gearTable
}

type Agent struct {
	name      string
	equipment string
}
