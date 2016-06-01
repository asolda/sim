package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/asolda/sim/framework"
	"github.com/asolda/sim/metrics"
	"github.com/asolda/sim/simulation"
	"github.com/asolda/sim/utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numAgents := 500
	numEdges := 1000
	numStep := 500

	exposedTime := 2
	infectedTime := 40

	seedSize := 1

	edgeMaxLifeSpan := 30

	p := 1.0

	g, model := simulation.ParseParams(numAgents, numEdges, numStep, exposedTime, infectedTime, seedSize, edgeMaxLifeSpan, p)

	simulation.PerformSim(model, numStep)

	rCounter := 0
	for _, agent := range model.Graph.GetAgents() {
		if agent.GetStatus() == framework.REMOVED {
			//fmt.Println(agent.GetID(), "-> REMOVED")
			rCounter++
		}
	}

	//fmt.Println("Initial seed size:", seedSize)
	fmt.Println("Final deaths count:", rCounter)

	fmt.Println("Performing page rank...")

	ranks := metrics.PageRank(g)
	dRanks := metrics.DynamicPageRank(g, numStep)

	fmt.Println("Page rank done... preparing new simulations")

	topRanks := utils.GetMax(ranks, seedSize)
	topDRanks := utils.GetMax(dRanks, seedSize)

	seed := make([]framework.Agent, seedSize)
	dSeed := make([]framework.Agent, seedSize)

	for i := 0; i < seedSize; i++ {
		seed[i] = framework.CreateAgent(topRanks[i], 0, 0, 0)
		dSeed[i] = framework.CreateAgent(topDRanks[i], 0, 0, 0)
	}

	_, model = simulation.ParseWithSeed(numAgents, numEdges, numStep, exposedTime, infectedTime, seed, edgeMaxLifeSpan, p)

	fmt.Println("Running first simulation (rank)")

	simulation.PerformSim(model, numStep)

	rCounter = 0
	for _, agent := range model.Graph.GetAgents() {
		if agent.GetStatus() == framework.REMOVED {
			//fmt.Println(agent.GetID(), "-> REMOVED")
			rCounter++
		}
	}

	fmt.Println("Final deaths count (max rank):", rCounter)

	_, model = simulation.ParseWithSeed(numAgents, numEdges, numStep, exposedTime, infectedTime, dSeed, edgeMaxLifeSpan, p)

	fmt.Println("Running second simulation (dRank)")

	simulation.PerformSim(model, numStep)

	rCounter = 0
	for _, agent := range model.Graph.GetAgents() {
		if agent.GetStatus() == framework.REMOVED {
			//fmt.Println(agent.GetID(), "-> REMOVED")
			rCounter++
		}
	}

	fmt.Println("Final deaths count (max dRank):", rCounter)

}
