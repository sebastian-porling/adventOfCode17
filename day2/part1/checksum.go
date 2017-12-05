 package main
 import (

     "fmt"
     "strconv"
     "io/ioutil"
     "strings"
 )
 // Check error and print error
 func check(e error) {
     if e != nil {
         panic(e)
     }
 }

 func main() {
     // Read file
     dat, err := ioutil.ReadFile("input.txt")
     // Check error and print
     check(err)
     // Create buffer and parse input to string
     dataBuffer := string(dat)
     // Split input to array
     dataArr := strings.Split(dataBuffer, "\n")
     // Declare Sum
     sum := 0

     for i := 0; i < len(dataArr)-1; i++ {
        // Declare min and max
        min := 13333337
        max := 0
        // Split Each "row" into array
        tmpArr := strings.Split(dataArr[i], " ")
        // Declare Diff
        diff := 0

        for j := 0; j < len(tmpArr); j++{
            // Don't know how to filter...
            if tmpArr[j] != ""{
                // Create buffer
                buffer, err := strconv.Atoi(tmpArr[j])
                check(err)
                // If max is smaller than buffer
                if max < buffer{
                    max = buffer
                }
                // If min is smaller than buffer
                if min > buffer{
                    min = buffer
                }
            }
        }
        // Difference of max and min
        diff = max - min
        // Sum of all rows
        sum += diff
     }
     // Print sum
     fmt.Print(" Sum: ")
     fmt.Print(sum)

 }
