package timer

import "time"


func GetNowTime() time.Time {
	return time.Now()
}

//为什么一套语言来一套时间 这些基础的封装
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), err
}
