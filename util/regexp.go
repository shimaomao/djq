package util

import (
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	RegexpChinese     = "\u4e00-\u9fa5"
	ErrNameFormat     = errors.New("名称为英文或数字或下划线，不能以数字开头长度为4到32")
	ErrMobileFormat   = errors.New("手机号码为11为数字")
	ErrPasswordFormat = errors.New("密码长度为6到32")
	ErrPriorityFormat = errors.New("优先权重范围为0到999")
	ErrWeightFormat   = errors.New("抽奖权重范围为0到999")
)

func MatchNonnegativeNumberWithErr(number int, name string) error {
	if number >= 0 {
		return nil
	}
	return errors.New(name + "应为非负数:" + strconv.Itoa(number))
}

func MatchPriority(priority int) error {
	if 0 <= priority && priority <= 999 {
		return nil
	}
	return ErrPriorityFormat
}

func MatchWeight(weight int) error {
	if 0 <= weight && weight <= 999 {
		return nil
	}
	return ErrWeightFormat
}

func MatchLenWithErr(str string, minInclude int, maxInclude int, name string) error {
	if !MatchLen(str, minInclude, maxInclude) {
		return errors.New(name + "长度为" + strconv.Itoa(minInclude) + "到" + strconv.Itoa(maxInclude) + ":" + str)
	}
	return nil
}

func MatchLen(str string, minInclude int, maxInclude int) bool {
	len := strings.Count(str, "") - 1
	return minInclude <= len && len <= maxInclude
}

func MatchDescription(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[\\s\\S]{0,200}$", s)
		if false == b {
			return b
		}
	}
	return b
}

func MatchName(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[a-zA-Z_][a-zA-Z0-9_]{3,31}$", s)
		if false == b {
			return b
		}
	}
	return b
}

func MatchMobile(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[1][0-9]{10}$", s)
		if false == b {
			return b
		}
	}
	return b
}

func MatchPassword(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[\\s\\S]{6,32}$", s)
		if false == b {
			return b
		}
	}
	return b
}
