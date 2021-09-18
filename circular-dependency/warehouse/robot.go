package warehouse

import (
	"circular-dependency/entity"
	"context"
)

// Robot navigates the warehouse floor and fetches items for packing.
type Robot struct {
	// various fields
}

// AcquireRobot blocks until a Robot becomes available or until the
// context expires.
func AcquireRobot(ctx context.Context) *Robot {
	//...
	return &Robot{}
}

// Pack instructs the robot to pick up an item from its shelf and place
// it into a box that will be shipped to the customer.
func (r *Robot) Pack(item *entity.Item, to *entity.Box) error {
	//...
	return nil
}
