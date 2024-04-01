package helpers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHideEmail(t *testing.T) {
	items := map[string]string{
		"xxsazbanx.com":        "****",
		"1@zbanx.com":          "**@zbanx.com",
		"12@zbanx.com":         "**@zbanx.com",
		"123@zbanx.com":        "1**@zbanx.com",
		"1234@zbanx.com":       "12**@zbanx.com",
		"12345@zbanx.com":      "12***@zbanx.com",
		"123456@zbanx.com":     "12****@zbanx.com",
		"1234567@zbanx.com":    "12****@zbanx.com",
		"12345678@zbanx.com":   "12****78@zbanx.com",
		"123456789@zbanx.com":  "12****89@zbanx.com",
		"1234567890@zbanx.com": "12****90@zbanx.com",
	}

	for email, hide := range items {
		assert.Equal(t, hide, HideEmail(email), "hide email not equal:"+email)
	}
}

func ExampleThrottle() {
	var start = time.Now().Unix()
	var throttle = Throttle(3 * time.Second)
	for i := 0; i < 1000; i++ {
		throttle(func() {
			fmt.Println("Throttle:", time.Now().Unix()-start)
		})
		time.Sleep(10 * time.Millisecond)
	}
	// Output:
	//Throttle: 0
	//Throttle: 3
	//Throttle: 6
	//Throttle: 9
}
