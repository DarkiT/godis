package cluster

import (
	"testing"

	"github.com/darkit/godis/lib/utils"
	"github.com/darkit/godis/redis/connection"
	"github.com/darkit/godis/redis/parser"
	"github.com/darkit/godis/redis/protocol/asserts"
)

func TestPublish(t *testing.T) {
	testNodeA := testCluster[0]
	channel := utils.RandString(5)
	msg := utils.RandString(5)
	conn := connection.NewFakeConn()
	Subscribe(testNodeA, conn, utils.ToCmdLine("SUBSCRIBE", channel))
	conn.Clean() // clean subscribe success
	Publish(testNodeA, conn, utils.ToCmdLine("PUBLISH", channel, msg))
	data := conn.Bytes()
	ret, err := parser.ParseOne(data)
	if err != nil {
		t.Error(err)
		return
	}
	asserts.AssertMultiBulkReply(t, ret, []string{
		"message",
		channel,
		msg,
	})

	// unsubscribe
	UnSubscribe(testNodeA, conn, utils.ToCmdLine("UNSUBSCRIBE", channel))
	conn.Clean()
	Publish(testNodeA, conn, utils.ToCmdLine("PUBLISH", channel, msg))
	data = conn.Bytes()
	if len(data) > 0 {
		t.Error("expect no msg")
	}

	// unsubscribe all
	Subscribe(testNodeA, conn, utils.ToCmdLine("SUBSCRIBE", channel))
	UnSubscribe(testNodeA, conn, utils.ToCmdLine("UNSUBSCRIBE"))
	conn.Clean()
	Publish(testNodeA, conn, utils.ToCmdLine("PUBLISH", channel, msg))
	data = conn.Bytes()
	if len(data) > 0 {
		t.Error("expect no msg")
	}
}
