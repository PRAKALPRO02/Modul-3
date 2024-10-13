package main

import (
	"fmt"
	"math"
)
func jarak(x1,y1,x2, y2 int) float64{
  return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

func main(){
  var cx1, cy1, r1 int
  var cx2, cy2, r2 int
  var x,y int

  fmt.Scan(&cx1,&cy1, &r1)
  fmt.Scan(&cx2,&cy2, &r2)
  fmt.Scan(&x,&y)

  jarakLingkaran1 := jarak(x,y,cx1,cy1)
  jarakLingkaran2 := jarak(x,y,cx2,cy2)

  if jarakLingkaran1 <= float64(r1) && jarakLingkaran2 <= float64(r2) {
    fmt.Print("Titik di dalam lingkaran 1 dan 2")
    }else if jarakLingkaran1 <= float64(r1){
      fmt.Print("Titik di dalam lingkaran 1")
    }else if jarakLingkaran2 <= float64(r2){
      fmt.Print("Titik di dalam lingkaran 2")
    }else{
      fmt.Print("Titik di luar lingkaran 1 dan 2")
    }
}

