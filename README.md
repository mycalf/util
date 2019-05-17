# util 工具集

#### ChineseNumber（ mode bool）string

##### 参数: mode bool = true：中文小写;  false：中文大写

阿拉伯数字转化为中文数字

正整数部分最大支持52位，小数部分最大支持52位。



```go
util.Text("999999999999999999999999999999999999999999999999999").ChineseNumber(false)
// Out:
// 玖佰玖拾玖极玖仟玖佰玖拾玖载玖仟玖佰玖拾玖正玖仟玖佰玖拾玖涧玖仟玖佰玖拾玖沟玖仟玖佰玖拾玖穰玖仟玖佰玖拾玖秭玖仟玖佰玖拾玖垓玖仟玖佰玖拾玖京玖仟玖佰玖拾玖兆玖仟玖佰玖拾玖亿玖仟玖佰玖拾玖万玖仟玖佰玖拾玖 


util.Text("12345").ChineseNumber(false)
// Out:
// 壹万贰仟叁佰肆拾伍
```

