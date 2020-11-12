package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sunset-/sunset-go/components/metricss"
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"regexp"
	"time"
	"unsafe"
)


//return GoString's buffer slice(enable modify string)
func StringBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

type Person struct{
	Name string
}

func main() {

	p := &Person{Name : "jack"}

	//pTemp :=

	pTempP := &*p

	pTempP.Name = "tom"

	fmt.Println(pTempP)

	outLoop:
		for i:=0;i<10;i++{
			for j:=0;j<10;j++{
				if j==2{
					continue outLoop
				}
				fmt.Println(i,j)
			}
	}


	//extra.RegisterFuzzyDecoders()
	var subscribeID map[string]string
	subscribeID = make(map[string]string)
	subscribeIdJson := `{"a13":"eqweqwewqeqweqweqwewqewqeqw"}`
	err := json.Unmarshal([]byte(subscribeIdJson), &subscribeID)
	fmt.Println(err)

	str := `([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼]{1}(([A-HJ-Z]{1}[A-HJ-NP-Z0-9]{5})|([A-HJ-Z]{1}(([DF]{1}[A-HJ-NP-Z0-9]{1}[0-9]{4})|([0-9]{5}[DF]{1})))|([A-HJ-Z]{1}[A-D0-9]{1}[0-9]{3}警)))|([0-9]{6}使)|((([沪粤川云桂鄂陕蒙藏黑辽渝]{1}A)|鲁B|闽D|蒙E|蒙H)[0-9]{4}领)|(WJ[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼·•]{1}[0-9]{4}[TDSHBXJ0-9]{1})|([VKHBSLJNGCE]{1}[A-DJ-PR-TVY]{1}[0-9]{5})`
	re,err := regexp.Compile(str)
	fmt.Println(err)
	fmt.Println(re.MatchString("12321陕AW31134"))

	timer := metricss.MonitorTimer("DagNodeMonitor", "abc")
	go func() {
		ticker := time.NewTicker(time.Millisecond * 10)
		for {
			<-ticker.C
			timer.Update(100*time.Millisecond)
		}
	}()

	go metrics.Log(metricss.Registry("DagNodeMonitor"),3*time.Second,log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	select {}
	//fmt.Println(casts.FloatToInt(casts.StrToFloat64("-32.663")))
	//
	//r,_ := regexp.Compile("^\\S+$")
	//fmt.Println(r.MatchString("123"))
	//user := &User{Name: "tom", Age: 123}
	//validates.ValidateStruct(user)
}
