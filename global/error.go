package global

import "fmt"

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError CustomError
	ValidateError CustomError
	TokenError    CustomError
}

var Errors = CustomErrors{
	BusinessError: CustomError{400, "业务错误"},
	ValidateError: CustomError{422, "请求参数错误"},
	TokenError:    CustomError{401, "登录授权失效"},
}

func (customErrors *CustomErrors) DoError() {
	fmt.Println("开始处理异常")
	// 获取异常信息
	if err := recover(); err != nil {
		//  输出异常信息
		App.Log.Error(err.(string))
	}
	fmt.Println("结束异常处理")
}
