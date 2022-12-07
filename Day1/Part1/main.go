package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main(){
    // get the input from the text file
    file, err := os.Open("./input.txt")
    if err != nil {
        fmt.Printf("Error: %s \n", err)
    }

    // Scan the file line by line
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)


    highest := 0
    count := 0;

    // Traverse the file
    for fileScanner.Scan(){

        // Every line that's not blank gets converted to an integer and added to count
        if (fileScanner.Text() != ""){
            n, err := strconv.Atoi(fileScanner.Text())
            if (err != nil){
                fmt.Printf("Error:  \n", err)
            }
            count += n

            // If the count is greater than the highest recorded number, count becomes the highest number
        }else if (count > highest){
            highest = count

        }
        // Reset count on every blank line
        if (fileScanner.Text() == ""){
            count = 0
        }

    }
    if err = file.Close(); err != nil{
        fmt.Printf("Error: %s \n", err)
    }

    fmt.Println(highest)
}