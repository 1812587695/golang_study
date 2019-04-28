package util

import (
	"github.com/mojocn/base64Captcha"
	"sync"
)

type CaptchaConfig struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

var (
	captchaConfig *CaptchaConfig
	captchaConfigOnce sync.Once
)

// 获取base64验证码基本配置
func GetCaptchaConfig() *CaptchaConfig {
	captchaConfigOnce.Do(func() {
		captchaConfig = &CaptchaConfig{
			Id:              "",
			CaptchaType:     "character",
			VerifyValue:     "",
			ConfigAudio:     base64Captcha.ConfigAudio{},
			ConfigCharacter: base64Captcha.ConfigCharacter{
				//Height:             60,
				//Width:              240,
				//Mode:               2,
				//IsUseSimpleFont:    false,
				//ComplexOfNoiseText: 0,
				//ComplexOfNoiseDot:  0,
				//IsShowHollowLine:   false,
				//IsShowNoiseDot:     false,
				//IsShowNoiseText:    false,
				//IsShowSlimeLine:    false,
				//IsShowSineLine:     false,
				//CaptchaLen:         0,
				Height:             60,
				Width:              240,
				//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
				Mode:               base64Captcha.CaptchaModeNumberAlphabet,
				ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
				ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
				IsShowHollowLine:   true,
				IsShowNoiseDot:     true,
				IsShowNoiseText:    true,
				//IsShowSlimeLine:    true,
				//IsShowSineLine:     true,
				CaptchaLen:         4,
			},
			ConfigDigit:     base64Captcha.ConfigDigit{},
		}
	})
	return captchaConfig
}
