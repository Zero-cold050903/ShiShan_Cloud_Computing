package api

import (
	"errors"

	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/config"
	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/types"
)

func (n types.Node) FORK() {
	//fork somenew resources for the node

}

const EXITED = false

func (n types.Node) EXIT() bool {
	return EXITED
}
func (n types.Node) INIT() *pool {
	//init the source pool
	handle = &types.Pool
	return handle
}
func (handle pool) Monitot() error {
	for _, node := range handle.SourcePool {
		if !node.Healthy {
			return errors.New()

		} else {
			return nil
		}
	}
}
func (handle pool) Scanner() *Node {
	for _, node := range handle.SourcePool {
		if node.Status == config.STATUS_CODE[RUNNING] || node.Status == config.STATUS_CODE[OCCUPIED] {
			continue
		} else {
			return &node
		}
	}
}
func (handle types.Pool) Adjust() {
	for _, node := range handle.SourcePool {
		if node.Status == config.STATUS_CODE["RUNNING"] {
			node.Status = config.STATUS_CODE["OCCUPIED"]
		} else {
			continue
		}
	}
}
func (handle pool) Require() {}
