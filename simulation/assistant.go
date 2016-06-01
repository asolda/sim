package simulation

import (
	"math/rand"

	"github.com/asolda/sim/framework"
)

func ParseParams(nAgents int, nEdges int, nStep int,
	expTime int, infTime int,
	sSize int, eMaxSpan int, prob float64) (framework.Graph, SimulationModel) {
	numAgents := nAgents
	numEdges := nEdges
	numStep := nStep

	exposedTime := expTime
	infectedTime := infTime

	seedSize := sSize

	edgeMaxLifeSpan := eMaxSpan

	p := prob

	g := framework.InitGraph()

	//pool := make([]framework.Agent, numAgents)

	for i := 0; i < numAgents; i++ {
		//pool = append(pool,framework.CreateAgent(i, framework.SAFE, exposedTime, infectedTime))

		g.AddAgent(framework.CreateAgent(i,
			framework.SAFE,
			exposedTime,
			infectedTime))
	}

	for i := 0; i < numEdges; i++ {
		fromID := rand.Int() % numAgents
		toID := rand.Int() % numAgents

		agentFrom := framework.CreateAgent(fromID, 0, 0, 0)
		agentTo := framework.CreateAgent(toID, 0, 0, 0)

		timeStamp := rand.Int() % (numStep - (exposedTime + infectedTime - 2))
		lifeSpan := rand.Int() % edgeMaxLifeSpan

		g.AddDoubleEdge(agentFrom, agentTo, timeStamp, lifeSpan)

	}

	var seed []framework.Agent
	//seed := make([]framework.Agent, 0)
	for i := 0; i < seedSize; i++ {
		seedID := rand.Int() % (numAgents - 1)
		seed = append(seed, framework.CreateAgent(seedID, 0, 0, 0))
	}
	model := InitSim(g, p, seed)

	return g, model
}

func ParseWithSeed(nAgents int, nEdges int, nStep int,
	expTime int, infTime int,
	seeds []framework.Agent, eMaxSpan int, prob float64) (framework.Graph, SimulationModel) {
	numAgents := nAgents
	numEdges := nEdges
	numStep := nStep

	exposedTime := expTime
	infectedTime := infTime

	edgeMaxLifeSpan := eMaxSpan

	p := prob

	g := framework.InitGraph()

	//pool := make([]framework.Agent, numAgents)

	for i := 0; i < numAgents; i++ {
		//pool = append(pool,framework.CreateAgent(i, framework.SAFE, exposedTime, infectedTime))

		g.AddAgent(framework.CreateAgent(i,
			framework.SAFE,
			exposedTime,
			infectedTime))
	}

	for i := 0; i < numEdges; i++ {
		fromID := rand.Int() % numAgents
		toID := rand.Int() % numAgents

		agentFrom := framework.CreateAgent(fromID, 0, 0, 0)
		agentTo := framework.CreateAgent(toID, 0, 0, 0)

		timeStamp := rand.Int() % (numStep - (exposedTime + infectedTime - 2))
		lifeSpan := rand.Int() % edgeMaxLifeSpan

		g.AddDoubleEdge(agentFrom, agentTo, timeStamp, lifeSpan)

	}

	model := InitSim(g, p, seeds)

	return g, model
}

func PerformSim(m SimulationModel, numStep int) {
	for step := 0; step < numStep; step++ {
		m.PerformStep(step)
	}
}
