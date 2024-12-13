# Go Slice Modification in Range Loop

This example demonstrates a potential bug in Go when modifying a slice within a range-based loop.

## Bug Description
Modifying the length of a slice during iteration using a range loop can lead to unexpected behavior and potentially index out of bound errors.   The standard range loop provides the index and value; when changing values, this might not be an issue.  But, if the loop modifies the slice's length, the loop's behavior becomes unpredictable.

## Bug Reproduction
1. Clone this repository.
2. Run `go run bug.go`.
3. Observe the unexpected output. 

## Bug Solution
The provided `bugSolution.go` file shows how to avoid the issue by either iterating over a copy of the slice or using a for loop that explicitly controls the index.
