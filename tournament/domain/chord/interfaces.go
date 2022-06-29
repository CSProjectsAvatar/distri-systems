package chord

// RingApi defines how nodes talk to the Chord ring.
type RingApi interface {
	// StartNode sets up node as a ring server. This is where transport implementation is initialized.
	StartNode(node *Node)

	FindSuccessor(entry *RemoteNode, id []byte) (*RemoteNode, error)
	GetSuccessor(entry *RemoteNode) (*RemoteNode, error)

	// Notify tells node about a possible predecessor update.
	Notify(node *RemoteNode, pred *RemoteNode) error

	GetPredecessor(node *RemoteNode) (*RemoteNode, error)

	// CheckNode pings node in order to check whether it's available.
	CheckNode(node *RemoteNode) error

	// StopNode sets down the ring server
	StopNode() error
	SendData(data []*Data, node *RemoteNode) error
	SetValue(node *RemoteNode, key []byte, value string) error
	GetValue(node *RemoteNode, key []byte) (string, error)
}

type DataInteract interface {
	// LowerEq returns the key-value pairs when key is <= the given value.
	LowerEq(upper []byte) []*Data
	Delete(data []*Data)

	// Save saves the given data. If some key exists, then no changes
	// are committed.
	Save(data []*Data) error

	// Get gets the value for the given key. If key doesn't exist,
	// then an error is returned.
	Get(key []byte) (string, error)
}
