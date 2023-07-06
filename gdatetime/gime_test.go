package gdatetime

import (
	"testing"
)

func TestGetCurrentTime(t *testing.T) {
	t.Log(GetCurrentTime())
}

func TestTime2StrAsFormat(t *testing.T) {
	timeStr := TimeToStrAsFormat(GetCurrentTime(), MYSec)
	t.Log("输出结果:" + timeStr)
}

func TestTimeStr2Time(t *testing.T) {
	timeStr := "2022-01-05 01:02:03"
	tm, err := TimeStrToTime(timeStr)
	if err == nil {
		t.Log(tm)
	} else {
		t.Error(err)
	}

}

func TestTimeToUTC(t *testing.T) {
	t.Log(TimeToUTC(GetCurrentTime()))
}

func TestTimeToTimeStampSecond(t *testing.T) {
	t.Log(TimeToTimeStampSecond(GetCurrentTime()))
}

func TestTimeToTimeStampNano(t *testing.T) {
	t.Log(TimeToTimeStampNano(GetCurrentTime()))
}
func TestTimeToTimeStampMill(t *testing.T) {
	t.Log(TimeToTimeStampMill(GetCurrentTime()))
}

func TestTimestampMilToTime(t *testing.T) {
	t.Log(TimestampMilToTime(1669611729485))
}

func TestTimestampSecToTime(t *testing.T) {
	t.Log(TimestampSecToTime(1669611726))
}

func TestTimeStrToTimestampMill(t *testing.T) {
	t.Log(TimeStrToTimestampMill("2022-11-01 15:32:36"))
}

func TestNumberToDate(t *testing.T) {
	t.Log(NumberToDate(20221101))
}

func TestNumStrToDate(t *testing.T) {
	t.Log(NumStrToDate("20221101"))
}
