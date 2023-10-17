package c3

import "fmt"

type Vertex struct {
	lat, long float64
}

// key is string and value is Vertax
var m map[string]Vertex

func Map() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["My home"] = Vertex{30.5132, -52.5646}

	fmt.Println(m["My home"])

	// In literals key is required
	mm := map[string]Vertex{
		"drdo": {30.2454, 50.3534},
		"isro": {45.3535, 60.2345},
	}
	fmt.Println(mm["drdo"])

	mx := make(map[string]int)
	mx["name"] = 25
	fmt.Println(mx["name"])
	mx["name"] = 20
	fmt.Println(mx["name"])
	mx["delete"] = 30
	// delete
	delete(mx, "delete")

	// get
	val, ok := mx["delete"]
	if !ok {
		fmt.Println("deleted")
	} else {
		fmt.Println(val)
	}
}

/*
	map zero value is null that is no key or value
*/
