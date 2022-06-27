# Eight Queen's Chess Puzzle Solver

This program attempts to solve the Eight Queens chess puzzle each time it is run.

The Eight Queens is a chess-based puzzle where eight queens must be placed on a chess board in a configuration where none of the queens threatens any of the others.

Additionally, this program will solve this puzzle for any square board size. The default is the standard chess board size of 8x8.

To solve this puzzle, this program uses:

- A two-dimensional array to track the state of the board
- A shrinking two-dimensional array to track the remaining valid locations on the board
- Recursion to allow backtracking when a solution can no longer be advanced
- RNG to randomly determine where to place the next queen amongst the remaining available spaces

# Running the program for the default board size (8x8)

`go run .`

# Running the program for a custom board size (NxN)

`go run . N` where `N` is the size of the board.

For example:
`go run . 10` will attempt to solve the puzzle on a 10x10 board.

Note:
While you CAN run this script with a very large board size, be warned the maximum time required to solve the puzzle increases dramatically as the board size increases.
