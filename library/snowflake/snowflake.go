package snowflake

import (
	"github.com/gogf/gf/util/gconv"
	"github.com/sony/sonyflake"
)

// GenerateSonyflakeID 生成雪花id
func GenerateSonyflakeID() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		return ""
	}
	return gconv.String(id)
}
