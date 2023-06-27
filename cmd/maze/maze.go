package main

import (
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/ronna-s/go-ood/pkg/maze"
	"github.com/ronna-s/go-ood/pkg/maze/travel"
	"github.com/ronna-s/go-ood/pkg/robot"
)

//go:embed resources/maze.tmpl
var tmpl []byte

//go:generate mockery --case=underscore --name=Gopher

// Gopher is an interface to a thing that can move around a maze
type Gopher interface {
	Finished() bool // Has the Gopher reached the target cell?
	Move() error    // The Gopher moves one step in its current direction if the gopher can't move it returns an error
	TurnLeft()      // The Gopher will turn left
	TurnRight()     // The Gopher will turn right
}

// SolveMaze is where your code goes to solve our maze
// It takes in g Gopher that knows how to travel a maze
// See the Gopher interface methods for more details
func SolveMaze(g Gopher) {
	lookRight := func(g Gopher) error {
		g.TurnRight()
		return g.Move()
	}

	for !g.Finished() {
		g.TurnLeft() // -90
		if err := g.Move(); err != nil {
			if err := lookRight(g); err != nil {
				// 90
				if err := lookRight(g); err != nil {
					// 180
					if err := lookRight(g); err != nil {
						// 270
						if err := lookRight(g); err != nil {
							// 360
							if err := lookRight(g); err != nil {
								panic(
									fmt.Errorf(
										"We did complete a whole circle. That's wrong! We're trapped! Call the Police!!! %w",
										err,
									),
								)
							}
						}
					}
				}
			}
		}
	}
}

// drawHTML writes the movement of the gopher through the maze to HTML
func drawHTML(g robot.Robot, w io.Writer) {
	res := struct {
		maze.Maze
		Steps []robot.Step
	}{
		Maze:  g.Maze,
		Steps: g.Steps(),
	}
	const maxMoves = 10000
	if len(res.Steps) > maxMoves {
		res.Steps = res.Steps[:maxMoves]
	}

	t, err := template.New("main").Funcs(template.FuncMap{
		"Upto": func(count int) (items []int) {
			for i := 0; i < count; i++ {
				items = append(items, i)
			}
			return items
		},
		"Incr": func(i int) int {
			return i + 1
		},
		"Passages": func(i int) (s string) {
			c := res.CoordsFromCell(i)
			if res.PathRight(i) {
				s += "right "
			} else if c[maze.X] != res.DimX-1 {
				s += "no-right "
			}
			if i == res.DimX*res.DimY-1 {
				s += "exit "
			}
			if res.PathDown(i) {
				s += "down"
			} else if c[maze.Y] != res.DimY-1 {
				s += "no-down"
			}
			return
		},
		"PathDown": func(i int) bool {
			return res.PathDown(i)
		},
		"Mul": func(x, y int) int {
			return x * y
		},
	}).Parse(string(tmpl))
	if err != nil {
		panic(err)
	}
	if err := t.ExecuteTemplate(w, "T", res); err != nil {
		panic(err)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	m := maze.New(rand.Intn(10)+1, rand.Intn(10)+1)
	g := robot.New(travel.New(m))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		SolveMaze(&g)
		cancel()
	}()
	<-ctx.Done()
	drawHTML(g, os.Stdout)
}

// code that will fail if the constants change value since the JS code depends on it
func _() {
	var x [1]struct{}
	_ = x[maze.Right-0]
	_ = x[maze.Down-1]
	_ = x[maze.Up-2]
	_ = x[maze.Left-3]
}
