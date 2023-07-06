package gtime

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/lytdev/go-mykit/gconv"
)

const (
	MYNano      = "2006-01-02 15:04:05.000000000"
	MYMicro     = "2006-01-02 15:04:05.000000"
	MYMil       = "2006-01-02 15:04:05.000"
	MYSec       = "2006-01-02 15:04:05"
	MYCST       = "2006-01-02 15:04:05 +0800 CST"
	MYUTC       = "2006-01-02 15:04:05 +0000 UTC"
	MYDate      = "2006-01-02"
	MYTime      = "15:04:05"
	FBTIME      = "2006-01-02T15:04:05+0800"
	APPTIME     = "2006-01-02T15:04:05.000"
	TWITTERTIME = "2006-01-02T15:04:05Z"
)

var TimeLocation *time.Location

func init() {
	var err error
	TimeLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}

// GetCurrentTime 当前时区的当前时间
func GetCurrentTime() time.Time {
	return time.Now().In(TimeLocation)
}

func getTimeDefault() time.Time {
	t, _ := time.ParseInLocation(MYSec, "", TimeLocation)
	return t
}

// TimeToStrAsFormat 按照指定的格式输出时间
func TimeToStrAsFormat(t time.Time, timeFormat string) string {
	// 先将输入的时间转换到指定的时区，然后再转换格式
	return t.In(TimeLocation).Format(timeFormat)
}

// TimeStrToTime 时间字符串转时间
func TimeStrToTime(timeStr string) (time.Time, error) {
	// 可能的转换格式
	useFormat := []string{
		MYNano, MYMicro, MYMil, MYSec, MYCST, MYUTC, MYDate, MYTime, FBTIME, APPTIME, TWITTERTIME,
		time.RFC3339,
		time.RFC3339Nano,
	}
	var t time.Time
	for _, useF := range useFormat {
		tt, err1 := time.ParseInLocation(useF, timeStr, TimeLocation)
		if err1 != nil {
			continue
		}
		t = tt
		break
	}
	if t == getTimeDefault() { // 0001-01-01 00:00:00 +0000 UTC
		return t, errors.New("时间字符串格式错误")
	}
	return t, nil
}

///////////////////////时间的时区转换//////////////////////

// TimeToUTC 本地时区时间与UTC时区时间转换
func TimeToUTC(t time.Time) time.Time {
	// 时间转换成 UTC时区的时间
	return t.UTC()
}

// Time2ToLocal 转成本地时区时间
func Time2ToLocal(t time.Time) time.Time {
	return t.Local()
}

//////////////////时间戳与时间的相关转换/////////////////////////

// TimeToTimeStampSecond 时间转秒级别时间戳
func TimeToTimeStampSecond(t time.Time) int64 {
	return t.Unix()
}

// TimeToTimeStampNano 时间转纳秒级别时间戳
func TimeToTimeStampNano(t time.Time) int64 {
	return t.UnixNano()
}

// TimeToTimeStampMill 时间转毫秒级别时间戳
func TimeToTimeStampMill(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// timestampToTime 时间戳转时间
func timestampToTime(stamp int64, nsec int64) time.Time {
	return time.Unix(stamp, nsec)
}

// TimestampNanoToTime 纳秒时间戳转时间
func TimestampNanoToTime(stamp int64) time.Time {
	return timestampToTime(0, stamp)
}

// TimestampMilToTime 毫秒时间戳转时间(毫秒 *1e6 先转成纳秒)
func TimestampMilToTime(stamp int64) time.Time {
	return timestampToTime(0, stamp*1e6)
}

// TimestampSecToTime 秒级别时间戳转时间
func TimestampSecToTime(stamp int64) time.Time {
	return timestampToTime(stamp, 0)
}

// TimeStrToTimestampMill 字符串转毫秒时间戳
func TimeStrToTimestampMill(timeStr string) (int64, error) {
	t, err := TimeStrToTime(timeStr)
	if err != nil {
		return -1., err
	}
	// 毫秒级别
	return (t.UnixNano()) / 1e6, nil
}

// ////////////////////////////////////////////////////////////

const TimeActivitiesLayout = "20060102"

// NumberToDate 时间数字int必须是：20210222
func NumberToDate(number int) time.Time {
	var year = number / 10000
	var month = number % 10000 / 100
	var day = number % 100
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, TimeLocation)
}

// NumStrToDate 时间字符串的格式必须是："20210222"
func NumStrToDate(s string) (time.Time, error) {
	timeRet, err := time.ParseInLocation(TimeActivitiesLayout, s, TimeLocation)
	if err != nil {
		return timeRet, err
	}
	return timeRet, nil
}

// GetDayBeginMoment 获取日期的最早时刻
func GetDayBeginMoment(t time.Time) time.Time {
	y, m, d := t.Date()
	n := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	return n
}

// GetDayEndMoment 获取日期的最晚时刻
func GetDayEndMoment(t time.Time) time.Time {
	y, m, d := t.Date()
	n := time.Date(y, m, d, 23, 59, 59, 999999999, time.Local)
	return n
}

type FormatTime struct {
	time.Time
}

// UnmarshalJSON 替换time的json反序列化
func (t *FormatTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+MYSec+`"`, string(data), TimeLocation)
	*t = FormatTime{
		now,
	}
	return
}

// MarshalJSON 替换time的json序列化
func (t FormatTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(MYSec))
	return []byte(formatted), nil
}

func (t FormatTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *FormatTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = FormatTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// FormatDurationToSecond 将持续时间转为秒数 1:01:03
func FormatDurationToSecond(duration string) (int64, error) {
	hms := regexp.MustCompile(`^(?P<hour>\d+):(?P<minute>\d{1,2}):(?P<second>\d{1,2})$`)
	ms := regexp.MustCompile(`^(?P<minute>\d{1,2}):(?P<second>\d{1,2})$`)
	s := regexp.MustCompile(`^(?P<second>\d{1,2})$`)
	var groupNames []string
	var matched []string
	if matched = hms.FindStringSubmatch(duration); len(matched) > 0 {
		groupNames = hms.SubexpNames()
	} else if matched = ms.FindStringSubmatch(duration); len(matched) > 0 {
		groupNames = ms.SubexpNames()
	} else if matched = s.FindStringSubmatch(duration); len(matched) > 0 {
		groupNames = s.SubexpNames()
	} else {
		return 0, errors.New("持续时间格式不正确")
	}
	result := make(map[string]string)
	// 转换为map
	for i, name := range groupNames {
		if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
			result[name] = matched[i]
		}
	}
	hour := result["hour"]
	minute := result["minute"]
	second := result["second"]
	return gconv.ToInt64(hour)*3600 + gconv.ToInt64(minute)*60 + gconv.ToInt64(second), nil
}
