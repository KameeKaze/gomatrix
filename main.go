package main

import (
	"math/rand"
	"time"

	"os"
	"os/signal"
	"syscall"
	
	"github.com/nsf/termbox-go"
	"github.com/rivo/tview"
)


func main(){

	// init termbox
	err := termbox.Init()
	if err != nil {
			panic(err)
	}

	// get the terminal's width and height
	width, height := termbox.Size()
	// close termbox when program finishes
	defer termbox.Close()
	
	// generate seed for random number generator
	rand.Seed(time.Now().UnixNano())

	// generate the base matrix
	matrix := generateMatrix(width, height)
	//display matrix
	

	CloseHandler()
	

	for {
		matrix=animateMatrix(matrix)
		printMatrix(matrix)
		
		time.Sleep(time.Millisecond * 200)


	}

	
	
}


func CloseHandler(){
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func(){
		<-c
		os.Exit(0)

	}()
}
func animateMatrix(matrix  [][]int32) [][]int32 {
	
	for x := range matrix[0]{
		for y := 1; y < len(matrix)-1; y++{
			if matrix[y][x] != 32 && matrix[y-1][x] == 32{
				matrix[y][x] = 32;
				y++
			}

			if matrix[y][x] == 32 && matrix[y-1][x] != 32{
				matrix[y][x] = int32(rand.Intn(126-33)+33);
				y++
			}

		}

	}

	return matrix
}



func printMatrix(matrix [][]int32){

	for y := range matrix{
		for x := range matrix[y]{
			
			// if the char is the head of the column, print in different color
			if y < len(matrix)-1 && matrix[y+1][x] == 32{
				termbox.SetCell(x, y, matrix[y][x], termbox.ColorRed, termbox.ColorBlack)

			}else{
				termbox.SetCell(x, y, matrix[y][x], termbox.ColorGreen, termbox.ColorBlack)
			}
			
		}
	}

	err := termbox.Flush()
	if err != nil {
        return
	}
	
}




// takes in two dimension and returns the matrix
func generateMatrix(width int, height int) (matrix [][]int32){

	//create matrix for the terminal 
	matrix = make([][]int32, height)
	
	for y := 0; y < height; y++ {
		matrix[y] = make([]int32, width)
	}

	//iterate over columns and rows to and generate matrix
	for x := 0; x < width; x++ {
		columnLenght := rand.Intn(height/4)+height/3 // random length for each column
		arrayLength := 0 // starting length 
		var isChar bool = rand.Intn(2) == 1 // start random with text or empty 
		

		for y := 0; y < height; y++{
			// matrix[y][x] = 32	// ASCII(32) == space	
			if arrayLength <= columnLenght{
				if isChar{ 
					matrix[y][x] = int32(rand.Intn(126-33)+33) // add a random value between 33-126 from the ascii table
				}else{
					matrix[y][x] = 32 // add a space ~ ASCII(32): space
					
				}
				arrayLength++ // increase lenght because the column had grown
			
			// if character array is too long, switch to the other type and reset new array size
			}else{
				if isChar{
					isChar = false
					columnLenght = rand.Intn(height/4)+height/2 // less characters than spaces looks better
				}else{
					isChar = true
					columnLenght = rand.Intn(height/4)+height/4
					
				}
			
				// new array size is 0 and create new random size
				arrayLength = 0
							
			
				// still need to add a char, so add a space
				matrix[y][x] = 32
			}	
			 
		}
	}
	return

}