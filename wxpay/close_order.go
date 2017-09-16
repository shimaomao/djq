package wxpay

//应用场景
//以下情况需要调用关单接口：商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
//注意：订单生成后不能马上调用关单接口，最短调用时间间隔为5分钟。
//接口地址
//接口链接：https://api.mch.weixin.qq.com/pay/closeorder
//是否需要证书
//不需要。
//请求参数
//字段名	变量名	必填	类型	示例值	描述
//公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
//商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
//商户订单号	out_trade_no	是	String(32)	1217752501201407033233368018	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
//随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
//签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
//签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5

//返回结果
//字段名	变量名	必填	类型	示例值	描述
//返回状态码	return_code	是	String(16)	SUCCESS	SUCCESS/FAIL
//返回信息	return_msg	否	String(128)	签名失败
//返回信息，如非空，为错误原因
//签名失败
//参数格式校验错误
//以下字段在return_code为SUCCESS的时候有返回
//
//字段名	变量名	必填	类型	示例值	描述
//公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID
//商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
//随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
//签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，验证签名算
//业务结果	result_code	是	String(16)	SUCCESS	SUCCESS/FAIL
//业务结果描述	result_msg	是	String(32)	OK	对于业务执行的详细描述
//错误代码	err_code	否	String(32)	SYSTEMERROR	详细参见第6节错误列表
//错误代码描述	err_code_des	否	String(128)	系统错误	结果信息描述

//错误码
//名称	描述	原因	解决方案
//ORDERPAID	订单已支付	订单已支付，不能发起关单	订单已支付，不能发起关单，请当作已支付的正常交易
//SYSTEMERROR	系统错误	系统错误	系统异常，请重新调用该API
//ORDERCLOSED	订单已关闭	订单已关闭，无法重复关闭	订单已关闭，无需继续调用
//SIGNERROR	签名错误	参数签名结果不正确	请检查签名参数和方法是否都符合签名算法要求
//REQUIRE_POST_METHOD	请使用post方法	未使用post传递参数 	请检查请求参数是否通过post方法提交
//XML_FORMAT_ERROR	XML格式错误	XML格式错误	请检查XML参数格式是否正确