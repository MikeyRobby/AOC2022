package main

import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    const win = 6
    const loss = 0
    const draw = 3
    const rock = 1
    const paper = 2
    const scissors = 3
    score := 0

    file, err := os.Open("./input.txt")
    if err != nil {
        fmt.Printf("[ERROR]: %s\n", err)
    }
    defer file.Close()

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan(){
        if fileScanner.Text() == "A Y"{
            score += (rock+draw)
        }else if fileScanner.Text() == "A X"{
            score += (scissors+loss)
        }else if fileScanner.Text() == "A Z"{
            score += (paper+win)
        }


        if fileScanner.Text() == "B Y"{
            score += (paper+draw)
        }else if fileScanner.Text() == "B X"{
            score += (rock+loss)
        }else if fileScanner.Text() == "B Z"{
            score += (scissors+win)
        }

        if fileScanner.Text() == "C Y"{
            score += (scissors+draw)
        }else if fileScanner.Text() == "C X"{
            score += (paper+loss)
        }else if fileScanner.Text() == "C Z"{
            score += (rock+win)
        }

    }

    fmt.Printf("%d\n", score)

}