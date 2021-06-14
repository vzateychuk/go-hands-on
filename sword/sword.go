package sword

import "fmt"

type Sword struct {
	name string // Important tip for RPG players: always name your swords!
}

// Damage returns the damage dealt by this sword.
func (Sword) Damage() int {
	return 2
}

func (s Sword) String() string {
	return fmt.Sprintf("%s is a sword that can deal %d points of damage to opponents", s.name, s.Damage())
}
