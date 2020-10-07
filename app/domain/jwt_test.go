package domain

// FIXME:
// import (
// 	"os"
// 	"testing"
// 	"time"
// )

// func TestExtractToken(t *testing.T) {
// 	j := JWT{}
// 	id := 1
// 	at, err := j.CreateAccessToken(id)
// 	if err != nil {
// 		t.Errorf("%v", err)
// 	}
// 	token := j.ExtractToken("Bearer " + at)
// 	if at != token {
// 		t.Errorf("actual:%v expected:%v", token, at)
// 	}
// }

// func TestVerifyToken(t *testing.T) {
// 	cases := []struct {
// 		j        JWT
// 		sct      string
// 		expected bool
// 	}{
// 		{
// 			j:        JWT{},
// 			sct:      os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
// 			expected: true,
// 		},
// 		{
// 			// signature is invalid
// 			j:        JWT{},
// 			sct:      "fail",
// 			expected: false,
// 		},
// 		{
// 			// Token is expired
// 			j: JWT{
// 				// yesterday
// 				AtExpires: time.Now().AddDate(0, 0, -1).Unix(),
// 			},
// 			sct:      os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
// 			expected: false,
// 		},
// 	}

// 	id := 1
// 	for _, c := range cases {
// 		at, err := c.j.CreateAccessToken(id)
// 		if err != nil {
// 			t.Errorf("%v", err)
// 		}
// 		token := c.j.ExtractToken("Bearer " + at)

// 		_, err = c.j.VerifyToken(token, c.sct)

// 		if c.expected {
// 			if err != nil {
// 				t.Errorf("%v", err)
// 			}
// 		} else {
// 			if err == nil {
// 				t.Errorf("%v", "This test should fail")
// 			}
// 		}
// 	}
// }
