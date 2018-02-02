package gsmap

type MapPos struct {
	X int16
	Y int16
}

var (
	GRID_DIR = [15][2]int16{
		{0, 0}, {1, 0}, {-1, 0},
		{0, 1}, {1, 1}, {-1, 1},
		{0, -1}, {1, -1}, {-1, -1},
		{0, 2}, {1, 2}, {-1, 2},
		{0, -2}, {1, -2}, {-1, -2},
	} // Grid邻接方向。

)

const (
	GRID_SIZE  = 4   // 地图网格大小
	MAP_WIDTH  = 512 // 地图度(x)
	MAP_HEIGHT = 512 // 地图高(y)

)

type GridPos struct {
	X int16
	Y int16
}

//--------------------------------------------------------- MapPos对应的grid pos
func getGridPos(pos MapPos) GridPos {
	return GridPos{pos.X / GRID_SIZE, pos.Y / GRID_SIZE}
}

//-------------------------------------------------------- mapPos对应grid包含的所有MapPos
func GetAllMapPosInGrid(pos MapPos) []MapPos {
	grid := getGridPos(pos)
	return getAllMapPosInGrid(grid)
}

func getAllMapPosInGrid(grid GridPos) []MapPos {
	var pos_list []MapPos
	for i := int16(0); i < GRID_SIZE; i++ {
		for j := int16(0); j < GRID_SIZE; j++ {
			new_pos := MapPos{i + grid.X, j + grid.Y}
			if !IsMapPosValid(new_pos) {
				break
			}
			pos_list = append(pos_list, new_pos)
		}
	}
	return pos_list
}

//--------------------------------------------------------- MapPos对应grid以及附近grid的所有MapPos
func GetAllMapPosNearyGrid(pos MapPos) []MapPos {
	var pos_list []MapPos
	grid := getGridPos(pos)
	for _, dir := range GRID_DIR {
		new_grid := GridPos{grid.X + dir[0], grid.Y + dir[1]}
		if isGridPosValid(new_grid) {
			pos_list = append(pos_list, getAllMapPosInGrid(new_grid)...)
		}
	}
	return pos_list
}

func IsMapPosValid(pos MapPos) bool {
	if pos.X < 0 || pos.X >= MAP_WIDTH {
		return false
	}

	if pos.Y < 0 || pos.Y >= MAP_HEIGHT {
		return false
	}
	return true
}

func isGridPosValid(grid GridPos) bool {
	if grid.X < 0 || grid.X >= MAP_WIDTH/GRID_SIZE {
		return false
	}

	if grid.Y < 0 || grid.Y >= MAP_HEIGHT/GRID_SIZE {
		return false
	}
	return true
}
