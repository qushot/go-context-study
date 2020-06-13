package main

import (
	"context"
	"fmt"
	"time"
)

type ctxKey string

var key ctxKey = "key"

func infiniteLoop(ctx context.Context) {
	innerCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	for {
		fmt.Printf("Help! [%s]\n", ctx.Value(key).(ctxKey))
		select {
		case <-innerCtx.Done():
			fmt.Println("Exit from hell.")
			return
		}
	}
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, key, "value")

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	go infiniteLoop(ctx)

	//ctx.Doneが飛んでくるまでmainプロセスの終了をブロックしている
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// WithDeadlineでのキャンセルは「2秒後をデッドラインにする」となっているので、絶対指定感は薄いが、
// time.Date(2020, time.June, 20, 10, 0, 0, 0, time.Local) のようなことができるという話
