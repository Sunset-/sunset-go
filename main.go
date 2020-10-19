package main

import "github.com/Sunset-/sunset-go/utils/validates"

type User struct {
	Name string `validate:"required" json:"name"`
	Age  int
}

func main() {
	//fmt.Println(casts.FloatToInt(casts.StrToFloat64("-32.663")))
	//
	//r,_ := regexp.Compile("^\\S+$")
	//fmt.Println(r.MatchString("123"))
	user := &User{Name: "tom", Age: 123}
	validates.ValidateStruct(user)
}
