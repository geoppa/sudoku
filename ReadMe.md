# 🧩 Sudoku Solver (Go)

A high-performance terminal application built with **Go** that solves any valid 9x9 Sudoku puzzle using a **Recursive Backtracking** algorithm.

---

## 🛠 Features

* **Backtracking Algorithm:** An efficient depth-first search approach to find the solution.
* **Memory Management:** Optimized use of **Pointers** to modify the board state in-place.
* **Robust Validation:** Detailed error reporting for row lengths, invalid characters, and input size.
* **Interactive UI:** A user-friendly flow that allows you to preview the board and choose to solve or quit (`Q`).



---

## 🚀 Usage

To run the solver, provide **9 strings** as command-line arguments. Use a dot (`.`) for empty cells:

```bash
go run . "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79" 
```
📂 Project Structure

    main.go: Handles CLI arguments, validation, and the interactive menu.

    solver.go: Contains the core recursive backtracking logic.

    utils.go: Helper functions for safety checks (isSafe) and board visualization.

🧩 Test Example

Copy and paste these command into your terminal:

```Bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
```


🧠 Logic in a Nutshell

    Validate: Ensures the input is a perfect 9x9 grid.

    Locate: Scans for the next empty cell (.).

    Try: Attempts numbers 1-9 using isSafe.

    Recurse: Moves forward if a number fits.

    Backtrack: Resets and tries again if a dead end is reached.