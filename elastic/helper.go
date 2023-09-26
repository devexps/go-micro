package elastic

import (
	"fmt"
	google_proto "github.com/devexps/go-micro/v2/elastic/type"
	"reflect"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	elasticPkg "gopkg.in/olivere/elastic.v6"
)

// IsEnumAll ...
func IsEnumAll(vv interface{}) bool {
	type enumInterface interface {
		EnumDescriptor() ([]byte, []int)
	}
	if _, ok := vv.(enumInterface); ok {
		return fmt.Sprintf("%d", vv) == "-1"
	}
	return false
}

// IsNil ...
func IsNil(field interface{}) bool {
	return field == nil || reflect.ValueOf(field).IsZero()
}

// IsZero ...
func IsZero(field interface{}) bool {
	if field != nil {
		if reflect.ValueOf(field).IsZero() {
			return true
		}
	}
	return false
}

// DoubleSlice ...
func DoubleSlice(s interface{}) []interface{} {
	v := reflect.ValueOf(s)
	items := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		items[i] = v.Index(i).Interface()
	}
	return items
}

// RangeQuery ...
type RangeQuery struct {
	MapQuery map[string]*elasticPkg.RangeQuery
}

// NewRangeQuery ...
func (r *RangeQuery) NewRangeQuery(name string) *elasticPkg.RangeQuery {
	if q, ok := r.MapQuery[name]; ok {
		return q
	}
	q := elasticPkg.NewRangeQuery(name)
	r.MapQuery[name] = q
	return q
}

// MapRangeDateSearch ...
type MapRangeSearch struct {
	MapRangeTmStampSearch map[string]*RangeTmstampSearch
}

type RangeTmstampSearch struct {
	From, To     int64
	Upper, Lower bool
}

// AddFrom ...
func (r *MapRangeSearch) AddFrom(name string, vv interface{}, lower bool) bool {
	tm, ok := convertToTime(vv, false)
	if !ok {
		return false
	}

	if q, ok := r.MapRangeTmStampSearch[name]; ok {
		q.From = tm.UnixNano() / int64(time.Millisecond)
	} else {
		r.MapRangeTmStampSearch[name] = &RangeTmstampSearch{From: tm.UnixNano() / int64(time.Millisecond), Lower: lower}
	}
	return true
}

// AddTo ...
func (r *MapRangeSearch) AddTo(name string, vv interface{}, upper bool) bool {

	tm, ok := convertToTime(vv, true)
	if !ok {
		return false
	}
	if q, ok := r.MapRangeTmStampSearch[name]; ok {
		q.To = tm.UnixNano() / int64(time.Millisecond)
	} else {
		r.MapRangeTmStampSearch[name] = &RangeTmstampSearch{To: tm.UnixNano() / int64(time.Millisecond), Upper: upper} //tm.UnixNano() / int64(`, c.getPacketName(c.timePkg), `.Millisecond)`)
	}
	return true
}

// MakeKeyEsMap ...
func MakeKeyEsMap(m map[string]interface{}, key string) map[string]interface{} {
	if t, ok := m[key]; ok {
		if t, ok := t.(map[string]interface{}); ok {
			return t
		}
	}
	t := map[string]interface{}{}
	m[key] = t
	return t
}

// MakeKeyMap ...
func MakeKeyMap(m *map[string]interface{}, key string) *map[string]interface{} {
	if t, ok := (*m)[key]; ok {
		if t, ok := t.(*map[string]interface{}); ok {
			return t
		}
	}
	t := &map[string]interface{}{}
	(*m)[key] = t
	return t
}

// CheckTimestampType ...
func CheckTimestampType(field interface{}) (*timestamp.Timestamp, bool) {
	if ts, ok := field.(*timestamp.Timestamp); ok {
		return ts, true
	}
	return nil, false
}

// CheckDateType ...
func CheckDateType(field interface{}) (*google_proto.Date, bool) {
	if date, ok := field.(*google_proto.Date); ok {
		return date, true
	}
	return nil, false
}

// GetTypeName ...
func GetTypeName(name string) string {
	typeNames := strings.Split(name, ".")
	return typeNames[len(typeNames)-1]
}

func convertToTime(vv interface{}, toNextDate bool) (time.Time, bool) {
	var (
		tm  time.Time
		err error
	)
	if ts, ok := vv.(*timestamp.Timestamp); ok {
		tm, err = ptypes.Timestamp(ts)
	} else if from, ok := vv.(*google_proto.Date); ok {
		if toNextDate {
			tm = from.GetNextDate()
		} else {
			tm = from.AsTime()
		}
	} else {
		return time.Time{}, false
	}

	if err != nil {
		return time.Time{}, false
	}
	return tm, true
}
