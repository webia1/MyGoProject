package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	r.LoadHTMLFiles("index.html", "solved.html")



	r.GET("/", func(c *gin.Context) {
		s := &Sudoku{}
		c.HTML(http.StatusOK, "index.html", s.Puzzle)
})

	r.POST("/solve", func(c *gin.Context) {
		s := &Sudoku{}
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				val := c.PostForm(strconv.Itoa(i) + strconv.Itoa(j))
				if val != "" {
					s.Puzzle[i][j], _ = strconv.Atoi(val)
				} else {
					s.Puzzle[i][j] = 0
				}
			}
		}
		s.solve()
		c.HTML(200, "solved.html", gin.H{"puzzle": s.Puzzle})
	})

	r.Run()
}
