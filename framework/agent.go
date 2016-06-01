package framework

import "math/rand"

//State describes the state of an agent
type State int

//Constants represents possible status
const (
	SAFE = iota
	EXPOSED
	INFECTED
	REMOVED
)

//Agent is a member of the simulation
type Agent struct {
	id            int
	status        State
	connections   []Edge
	exposedTime   int
	infectedTime  int
	statusCounter int
}

//UpdateStatus does extacltly what the name tells
func (a *Agent) UpdateStatus(new State) {
	//fmt.Println("calling update status on:", a, "new:", new)
	if a.GetStatus() == SAFE && new == EXPOSED {
		a.status = EXPOSED
		a.statusCounter = 0
		return
	}
	if a.GetStatus() == EXPOSED && new == INFECTED {
		a.status = INFECTED
		a.statusCounter = 0
		return
	}
	if a.GetStatus() == INFECTED && new == REMOVED {
		a.status = REMOVED
		return
	}
	return
}

//GetID returns agent's id
func (a Agent) GetID() int {
	return a.id
}

//GetStatus returns the current status of the agent
func (a Agent) GetStatus() State {
	return a.status
}

//GetActiveConnections returns all edges active at a given step
func (a Agent) GetActiveConnections(step int) []Edge {
	var toReturn []Edge

	for _, e := range a.connections {
		if e.IsActive(step) {
			toReturn = append(toReturn, e)
		}
	}

	return toReturn
}

//PerformStep does everithing required to an agent
func (a *Agent) PerformStep(step int, p float64) {
	if a.status == REMOVED || a.status == SAFE {
		return
	} else if a.GetStatus() == EXPOSED {
		if a.statusCounter == a.exposedTime {
			a.UpdateStatus(INFECTED)
			a.statusCounter = 0
		} else {
			a.statusCounter++
		}
	} else if a.GetStatus() == INFECTED {
		if a.statusCounter == a.infectedTime {
			a.UpdateStatus(REMOVED)
			return
		}
		a.statusCounter++
		for _, target := range a.GetActiveConnections(step) {
			if rand.Float64() <= p {
				target.To.UpdateStatus(EXPOSED)
				//fmt.Println("Sto provando a infettare", target.to.GetID())
			}
		}
	}
}

//GetExposedTime does extactly what the name tells
func (a Agent) GetExposedTime() int {
	return a.exposedTime
}

//GetInfectedTime does extactly what the name tells
func (a Agent) GetInfectedTime() int {
	return a.infectedTime
}

//GetCurrentCounter does extactly what the name tells
func (a Agent) GetCurrentCounter() int {
	return a.statusCounter
}

//Compare two different agents to check they are the same
func (a Agent) Compare(other Agent) bool {
	return (a.id == other.id)
}

//GetConnections returns all connections
func (a Agent) GetConnections() []Edge {
	return a.connections
}

//GetOutDegree returns the out degree of the agent
func (a Agent) GetOutDegree() int {
	return len(a.connections)
}

//AddConnection does stuff
func (a Agent) AddConnection(to *Agent, ts, ls int) {
	a.connections = append(a.connections, Edge{&a, to, ts, ls})
}

//CreateAgent creates a new Agent without connections
func CreateAgent(id int, status State, exposed int, infected int) Agent {
	return Agent{id, status, make([]Edge, 0), exposed, infected, 0}
}
