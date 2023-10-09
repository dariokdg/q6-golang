package models

type Result struct {
	GameType       GameType `json:"gameType"`
	DrawingResults []int    `json:"drawingResults"`
}
