package api

import (
	"fmt"
	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/config"
	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/types"
)
func (handle pool) Monitor() (int ,int ,int){
	for _ , node := range handle.SourcePool {
		if node.Status 
	}
}
func (handle pool) Scanner() &Node {
	for _ , node := range handle.SourcePool {
		if node.Status = config.STATUS_CODE[RUNNING]
		|| node.Status = config.STATUS_CODE[OCCUPIED]{
			continue
		} else {
			return &node
		}
	}
}
func (handle pool) Adjust() {
	for _, node := range handle.SourcePool {
		if node.Status = config.STATUS_CODE[RUNNING] {
			node.Status = config.STATUS_CODE[OCCUPIED]
		} else {
			continue
		}
	}
}
func (handle pool)Require() {}
