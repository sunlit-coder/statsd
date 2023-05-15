####业务和性能监控系统

#####框架：
1. statsd：数据收集。
2. graphite：数据源，同时也可用作前端展示。
3. grafana：前端展示。

#####消息格式
`proj.module.func.<count | value | timing>`

* proj：项目名，如user，order等
* module：模块名，项目中的具体模块
* func：模块完成的功能，或者需监控的切入点，推荐以具体功能作为名称

#####备注
统计分为步进(Count)，数值(Gauge)和时长(Timing)

#####应用场景
* 登录次数：应为步进统计，消息格式可为 `user.login.count`
* 商品价格：应为数值统计，消息格式可为 `goods.price.value`
* 下单时长：应为时长统计，消息格式可为 `order.place.timing`

#####使用示例

	import "github.com/sunlit-coder/statsd"

	func login(){
		// user login

		// becase the prefix (usually set as project name) is retrieved from config file
		// so we just set the module & function name here
		statsd.Incr("user.login.count")
	}

	func getGoodsPrice() {
		// get goods price

		statsd.FGauge("goods.price.value", 100.5)
		// or
		statsd.Gauge("goods.price.value", 100)
	}

	func placeOrder() {
		t1 := statsd.Now()

		// heavy work

		t2 := statsd.Now()
		statsd.TimingByValue("order.place.timing", t2.Sub(t1))
		// or
		statsd.Timing("order.place.timing", t1, t2)
	}

#####后续
