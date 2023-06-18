package main

type gameState int

const (
	gameStateMenu gameState = iota
	gameStatePlaying
	gameStateEnded
)
