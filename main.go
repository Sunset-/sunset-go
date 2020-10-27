package main

import (
	"fmt"
)

type User struct {
	Name string `validate:"required" json:"name"`
	Age  int
}

func (u *User) setName(a string){
	u.Name = a
}
func test() (ret int) {
	ret = 1
	defer func() {
		ret += 100
	}()
	return ret
}

func main() {

	fmt.Println((test()))
	c := make(chan int ,10)

	c<-1
	c<-1
	c<-1
	c<-1
	close(c)
	for {
		t,ok := <-c
		fmt.Println(t,ok)
		if !ok{
			break
		}
	}

	fmt.Println("END")

	//fmt.Println(casts.FloatToInt(casts.StrToFloat64("-32.663")))
	//
	//r,_ := regexp.Compile("^\\S+$")
	//fmt.Println(r.MatchString("123"))
	//user := &User{Name: "tom", Age: 123}
	//validates.ValidateStruct(user)
}
