package main

import (
	"log"
	"net/url"
	"strconv"
)

func main() {
	q := url.Values{}
	q.Add("text", "text1")
	q.Add("text", "text2")
	q.Add("num", strconv.Itoa(2)) 
	log.Println(q["num"])
	
}