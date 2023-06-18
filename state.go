package main

type state int

const (
	stateMovingHoritzontally state = iota

	stateMovingDown
	stateDead
)
