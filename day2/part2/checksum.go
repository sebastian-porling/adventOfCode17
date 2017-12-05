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
     // Declare sum
     sum := 0

     for i := 0; i < len(dataArr)-1; i++ {
        // Make array from "row" i
        tmpArr := strings.Split(dataArr[i], " ")
        // Label for breaking out of loop 
        // When finding the even division
        Loop:
        for j := 0; j < len(tmpArr); j++{
            // Have to check empty string... Don't know how to filter
            if tmpArr[j] != ""{
                // Create buffer
                buffer, err := strconv.Atoi(tmpArr[j])
                check(err)

                for k := j+1; k < len(tmpArr); k++ {
                    // Second buffer that is j+1
                    buffer2, err2 := strconv.Atoi(tmpArr[k])
                    check(err2)
                    //Check if division with first buffer and second buffer is even
                    t := buffer % buffer2
                    if t == 0{
                        sum += buffer / buffer2
                        break Loop
                    }
                    //Check if division with second buffer and first buffer is even
                    t = buffer2 % buffer
                    if t == 0{
                        sum += buffer2 / buffer
                        break Loop
                    }
                }
            }
        }
     }
     // Print sum
     fmt.Print(" Sum: ")
     fmt.Print(sum)

 }
