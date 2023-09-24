package utils

func IsTimeoutCheckoutByMilli(pre, expireTime, cur int64) bool {
	return pre+expireTime < cur // true: 过期, false: 正常
}
