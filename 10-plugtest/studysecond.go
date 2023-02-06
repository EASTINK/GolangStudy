package main

import (
	"time"
)

func studysecond(current, rdate time.Time) int {
	//Current := time.Now()
	//2月
	if current.Month() == 2 {
		if current.Day() <= rdate.Day() {

			if current.Day() == rdate.Day() {
				hour := rdate.Hour() - current.Hour()
				miniute := rdate.Minute() - current.Minute()
				seconds := rdate.Second() - current.Second()
				num := (hour*60+miniute)*60 + seconds
				return num
			}

			if current.Day() < rdate.Day() {
				day := rdate.Day() - current.Day()           //1day
				hour := rdate.Hour() - current.Hour()        //-10h
				miniute := rdate.Minute() - current.Minute() //-22m
				seconds := rdate.Second() - current.Second() //-32s
				num := ((day*24+hour)*60+miniute)*60 + seconds
				return num
			}

		}
	}
	return 0
}

func main() {

	timezone, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(timezone)

	s.Every(1).Day().At("23:13").Do(func() {
		go cron1()
	})
	//开学时间 2022年2月18号 8点
	//fmt.Println(studysecond(time.Date(2023, 2, 18, 0, 32, 00, 00, time.Local), time.Date(2023, 2, 18, 8, 00, 00, 00, time.Local)))
	s.StartBlocking()
}
