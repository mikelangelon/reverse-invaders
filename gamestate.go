package main

type gameState int

const (
	gameStateMenu       gameState = iota
	gameStatePrePlaying gameState = iota
	gameStatePlaying
	gameRestarts
	gameStateEnded
)
