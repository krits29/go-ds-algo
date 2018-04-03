package main

import (
	"fmt"
)

// given a sorted list of events containing start, end time and id for an event.
func findConflicts(events [][]int) []int {

	conflicts := []int{}
	if len(events) <= 1 {
		return conflicts
	}

	// prime the loop
	end := events[0][1]                 // end time for the first event
	tmpConflicts := []int{events[0][2]} // add the id of the first event

	for i := 1; i < len(events); i++ { // go over the list onces as it is sorted on start time
		if events[i][0] >= end { // no conflicts
			if len(tmpConflicts) > 1 { // more than one event in the tmp means there is a conflict, remember first element is itself
				conflicts = append(conflicts, tmpConflicts...)
			}
			tmpConflicts = tmpConflicts[:0] // reset the tempConflicts
		}
		end = max(events[i][1], end)                      // takes the max
		tmpConflicts = append(tmpConflicts, events[i][2]) // add to tempConflicts
	}
	if len(tmpConflicts) > 1 { // handle the last case
		conflicts = tmpConflicts
	}
	return conflicts
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	events := [][]int{{1, 2, 1}, {3, 4, 2}, {4, 6, 3}, {5, 6, 4}, {6, 7, 5}}
	fmt.Println(findConflicts(events))
}
