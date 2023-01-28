package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Sudoku struct {
    Puzzle [9][9]int
}

func (s *Sudoku) solve() bool {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if s.Puzzle[i][j] == 0 {
                for k := 1; k <= 9; k++ {
                    if isValid(s.Puzzle, i, j, k) {
                        s.Puzzle[i][j] = k
                        if s.solve() {
                            return true
                        } else {
                            s.Puzzle[i][j] = 0
                        }
                    }
                }
                return false
            }
        }
    }
    return true
}

func isValid(puzzle [9][9]int, row int, col int, num int) bool {
    // Check row
    for i := 0; i < 9; i++ {
        if puzzle[row][i] == num {
            return false
        }
    }

    // Check col
    for i := 0; i < 9; i++ {
        if puzzle[i][col] == num {
            return false
        }
    }

    // Check small grid
    startRow := row - row%3
    startCol := col - col%3
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if puzzle[i+startRow][j+startCol] == num {
                return false
            }
        }
    }

    return true
}

func main() {
    http.HandleFunc("/sudoku", func(w http.ResponseWriter, r *http.Request) {



        if r.Method == "POST" {
            s := &Sudoku{}
            for i := 0; i < 9; i++ {
                for j := 0; j < 9; j++ {
                    val, _ := strconv.Atoi(r.FormValue(fmt.Sprintf("%d-%d", i, j)))
                    s.Puzzle[i][j] = val
                }
            }
            s.solve()
            t, _ := template.ParseFiles("./sudoku.gohtml")
            t.Execute(w, s.Puzzle)
        } else {
            t, _ := template.ParseFiles("./sudoku.gohtml")
            t.Execute(w, [9][9]int{})
        }

				t, err := template.ParseFiles("sudoku.gohtml")
				if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
				}
				t.Execute(w, nil)
    })

	

    http.ListenAndServe(":8080", nil)
}
