package metrics

import (
	"github.com/asolda/sim/framework"
	"github.com/gonum/matrix/mat64"
)

import "math"

func ToMatrix(g framework.Graph) *mat64.Dense {
	n := len(g.GetAgents())
	data := make([]float64, n*n)

	for _, agent := range g.GetAgents() {
		for _, edge := range agent.GetConnections() {
			data[edge.From.GetID()*edge.To.GetID()] = 1
		}
	}

	adj := mat64.NewDense(n, n, data)
	return adj
}

func ComputeKatz(d *mat64.Dense, confidence float64, alpha float64, maxStep int) []float64 {
	step := 0
	d.Inverse(d)

	n, _ := d.Dims()

	data := make([]float64, n*n)

	for i := 0; i < n*n; i += n {
		data[i] = 1
	}

	data = make([]float64, n)
	x1 := mat64.NewVector(n, data)
	for i := 0; i < n; i++ {
		data[i] = 1
	}
	x := mat64.NewVector(n, data)
	done := false
	for !done {
		step++
		if step == maxStep {
			done = true
		}
		data = make([]float64, n*n)

		for i := 0; i < n; i++ {
			data[i*i] = 1
		}
		identity := mat64.NewDense(n, n, data)

		d.Scale(alpha, d)
		identity.Sub(identity, d)
		identity.Inverse(identity)
		data = make([]float64, n)
		for i := 0; i < n; i++ {
			data[i] = 1
		}
		idenvector := mat64.NewDense(n, 1, data)
		for i := 0; i < n; i++ {
			data[i] = 0
		}
		res := mat64.NewDense(n, 1, data)
		res.Mul(identity, idenvector)
		j := 0
		for i := 0; i < n; i++ {
			x1.SetVec(j, res.At(i, 0))
			j++
		}

		for i := 0; i < n; i++ {
			if math.Abs(x.At(i, 0)-x1.At(i, 0)) > confidence {
				done = false
				break
			} else {
				done = true
			}
		}
		x = x1
	}

	toReturn := make([]float64, n)
	for i := 0; i < n; i++ {
		toReturn[i] = x1.At(i, 0)
	}

	return toReturn
}
