package cluster

import (
	"github.com/hdt3213/godis/interface/redis"
)

func ping(cluster *Cluster, c redis.Connection, cmdLine CmdLine) redis.Reply {
	return cluster.db.Exec(c, cmdLine)
}

/*----- utils -------*/

func makeArgs(cmd string, args ...string) [][]byte {
	result := make([][]byte, len(args)+1)
	result[0] = []byte(cmd)
	for i, arg := range args {
		result[i+1] = []byte(arg)
	}
	return result
}

// return node -> writeKeys
func (cluster *Cluster) groupBy(keys []string) map[string][]string {
	result := make(map[string][]string)
	for _, key := range keys {
		peer := cluster.pickNodeAddrByKey(key)
		group, ok := result[peer]
		if !ok {
			group = make([]string, 0)
		}
		group = append(group, key)
		result[peer] = group
	}
	return result
}

// pickNode returns the node id hosting the given slot.
// If the slot is migrating, return the node which is importing the slot
func (cluster *Cluster) pickNode(slotID uint32) *Node {
	slot := cluster.topology.GetSlots()[int(slotID)]
	return cluster.topology.GetNode(slot.NodeID)
}

func (cluster *Cluster) pickNodeAddrByKey(key string) string {
	slotId := getSlot(key)
	return cluster.pickNode(slotId).Addr
}
