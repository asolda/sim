package framework

//Edge is a generic edge inside the dynamic simulation graph
type Edge struct {
	From      *Agent
	To        *Agent
	timestamp int
	lifespan  int
}

//IsActive returns true if the edge is active at the given step
func (e Edge) IsActive(step int) bool {
	if step >= e.timestamp && step <= e.timestamp+e.lifespan {
		return true
	}
	return false
}
