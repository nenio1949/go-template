package common

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// 自定义字段类型

/****************************8日期****************************/
type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte(fmt.Sprintf("\"%v\"", "")), nil
	} else {
		tTime := time.Time(t.Time)
		return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
	}
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t.Time)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

/******************************日期*********************************/

/*******************************字符数组****************************/
type Strs []string

func (m *Strs) Scan(val interface{}) error {
	s := val.([]uint8)
	ss := strings.Split(string(s), "|")
	*m = ss
	return nil
}

func (m Strs) Value() (driver.Value, error) {
	str := strings.Join(m, "|")
	return str, nil
}

/*******************************字符数组****************************/
