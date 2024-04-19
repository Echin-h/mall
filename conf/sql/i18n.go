package sql

import (
	"strings"
)

// 校验器的国际化 -- 适配不同国家的语法
var Dictinary *map[interface{}]interface{}

//func LoadLocales(path string) error{
//	data , err := ioutil.ReadDir(path)
//	if err != nil {
//		return err
//	}
//	m := make(map[interface{}]interface{})
//	err = yaml.Unmarshal([]byte(data), &m)
//	if err != nil {
//		return err
//	}
//	Dictinary = &m
//	return nil
//}

// 自定义翻译函数
func T(key string) string {
	dic := *Dictinary
	keys := strings.Split(key, ".")
	for index, path := range keys {
		if len(keys) == (index + 1) {
			for k, v := range dic {
				if k, ok := k.(string); ok {
					if k == path {
						if value, ok := v.(string); ok {
							return value
						}
					}
				}
			}
			return path
		}
		for k, v := range dic {
			if ks, ok := k.(string); ok {
				if ks == path {
					if dic, ok = v.(map[interface{}]interface{}); !ok {
						return path
					}
				}
			} else {
				return ""
			}
		}
	}
	return ""
}
