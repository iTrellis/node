// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package node

import (
	"fmt"

	"github.com/gogap/config"
)

// NewNodesFromConfig 同步配置文件
func NewNodesFromConfig(filepath string) (map[string]Manager, error) {
	return NewNodes(config.NewConfig(config.ConfigFile(filepath)))
}

// NewNodes 增加Nodes节点
func NewNodes(cfg config.Configuration) (map[string]Manager, error) {
	mapManager := make(map[string]Manager)
	for _, key := range cfg.GetConfig("node").Keys() {
		itemCfg := cfg.GetConfig("node." + key)
		fmt.Println(Type(itemCfg.GetInt32("type")), key)
		m := New(Type(itemCfg.GetInt32("type")), key)

		nodesCfg := itemCfg.GetConfig("nodes")
		for _, nodeID := range nodesCfg.Keys() {
			item := &Node{
				ID:     nodeID,
				Weight: uint32(nodesCfg.GetInt32(nodeID + ".weight")),
				Value:  nodesCfg.GetString(nodeID + ".value"),
			}
			fmt.Println(key, item)
			m.Add(&Node{
				ID:     nodeID,
				Weight: uint32(nodesCfg.GetInt32(nodeID + ".weight")),
				Value:  nodesCfg.GetString(nodeID + ".value"),
			})
		}
		mapManager[key] = m
	}
	return mapManager, nil
}
