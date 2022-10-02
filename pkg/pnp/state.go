package pnp

import (
	"math/rand"
)

// Rand - random function
var Rand = rand.Intn

//go:generate stringer -type=State

const (
	Calm State = iota
	Annoyed
	Enraged
	Legacy
)

var chances map[State]map[Action]int

func init() {
	chances = map[State]map[Action]int{
		Calm: {
			Banana:          10,
			DuckTyping:      80,
			TypeSafety:      90,
			Inheritance:     80,
			Interfaces:      90,
			Modules:         70,
			Reflect:         70,
			MetaProgramming: 50,
			Generics:        70,
			DarkMagic:       30,
			Boredom:         100,
		},
		Annoyed: make(map[Action]int),
		Enraged: make(map[Action]int),
		Legacy:  make(map[Action]int),
	}
	for k, v := range chances[Calm] {
		chances[Annoyed][k] = v - 10
	}
	for k, v := range chances[Calm] {
		chances[Enraged][k] = v - 20
	}
	for k, v := range chances[Calm] {
		chances[Legacy][k] = v - 30
	}
	for _, v := range chances {
		v[Banana] = 10
	}
	// dark magic is surprisingly effective against legacy
	chances[Legacy][DarkMagic] = 80
	// boredom is surprisingly uneffective against legacy
	chances[Legacy][Boredom] = 10
}

func (s State) Chances(a Action) int {
	return chances[s][a]
}

// React returns XP gained, Health gained and the new Production state
func (s State) React(a Action) (int, int, State) {
	chanceOfSuccess := chances[s][a]
	success := Rand(100) < chanceOfSuccess
	next := s.nextState(success)
	xp := Rand(10*int(a)/2+1) + 1 + int(s)*10
	if success {
		return xp, Rand(10) + 1, next
	}
	if s == Legacy {
		return xp, -100, next
	}
	return xp, -Rand(10) - 1, next
}

func (s State) nextState(success bool) State {
	if success {
		if s == Calm {
			return s
		}
		return s - 1
	}
	if s == Legacy {
		return s
	}
	return s + 1
}
