# Testing: Spies

Здесь пример использования Spy в тестировании.
--
This test scenario involves creating a new room, adding some users to it,
and broadcasting a message to everyone who has joined the room. The
test runner's task is to verify that the call to Broadcast did in fact broadcast
the message to all the users. We can achieve this by examining the list of
messages that have been recorded by our injected spy