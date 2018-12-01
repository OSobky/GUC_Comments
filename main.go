package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

func main() {
	res, _ := fb.Get("582313518881669_582751342171220/comments", fb.Params{
		"fields":       "message",
		"access_token": "EAAFDrTDhvyMBACnZBa6YToQngCZBpgk4DmMNto5FfNmFBUKpSGvgRhGIFVFvVaaRcvG3Tx3MsKZAvq5V25FZBzvHjng1ZAj6nOC4EVqMlNrwGHChe3EtTLDZBWUFC15P5RvTkhZAE6AUYMU0Jt15wsvVRXYF4ZCVFxVGaMvZCxkinD5qD154sCZASz9QptYBcMdZA4ZD",
	})
	fmt.Println(res)
	fmt.Println()
	fmt.Println()

	res2, _ := fb.Get("582313518881669_582750698837951/comments", fb.Params{
		"fields":       "message",
		"access_token": "EAAFDrTDhvyMBACnZBa6YToQngCZBpgk4DmMNto5FfNmFBUKpSGvgRhGIFVFvVaaRcvG3Tx3MsKZAvq5V25FZBzvHjng1ZAj6nOC4EVqMlNrwGHChe3EtTLDZBWUFC15P5RvTkhZAE6AUYMU0Jt15wsvVRXYF4ZCVFxVGaMvZCxkinD5qD154sCZASz9QptYBcMdZA4ZD",
	})
	fmt.Println(res2)
	fmt.Println()
	fmt.Println()

	res3, _ := fb.Get("582313518881669_582751015504586/comments", fb.Params{
		"fields":       "message",
		"access_token": "EAAFDrTDhvyMBACnZBa6YToQngCZBpgk4DmMNto5FfNmFBUKpSGvgRhGIFVFvVaaRcvG3Tx3MsKZAvq5V25FZBzvHjng1ZAj6nOC4EVqMlNrwGHChe3EtTLDZBWUFC15P5RvTkhZAE6AUYMU0Jt15wsvVRXYF4ZCVFxVGaMvZCxkinD5qD154sCZASz9QptYBcMdZA4ZD",
	})
	fmt.Println(res3)
}
