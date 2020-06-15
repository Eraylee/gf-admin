package boot

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
