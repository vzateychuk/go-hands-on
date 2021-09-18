package chat

import (
	"reflect"
	"testing"
)

/*
This test scenario involves creating a new room, adding some users to it,
and broadcasting a message to everyone who has joined the room. The
test runner's task is to verify that the call to Broadcast did in fact broadcast
the message to all the users. We can achieve this by examining the list of
messages that have been recorded by our injected spy
*/
func TestChatRoomBroadcast(t *testing.T) {
	pub := new(spyPublisher)
	room := NewRoom(pub)
	room.AddUser("bob")
	room.AddUser("alice")
	_ = room.Broadcast("hi")

	// Check published entries
	exp := []entry{
		{user: "bob", message: "hi"},
		{user: "alice", message: "hi"},
	}

	if got := pub.published; !reflect.DeepEqual(got, exp) {
		t.Fatalf("expected the following messages:\n%#+v\ngot:\n%#+v", exp, got)
	}
}
