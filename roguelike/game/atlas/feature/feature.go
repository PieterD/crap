package feature

type Feature struct {
	Name     string
	Passable bool
	Wallable bool
}

var Void = Feature{}
var Wall = Feature{
	Name:     "wall",
	Wallable: true,
}
var Floor Feature = Feature{
	Name:     "floor",
	Passable: true,
}
