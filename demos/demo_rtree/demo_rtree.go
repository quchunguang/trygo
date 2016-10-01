package main

import (
	"fmt"
	"github.com/dhconnelly/rtreego"
)

// Thing describes object with 2D bound
type Thing struct {
	where *rtreego.Rect
	name  string
}

func (t *Thing) Bounds() *rtreego.Rect {
	return t.where
}

// Somewhere describes object with a position,
// with is centered at location and 2*tol weigh and length.
var tol = 0.01

type Somewhere struct {
	location rtreego.Point
	name     string
	wormhole chan int
}

func (s *Somewhere) Bounds() *rtreego.Rect {
	// define the bounds of s to be a Rectangle centered at s.location
	// with side lengths 2 * tol:
	return s.location.ToRect(tol)
}

func main() {
	// Demo rtree describes rect
	rt := rtreego.NewTree(2, 25, 50)
	p1 := rtreego.Point{0.4, 0.5}
	p2 := rtreego.Point{6.2, -3.4}
	r1, _ := rtreego.NewRect(p1, []float64{1, 2})
	r2, _ := rtreego.NewRect(p2, []float64{1.7, 2.7})
	rt.Insert(&Thing{r1, "foo"})
	rt.Insert(&Thing{r2, "bar"})

	fmt.Println(rt.Size()) // returns 2
	// rt.Delete(thing2)
	// rt.Insert(anotherThing

	// Bounding-box queries
	bb, _ := rtreego.NewRect(rtreego.Point{0, 0}, []float64{4, 4})
	// Get a slice of the objects in rt that intersect bb:
	results := rt.SearchIntersect(bb)
	fmt.Println(results[0].Bounds())

	// K-nearest-neighbors queries
	q := rtreego.Point{6.5, -2.47}
	k := 1
	// Get a slice of the k objects in rt closest to q:
	fmt.Println(rt.NearestNeighbors(k, q)[0].Bounds())

	// Demo rtree describes position
	rtp := rtreego.NewTree(2, 25, 50)
	rtp.Insert(&Somewhere{rtreego.Point{0, 0}, "Someplace", nil})
	fmt.Println(rtp.Size()) // returns 2
}
