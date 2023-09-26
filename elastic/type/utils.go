package google_proto

import "time"

var (
	secondsOfVietNam = int((7 * time.Hour).Seconds())
	timeZone         = time.FixedZone("Vietnam", secondsOfVietNam)
)

func SetTimeZone(location *time.Location) {
	timeZone = location
}

func (d *Date) AsTime() time.Time {
	return time.Date(int(d.GetYear()), time.Month(d.GetMonth()), int(d.GetDay()), 0, 0, 0, 0, timeZone)
}

func (d *Date) GetNextDate() time.Time {
	return d.AsTime().AddDate(0, 0, 1)
}

func TimeToDate(t time.Time) *Date {
	t = t.In(timeZone)
	return &Date{Year: int32(t.Year()), Month: int32(t.Month()), Day: int32(t.Day())}
}
