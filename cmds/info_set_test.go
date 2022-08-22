package cmds

import (
	"errors"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"testing"
	"time"
)

// TestParseSetInfoData tests parseSetInfoData to verify that the function
// correctly parses and validates the data.
func TestParseSetInfoData(t *testing.T) {
	tt := []struct {
		name         string
		premisesUUID string
		startDate    string
		endDate      string
		err          error
		u            uuid.UUID
		s            time.Time
		e            time.Time
	}{
		{
			name:         "invalid args",
			premisesUUID: "",
			startDate:    "",
			endDate:      "",
			err:          errors.New("invalid args: at least one argument is required"),
		},
		{
			name:         "invalid uuid",
			premisesUUID: "aa53-4be6-a63e-2758758022bf",
			startDate:    "",
			endDate:      "",
			err:          errors.New("invalid UUID length: 27"),
		},
		{
			name:         "invalid start date",
			premisesUUID: "",
			startDate:    "25-05-1993",
			endDate:      "",
			err:          errors.New("parsing time \"25-05-1993\" as \"2006-01-02\": cannot parse \"5-1993\" as \"2006\""),
		},
		{
			name:         "invalid end date",
			premisesUUID: "",
			startDate:    "",
			endDate:      "25-05-1993",
			err:          errors.New("parsing time \"25-05-1993\" as \"2006-01-02\": cannot parse \"5-1993\" as \"2006\""),
		},
		{

			name:         "valid uuid",
			premisesUUID: "62822578-aa53-4be6-a63e-2758758022bf",
			startDate:    "",
			endDate:      "",
			err:          nil,
			u:            uuid.MustParse("62822578-aa53-4be6-a63e-2758758022bf"),
		},
		{
			name:         "valid start date",
			premisesUUID: "",
			startDate:    "2022-05-11",
			endDate:      "",
			err:          nil,
			s:            MustParseTime("2006-01-02", "2022-05-11"),
		},
		{
			name:         "valid start date",
			premisesUUID: "",
			startDate:    "",
			endDate:      "2022-05-11",
			err:          nil,
			e:            MustParseTime("2006-01-02", "2022-05-11"),
		},
		{
			name:         "all valid args",
			premisesUUID: "46828013-14ae-46e6-a8e1-fcdda54dd521",
			startDate:    "2022-04-20",
			endDate:      "2022-05-11",
			err:          nil,
			u:            uuid.MustParse("46828013-14ae-46e6-a8e1-fcdda54dd521"),
			s:            MustParseTime("2006-01-02", "2022-04-20"),
			e:            MustParseTime("2006-01-02", "2022-05-11"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			u, s, e, err := parseSetInfoData(&tc.premisesUUID, &tc.startDate, &tc.endDate)
			//fmt.Println(err, tc.err, cli.ErrorNotEqual(err, tc.err))
			if cli.ErrorNotEqual(err, tc.err) {
				t.Errorf("expected error '%v' got '%v'", tc.err, err)
			}
			if u != tc.u {
				t.Errorf("expected uuid %v got %v", tc.u, u)
			}
			if s != tc.s {
				t.Errorf("expected start date %s got %s", s.Format("2006-01-02"), s.Format("2006-01-02"))
			}
			if e != tc.e {
				t.Errorf("expected end date %s got %s", s.Format("2006-01-02"), s.Format("2006-01-02"))
			}
		})
	}
}
