package models

type GameTypeResult struct {
	GameType       GameType `json:"gameType"`
	DrawingResults []int    `json:"drawingResults"`
}
