package metrics

import "github.com/asolda/sim/simulation"

func DeepPageRank(m simulation.SimulationModel, steps int, s float64) []float64 {
	agents := m.Graph.Agents

	n := len(agents)

	done := false
	time := 0

	ranks := make([]float64, n)
	tmp := make([]float64, n)

	for i := range agents {
		ranks[i] = float64(1 / n)
	}

	for !done && time < steps {
		time++

		for i, agent := range agents {
			tmp[i] = (1 - s) / float64(n)
			for _, conn := range agent.GetActiveConnections(time) {
				tmp[conn.To.GetID()] += s * ranks[i] / float64(len(agent.GetActiveConnections(time)))
			}
		}
	}

	ranks = tmp

	return ranks
}
