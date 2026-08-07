package main

import(
     "fmt"
     "runtime"
)	

func main(){
     fmt.Print("Go runs on ")
     switch os:=runtime.GOOS;os{
     case "darwin":
	   fmt.Println("MacOS.")
     case "linux":
	   fmt.Println("linux.")
     default:
           //freebad,openbsd,
	   //plan9,windows
	   fmt.Printf("%s\n",os)
     }	   
}
