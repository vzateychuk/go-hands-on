package entity

import (
	"context"
)

type Item struct {
	// various fields
}

// Box contains a list of items that are shipped to the customer.
type Box struct {
	// various fields
}

// Pack qty items of type i into the box.
func (b *Box) Pack(i *Item, qty int) error {
	p := AcquirePacker(context.Background())
	for j := 0; j < qty; j++ {
		if err := p.Pack(i, b); err != nil {
			return err
		}
	}
	return nil
}
