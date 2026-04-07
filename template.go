package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"time"
)

//go:embed index.html
var pageTemplate string

func plural(w string, d time.Duration) string {
	if d > 1 {
		return w + "s"
	}
	return w
}

const (
	dSecond = time.Second
	dMinute = time.Minute
	dHour = time.Hour
	dDay = 24 * time.Hour
	dWeek = 7 * dDay
	dMonth = 30 * dDay
	dYear = 365 * dDay
)

func renderTemplate() *template.Template {
	funcMap := template.FuncMap{
	    "percent": func(x uint, y uint) uint {
	        return (x * 100) / y
	    },
		"format": func(d time.Duration) string {
			i := d/time.Millisecond
			w := "ms"
			switch {
				case d >= dYear:
					i = d/dYear
					w = plural("year", i)
				case d >= dMonth:
					i = d/dMonth
					w = plural("month", i)
				case d >= dWeek:
					i = d/dWeek
					w = plural("week", i)
				case d >= dDay + time.Hour:
					i = d/dDay
					w = plural("day", i)
				case d >= dHour:
					i = d/dHour
					w = plural("hour", i)
				case d >= dMinute:
					i = d/dMinute
					w = plural("minute", i)
				case d >= dSecond:
					i = d/dSecond
					w = plural("second", i)
			}

			return fmt.Sprintf("%d %s", i, w)
		},
	}

	t := template.Must(
		template.New("page").Funcs(funcMap).Parse(pageTemplate),
	)

	return t
}
