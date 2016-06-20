package main

import (
	"math/rand"
	"time"

	"github.com/asolda/sim/framework"
	"github.com/asolda/sim/metrics"
	"github.com/asolda/sim/simulation"
	"github.com/asolda/sim/utils"

	"github.com/fatih/color"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numAgents := 10000
	numEdges := 25000
	numStep := 5000

	exposedTime := 2
	infectedTime := 150

	seedSize := 5

	edgeMaxLifeSpan := 100

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
	//fmt.Println("Final deaths count:", rCounter)
	color.Yellow("Random seeds:\t %d", rCounter)

	//fmt.Println("Performing page rank...")

	ranks := metrics.PageRank(g)
	dRanks := metrics.DynamicPageRank(g, numStep)

	//fmt.Println("Page rank done... preparing new simulations")

	topRanks := utils.GetMax(ranks, seedSize)
	topDRanks := utils.GetMax(dRanks, seedSize)

	seed := make([]framework.Agent, seedSize)
	dSeed := make([]framework.Agent, seedSize)

	for i := 0; i < seedSize; i++ {
		seed[i] = framework.CreateAgent(topRanks[i], 0, 0, 0)
		dSeed[i] = framework.CreateAgent(topDRanks[i], 0, 0, 0)
	}

	_, model = simulation.ParseWithSeed(numAgents, numEdges, numStep, exposedTime, infectedTime, seed, edgeMaxLifeSpan, p)

	//fmt.Println("Running first simulation (rank)")

	simulation.PerformSim(model, numStep)

	rCounter = 0
	for _, agent := range model.Graph.GetAgents() {
		if agent.GetStatus() == framework.REMOVED {
			//fmt.Println(agent.GetID(), "-> REMOVED")
			rCounter++
		}
	}

	color.Green("Top pageRank:\t %d", rCounter)

	_, model = simulation.ParseWithSeed(numAgents, numEdges, numStep, exposedTime, infectedTime, dSeed, edgeMaxLifeSpan, p)

	//fmt.Println("Running second simulation (dRank)")

	simulation.PerformSim(model, numStep)

	rCounter = 0
	for _, agent := range model.Graph.GetAgents() {
		if agent.GetStatus() == framework.REMOVED {
			//fmt.Println(agent.GetID(), "-> REMOVED")
			rCounter++
		}
	}

	//fmt.Println("Final deaths count (max dRank):", rCounter)
	color.Red("Top dRank:\t %d", rCounter)

	//fmt.Println("Computing trueRank...")

	simulation.ClearSeeds(&model)

	trueRanks := make([]float64, numAgents)

	for i, agent := range model.Graph.Agents {
		trueRanks[i] = metrics.ComputeRank(agent, model, numStep)
	}

	//fmt.Println("trueRank done... preparing new simulation")

	topTrueRanks := utils.GetMax(trueRanks, seedSize)

	trueSeed := make([]framework.Agent, seedSize)

	for i := 0; i < seedSize; i++ {
		trueSeed[i] = framework.CreateAgent(topTrueRanks[i], 0, 0, 0)
	}

	_, model = simulation.ParseWithSeed(numAgents, numEdges, numStep, exposedTime, infectedTime, trueSeed, edgeMaxLifeSpan, p)

	//fmt.Println("Running simulation (trueRank)")

	simulation.PerformSim(model, numStep)

	rCounter = 0
	for _, agent := range model.Graph.GetAgents() {
		if agent.GetStatus() == framework.REMOVED {
			//fmt.Println(agent.GetID(), "-> REMOVED")
			rCounter++
		}
	}

	//fmt.Println("Final deaths count (max trueRank):", rCounter)
	color.Blue("Top bcRank:\t %d", rCounter)

	//fmt.Println("Computing deep page rank...")

	deepRanks := metrics.DeepPageRank(model, numStep, 0.85)

	deepSeeds := utils.GetMax(deepRanks, seedSize)

	deepSeed := make([]framework.Agent, seedSize)

	for i := 0; i < seedSize; i++ {
		deepSeed[i] = framework.CreateAgent(deepSeeds[i], 0, 0, 0)
	}

	_, model = simulation.ParseWithSeed(numAgents, numEdges, numStep, exposedTime, infectedTime, deepSeed, edgeMaxLifeSpan, p)

	//fmt.Println("Running simulation (deepRank)")

	simulation.PerformSim(model, numStep)

	rCounter = 0
	for _, agent := range model.Graph.GetAgents() {
		if agent.GetStatus() == framework.REMOVED {
			//fmt.Println(agent.GetID(), "-> REMOVED")
			rCounter++
		}
	}

	color.Yellow("Top deepRank:\t %d", rCounter)

	adj := metrics.ToMatrix(g)
	katzRanks := metrics.ComputeKatz(adj, 0.01, 0.8, 3000)
	katzSeeds := utils.GetMax(katzRanks, seedSize)

	katzSeed := make([]framework.Agent, seedSize)

	for i := 0; i < seedSize; i++ {
		katzSeed[i] = framework.CreateAgent(katzSeeds[i], 0, 0, 0)
	}

	_, model = simulation.ParseWithSeed(numAgents, numEdges, numStep, exposedTime, infectedTime, katzSeed, edgeMaxLifeSpan, p)

	//fmt.Println("Running simulation (deepRank)")

	simulation.PerformSim(model, numStep)

	rCounter = 0
	for _, agent := range model.Graph.GetAgents() {
		if agent.GetStatus() == framework.REMOVED {
			//fmt.Println(agent.GetID(), "-> REMOVED")
			rCounter++
		}
	}

	color.Green("Top katzRank:\t %d", rCounter)

}
