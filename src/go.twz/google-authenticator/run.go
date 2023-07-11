package main

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
	"os"
	"time"
)

func main() {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "TongWz Google Authenticator",
		AccountName: "www.sscf.gushi.com",
	})

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error generating key : %v \n", err)
		os.Exit(1)
	}

	// 打印密钥
	fmt.Printf("Key: %s \n", key.Secret())

	// 生成一个基于时间的一次性密码
	code, _ := totp.GenerateCode(key.Secret(), time.Now())
	fmt.Printf("Code: %s\n", code)

	fmt.Printf("isTrue: %#v \n", totp.Validate(code, key.Secret()))
	// 生成一个包含密钥信息的URL
	url := key.URL()
	fmt.Printf("URL: %s\n", url)

	// 将URL转换为QR码并展示
	err = qrcode.WriteFile(url, qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating QR code: %v\n", err)
		os.Exit(1)
	}
}
