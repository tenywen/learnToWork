package gsmap

type Manager struct {
	Tiles   map[MapPos]*Tile
	Marches map[int32]*March
}
