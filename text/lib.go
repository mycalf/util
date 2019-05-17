package textutil

import (
	"math"
	"math/big"
	"regexp"
	"unicode"
)

// NumberToChinese ...
// 数字转汉字
// 第二个参数为大小写开关
func NumberToChinese(str string) string {
	var nums []int
	ten := big.NewInt(10)

	num, ok := big.NewInt(math.MaxInt64).SetString(str, 0)
	if ok {
		for ; num.Cmp(big.NewInt(0)) > 0; num.Set(num.Div(num, ten)) {
			nums = append(nums, int(big.NewInt(0).Rem(num, ten).Int64()))
		}
	}

	if len(nums) > 52 {
		return "超出可处理范围！"
	}

	number := []rune("零壹贰叁肆伍陆柒捌玖")
	// number := []rune("零一二三四五六七八九")
	unit := []rune(" 拾佰仟万亿兆京垓秭穰沟涧正载极")
	// unit := []rune(" 十百千万亿兆")

	var u []string

	for i := 0; i < len(nums); i++ {

		t := Text()

		if nums[i] != 0 {
			t.Add(string(number[nums[i]]))
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
					t.Add(string(unit[3+i/4]))
				}

			}
		} else {
			if nums[i] != 0 {
				t.Add(string(unit[i%4]))
			}
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
			t.Add(string(number[0]))
		}
		u = append(u, t.Text)
	}

	t := Text()

	for i := len(u); i > 0; i-- {
		if u[i-1] != "" {
			t.Add(u[i-1])
		}
	}

	return t.Text
}

// IsHan 判断是否为中文...
func IsHan(text string) bool {
	for _, r := range text {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
