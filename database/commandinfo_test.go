package database

import (
	"testing"

	"github.com/darkit/godis/lib/utils"
	"github.com/darkit/godis/redis/connection"
	"github.com/darkit/godis/redis/protocol/asserts"
)

func TestCommandInfo(t *testing.T) {
	c := connection.NewFakeConn()
	ret := testServer.Exec(c, utils.ToCmdLine("command"))
	asserts.AssertNotError(t, ret)
	ret = testServer.Exec(c, utils.ToCmdLine("command", "info", "mset"))
	asserts.AssertNotError(t, ret)
	ret = testServer.Exec(c, utils.ToCmdLine("command", "getkeys", "mset", "a", "a", "b", "b"))
	asserts.AssertMultiBulkReply(t, ret, []string{"a", "b"})
	ret = testServer.Exec(c, utils.ToCmdLine("command", "foobar"))
	asserts.AssertErrReply(t, ret, "Unknown subcommand 'foobar'")
}
