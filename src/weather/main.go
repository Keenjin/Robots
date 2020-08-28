package main

import (
	"fmt"
	"github.com/keenjin/gomini/khttp"
	"github.com/robfig/cron"
	"github.com/tidwall/gjson"
	"strings"
)

func sendWeatherToQiyeWechat() {
	resp, err := khttp.KHttpPostJson("http://www.tianqiapi.com/api?version=v1&appid=94393695&appsecret=Di8BvZ4c", "")
	if err != nil {
		return
	}
	fmt.Println(resp)

	// 发送给企业微信机器人
	msg := `{
			"msgtype": "markdown",
			"markdown": {
				"content": "${city}今日（${week}）天气：${tem}（${tem_l}~${tem_h}），${wea}，${air_tips}\n> ${week1}天气：${tem1}，${wea1}\n> ${week2}天气：${tem2}，${wea2}\n> ${week3}天气：${tem3}，${wea3}",
				"mentioned_list": ["@all"]
			}
		}`

	city := gjson.Get(resp, "city").String()
	week := gjson.Get(resp, "data.0.week").String()
	tem := gjson.Get(resp, "data.0.tem").String()
	tem_l := gjson.Get(resp, "data.0.tem2").String()
	tem_h := gjson.Get(resp, "data.0.tem1").String()
	wea := gjson.Get(resp, "data.0.wea").String()
	air_tips := gjson.Get(resp, "data.0.air_tips").String()

	week1 := gjson.Get(resp, "data.1.week").String()
	tem1 := gjson.Get(resp, "data.1.tem").String()
	wea1 := gjson.Get(resp, "data.1.wea").String()

	week2 := gjson.Get(resp, "data.2.week").String()
	tem2 := gjson.Get(resp, "data.2.tem").String()
	wea2 := gjson.Get(resp, "data.2.wea").String()

	week3 := gjson.Get(resp, "data.3.week").String()
	tem3 := gjson.Get(resp, "data.3.tem").String()
	wea3 := gjson.Get(resp, "data.3.wea").String()

	msg = strings.ReplaceAll(msg, "${city}", city)
	msg = strings.ReplaceAll(msg, "${week}", week)
	msg = strings.ReplaceAll(msg, "${tem}", tem)
	msg = strings.ReplaceAll(msg, "${tem_l}", tem_l)
	msg = strings.ReplaceAll(msg, "${tem_h}", tem_h)
	msg = strings.ReplaceAll(msg, "${wea}", wea)
	msg = strings.ReplaceAll(msg, "${air_tips}", air_tips)
	msg = strings.ReplaceAll(msg, "${week1}", week1)
	msg = strings.ReplaceAll(msg, "${tem1}", tem1)
	msg = strings.ReplaceAll(msg, "${wea1}", wea1)
	msg = strings.ReplaceAll(msg, "${week2}", week2)
	msg = strings.ReplaceAll(msg, "${tem2}", tem2)
	msg = strings.ReplaceAll(msg, "${wea2}", wea2)
	msg = strings.ReplaceAll(msg, "${week3}", week3)
	msg = strings.ReplaceAll(msg, "${tem3}", tem3)
	msg = strings.ReplaceAll(msg, "${wea3}", wea3)

	_, err = khttp.KHttpPostJson("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=9a0164d3-dd1b-4415-a02c-646699d71bb3", msg)
	if err != nil {
		return
	}
}

func main() {
	c := cron.New()
	err := c.AddFunc(`0 0 9 * * ?`, func() {
		sendWeatherToQiyeWechat()
	})
	if err != nil {
		return
	}

	c.Start()

	select {}
}
