// This file has automatically been generated on Wed Feb 26 02:10:17 +05 2020.
// DO NOT EDIT.
package time

import (
	"time"
	_ "unsafe"
)

//go:linkname timegobencode time.sub_timegobencode
func timegobencode(t time.Time) ([]byte, error) {
	return t.GobEncode()
}

//go:linkname TimeGobEncode time.sub_timegobencode
//go:noescape
func TimeGobEncode(t time.Time) ([]byte, error)

//go:linkname timeisoweek time.sub_timeisoweek
func timeisoweek(t time.Time) int {
	return t.ISOWeek()
}

//go:linkname TimeISOWeek time.sub_timeisoweek
//go:noescape
func TimeISOWeek(t time.Time) int

//go:linkname timeutc time.sub_timeutc
func timeutc(t time.Time) time.Time {
	return t.UTC()
}

//go:linkname TimeUTC time.sub_timeutc
//go:noescape
func TimeUTC(t time.Time) time.Time

//go:linkname timeunix time.sub_timeunix
func timeunix(t time.Time) int64 {
	return t.Unix()
}

//go:linkname TimeUnix time.sub_timeunix
//go:noescape
func TimeUnix(t time.Time) int64

//go:linkname AfterFunc time.AfterFunc
//go:noescape
func AfterFunc(d time.Duration, f func()) *time.Timer

//go:linkname ParseDuration time.ParseDuration
//go:noescape
func ParseDuration(s string) (time.Duration, error)

//go:linkname durationnanoseconds time.sub_durationnanoseconds
func durationnanoseconds(d time.Duration) int64 {
	return d.Nanoseconds()
}

//go:linkname DurationNanoseconds time.sub_durationnanoseconds
//go:noescape
func DurationNanoseconds(d time.Duration) int64

//go:linkname timein time.sub_timein
func timein(t time.Time, loc *time.Location) time.Time {
	return t.In(loc)
}

//go:linkname TimeIn time.sub_timein
//go:noescape
func TimeIn(t time.Time, loc *time.Location) time.Time

//go:linkname timelocal time.sub_timelocal
func timelocal(t time.Time) time.Time {
	return t.Local()
}

//go:linkname TimeLocal time.sub_timelocal
//go:noescape
func TimeLocal(t time.Time) time.Time

//go:linkname Unix time.Unix
//go:noescape
func Unix(sec int64, nsec int64) time.Time

//go:linkname timeappendformat time.sub_timeappendformat
func timeappendformat(t time.Time, b []byte, layout string) []byte {
	return t.AppendFormat(b, layout)
}

//go:linkname TimeAppendFormat time.sub_timeappendformat
//go:noescape
func TimeAppendFormat(t time.Time, b []byte, layout string) []byte

//go:linkname timenanosecond time.sub_timenanosecond
func timenanosecond(t time.Time) int {
	return t.Nanosecond()
}

//go:linkname TimeNanosecond time.sub_timenanosecond
//go:noescape
func TimeNanosecond(t time.Time) int

//go:linkname NewTimer time.NewTimer
//go:noescape
func NewTimer(d time.Duration) *time.Timer

//go:linkname timelocation time.sub_timelocation
func timelocation(t time.Time) *time.Location {
	return t.Location()
}

//go:linkname TimeLocation time.sub_timelocation
//go:noescape
func TimeLocation(t time.Time) *time.Location

//go:linkname timerstop time.sub_timerstop
func timerstop(t *time.Timer) bool {
	return t.Stop()
}

//go:linkname TimerStop time.sub_timerstop
//go:noescape
func TimerStop(t *time.Timer) bool

//go:linkname timestring time.sub_timestring
func timestring(t time.Time) string {
	return t.String()
}

//go:linkname TimeString time.sub_timestring
//go:noescape
func TimeString(t time.Time) string

//go:linkname timeweekday time.sub_timeweekday
func timeweekday(t time.Time) time.Weekday {
	return t.Weekday()
}

//go:linkname TimeWeekday time.sub_timeweekday
//go:noescape
func TimeWeekday(t time.Time) time.Weekday

//go:linkname FixedZone time.FixedZone
//go:noescape
func FixedZone(name string, offset int) *time.Location

//go:linkname timeadd time.sub_timeadd
func timeadd(t time.Time, d time.Duration) time.Time {
	return t.Add(d)
}

//go:linkname TimeAdd time.sub_timeadd
//go:noescape
func TimeAdd(t time.Time, d time.Duration) time.Time

//go:linkname timebefore time.sub_timebefore
func timebefore(t time.Time, u time.Time) bool {
	return t.Before(u)
}

//go:linkname TimeBefore time.sub_timebefore
//go:noescape
func TimeBefore(t time.Time, u time.Time) bool

//go:linkname timemonth time.sub_timemonth
func timemonth(t time.Time) time.Month {
	return t.Month()
}

//go:linkname TimeMonth time.sub_timemonth
//go:noescape
func TimeMonth(t time.Time) time.Month

//go:linkname Now time.Now
//go:noescape
func Now() time.Time

//go:linkname timeday time.sub_timeday
func timeday(t time.Time) int {
	return t.Day()
}

//go:linkname TimeDay time.sub_timeday
//go:noescape
func TimeDay(t time.Time) int

//go:linkname timetruncate time.sub_timetruncate
func timetruncate(t time.Time, d time.Duration) time.Time {
	return t.Truncate(d)
}

//go:linkname TimeTruncate time.sub_timetruncate
//go:noescape
func TimeTruncate(t time.Time, d time.Duration) time.Time

//go:linkname timeunmarshaltext time.sub_timeunmarshaltext
func timeunmarshaltext(t *time.Time, data []byte) error {
	return t.UnmarshalText(data)
}

//go:linkname TimeUnmarshalText time.sub_timeunmarshaltext
//go:noescape
func TimeUnmarshalText(t *time.Time, data []byte) error

//go:linkname After time.After
//go:noescape
func After(d time.Duration) <-chan time.Time

//go:linkname Tick time.Tick
//go:noescape
func Tick(d time.Duration) <-chan time.Time

//go:linkname durationseconds time.sub_durationseconds
func durationseconds(d time.Duration) float64 {
	return d.Seconds()
}

//go:linkname DurationSeconds time.sub_durationseconds
//go:noescape
func DurationSeconds(d time.Duration) float64

//go:linkname locationstring time.sub_locationstring
func locationstring(l *time.Location) string {
	return l.String()
}

//go:linkname LocationString time.sub_locationstring
//go:noescape
func LocationString(l *time.Location) string

//go:linkname Since time.Since
//go:noescape
func Since(t time.Time) time.Duration

//go:linkname monthstring time.sub_monthstring
func monthstring(m time.Month) string {
	return m.String()
}

//go:linkname MonthString time.sub_monthstring
//go:noescape
func MonthString(m time.Month) string

//go:linkname timeafter time.sub_timeafter
func timeafter(t time.Time, u time.Time) bool {
	return t.After(u)
}

//go:linkname TimeAfter time.sub_timeafter
//go:noescape
func TimeAfter(t time.Time, u time.Time) bool

//go:linkname timeclock time.sub_timeclock
func timeclock(t time.Time) int {
	return t.Clock()
}

//go:linkname TimeClock time.sub_timeclock
//go:noescape
func TimeClock(t time.Time) int

//go:linkname timeyear time.sub_timeyear
func timeyear(t time.Time) int {
	return t.Year()
}

//go:linkname TimeYear time.sub_timeyear
//go:noescape
func TimeYear(t time.Time) int

//go:linkname durationmicroseconds time.sub_durationmicroseconds
func durationmicroseconds(d time.Duration) int64 {
	return d.Microseconds()
}

//go:linkname DurationMicroseconds time.sub_durationmicroseconds
//go:noescape
func DurationMicroseconds(d time.Duration) int64

//go:linkname NewTicker time.NewTicker
//go:noescape
func NewTicker(d time.Duration) *time.Ticker

//go:linkname Date time.Date
//go:noescape
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time

//go:linkname timeformat time.sub_timeformat
func timeformat(t time.Time, layout string) string {
	return t.Format(layout)
}

//go:linkname TimeFormat time.sub_timeformat
//go:noescape
func TimeFormat(t time.Time, layout string) string

//go:linkname durationstring time.sub_durationstring
func durationstring(d time.Duration) string {
	return d.String()
}

//go:linkname DurationString time.sub_durationstring
//go:noescape
func DurationString(d time.Duration) string

//go:linkname ParseInLocation time.ParseInLocation
//go:noescape
func ParseInLocation(layout, value string, loc *time.Location) (time.Time, error)

//go:linkname timesecond time.sub_timesecond
func timesecond(t time.Time) int {
	return t.Second()
}

//go:linkname TimeSecond time.sub_timesecond
//go:noescape
func TimeSecond(t time.Time) int

//go:linkname timeunmarshaljson time.sub_timeunmarshaljson
func timeunmarshaljson(t *time.Time, data []byte) error {
	return t.UnmarshalJSON(data)
}

//go:linkname TimeUnmarshalJSON time.sub_timeunmarshaljson
//go:noescape
func TimeUnmarshalJSON(t *time.Time, data []byte) error

//go:linkname timegobdecode time.sub_timegobdecode
func timegobdecode(t *time.Time, data []byte) error {
	return t.GobDecode(data)
}

//go:linkname TimeGobDecode time.sub_timegobdecode
//go:noescape
func TimeGobDecode(t *time.Time, data []byte) error

//go:linkname timemarshalbinary time.sub_timemarshalbinary
func timemarshalbinary(t time.Time) ([]byte, error) {
	return t.MarshalBinary()
}

//go:linkname TimeMarshalBinary time.sub_timemarshalbinary
//go:noescape
func TimeMarshalBinary(t time.Time) ([]byte, error)

//go:linkname timeunmarshalbinary time.sub_timeunmarshalbinary
func timeunmarshalbinary(t *time.Time, data []byte) error {
	return t.UnmarshalBinary(data)
}

//go:linkname TimeUnmarshalBinary time.sub_timeunmarshalbinary
//go:noescape
func TimeUnmarshalBinary(t *time.Time, data []byte) error

//go:linkname timehour time.sub_timehour
func timehour(t time.Time) int {
	return t.Hour()
}

//go:linkname TimeHour time.sub_timehour
//go:noescape
func TimeHour(t time.Time) int

//go:linkname timeiszero time.sub_timeiszero
func timeiszero(t time.Time) bool {
	return t.IsZero()
}

//go:linkname TimeIsZero time.sub_timeiszero
//go:noescape
func TimeIsZero(t time.Time) bool

//go:linkname timemarshaltext time.sub_timemarshaltext
func timemarshaltext(t time.Time) ([]byte, error) {
	return t.MarshalText()
}

//go:linkname TimeMarshalText time.sub_timemarshaltext
//go:noescape
func TimeMarshalText(t time.Time) ([]byte, error)

//go:linkname timeunixnano time.sub_timeunixnano
func timeunixnano(t time.Time) int64 {
	return t.UnixNano()
}

//go:linkname TimeUnixNano time.sub_timeunixnano
//go:noescape
func TimeUnixNano(t time.Time) int64

//go:linkname durationminutes time.sub_durationminutes
func durationminutes(d time.Duration) float64 {
	return d.Minutes()
}

//go:linkname DurationMinutes time.sub_durationminutes
//go:noescape
func DurationMinutes(d time.Duration) float64

//go:linkname durationtruncate time.sub_durationtruncate
func durationtruncate(d time.Duration, m time.Duration) time.Duration {
	return d.Truncate(m)
}

//go:linkname DurationTruncate time.sub_durationtruncate
//go:noescape
func DurationTruncate(d time.Duration, m time.Duration) time.Duration

//go:linkname timeadddate time.sub_timeadddate
func timeadddate(t time.Time, years int, months int, days int) time.Time {
	return t.AddDate(years, months, days)
}

//go:linkname TimeAddDate time.sub_timeadddate
//go:noescape
func TimeAddDate(t time.Time, years int, months int, days int) time.Time

//go:linkname timedate time.sub_timedate
func timedate(t time.Time) (int, time.Month, int) {
	return t.Date()
}

//go:linkname TimeDate time.sub_timedate
//go:noescape
func TimeDate(t time.Time) (int, time.Month, int)

//go:linkname timeequal time.sub_timeequal
func timeequal(t time.Time, u time.Time) bool {
	return t.Equal(u)
}

//go:linkname TimeEqual time.sub_timeequal
//go:noescape
func TimeEqual(t time.Time, u time.Time) bool

//go:linkname timeminute time.sub_timeminute
func timeminute(t time.Time) int {
	return t.Minute()
}

//go:linkname TimeMinute time.sub_timeminute
//go:noescape
func TimeMinute(t time.Time) int

//go:linkname durationhours time.sub_durationhours
func durationhours(d time.Duration) float64 {
	return d.Hours()
}

//go:linkname DurationHours time.sub_durationhours
//go:noescape
func DurationHours(d time.Duration) float64

//go:linkname LoadLocation time.LoadLocation
//go:noescape
func LoadLocation(name string) (*time.Location, error)

//go:linkname parseerrorerror time.sub_parseerrorerror
func parseerrorerror(e *time.ParseError) string {
	return e.Error()
}

//go:linkname ParseErrorError time.sub_parseerrorerror
//go:noescape
func ParseErrorError(e *time.ParseError) string

//go:linkname Parse time.Parse
//go:noescape
func Parse(layout, value string) (time.Time, error)

//go:linkname timeround time.sub_timeround
func timeround(t time.Time, d time.Duration) time.Time {
	return t.Round(d)
}

//go:linkname TimeRound time.sub_timeround
//go:noescape
func TimeRound(t time.Time, d time.Duration) time.Time

//go:linkname durationmilliseconds time.sub_durationmilliseconds
func durationmilliseconds(d time.Duration) int64 {
	return d.Milliseconds()
}

//go:linkname DurationMilliseconds time.sub_durationmilliseconds
//go:noescape
func DurationMilliseconds(d time.Duration) int64

//go:linkname weekdaystring time.sub_weekdaystring
func weekdaystring(d time.Weekday) string {
	return d.String()
}

//go:linkname WeekdayString time.sub_weekdaystring
//go:noescape
func WeekdayString(d time.Weekday) string

//go:linkname durationround time.sub_durationround
func durationround(d time.Duration, m time.Duration) time.Duration {
	return d.Round(m)
}

//go:linkname DurationRound time.sub_durationround
//go:noescape
func DurationRound(d time.Duration, m time.Duration) time.Duration

//go:linkname timemarshaljson time.sub_timemarshaljson
func timemarshaljson(t time.Time) ([]byte, error) {
	return t.MarshalJSON()
}

//go:linkname TimeMarshalJSON time.sub_timemarshaljson
//go:noescape
func TimeMarshalJSON(t time.Time) ([]byte, error)

//go:linkname timesub time.sub_timesub
func timesub(t time.Time, u time.Time) time.Duration {
	return t.Sub(u)
}

//go:linkname TimeSub time.sub_timesub
//go:noescape
func TimeSub(t time.Time, u time.Time) time.Duration

//go:linkname timerreset time.sub_timerreset
func timerreset(t *time.Timer, d time.Duration) bool {
	return t.Reset(d)
}

//go:linkname TimerReset time.sub_timerreset
//go:noescape
func TimerReset(t *time.Timer, d time.Duration) bool

//go:linkname Until time.Until
//go:noescape
func Until(t time.Time) time.Duration

//go:linkname LoadLocationFromTZData time.LoadLocationFromTZData
//go:noescape
func LoadLocationFromTZData(name string, data []byte) (*time.Location, error)

//go:linkname timeyearday time.sub_timeyearday
func timeyearday(t time.Time) int {
	return t.YearDay()
}

//go:linkname TimeYearDay time.sub_timeyearday
//go:noescape
func TimeYearDay(t time.Time) int

//go:linkname timezone time.sub_timezone
func timezone(t time.Time) (string, int) {
	return t.Zone()
}

//go:linkname TimeZone time.sub_timezone
//go:noescape
func TimeZone(t time.Time) (string, int)
