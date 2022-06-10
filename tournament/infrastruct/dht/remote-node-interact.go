package dht

// Helper for RingApi.Notify().
func (node *RemoteNode) notify(pred *RemoteNode, ring RingApi) error {
	return ring.Notify(node, pred)
}

// Helper for RingApi.GetPredecessor().
func (node *RemoteNode) getPredecessor(ring RingApi) (*RemoteNode, error) {
	return ring.GetPredecessor(node)

}

// Helper for RingApi.FindSuccessor().
func (node *RemoteNode) findSuccessor(id []byte, ring RingApi) (*RemoteNode, error) {
	return ring.FindSuccessor(node, id)
}

// Helper for RingApi.GetSuccessor().
func (node *RemoteNode) getSuccessor(ring RingApi) (*RemoteNode, error) {
	return ring.GetSuccessor(node)
}
