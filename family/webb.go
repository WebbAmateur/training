package main

import "fmt"

type Webb struct {
	Index  int
	First  string
	Middle string
	Last   string
	Day    int
	Month  int
	Year   int
}

func (w Webb) String() string {

	return fmt.Sprintf("Index: %d    %s, %s %s - %d/%d/%d", w.Index, w.Last, w.First, w.Middle, w.Month, w.Day, w.Year)

}
