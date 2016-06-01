package framework

//Graph is a dynamic graph
type Graph struct {
	Agents []Agent
}

//InitGraph initializes a new graph
func InitGraph() Graph {
	var g Graph
	g.Agents = make([]Agent, 0)
	return g
}

//AddAgent adds a new member to the simulation
func (g *Graph) AddAgent(a Agent) {
	g.Agents = append(g.Agents, a)
}

//AddDoubleEdge adds a bidirectional edge
func (g *Graph) AddDoubleEdge(first Agent, second Agent, ts int, ls int) {
	g.AddEdge(first, second, ts, ls)
	g.AddEdge(second, first, ts, ls)
}

//AddEdge adds a new edge with specified fields
func (g *Graph) AddEdge(from Agent, to Agent, ts int, ls int) {
	var memberFrom *Agent
	var memberTo *Agent
	//for _, member := range g.GetAgents() {
	for i := 0; i < len(g.Agents); i++ {
		member := &g.Agents[i]
		if member.Compare(from) {
			memberFrom = member
		}
		if member.Compare(to) {
			memberTo = member
		}
	}

	//memberFrom.AddConnection(memberTo, ts, ls)

	memberFrom.connections = append(memberFrom.connections, Edge{memberFrom, memberTo, ts, ls})
}

//GetAgents returns all agents
func (g Graph) GetAgents() []Agent {
	return g.Agents
}
