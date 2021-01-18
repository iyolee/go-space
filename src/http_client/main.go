package main

import (
	"fmt"
	"net/http"
)

func main() {
	httpNewRequest()
}

func httpNewRequest() {
	client := http.Client{}
	request, err := http.NewRequest("GET", "https://cnodejs.org/api/v1/topics/", nil)
	checkErr(err)
	cookName := &http.Cookie{Name: "username", Value: "lee"}
	request.AddCookie(cookName)
	response, err := client.Do(request)
	checkErr(err)
	request.Header.Set("Accept-Language", "zh-cn")
	defer response.Body.Close()
	fmt.Printf("Header: %+v\n", request.Header)
	fmt.Printf("响应状态码：%v\n", response.StatusCode)
	if response.StatusCode == 200 {
		fmt.Println("success")
		// data, err := ioutil.ReadAll(response.Body)
		// checkErr(err)
		// fmt.Println(string(data))
	}
}

func checkErr(err error) {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("error:", ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}
