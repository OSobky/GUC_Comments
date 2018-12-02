package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

func main() {
	res, _ := fb.Get("582313518881669_582751342171220/comments", fb.Params{
		"fields":       "message",
		"access_token": "EAAFDrTDhvyMBAE7A6bIPt92NdOC0RVqDaHAGVbiWJeLNDATBZACvjaKZCOaNn9W2ZBAtNYQOPgouDVgGG1hcmcfESgts8KBSCNz7ZABzzqShi4omU2wyED1kBqieIClFTSMM88yHwaHidME04ajLpZCG311tSEjrOZBsaK0HLJPiGt8V6bEFtJb9ZAl2G37z54ZD",
	})
	//var message string
	x, mOk := res["data"].([]map[string]interface{})
	s, sOk := res["data"].([]interface{})

	if mOk {
		y := x[0]
		fmt.Println(y["message"])
	} else {
		if sOk {
			msg := s[0].(map[string]interface{})
			fmt.Println(msg["message"])
		}
	}

	fmt.Println()
	res2, _ := fb.Get("582313518881669_582750698837951/comments", fb.Params{
		"fields":       "message",
		"access_token": "EAAFDrTDhvyMBAE7A6bIPt92NdOC0RVqDaHAGVbiWJeLNDATBZACvjaKZCOaNn9W2ZBAtNYQOPgouDVgGG1hcmcfESgts8KBSCNz7ZABzzqShi4omU2wyED1kBqieIClFTSMM88yHwaHidME04ajLpZCG311tSEjrOZBsaK0HLJPiGt8V6bEFtJb9ZAl2G37z54ZD",
	})
	e, mOk := res2["data"].([]map[string]interface{})
	r, sOk := res2["data"].([]interface{})

	if mOk {
		y := e[0]
		fmt.Println(y["message"])
	} else {
		if sOk {
			msg := r[0].(map[string]interface{})
			fmt.Println(msg["message"])
		}
	}

	res3, _ := fb.Get("582313518881669_582751015504586/comments", fb.Params{
		"fields":       "message",
		"access_token": "EAAFDrTDhvyMBAE7A6bIPt92NdOC0RVqDaHAGVbiWJeLNDATBZACvjaKZCOaNn9W2ZBAtNYQOPgouDVgGG1hcmcfESgts8KBSCNz7ZABzzqShi4omU2wyED1kBqieIClFTSMM88yHwaHidME04ajLpZCG311tSEjrOZBsaK0HLJPiGt8V6bEFtJb9ZAl2G37z54ZD",
	})
	c, mOk := res3["data"].([]map[string]interface{})
	v, sOk := res3["data"].([]interface{})

	if mOk {
		y := c[0]
		fmt.Println(y["message"])
	} else {
		if sOk {
			msg := v[0].(map[string]interface{})
			fmt.Println(msg["message"])
		}
	}
}
