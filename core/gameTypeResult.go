package core

import (
	e "q6-golang/enumerators"
)

type GameTypeResult struct {
	Players        []Player   `json:"players"`
	GameType       e.GameType `json:"gameType"`
	DrawingResults []int      `json:"drawingResults"`
}
