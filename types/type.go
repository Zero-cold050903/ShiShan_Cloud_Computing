package types

import (
	"errors"

	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/config"
)

/*
This is the node structure I defined.
*/
type Node struct {
	NodeID     string
	Status     int
	MemSize    int
	CPUCore    int //fix how many cores will this node have.Default 1 core for every node
	CPU_Usage  int
	Memo_Usage int  //the used percentage of the node's memory
	Healthy    bool //Check whether the node is in healthy satatus
}

// The POOL_SIZE is a contant defined in the @config pkg.
// It specified how many nodes are in the pool
// You can use the @adj_Pool_size() method to adjust it.

type Pool struct {
	SourcePool [config.POOL_SIZE]Node
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

func (n Node) FORK() {
	//fork somenew resources for the node

}

//============================

const EXITED = false

func (n Node) EXIT() bool {
	return EXITED
}
func (handle Pool) INIT() *Pool {
	//init the source pool
	ptr := new(Pool)
	return ptr
}
func (handle Pool) Monitot() error {
	for _, node := range handle.SourcePool {
		if !node.Healthy {
			return errors.New("")

		} else {
			return nil
		}
	}
	return nil
}
func (handle Pool) Scanner() *Node {
	for _, node := range handle.SourcePool {
		if node.Status == config.STATUS_CODE["RUNNING"] || node.Status == config.STATUS_CODE["OCCUPIED"] {
			continue
		} else {
			return &node
		}
	}
}
func (handle Pool) Adjust() {
	for _, node := range handle.SourcePool {
		if node.Status == config.STATUS_CODE["RUNNING"] {
			node.Status = config.STATUS_CODE["OCCUPIED"]
		} else {
			continue
		}
	}
}
func (handle Pool) Require() {}
