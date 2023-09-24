package cache

import (
	"VideoClipSystem/app/utils"
	"github.com/linyerun/GeeCache/cache"
)

func GetTimestampByAutoCode(key string) (val int64, ok bool) {
	view, err := authCodeGroup.Get(key)
	if err != nil {
		utils.Logger().Errorln(err)
		return 0, false
	}
	return utils.ByteArrayToInt64(view.ByteSlice()), true
}

func DeleteTimestampByAutoCode(key string) {
	err := authCodeGroup.Delete(key)
	if err != nil {
		utils.Logger().Error(err)
	}
}

func AddAutoCodeTimestamp(authCode string, timestamp int64) {
	err := authCodeGroup.AddOrUpdate(authCode, cache.NewByteView(utils.Int64ToByteArray(timestamp)))
	if err != nil {
		utils.Logger().Error(err)
	}
}

func GetAutoCodeByEmail(email string) (authCode string, ok bool) {
	view, err := emailGroup.Get(email)
	if err != nil {
		return
	}
	return view.String(), true
}

func DeleteAutoCodeByEmail(email string) {
	err := emailGroup.Delete(email)
	if err != nil {
		utils.Logger().Error(err)
	}
}

func AddEmailAuthCode(email, authCode string) {
	err := emailGroup.AddOrUpdate(email, cache.NewByteView([]byte(authCode)))
	if err != nil {
		utils.Logger().Error(err)
	}
}
