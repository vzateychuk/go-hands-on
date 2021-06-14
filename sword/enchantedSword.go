package sword

import "fmt"

type EnchantedSword struct {
	Sword // Embed the Sword, все методы Sword теперь в EnchantedSword
}

func (EnchantedSword) Damage() int {
	return 42 // Damage returns big damage dealt by the enchanted sword.
}

func (s EnchantedSword) String() string {
	return fmt.Sprintf("%s is a sword that can deal %d points of damage to opponents", s.name, s.Damage())
}
