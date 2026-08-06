package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64)float64{
     z:=1.0
     for {
          prev:=z

         z-=(z*z-x)/(2*z)
         if  math.Abs(z - prev) < 1e-10{
             break
         }
         fmt.Println(z)    
     }    
         return z
}    

func main(){
     fmt.Println(Sqrt(2))
}    
