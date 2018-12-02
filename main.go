package main

import (
	"fmt"
	"reflect"

	fb "github.com/huandu/facebook"
)

const ac = "EAAa6VkKVZBVUBAKZAZBpHTSn1nVG6ZBRLOyVWrAFPmBe7AJWr6RAH64fq7kaMHifjjTx8bCoIelxC3860gLRO7AozL1vZB3vPlZAyNzxmCS7Ln4ehhKgbOBm8LOyRx7ZB1UOlIHpddQZCZAir6BwC1JR1DbiUg5JNrH6HmdQYQbv5wTFGroNt9mSlW1SRzXCz1bEZD"

func main() {
	res, _ := fb.Get("582313518881669_582751342171220/comments", fb.Params{
		"fields":       "message",
		"access_token": ac,
	})

	s, _ := res["data"].([]interface{})

	for i := 0; i < reflect.ValueOf(s).Len(); i++ {

		msg := s[i].(map[string]interface{})

		fmt.Println(msg["message"])
	}

	fmt.Println("")

	res2, _ := fb.Get("582313518881669_582750698837951/comments", fb.Params{
		"fields":       "message",
		"access_token": ac})

	r, _ := res2["data"].([]interface{})

	for i := 0; i < reflect.ValueOf(r).Len(); i++ {

		msg := r[i].(map[string]interface{})

		fmt.Println(msg["message"])
	}

	res3, _ := fb.Get("582313518881669_582751015504586/comments", fb.Params{
		"fields":       "message",
		"access_token": ac})

	v, _ := res3["data"].([]interface{})

	for i := 0; i < reflect.ValueOf(v).Len(); i++ {

		msg := v[i].(map[string]interface{})

		fmt.Println(msg["message"])
	}
}
