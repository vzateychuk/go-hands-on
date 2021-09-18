package entity

import "context"

// AcquirePacker returns a Packer instance.
var AcquirePacker func(context.Context) Packer

// Packer is implemented by objects that can pack an Item into a Box.
type Packer interface {
	Pack(*Item, *Box) error
}
