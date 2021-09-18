package main

import (
	"circular-dependency/entity"
	"circular-dependency/warehouse"
	"context"
)

func wireComponents() {
	entity.AcquirePacker = func(ctx context.Context) entity.Packer {
		return warehouse.AcquireRobot(ctx)
	}
}

func main() {
	wireComponents()
	// ....
}
