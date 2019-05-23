# util 工具集

#### Text（ str ...string）*Textutil

##### 参数: str string

*Text类初始化*

```go
t := util.Text() 
t.Text // nil
t := util.Text("Hello Word!") 
t.Text // Hello Word!
```



#### (Text)  Add（ str string）*Textutil

##### 参数: str string

*在文本中追加文字*

```go
t := util.Text().Add("Hello Word!")
t.Text // Hello Word!

t := util.Text().Add("Hello").Add(" Word!")
t.Text // Hello Word!

t := util.Text("Hello ").Add("Word!")
t.Text // Hello Word!
```



#### (Text) Parse（ data interface{}）string

##### 参数: data interface{}

*根据模版解析文本*

```go
//Text : Hell{{.d}} Wor{{.o}}
//data : map[string]string{ "d": "o", "o":"d" }

//Text : Hello {{.}}
//data : Word*

//Text : {{.H}} {{.W}}*
//&data{H: "Hello", W: "Word"}


util.Text("Hell{{.d}}").Add(" Wor{{.o}}").Parse(map[string]string{"d": "o", "o": "d"})
// Out: Hello Word

```



#### (Text) ChineseNumber（ mode bool）string

##### 参数: mode bool = true：中文小写;  false：中文大写

*阿拉伯数字转化为中文数字*

*正整数部分最大支持52位，小数部分最大支持52位。*



```go
util.Text("999999999999999999999999999999999999999999999999999").ChineseNumber(false)
// Out:
// 玖佰玖拾玖极玖仟玖佰玖拾玖载玖仟玖佰玖拾玖正玖仟玖佰玖拾玖涧玖仟玖佰玖拾玖沟玖仟玖佰玖拾玖穰玖仟玖佰玖拾玖秭玖仟玖佰玖拾玖垓玖仟玖佰玖拾玖京玖仟玖佰玖拾玖兆玖仟玖佰玖拾玖亿玖仟玖佰玖拾玖万玖仟玖佰玖拾玖 


util.Text("12345.54321").ChineseNumber(false)
// Out:
// 壹万贰仟叁佰肆拾伍點伍肆叁贰壹


util.Text("12345.54321").ChineseNumber(true)
// Out:
// 一万二千三百四十五点五四三二一

```

