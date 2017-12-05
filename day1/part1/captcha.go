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
    dataArr := strings.Split(dataBuffer, "")
    // Declare sum
    sum := 0

    for i := 0; i < len(dataArr); i++ {
        
        if i <  len(dataArr)-1{
            // Add to sum if i and i+1 is the same
            if dataArr[i] == dataArr[i+1]{
                fmt.Print(dataArr[i])
                buffer, err := strconv.Atoi(dataArr[i])
                check(err) 
                sum += buffer
            }

        } else if i < len(dataArr) {
            // Add to sum if index 0 and last index is the same
            if dataArr[0] == dataArr[len(dataArr)-2]{
                buffer, err := strconv.Atoi(dataArr[0])
                check(err)
                sum += buffer
            }
        }
    }
    // Print Sum
    fmt.Print("Sum:  " )
    fmt.Print(sum)
    

}
