// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package node

import (
	"github.com/go-trellis/config"
)

// NewNodesFromConfig 同步配置文件
func NewNodesFromConfig(filepath string) (map[string]Manager, error) {
	cfg, err := config.NewConfigOptions(config.OptionFile(filepath))
	if err != nil {
		return nil, err
	}
	return NewNodes(cfg)
}

// NewNodes 增加Nodes节点
func NewNodes(cfg config.Config) (ms map[string]Manager, err error) {
	mapManager := make(map[string]Manager)

	valConfigs := cfg.GetValuesConfig("node")
	for _, key := range valConfigs.GetKeys() {
		m := New(Type(valConfigs.GetInt(key+".type")), key)
		nodesCfg := valConfigs.GetValuesConfig(key + ".nodes")

		for _, nKey := range nodesCfg.GetKeys() {
			item := &Node{
				ID:       nKey,
				Value:    nodesCfg.GetString(nKey + ".value"),
				Weight:   uint32(nodesCfg.GetInt(nKey + ".weight")),
				Metadata: nodesCfg.GetMap(nKey + ".metadata"),
			}
			m.Add(item)
		}
		mapManager[key] = m
	}
	return mapManager, nil
}
