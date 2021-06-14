package sword

import (
	"fmt"
	"testing"
)

func TestSwordDamage(t *testing.T) {
	specs := []struct {
		sword interface{ Damage() int }
		exp   int
	}{
		{
			sword: Sword{name: "Silver Saber"},
			exp:   2,
		},
		{
			sword: EnchantedSword{Sword{name: "Dragon's Greatsword"}},
			exp:   42,
		},
	}

	for specIndex, spec := range specs {
		if got := spec.sword.Damage(); got != spec.exp {
			t.Errorf("[spec %d] expected to get damage %d; got %d", specIndex, spec.exp, got)
		}
	}
}

func TestSword_String(t *testing.T) {
	specs := []struct {
		sword fmt.Stringer
		exp   string
	}{
		{
			sword: Sword{name: "Silver Saber"},
			exp:   "Silver Saber is a sword that can deal 2 points of damage to opponents",
		},
		{
			sword: EnchantedSword{Sword{name: "Dragon's Greatsword"}},
			exp:   "Dragon's Greatsword is a sword that can deal 42 points of damage to opponents",
		},
	}

	for specIndex, spec := range specs {
		got := spec.sword.String()

		if got != spec.exp {
			t.Errorf("[spec %d] expected to get\n%q\ngot:\n%q",
				specIndex, spec.exp, got)
		}
	}
}
