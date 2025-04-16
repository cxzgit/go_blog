package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/model/response"
)

/*
*
API层
Captcha

	调用第三方 base64Captcha 生成验证码图片和 ID，并存储在内存中（DefaultMemStore）。
	返回 { captcha_id, pic_path } 给前端。

SendEmailVerificationCode

	先校验前端传来的图形验证码是否正确；
	正确后调用 baseService.SendEmailVerificationCode 发送邮件；
	失败时统一返回失败消息。

QQLoginURL

	直接从全局配置中获取 QQ 登录跳转链接并返回。
*/
type BaseApi struct {
}

var store = base64Captcha.DefaultMemStore

// Captcha 生成数字验证码
func (baseApi *BaseApi) Captcha(c *gin.Context) {
	// 创建数字验证码的驱动
	driver := base64Captcha.NewDriverDigit(
		global.Config.Captcha.Height,
		global.Config.Captcha.Width,
		global.Config.Captcha.Length,
		global.Config.Captcha.MaxSkew,
		global.Config.Captcha.DotCount,
	)

	// 创建验证码对象
	captcha := base64Captcha.NewCaptcha(driver, store)

	// 生成验证码
	id, b64s, _, err := captcha.Generate()

	if err != nil {
		global.Log.Error("Failed to generate captcha:", zap.Error(err))
		response.FailWithMessage("Failed to generate captcha", c)
		return
	}
	response.OkWithData(response.Captcha{
		CaptchaID: id,
		PicPath:   b64s,
	}, c)
}

// SendEmailVerificationCode 发送邮箱验证码
func (baseApi *BaseApi) SendEmailVerificationCode(c *gin.Context) {
	var req request.SendEmailVerificationCode
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(req.CaptchaID, req.Captcha, true) {
		err = baseService.SendEmailVerificationCode(c, req.Email)
		if err != nil {
			global.Log.Error("Failed to send email:", zap.Error(err))
			response.FailWithMessage("Failed to send email", c)
			return
		}
		response.OkWithMessage("Successfully sent email", c)
		return
	}
	response.FailWithMessage("Incorrect verification code", c)
}

// QQLoginURL 返回 QQ 登录链接
func (baseApi *BaseApi) QQLoginURL(c *gin.Context) {
	url := global.Config.QQ.QQLoginURL()
	response.OkWithData(url, c)
}
