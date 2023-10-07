package core

type GameTypeResult struct {
	Players        []Player `json:"players"`
	GameType       GameType `json:"gameType"`
	DrawingResults []int    `json:"drawingResults"`
}
