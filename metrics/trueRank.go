package metrics

import (
	"github.com/asolda/sim/framework"
	"github.com/asolda/sim/simulation"
)

func ComputeRank(target framework.Agent, model simulation.SimulationModel, step int) float64 {
	simulation.ClearSeeds(&model)
	simulation.AddSeed(&model, target)
	simulation.PerformSim(model, step)

	rCounter := 0
	for _, agent := range model.Graph.GetAgents() {
		if agent.GetStatus() == framework.REMOVED {
			rCounter++
		}
	}

	return float64(rCounter / len(model.Graph.Agents))

}
