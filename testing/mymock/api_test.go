package mymock_test

import (
	"github.com/golang/mock/gomock"
	"mymock/mock"
	"testing"
)

func TestAPI(t *testing.T) {
	// Create a controller to manage all our mock objects and make sure
	// that all expectations were met before completing the test
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Obtain a mock instance that implements API and associate it with the controller.
	api := mock_mymock.NewMockAPI(ctrl)

	/* Under normal circumstances, gomock would just check that the method call
	expectations are met, regardless of the order that they were invoked in.
	However, if a test relies on a sequence of method calls being performed
	in a particular order, it can specify this to gomock by invoking
	the gomock.InOrder helper function with an ordered list of expectations as
	arguments. */
	gomock.InOrder(
		api.EXPECT().SendMessage("msg").Times(1).Return(nil),
		api.EXPECT().ConsumeMessage().Times(1).Return("msg", nil),
	)

	// With the mock expectations in place, we can complete our unit by
	//introducing the necessary logic to wire everything together, invoke the actual
	// method used the API
}
