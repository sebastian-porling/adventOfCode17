 package main
 import (

     "fmt"
     "strconv"
     "io/ioutil"
     "strings"
 )

 func check(e error) {
     if e != nil {
         panic(e)
     }
 }

 func main() {
    // Read input file
    dat, err := ioutil.ReadFile("input.txt")
    // Check if there are error
    check(err)
    // Create buffer for parsing
    dataBuffer := string(dat)
    // Split string to array
    dataArr := strings.Split(dataBuffer, "")
    // Declare sum
    sum := 0
    // Print the half lenght of the array
    half := len(dataArr)/2
    fmt.Print(half)
    fmt.Print("<- half lenght  ")
    // Print the full lenght of the array
    fmt.Print(len(dataArr))
    fmt.Print("<- Full length   ")
    
    for i := 0; i < len(dataArr); i++ {

        if i <  half{
            // Add to sum if i and i+half is the same
            if dataArr[i] == dataArr[i + half]{
                buffer, err := strconv.Atoi(dataArr[i])
                check(err)
                sum += buffer
            }
        
        } else if i < len(dataArr) {
            // Add to sum if i and i - half is the same
            if dataArr[i] == dataArr[i - half]{
               buffer, err := strconv.Atoi(dataArr[i])
               check(err)
               sum += buffer
            }
        }
    }
    // Print sum
    fmt.Print("  Sum:  " )
    fmt.Print(sum)
 }
