package main

import (
	"fmt"
	"time"
)

type HandlerFunc func(ctx *Context)

type FilterBuilder func(next Filter) Filter

type Filter func(ctx *Context)

// 验证类型是否正确
var _ FilterBuilder = MetricFilterBuilder

func MetricFilterBuilder(next Filter) Filter {
	return func(ctx *Context) {
		start := time.Now().Nanosecond()
		next(ctx)
		end := time.Now().Nanosecond()
		fmt.Printf("用了 %d 时间", end-start)
	}
}
