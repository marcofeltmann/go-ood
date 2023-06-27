package pnpdev

import (
	"fmt"

	"github.com/ronna-s/go-ood/pkg/pnp"
)

// Minion represents a minion P&P player
// The zero value is a dead minion player.
type Rubyist struct {
	X, H int //X=XP, H=Health
}

// NewMinion returns a minion with 100 Health and 0 XP
func NewRubyist() *Rubyist {
	return &Rubyist{H: 100}
}

func (m Rubyist) String() string {
	return "Master Rubyist"
}

// Alive checks if the player is (still) alive
func (m Rubyist) Alive() bool {
	return m.H > 0
}

// Health returns the player's health level
func (m Rubyist) Health() int {
	return m.H
}

// XP returns the player's xp level
func (m Rubyist) XP() int {
	return m.X
}

// ApplyXPDiff adds the given xp to the player's xp down to a minimum of 0
func (m *Rubyist) ApplyXPDiff(xp int) int {
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
func (m *Rubyist) ApplyHealthDiff(health int) int {
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

var allRubyistSkills = []pnp.Skill{pnp.DuckTyping, pnp.Inheritance, pnp.Modules, pnp.MetaProgramming, pnp.DarkMagic}

// Skills returns the minion's skills which are Banana.
func (m Rubyist) Skills() []pnp.Skill {
	xp := m.XP()
	switch {
	case xp < 11:
		return allRubyistSkills[:3]
	case xp >= 11 && xp < 100:
		return allRubyistSkills[:4]
	default:
		return allRubyistSkills[:5]
	}
}

// AsciiArt returns the minion's ascii-art
func (m Rubyist) AsciiArt() string {
	return fmt.Sprintf(rubyistArt, m.H, m.X)
}
