package controllers

import (
	"testing"
)

func TestAdd(t *testing.T)  {
	got := Add[int](1, 3)
	
	want := 4

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMinus(t *testing.T)  {
	got := Subtract[int](10, 4)
	
	want := 6

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}