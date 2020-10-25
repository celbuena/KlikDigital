package times

import "time"

func Now(hourOffset int, format string) string  {
	return time.Now().UTC().Add(time.Hour * time.Duration(hourOffset)).Format(format)
}

func GenerateUnixTimeNow(hourOffSet int) int64{
	today := time.Now().UTC().Add(time.Hour * time.Duration(hourOffSet))
	return today.Unix()
}