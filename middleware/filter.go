package middleware

import (
	"fmt"
	"go-web/common"
	"time"
)

type HandlerFunc func(ctx *common.Context)

type FilterBuilder func(next Filter) Filter

type Filter func(ctx *common.Context)

// 验证类型是否正确
var _ FilterBuilder = MetricFilterBuilder

func MetricFilterBuilder(next Filter) Filter {
	return func(ctx *common.Context) {
		start := time.Now().Nanosecond()
		next(ctx)
		end := time.Now().Nanosecond()
		fmt.Printf("用了 %d 时间", end-start)
	}
}
