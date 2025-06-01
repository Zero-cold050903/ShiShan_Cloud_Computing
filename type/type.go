package types

/*
This is the node structure I defined.
*/
type Node struct {
	NodeID  string
	Status  int
	MemSize int
}

// The POOL_SIZE is a contant defined in the @config pkg.
// It specified how many nodes are in the pool
// You can use the @adj_Pool_size() method to adjust it.
type Pool struct {
	SourcePool [POOL_SIZE]Node
}
type ResourceInvoke interface {
	FORK()
	EXIT()
	INIT()
}
type Schedular interface {
	Monitor()
	Scanner()
	Adjust()
	Require()
}
