package simulation

import "github.com/asolda/sim/framework"

//SimulationModel represents a generic simulation model
type SimulationModel struct {
	Graph       framework.Graph
	workingCopy framework.Graph
	p           float64
}

//InitSim creates a new SimulationModel
func InitSim(g framework.Graph,
	p float64,
	seed []framework.Agent) SimulationModel {
	//for _, a := range g.GetAgents() {
	for i := 0; i < len(g.Agents); i++ {
		a := &g.Agents[i]
		for _, s := range seed {
			if a.Compare(s) {
				a.UpdateStatus(framework.EXPOSED)
				//fmt.Println("Infecting node", a)
				break
			}
		}
	}

	wc := g

	return SimulationModel{g, wc, p}
}

//PerformStep performs a step (fuck you golint!!)
func (m SimulationModel) PerformStep(step int) {
	m.workingCopy = m.Graph

	//for _, agent := range m.workingCopy.GetAgents() {
	for i := 0; i < len(m.workingCopy.Agents); i++ {
		agent := &m.workingCopy.Agents[i]
		agent.PerformStep(step, m.p)
	}

	m.Graph = m.workingCopy
}
