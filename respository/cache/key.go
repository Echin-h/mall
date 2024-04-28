package cache

import (
	"fmt"
	"strconv"
)

const (
	SkillProductKey     = "skill:product:%d"
	SkillProductListKey = "skill:product_list"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
