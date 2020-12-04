package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/Sunset-/sunset-go/utils/casts"
	"github.com/influxdata/influxdb/client/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/shirou/gopsutil/net"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	MyDB     = "dag_monitor"
	username = "root"
	password = "Netposa123"
)

func getIo() map[string][]uint64{
	m := make(map[string][]uint64)
	stats,_ := net.IOCounters(true)
	for _,s := range stats{
		if s.Name!="以太网"{
			continue
		}
		m[s.Name] = []uint64{s.BytesRecv,s.BytesSent}
	}

}

func main() {
	conn, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://172.16.129.124:18086",
		Username: username,
		Password: password,
		Timeout:3*time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn)

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "ns",
		//RetentionPolicy: "one_week",
	})

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()


	t := time.Now()
	for i:=0;i<1;i++{
		tags := map[string]string{"nodeId": "n123445678","category": "abc"}
		fields := map[string]interface{}{
			"min":  1.0,
			"max":  10.0,
			"mean": float64(i),
			"p50":  1.8,
		}
		pt, err := client.NewPoint("test_mean", tags, fields, t)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}
	for{
		if err := conn.Write(bp); err != nil {
			fmt.Println(err)
			time.Sleep(5*time.Second)
			continue
		}
		break
	}
	fmt.Println("cost:",time.Since(start))
	return
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   false, //false 长链接 true 短连接
			Proxy:               http.ProxyFromEnvironment,
			MaxIdleConns:        20, //client对与所有host最大空闲连接数总和
			MaxConnsPerHost:     10,
			MaxIdleConnsPerHost: 10,               //连接池对每个host的最大连接数量,当超出这个范围时，客户端会主动关闭到连接
			IdleConnTimeout:     60 * time.Second, //空闲连接在连接池中的超时时间
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 5 * time.Second, //粗粒度 时间计算包括从连接(Dial)到读完response body
	}
	url := `http://14.122.0.12:9000/cpbs/download?picid=%40ft%3Djpg%26sz%3D904475%26st%3Dc%26cid%3DgroupId_test%26fp%3D%252Fhome%252Fcloud%252F20201113%252F3_20201113192355_042190356905100000050.cts%26ro%3D3920390%26po%3D3921893`

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("ERR1:", err)
		return
	}
	req.Header.Set("Connection", "keep-alive")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("ERR2:", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println("ERR3:", res.StatusCode)
	}
	resBytes, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		fmt.Println("ERR4:", err)
	}
	fmt.Println("LEN:", len(resBytes))
	err = ioutil.WriteFile("./"+casts.Int64ToStr(time.Now().Unix())+".jpg", resBytes, 0666)
	if res.StatusCode != http.StatusOK {
		fmt.Println("ERR5:", err)
	}
	return

}
