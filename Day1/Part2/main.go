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
    secondHighest := 0
    thirdHighest := 0
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
        }
        if (fileScanner.Text() == ""){
            // If the count is greater than the highest recorded number, count becomes the highest number
            // The highest gets passes down to the secondHighest and thirdHighest as it gets replaced
         if (count > highest){
             secondHighest = highest
             highest = count

        }else if(count > secondHighest){
            thirdHighest = secondHighest
            secondHighest = count
        }else if(count > thirdHighest){
            thirdHighest = count
        }

        // Reset count on every blank line
        fmt.Printf("highest: %d secondHighest: %d thirdHighest: %d count: %d\n",highest,secondHighest,thirdHighest,count)
            count = 0
        }

    }
    if err = file.Close(); err != nil{
        fmt.Printf("Error: %s \n", err)
    }

    fmt.Printf("highest: %d\n",highest)
    fmt.Printf("secondHighest: %d\n",secondHighest)
    fmt.Printf("thirdHighest: %d\n",thirdHighest)
   // total := highest+secondHighest+66172
    fmt.Println(highest+secondHighest+thirdHighest)
}