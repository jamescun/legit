package legit

import (
	"testing"
)

func TestEmail(t *testing.T) {
	testString(t, Email("foo@example.org"), Email("foo@@bar@!"), errEmail)
}

func TestCreditCard(t *testing.T) {
	testString(t, CreditCard("375556917985515"), CreditCard("foo"), errCreditCard)
}

func TestUUID(t *testing.T) {
	testString(t, UUID("a987fbc9-4bed-3078-cf07-9141ba07c9f3"), UUID("xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3"), errUUID)
}

func TestUUID3(t *testing.T) {
	testString(t, UUID3("a987fbc9-4bed-3078-cf07-9141ba07c9f3"), UUID3("xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3"), errUUID3)
}

func TestUUID4(t *testing.T) {
	testString(t, UUID4("625e63f3-58f5-40b7-83a1-a72ad31acffb"), UUID4("xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3"), errUUID4)
}

func TestUUID5(t *testing.T) {
	testString(t, UUID5("987fbc97-4bed-5078-9f07-9141ba07c9f3"), UUID5("xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3"), errUUID5)
}
