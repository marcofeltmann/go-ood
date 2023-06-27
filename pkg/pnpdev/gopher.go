package pnpdev

import (
	"fmt"

	"github.com/ronna-s/go-ood/pkg/pnp"
)

// Minion represents a minion P&P player
// The zero value is a dead minion player.
type Gopher struct {
	X, H int //X=XP, H=Health
}

type ZombieGopher struct {
	Gopher
}

func (z ZombieGopher) String() string {
	return "BRAINS!!!"
}
func (z ZombieGopher) Alive() bool {
	return true
}
func (z *ZombieGopher) ApplyHealthDiff(health int) int {
	// Zombie cannot heal.
	// But it shouldn't kill it, because that will make the pizza-feeding
	// boss a really ugly person, will it?
	if health >= 0 {
		return z.H
	}
	z.H += health
	return z.H
}

func NewZombieGopher() *ZombieGopher {
	g := Gopher{}
	return &ZombieGopher{g}
}

// NewMinion returns a minion with 100 Health and 0 XP
func NewGopher() *Gopher {
	return &Gopher{H: 100}
}

// Alive checks if the player is (still) alive
func (m Gopher) Alive() bool {
	return m.H > 0
}
func (m Gopher) String() string {
	return "Go Gopher"
}

// Health returns the player's health level
func (m Gopher) Health() int {
	return m.H
}

// XP returns the player's xp level
func (m Gopher) XP() int {
	return m.X
}

// ApplyXPDiff adds the given xp to the player's xp down to a minimum of 0
func (m *Gopher) ApplyXPDiff(xp int) int {
	sum := m.X + xp
	if sum < 0 {
		xp = -m.X
		m.X = 0
		return xp
	}
	m.X = sum
	return xp
}

// ApplyHealthDiff adds the given health to the player's health down to a minimum of 0 and up to 100
func (m *Gopher) ApplyHealthDiff(health int) int {
	sum := m.H + health
	if sum > 100 {
		health = 100 - m.H
		m.H = 100
		return health
	}
	if sum < 0 {
		health = -m.H
		m.H = 0
		return health
	}
	m.H = sum
	return health
}

var allGopherSkills = []pnp.Skill{pnp.TypeSafety, pnp.Interfaces, pnp.Reflect, pnp.Generics, pnp.Boredom, 23, 42}

// Skills returns the minion's skills which are Banana.
func (m Gopher) Skills() []pnp.Skill {
	xp := m.XP()
	switch {
	case xp < 11:
		return allGopherSkills[:3]
	case xp >= 11 && xp < 100:
		return allGopherSkills[:4]
	case xp == 101:
		return allGopherSkills[:5]
	default:
		return allGopherSkills
	}
}

// AsciiArt returns the minion's ascii-art
func (m Gopher) AsciiArt() string {
	return fmt.Sprintf(gopherArt, m.H, m.X)
}
