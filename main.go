package main

import (
	"fmt"
	"github.com/Sunset-/sunset-go/utils/casts"
)

func main() {
	fmt.Println(casts.FloatToInt(casts.StrToFloat64("-32.663")))

}
