package gsmap

import (
	"timer"
)

const (
	TILE_FOREST            = 1  // 绿地(basetitle)
	TILE_WASTELAND         = 2  // 废墟(basetitle)
	TILE_BOG               = 3  // 泥地(basetitle)
	TILE_MOUNTAIN          = 4  // 山脉(basetitle)
	TILE_MINE              = 5  // 乱石岗(basetitle)
	TILE_RUINS             = 6  // 城市遗址(basetitle)
	TILE_EMPTY             = 7  // 空地
	TILE_NULL              = 8  // 装饰物
	TILE_MONUMENT          = 9  // 遗迹
	TILE_FARM              = 11 // 农场(res)
	TILE_ENERGY            = 12 // 能源工厂(res)
	TILE_OIL               = 13 // 油田(res)
	TILE_STEEL             = 14 // 矿场(res)
	TILE_CASH              = 15 // 村落(res)
	TILE_ITEM              = 16 // 道具(res)
	TILE_WONDER            = 17 // 奇观
	TILE_WONDER_LAND       = 18 // 奇观领土
	TILE_CITY              = 20 // 城市
	TILE_FORT              = 21 // 据点
	TILE_LAND              = 22 // 领土
	TILE_REBEL             = 23 // 叛军
	TILE_SUPER_FARM        = 31 // 超级矿之农场
	TILE_SUPER_ENERGY      = 32 // 超级矿之能源矿
	TILE_SUPER_OIL         = 33 // 超级矿之油田
	TILE_SUPER_STEEL       = 34 // 超级矿之钢铁
	TILE_SUPER_CASH        = 35 // 超级矿之现金矿
	TILE_SUPER_DIAMOND     = 36 // 超级矿之砖石矿
	TILE_SUPER_DOMINO      = 37 // 超级矿之黑市币
	TILE_SUPER_CREDIT      = 38 // 超级矿之军演币
	TILE_ALLIANCE_RESEARCH = 50 // 科技矿
	TILE_ALLIANCE_MB       = 51 // 演习基地
)

type Tile struct {
	// base field
	BaseType   int8
	TileType   int8
	Pos        MapPos
	Level      int8
	Owner      int64            `bson:",omitempty"`
	AllianceId int32            `bson:",omitempty"`
	TroopGroup int32            `bson:",omitempty"`
	MarchIds   []int32          `bson:",omitempty"` // 经过这块tile的行军ids
	Status     int8             `bson:",omitempty"`
	Character  int8             `bson:",omitempty"` // 领土自定义A-Z字符
	Start      int64            `bson:",omitempty"`
	End        int64            `bson:",omitempty"`
	Fighting   int64            `bson:",omitempty"`
	TimeEvent  *timer.TimeEvent `bson:"-"`

	// extra field
	// ResType     int8
	// Invater     int64
	// LoadBonus   float64
	// GatherBonus float64
	Extra map[string]string `bson:",omitempty"`
}
