package main

import "fmt"

func countNegatives(grid [][]int) int {
    
    negativeCounter := 0

    for i, _ := range grid {
        for j, _ := range grid[0]{
            if grid[i][j] < 0 {
                negativeCounter++
            }
        }
    }

   return negativeCounter
}