package core

import "sync"

type Grid struct {
	Gid int
	/*
		x1, y1 ------ x2, y1
			|			|
			|			|
			|			|
			|			|
		x1, y2 ------- x2, y2
	*/
	X1 int
	X2 int
	Y1 int
	Y2 int
	// 当前格子内玩家id
	playerIds map[int]bool
	// map锁
	idLock sync.RWMutex
}

func NewGrid(gid, x1, x2, y1, y2 int) *Grid {
	return &Grid{
		Gid:       gid,
		X1:        x1,
		X2:        x2,
		Y1:        y1,
		Y2:        y2,
		playerIds: make(map[int]bool),
	}
}

// 给格子添加一个玩家
func (g *Grid) AddPlayer(playerId int) {
	g.idLock.Lock()
	defer g.idLock.Unlock()
	g.playerIds[playerId] = true
}

// 删除
func (g *Grid) RemovePlayer(playerId int) {
	g.idLock.Lock()
	defer g.idLock.Unlock()
	delete(g.playerIds, playerId)
}

func (g *Grid) GetPlayerIds() []int {
	g.idLock.RLock()
	defer g.idLock.RUnlock()
	ids := make([]int, 0)
	for id, _ := range g.playerIds {
		ids = append(ids, id)
	}
	return ids
}
