package cache

import (
	"errors"
	"github.com/linyerun/GeeCache/GeeCache"
)

var authCodeGroup GeeCache.IGroup // key: auth_code, value: timestamp
var emailGroup GeeCache.IGroup    // key: email, value: auth_code

func Init() {
	authCodeGroup = GeeCache.NewGroupByGetterFunc("auto_codes", (1<<20)*50, func(key string) ([]byte, error) {
		return nil, errors.New("can not find key=" + key)
	}, nil)

	emailGroup = GeeCache.NewGroupByGetterFunc("emails", (1<<20)*50, func(key string) ([]byte, error) {
		return nil, errors.New("can not find key=" + key)
	}, nil)
}
