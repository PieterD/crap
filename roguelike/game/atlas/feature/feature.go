package feature

type Feature struct {
	Name     string
	Passable bool
}

var Wall = Feature{
	Name:     "wall",
	Passable: false,
}
var Floor Feature = Feature{
	Name:     "floor",
	Passable: true,
}
