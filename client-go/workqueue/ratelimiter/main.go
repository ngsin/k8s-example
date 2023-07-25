package main

import (
	"fmt"
	"k8s.io/client-go/util/workqueue"
	"time"
)

func main() {
	// 混合模式
	limiter := workqueue.DefaultControllerRateLimiter()
	// 默认排队指数算法 baseDelay: 1ms MaxDelay: 1000s
	//limiter := workqueue.DefaultItemBasedRateLimiter()
	// 令牌桶算法（每秒进去桶的令牌数，桶容量）
	//limiter := &workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(1), 1)}
	// 自定义排队指数法
	//limiter := workqueue.NewItemExponentialFailureRateLimiter(time.Millisecond*5, time.Second*100)
	// 计算器算法 failures 达到后切换到慢速率
	//limiter := workqueue.NewItemFastSlowRateLimiter(time.Millisecond*5, time.Second*5, 6)
	queue := workqueue.NewRateLimitingQueue(limiter)

	for {
		if queue.Len() == 0 {
			queue.AddRateLimited(9999)
		}
		if limiter.NumRequeues(9999) > 12 {
			queue.Forget(9999)
		}

		startTime := time.Now()
		item, _ := queue.Get()
		finishTime := time.Now()
		cost := finishTime.Sub(startTime)
		queue.Done(item)

		fmt.Printf("Cost: %v\tWhen %v\tNumRequeues:%v\n",
			cost.Seconds(), limiter.When(item), limiter.NumRequeues(item))
	}

}
