package main
import (

    "fmt"
    //"strconv"
    "io/ioutil"
    "strings"

)

  // Check error and print error
  func check(e error) {
      if e != nil {
          panic(e)
      }
  }

func main(){
    // Read file
    dat, err := ioutil.ReadFile("input.txt")

    check(err)
    
    dataBuffer := string(dat)
    
    dataArr := strings.Split(dataBuffer, "\n")
    count := 0
    fmt.Print("Before first arr")
    for i:=0; i < len(dataArr); i++ {
        fmt.Print(dataArr)
        tmpArr := strings.Split(dataArr[i], " ")
        
        for j:=0; j < len(tmpArr); j++{
        fmt.Print(tmpArr)

        tmpWord := tmpArr[j]

        left := 0

        right := len(tmpArr)

        boo := true


        for boo == true {
            //fmt.Print("printing in bsearch")
            if left > right {
                count++
                boo = false
            }
            mid := (left + right) / 2
            if tmpArr[mid] < tmpWord {
                left = mid + 1
            }else if tmpArr[mid] > tmpWord {
                right = mid + 1
            }else{
                boo = false
            }

        }
    }

    }
    fmt.Print("Number of valid passphrases: ")
    fmt.Print(count)
}
