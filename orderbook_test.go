package main

import "testing"

func TestNewLimit(t *testing.T) {
	limit := NewLimit(10_000)
	buyOrder := NewOrder(true, 5)
}

func TestOrderBook(t *testing.T) {
	// TODO
}
