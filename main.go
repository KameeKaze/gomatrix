package main

import (
	"fmt"
	"math/rand"
	"time"
	

	"golang.org/x/term"
	"github.com/fatih/color"
)


func main(){

	// get the terminal's width and height
	width, height, err := term.GetSize(0)
    if err != nil {
        return
	}

	// generate seed for random number generator
	rand.Seed(time.Now().UnixNano())

	// generate the base matrix
	matrix := generateMatrix(width, height)

	// set the color to haxor green
	color.Set(color.FgGreen)

	// print the matrix rotated 90˚ (matrix is horizontal)
	for x := range matrix[0]{
		for y := range matrix{
			fmt.Printf("%c",matrix[y][x])
			
		}
		fmt.Println()

	}

}


// takes in two dimension and returns the matrix (vertical)
func generateMatrix(width int, height int) (matrix [][]int){
	//create matrix for the terminal 
	matrix = make([][]int, width)
	for i := 0; i < width; i++ {
		matrix[i] = make([]int, height)
	}

	// iterate over rows
	/*
	>[][][]
	>[][][]
	>[][][]
	*/
	for i := 0; i < width; i++ {
		columnLenght := rand.Intn(10)+4 // length of each column 4-14
		arrayLength := 0 // starting length 
		var isChar bool = rand.Intn(2) == 1 // start random with text or empty 

		// iterate over columns
		/*
		V V V
		[][][]
		*/
		for j := 0; j < height; j++{
			// check if length is smaller than column length, then add eather a -char- or -null-
			if arrayLength <= columnLenght{
				if isChar == false{ 
					matrix[i][j] = 32 // add a space ~ ASCII(32): space
				}else{
					matrix[i][j] = rand.Intn(126-33)+33 // add a random value between 33-126 from the ascii table
				}
				arrayLength++ // increase lenght because the column has grown

			// if character array is too long, switch to the other type and reset new array size
			}else{
				if isChar{
					isChar = false
				}else{
					isChar = true
				}

				// new array size is 0 and create new random size
				arrayLength = 0
				columnLenght = rand.Intn(10)+4

				// still need to add a char, so add a space
				matrix[i][j] = 32
			}
			
		}
		
	}
	return

}