package pkg

import "time"

func SetTimer(t string) time.Duration {
	setTime, err := time.ParseDuration(t)
	if err != nil {
		panic(err)
	}
	return setTime
}

func SetSeconds(t string) time.Duration {
	return time.Duration(SetTimer(t).Seconds())
}

func SetMinuts(t string) time.Duration {
	return time.Duration(SetTimer(t).Minutes())
}

func SetHours(t string) time.Duration {
	return time.Duration(SetTimer(t).Hours())
}
