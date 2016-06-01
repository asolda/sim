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

	numAgents := 1000
	numEdges := 3000
	numStep := 500

	exposedTime := 2
	infectedTime := 40

	seedSize := 1

	edgeMaxLifeSpan := 50

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

	fmt.Println("Computing trueRank...")

	simulation.ClearSeeds(&model)

	trueRanks := make([]float64, numAgents)

	for i, agent := range model.Graph.Agents {
		trueRanks[i] = metrics.ComputeRank(agent, model, numStep)
	}

	fmt.Println("trueRank done... preparing new simulation")

	topTrueRanks := utils.GetMax(trueRanks, seedSize)

	trueSeed := make([]framework.Agent, seedSize)

	for i := 0; i < seedSize; i++ {
		trueSeed[i] = framework.CreateAgent(topTrueRanks[i], 0, 0, 0)
	}

	_, model = simulation.ParseWithSeed(numAgents, numEdges, numStep, exposedTime, infectedTime, trueSeed, edgeMaxLifeSpan, p)

	fmt.Println("Running simulation (trueRank)")

	simulation.PerformSim(model, numStep)

	rCounter = 0
	for _, agent := range model.Graph.GetAgents() {
		if agent.GetStatus() == framework.REMOVED {
			//fmt.Println(agent.GetID(), "-> REMOVED")
			rCounter++
		}
	}

	fmt.Println("Final deaths count (max trueRank):", rCounter)

}
