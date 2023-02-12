package rainforestCommon

import (
	"flag"
	"os"
	"testing"
	"time"
)

var targetTimeM time.Time
var targetTimeU time.Time
var utime time.Time
var location *time.Location

//var locP *time.Location

func TestUnixTime(t *testing.T) {

	var err error
	utime, err = UnixTime("0x1C96BB5D")

	t.Log(err)

	if !targetTimeU.Equal(utime) {
		t.Error("Expected ", targetTimeU, " got ", utime.In(location))
	}
}

func TestMeterTime(t *testing.T) {

	mtime := MeterTime(utime)

	t.Log(mtime)
	if !targetTimeM.Equal(mtime) {
		t.Error("Expected ", targetTimeM, " got ", mtime.In(location))
	}
}

func TestCalcVal(t *testing.T) {
	f, err := CalcVal("0x00042d", "0x00000001", "0x000003e8")

	t.Log(err)

	if f != 1.069 {
		t.Log(f)
	}
}

func TestMain(m *testing.M) {
	flag.Parse()

	targetTimeU, _ = time.Parse("Mon, 02 Jan 2006 15:04:05 MST", "Sat, 14 Mar 2015 09:26:53 UTC")
	targetTimeM, _ = time.Parse("Mon, 02 Jan 2006 15:04:05 MST", "Sat, 14 Mar 1985 09:26:53 UTC")
	location, _ = time.LoadLocation("UTC")

	os.Exit(m.Run())
}
