package cluster

import (
	"testing"

	"github.com/darkit/godis/redis/connection"
	"github.com/darkit/godis/redis/protocol/asserts"
)

func TestDel(t *testing.T) {
	conn := connection.NewFakeConn()
	allowFastTransaction = false
	testNodeA := testCluster[0]
	testNodeA.Exec(conn, toArgs("SET", "a", "a"))
	ret := Del(testNodeA, conn, toArgs("DEL", "a", "b", "c"))
	asserts.AssertNotError(t, ret)
	ret = testNodeA.Exec(conn, toArgs("GET", "a"))
	asserts.AssertNullBulk(t, ret)
}
