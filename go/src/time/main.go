package main

import (
	"fmt"
	"time"
)
func main() {
	//var time1 int64 =2021/8/28
	the_time,_:= time.Parse("2006/01/02","2021/09/08")
	//fmt.Println(the_time)
		unix_time := the_time.Unix()
		//fmt.Println(unix_time)
		expirttime :=(unix_time-time.Now().Unix()  ) /86400
		fmt.Println(expirttime)
		//fmt.Println(time.Now().Unix())














}

