// Copyright 2016 Tom Messick. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package rainforestCommon

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// The difference in seconds between the unix epoch
// and the meter epoch
var offset int64 = findOffset()

// A regular expression to match hex strings
var hexPattern *regexp.Regexp

func init() {
	offset = findOffset()
	hexPattern = regexp.MustCompile(`^0[xX]([0-9a-fA-F]+)$`)
}

func findOffset() int64 {
	unix := time.Unix(0, 0)
	meter, _ := time.Parse("2006-Jan-02", "2000-Jan-01")
	return int64(meter.Sub(unix).Seconds())
}

func getOffset() int64 {
	return offset
}

// MeterTime converts time from the Epoch used by Unix/Windows to
// the epoch used by the utility meter (Jan 1, 2000)
func MeterTime(t time.Time) time.Time {
	return time.Unix(t.Unix()-offset, 0)
}

// UnixTime converts the hex string in an XML file to time in seconds
// since the Unix Epoch (Jan 1, 1970)
func UnixTime(s string) (time.Time, error) {
	if hexPattern.MatchString(s) {
		intVal, err := strconv.ParseInt(s[2:len(s)], 16, 32)
		if err != nil {
			return time.Unix(0, 0), err
		}
		intVal += offset
		return time.Unix(int64(intVal), 0), nil
	} else {
		return time.Unix(0, 0), fmt.Errorf("Invalid hex value %s", s)
	}
}

// Hex2Float converts the hex value from an XML file to floating point
func Hex2Float(in string) (float32, error) {
	if hexPattern.MatchString(in) {
		intVal, err := strconv.ParseUint(in[2:len(in)], 16, 32)
		if err != nil {
			return 0.0, err
		}
		return float32(intVal), nil
	} else {
		return 0.0, fmt.Errorf("Invalid hex value %s", in)
	}
}

// CalcVal converts hex values from an XML file to floating point
// and then scales the input value
func CalcVal(input, mult, div string) (float32, error) {
	inputf, err := Hex2Float(input)
	if err != nil {
		return 0.0, err
	}

	multf, err := Hex2Float(mult)
	if err != nil {
		return 0.0, err
	}

	divf, err := Hex2Float(div)
	if err != nil {
		return 0.0, err
	}
	return float32(inputf * multf / divf), nil
}

// getval turns a xml hex string into a decimal int
func getval(s string) int {
	i, err := strconv.ParseInt(s, 0, 32)
	if err == nil {
		return int(i)
	} else {
		return -1
	}

}

// gettime turns a rainforest timestamp into a readable string
func gettime(s string) string {
	utime, err := UnixTime(s)
	if err == nil {
		return fmt.Sprintf("%s", utime)
	} else {
		return err.Error()
	}
}
