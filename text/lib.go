package textutil

import (
	"math"
	"math/big"
	"regexp"
	"strings"
	"unicode"
)

// chineseFloat 阿拉伯数字小数点后的数字转为中文
func (t *Textutil) chineseFloat(mode bool) string {
	nums := t.floatSplit()
	n := chineseNums(mode)

	if nums == nil || len(nums) > 52 {
		return string(n[0])
	}

	var u []string
	for i := 0; i < len(nums); i++ {
		a := Text()
		a.Add(string(n[nums[i]]))
		u = append(u, a.Text)
	}

	if u == nil {
		return ""
	}

	a := Text()

	for i := len(u); i > 0; i-- {
		if u[i-1] != "" {
			a.Add(u[i-1])
		}
	}

	return a.Text

}

// ChineseInt ...
// 数字转汉字
// 参数为大小写开关
func (t *Textutil) chineseInt(mode bool) string {
	nums := t.intSplit()

	number := chineseNums(mode)

	if nums == nil || len(nums) > 52 {
		return string(number[0])
	}

	unit := chineseUnit(mode)

	var u []string

	for i := 0; i < len(nums); i++ {

		b := Text()

		if nums[i] != 0 {
			b.Add(string(number[nums[i]]))
		}

		if i%4 == 0 && i != 0 {
			// 万/亿/兆，单位
			if i/4 != 0 {
				n := 0

				if len(nums)-i > 4 {
					n = 4
				} else {
					n = len(nums) - i
				}

				e := 0

				for a := 0; a < n; a++ {
					e = e + nums[i+a]
				}

				if e != 0 {
					b.Add(string(unit[3+i/4]))
				}

			}
		} else if i%4 != 0 && nums[i] != 0 {
			b.Add(string(unit[i%4]))
		}

		e := 0
		z := 0

		// 当前数字不等于0
		// if nums[i] != 0 {
		for a := i; a >= 0; a-- {
			// 相同区间内
			if i/4 == a/4 && a != i {
				// 获取当前区间内有多少个连续的0
				if nums[a] == 0 {
					z++
				} else {
					break
				}
			}
		}

		//下一个区间是否存在连续的0
		for a := i; a >= 0; a-- {
			// 不相同区间内
			if i/4 != a/4 && a != i {
				// 获取当前区间内有多少个连续的0
				if nums[a] == 0 {
					e++
				} else {
					break
				}
			}
		}

		// 如果下一个区间开头存在0，那么则在大单位后加0
		// 如果当前数字区间内，当前数字之后存在0，则补0
		if i%4 == 0 && e > 0 && e < 4 || z > 0 && nums[i] != 0 && i%4 != 1 && z != i%4 {
			b.Add(string(number[0]))
		}
		u = append(u, b.Text)
	}

	if u == nil {
		return string(number[0])
	}

	b := Text()

	for i := len(u); i > 0; i-- {
		if u[i-1] != "" {
			b.Add(u[i-1])
		}
	}

	return b.Text
}

// floatSplit 浮点分割
func (t *Textutil) floatSplit() []int {
	a := t.intSplit()
	x, y := len([]rune(t.Text)), len(a)

	if x > y {
		for i := 0; i < x-y; i++ {
			a = append(a, 0)
		}
	}

	return a
}

// intSplit 整数分割
func (t *Textutil) intSplit() []int {

	if num, ok := big.NewInt(math.MaxInt64).SetString(strings.TrimLeft(t.Text, "0"), 0); ok {
		// fmt.Println(num)
		var nums []int
		ten := big.NewInt(10)
		for ; num.Cmp(big.NewInt(0)) > 0; num.Set(num.Div(num, ten)) {
			nums = append(nums, int(big.NewInt(0).Rem(num, ten).Int64()))
		}
		return nums
	}

	return nil
}

// 中文，数字
func chineseNums(mode bool) []rune {
	if mode == true {
		return []rune("〇一二三四五六七八九")
	}
	return []rune("零壹贰叁肆伍陆柒捌玖")
}

// 中文，单位
func chineseUnit(mode bool) []rune {
	if mode == true {
		return []rune(" 十百千万亿兆京垓秭穰沟涧正载极")
	}
	return []rune(" 拾佰仟万亿兆京垓秭穰沟涧正载极")
}

// 中文，点
func chineseDot(mode bool) string {
	if mode == true {
		return "点"
	}
	return "點"
}

// IsHan 判断是否为中文...
func (t *Textutil) IsHan(text string) bool {
	for _, r := range text {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
