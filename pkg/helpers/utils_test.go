package helpers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHideEmail(t *testing.T) {
	items := map[string]string{
		"xxsaexample.com":        "****",
		"1@example.com":          "**@example.com",
		"12@example.com":         "**@example.com",
		"123@example.com":        "1**@example.com",
		"1234@example.com":       "12**@example.com",
		"12345@example.com":      "12***@example.com",
		"123456@example.com":     "12****@example.com",
		"1234567@example.com":    "12****@example.com",
		"12345678@example.com":   "12****78@example.com",
		"123456789@example.com":  "12****89@example.com",
		"1234567890@example.com": "12****90@example.com",
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
