package timeseries

import (
	"GoCQLSockets/server/database"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

type Measurements struct {
	Measurements []float32 `json:"measurements"`
	CreationDate time.Time `json:"creation_date"`
}

type Request struct {
	Type      int64      `json:"name"`
	ID        gocql.UUID `json:"id"`
	BeginDate time.Time  `json:"begin_date"`
	EndDate   time.Time  `json:"end_date"`
}

func RequestAverage(value []byte) ([]byte, error) {
	request, err := bytesToRequest(value)
	if err != nil {
		return []byte(""), err
	} else {
		switch request.Type {
		case 0:
			return getDailyAverage(request.ID, request.BeginDate, request.EndDate)
		case 1:
			return getWeeklyAverage(request.ID, request.BeginDate, request.EndDate)
		case 2:
			return getMonthlyAverage(request.ID, request.BeginDate, request.EndDate)
		default:
			return []byte(""), errors.New("request type does not exists")
		}
	}

}

func bytesToRequest(value []byte) (Request, error) {
	var request Request
	value = bytes.Trim(value, "\x00")
	err := json.Unmarshal(value, &request)
	return request, err

}

func getDailyAverage(uuid gocql.UUID, beginDate time.Time, endDate time.Time) ([]byte, error) {
	return getCustomAverage(0, 0, 1, uuid, beginDate, endDate)
}
func getWeeklyAverage(uuid gocql.UUID, beginDate time.Time, endDate time.Time) ([]byte, error) {
	return getCustomAverage(0, 0, 7, uuid, beginDate, endDate)
}
func getMonthlyAverage(uuid gocql.UUID, beginDate time.Time, endDate time.Time) ([]byte, error) {
	return getCustomAverage(0, 1, 0, uuid, beginDate, endDate)
}

func getCustomAverage(x, y, z int, uuid gocql.UUID, beginDate, endDate time.Time) ([]byte, error) {

	var test Measurements
	var result []float32
	m := map[string]interface{}{}
	firstDayOfYear := time.Date(beginDate.Year(), 1, 1, 0, 0, 0, 0, beginDate.Location())
	for tempDate := beginDate; endDate.Sub(tempDate) > 0; tempDate = tempDate.AddDate(x, y, z) {
		query := "SELECT avg(pow_fac) FROM get_power_factor_and_usage_by_stone_id WHERE stone_id = ? and range >= ? and range <= ?;"
		nextDate := tempDate.AddDate(x, y, z)
		iterable := cassandra.Session.Query(query, uuid, int(tempDate.Sub(firstDayOfYear).Minutes()), int(nextDate.Sub(firstDayOfYear).Minutes())).Iter()

		for iterable.MapScan(m) {
			result = append(result, m["system.avg(pow_fac)"].(float32))
		}
		m = map[string]interface{}{}

	}
	test.CreationDate = firstDayOfYear
	test.Measurements = result
	fmt.Println(test)
	b, err := json.Marshal(test)
	return b, err
}
