package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2020, 10,6,14,35, 0,0, time.Local)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	/*p(now.Sub(then))
	p(now.Sub(then).Hours())
	p(now.Sub(then).Minutes())
	p(now.Sub(then).Seconds())*/
	/*p(then.Add(now.Sub(then)))
	p(then.Add(-now.Sub(then)))*/

	p(then.Local())

}
