package helpers

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"

	// "net/http"
	// "net/url"
	"strings"
)

func GetRandomOTP() string {
	maxDigits := 4
	bi, err := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(maxDigits)))),
	)
	if err != nil {
		panic(err)
	}
	newOtp := fmt.Sprintf("%0*d", maxDigits, bi)
	if strings.Contains(newOtp, "0") {
		GetRandomOTP()
	}
	return newOtp
}

func GetSixDigitRandomString() string {
	maxDigits := 6
	bi, err := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(maxDigits)))),
	)
	if err != nil {
		panic(err)
	}
	newOtp := fmt.Sprintf("%0*d", maxDigits, bi)
	if strings.Contains(newOtp, "0") {
		GetSixDigitRandomString()
	}
	return newOtp
}

func SendLoginOtp(MobileNo string) string {
	otp := GetRandomOTP()
	// baseUrl := "https://40bb89adf1be339bf74463420b0d83e41f10537d628845ba:16b50842dc77883a058b3aabd84aac2484bcec37baf98761@api.exotel.com/v1/Accounts/readyassist/Sms/send?Body="
	// bodyString := "Your OTP verification code is " + otp + ". Do not share it with anyone. ReadyAssist.&From=RDYAST&To=0" + MobileNo
	// apiUrl := baseUrl + bodyString
	// data := url.Values{}
	// req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return ""
	// }
	// client := &http.Client{}
	// _, err = client.Do(req)
	return otp
}
