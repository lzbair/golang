package main

import ( 
  "fmt"
  "math"
  "os"
  "strconv"
)

func main(){
  args := os.Args[1:]

  if(len(args) != 2) {
    fmt.Println("Program is used to compare integers, please provide two of them ...")
    return
  }

  
  var x, Xerr = strconv.ParseFloat(args[0], 64)
  var y, Yerr = strconv.ParseFloat(args[1], 64)
  if(Xerr != nil || Yerr != nil){
    fmt.Printf("Program failed, check if the arguments you provided please 1st argument : %v, 2nd argument : %v ....", Xerr, Yerr)
    return
  }
  fmt.Printf("The big one is: %v", math.Max(x, y))
}
