// Copyright 2016 Tom Messick. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package rainforestCommon

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type LocalCommand struct {
	XMLName   xml.Name `xml:"LocalCommand"`
	Name      string   `xml:"Name"`
	MacId     string
	StartTime string
	EndTime   string
	Frequency string
}

// All the different packets that might be sent from the eagle
type Root struct {
	XMLName     xml.Name `xml:"rainforest"`
	Current     CurrentSummationDelivered
	Device      DeviceInfo
	Demand      InstantaneousDemand
	History     HistoryData
	Message     MessageCluster
	Meter       MeterInfo
	Net         NetworkInfo
	Poll        FastPollStatus
	Price       PriceCluster
	PriceDetail BlockPriceDetail
	Profile     ProfileData
	Schedule    ScheduleInfo
	Time        TimeCluster
}

// Not in uploader API manual
type BlockPriceDetail struct {
	XMLName                          xml.Name `xml:"BlockPriceDetail"`
	DeviceMacId                      string
	MeterMacId                       string
	TimeStamp                        string
	CurrentStart                     string
	CurrentDuration                  string
	BlockPeriodConsumption           string
	BlockPeriodConsumptionMultiplier string
	BlockPeriodConsumptionDivisor    string
	NumberOfBlocks                   string
	Multiplier                       string
	Divisor                          string
	Currency                         string
	TrailingDigits                   string
	Port                             string
}

type CurrentSummationDelivered struct {
	XMLName             xml.Name `xml:"CurrentSummationDelivered"`
	DeviceMacId         string
	MeterMacId          string
	TimeStamp           string
	SummationDelivered  string
	SummationReceived   string
	Multiplier          string
	Divisor             string
	DigitsRight         string
	DigitsLeft          string
	SuppressLeadingZero string
	Port                string
}

type DeviceInfo struct {
	XMLName      xml.Name `xml:"DeviceInfo"`
	DeviceMacId  string
	InstallCode  string
	LinkKey      string
	FWVersion    string
	HWVersion    string
	ImageType    string
	Manufacturer string
	ModelId      string
	DateCode     string
	Port         string
}

type FastPollStatus struct {
	XMLName     xml.Name `xml:"FastPollStatus"`
	DeviceMacId string
	MeterMacId  string
	Frequency   string
	EndTime     string
	Port        string
}

type HistoryData struct {
	XMLName       xml.Name           `xml:"HistoryData"`
	SummationList []CurrentSummation `xml:"CurrentSummation"`
}

type CurrentSummation struct {
	XMLName             xml.Name `xml:"CurrentSummation"`
	DeviceMacId         string
	MeterMacId          string
	TimeStamp           string
	SummationDelivered  string
	SummationReceived   string
	Multiplier          string
	Divisor             string
	DigitsRight         string
	DigitsLeft          string
	SuppressLeadingZero string
}

type InstantaneousDemand struct {
	XMLName             xml.Name `xml:"InstantaneousDemand"`
	DeviceMacId         string
	MeterMacId          string
	TimeStamp           string
	Demand              string
	Multiplier          string
	Divisor             string
	DigitsRight         string
	DigitsLeft          string
	SuppressLeadingZero string
	Port                string
}

type MessageCluster struct {
	XMLName              xml.Name `xml:"MessageCluster"`
	DeviceMacId          string
	MeterMacId           string
	TimeStamp            string
	Id                   string
	Text                 string
	Priority             string
	StartTime            string
	Duration             string
	ConfirmationRequired string
	Confirmed            string
	Queue                string
	Port                 string
}

type MeterInfo struct {
	XMLName     xml.Name `xml:"MeterInfo"`
	DeviceMacId string
	MeterMacId  string
	Type        string
	NickName    string
	Account     string
	Auth        string
	Host        string
	Enabled     string
}

type NetworkInfo struct {
	XMLName      xml.Name `xml:"NetworkInfo"`
	DeviceMacId  string
	CoordMacId   string
	Status       string
	Description  string
	ExtPanId     string
	Channel      string
	ShortAddr    string
	LinkStrength string
	Port         string
}

type PriceCluster struct {
	XMLName        xml.Name `xml:"PriceCluster"`
	DeviceMacId    string
	MeterMacId     string
	TimeStamp      string
	Price          string
	Currency       string
	TrailingDigits string
	Tier           string
	StartTime      string
	Duration       string
	RateLabel      string
	Port           string
}

type ProfileData struct {
	XMLName                  xml.Name `xml:"ProfileData"`
	DeviceMacId              string
	MeterMacId               string
	EndTime                  string
	Status                   string
	ProfileIntervalPeriod    string
	NumberOfPeriodsDelivered string
	IntervalData1            string
	IntervalData2            string
	IntervalData3            string
	IntervalData4            string
	IntervalData5            string
	IntervalData6            string
	IntervalData7            string
	IntervalData8            string
	IntervalData9            string
	IntervalData10           string
	IntervalData11           string
	IntervalData12           string
	Port                     string
}

// Not in uploader API manual
type ScheduleInfo struct {
	XMLName     xml.Name `xml:"ScheduleInfo"`
	DeviceMacId string
	MeterMacId  string
	Event       string
	Frequency   string
	Enabled     string
}

// Not in uploader API manual
type TimeCluster struct {
	XMLName     xml.Name `xml:"TimeCluster"`
	DeviceMacId string
	MeterMacId  string
	UTCTime     string
	LocalTime   string
	Port        string
}

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

func (r Root) String() string {
	result := ""
	result1 := ""

	if r.XMLName.Local == "rainforest" {
		result1 += r.Current.String()
		result1 += r.Device.String()
		result1 += r.Demand.String()
		result += r.History.String()
		result1 += r.Message.String()
		result += r.Meter.String()
		result1 += r.Net.String()
		result1 += r.Poll.String()
		result1 += r.Price.String()
		result1 += r.PriceDetail.String()
		result1 += r.Profile.String()
		result += r.Schedule.String()
		result1 += r.Time.String()
	} else {
		result = fmt.Sprintf("Unexpected name: %s", r.XMLName.Local)
	}

	if len(result) == 0 && len(result1) == 0 {
		result = fmt.Sprintf("Unknown xml")
	}
	return result
}

func (c CurrentSummationDelivered) String() string {
	if c.XMLName.Local != "" {
		dval, _ := CalcVal(
			c.SummationDelivered,
			c.Multiplier,
			c.Divisor)
		rval, _ := CalcVal(
			c.SummationReceived,
			c.Multiplier,
			c.Divisor)
		return fmt.Sprintf("%s DeviceMacId          %s\n"+
			"                          MeterMacId           %s\n"+
			"                          TimeStamp            %s\n"+
			"                          SummationDelivered   %d %*.*f\n"+
			"                          SummationReceived    %d %*.*f\n"+
			"                          Multiplier           %d\n"+
			"                          Divisor              %d\n"+
			"                          DigitsRight          %d\n"+
			"                          DigitsLeft           %d\n"+
			"                          SuppressLeadingZero  %s\n",
			c.XMLName.Local,
			c.DeviceMacId,
			c.MeterMacId,
			gettime(c.TimeStamp),
			getval(c.SummationDelivered),
			getval(c.DigitsLeft)+getval(c.DigitsRight),
			getval(c.DigitsRight),
			dval,
			getval(c.SummationReceived),
			getval(c.DigitsLeft)+getval(c.DigitsRight),
			getval(c.DigitsRight),
			rval,
			getval(c.Multiplier),
			getval(c.Divisor),
			getval(c.DigitsRight),
			getval(c.DigitsLeft),
			c.SuppressLeadingZero)
	} else {
		return ""
	}
}

func (d DeviceInfo) String() string {
	if d.XMLName.Local != "" {
		return fmt.Sprintf("%s                DeviceMacId          %s\n"+
			"                          InstallCode          %s\n"+
			"                          LinkKey              %s\n"+
			"                          FWVersion            %s\n"+
			"                          HWVersion            %s\n"+
			"                          ImageType            %s\n"+
			"                          Manufacturer         %s\n"+
			"                          ModelId              %s\n"+
			"                          DateCode             %s\n"+
			"                          Port                 %s\n",
			d.XMLName.Local,
			d.DeviceMacId,
			d.InstallCode,
			d.LinkKey,
			d.FWVersion,
			d.HWVersion,
			d.ImageType,
			d.Manufacturer,
			d.ModelId,
			d.DateCode,
			d.Port)
	} else {
		return ""
	}
}

func (d InstantaneousDemand) String() string {
	if d.XMLName.Local != "" {
		val, _ := CalcVal(
			d.Demand,
			d.Multiplier,
			d.Divisor)
		return fmt.Sprintf("%s       DeviceMacId          %s\n"+
			"                          MeterMacId           %s\n"+
			"                          TimeStamp            %s\n"+
			"                          Demand               %d %*.*f\n"+
			"                          Multiplier           %d\n"+
			"                          Divisor              %d\n"+
			"                          DigitsRight          %d\n"+
			"                          DigitsLeft           %d\n"+
			"                          SuppressLeadingZero  %s\n"+
			"                          Port                 %s\n",
			d.XMLName.Local,
			d.DeviceMacId,
			d.MeterMacId,
			gettime(d.TimeStamp),
			getval(d.Demand),
			getval(d.DigitsLeft)+getval(d.DigitsRight),
			getval(d.DigitsRight),
			val,
			getval(d.Multiplier),
			getval(d.Divisor),
			getval(d.DigitsRight),
			getval(d.DigitsLeft),
			d.SuppressLeadingZero,
			d.Port)

	} else {
		return ""
	}
}

func (h HistoryData) String() string {
	if h.XMLName.Local != "" {
		return fmt.Sprintf("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++%s\n",
			h.XMLName.Local)
	} else {
		return ""
	}
}

func (m MessageCluster) String() string {
	if m.XMLName.Local != "" {
		return fmt.Sprintf("%s            DeviceMacId          %s\n"+
			"                          MeterMacId           %s\n"+
			"                          TimeStamp            %s\n"+
			"                          Id                   %s\n"+
			"                          Text                 %s\n"+
			"                          Priority             %s\n"+
			"                          StartTime            %s\n"+
			"                          Duration             %s\n"+
			"                          ConfirmationRequired %s\n"+
			"                          Confirmed            %s\n"+
			"                          Queue                %s\n"+
			"                          Port                 %s\n",
			m.XMLName.Local,
			m.DeviceMacId,
			m.MeterMacId,
			m.TimeStamp,
			m.Id,
			m.Text,
			m.Priority,
			m.StartTime,
			m.Duration,
			m.ConfirmationRequired,
			m.Confirmed,
			m.Queue,
			m.Port)
	} else {
		return ""
	}
}

func (m MeterInfo) String() string {
	if m.XMLName.Local != "" {
		return fmt.Sprintf("%s            DeviceMacId          %s\n"+
			"                          MeterMacId           %s\n"+
			"                          Type                 %s\n"+
			"                          NickName             %s\n"+
			"                          Account              %s\n"+
			"                          Auth                 %s\n"+
			"                          Host                 %s\n"+
			"                          Enabled              %s\n",
			m.XMLName.Local,
			m.DeviceMacId,
			m.MeterMacId,
			m.Type,
			m.NickName,
			m.Account,
			m.Auth,
			m.Host,
			m.Enabled)
	} else {
		return ""
	}
}

func (n NetworkInfo) String() string {
	if n.XMLName.Local != "" {
		return fmt.Sprintf("%s               DeviceMacId          %s\n"+
			"                          CoordMacId           %s\n"+
			"                          Status               %s\n"+
			"                          Description          %s\n"+
			"                          ExtPanId             %s\n"+
			"                          Channel              %s\n"+
			"                          ShortAddr            %s\n"+
			"                          LinkStrength         %s\n"+
			"                          Port                 %s\n",
			n.XMLName.Local,
			n.DeviceMacId,
			n.CoordMacId,
			n.Status,
			n.Description,
			n.ExtPanId,
			n.Channel,
			n.ShortAddr,
			n.LinkStrength,
			n.Port)
	} else {
		return ""
	}
}

func (f FastPollStatus) String() string {
	if f.XMLName.Local != "" {
		return fmt.Sprintf("%s            DeviceMacId          %s\n"+
			"                          MeterMacId           %s\n"+
			"                          Frequency            %d\n"+
			"                          EndTime              %s\n"+
			"                          Port                 %s\n",
			f.XMLName.Local,
			f.DeviceMacId,
			f.MeterMacId,
			getval(f.Frequency),
			gettime(f.EndTime),
			f.Port)
	} else {
		return ""
	}
}

func (p PriceCluster) String() string {
	if p.XMLName.Local != "" {
		return fmt.Sprintf("%s              DeviceMacId          %s\n"+
			"                          MeterMacId           %s\n"+
			"                          TimeStamp            %s\n"+
			"                          Price                %s\n"+
			"                          Currency             %d\n"+
			"                          TrailingDigits       %d\n"+
			"                          Tier                 %d\n"+
			"                          StartTime            %d\n"+
			"                          Duration             %d\n"+
			"                          RateLabel            %s\n"+
			"                          Port                 %s\n",
			p.XMLName.Local,
			p.DeviceMacId,
			p.MeterMacId,
			p.TimeStamp,
			p.Price,
			getval(p.Currency),
			getval(p.TrailingDigits),
			getval(p.Tier),
			getval(p.StartTime),
			getval(p.Duration),
			p.RateLabel,
			p.Port)
	} else {
		return ""
	}
}

func (b BlockPriceDetail) String() string {
	if b.XMLName.Local != "" {
		cval, _ := CalcVal(
			b.BlockPeriodConsumption,
			b.BlockPeriodConsumptionMultiplier,
			b.BlockPeriodConsumptionDivisor)
		bval, _ := CalcVal(
			b.NumberOfBlocks,
			b.Multiplier,
			b.Divisor)
		return fmt.Sprintf("%s          DeviceMacId                      %s\n"+
			"                          MeterMacId                       %s\n"+
			"                          TimeStamp                        %s\n"+
			"                          CurrentStart                     %d\n"+
			"                          CurrentDuration                  %d\n"+
			"                          BlockPeriodConsumption           %d %6.*f\n"+
			"                          BlockPeriodConsumptionMultiplier %d\n"+
			"                          BlockPeriodConsumptionDivisor    %d\n"+
			"                          NumberOfBlocks                   %d %6.*f\n"+
			"                          Multiplier                       %d\n"+
			"                          Divisor                          %d\n"+
			"                          Currency                         %d\n"+
			"                          TrailingDigits                   %d\n"+
			"                          Port                             %s\n",
			b.XMLName.Local,
			b.DeviceMacId,
			b.MeterMacId,
			gettime(b.TimeStamp),
			getval(b.CurrentStart),
			getval(b.CurrentDuration),
			getval(b.BlockPeriodConsumption),
			getval(b.TrailingDigits),
			cval,
			getval(b.BlockPeriodConsumptionMultiplier),
			getval(b.BlockPeriodConsumptionDivisor),
			getval(b.NumberOfBlocks),
			getval(b.TrailingDigits),
			bval,
			getval(b.Multiplier),
			getval(b.Divisor),
			getval(b.Currency),
			getval(b.TrailingDigits),
			b.Port)
	} else {
		return ""
	}
}

func (p ProfileData) String() string {
	if p.XMLName.Local != "" {
		return fmt.Sprintf("%s               DeviceMacId              %s\n"+
			"                          MeterMacId               %s\n"+
			"                          EndTime                  %s\n"+
			"                          Status                   %d\n"+
			"                          ProfileIntervalPeriod    %s\n"+
			"                          NumberOfPeriodsDelivered %d\n"+
			"                          IntervalData1            %d\n"+
			"                          IntervalData2            %d\n"+
			"                          IntervalData3            %d\n"+
			"                          IntervalData4            %d\n"+
			"                          IntervalData5            %d\n"+
			"                          IntervalData6            %d\n"+
			"                          IntervalData7            %d\n"+
			"                          IntervalData8            %d\n"+
			"                          IntervalData9            %d\n"+
			"                          IntervalData10           %d\n"+
			"                          IntervalData11           %d\n"+
			"                          IntervalData12           %d\n"+
			"                          Port                     %s\n",
			p.XMLName.Local,
			p.DeviceMacId,
			p.MeterMacId,
			gettime(p.EndTime),
			getval(p.Status),
			p.ProfileIntervalPeriod,
			getval(p.NumberOfPeriodsDelivered),
			getval(p.IntervalData1),
			getval(p.IntervalData2),
			getval(p.IntervalData3),
			getval(p.IntervalData4),
			getval(p.IntervalData5),
			getval(p.IntervalData6),
			getval(p.IntervalData7),
			getval(p.IntervalData8),
			getval(p.IntervalData9),
			getval(p.IntervalData10),
			getval(p.IntervalData11),
			getval(p.IntervalData12),
			p.Port)
	} else {
		return ""
	}
}

func (s ScheduleInfo) String() string {
	if s.XMLName.Local != "" {
		return fmt.Sprintf("%s               DeviceMacId          %s\n"+
			"                          MeterMacId           %s\n"+
			"                          Event                %s\n"+
			"                          Frequency            %s\n"+
			"                          Enabled              %s\n",
			s.XMLName.Local,
			s.DeviceMacId,
			s.MeterMacId,
			s.Event,
			s.Frequency,
			s.Enabled)
	} else {
		return ""
	}
}

func (t TimeCluster) String() string {
	if t.XMLName.Local != "" {
		return fmt.Sprintf("%s               DeviceMacId          %s\n"+
			"                          MeterMacId           %s\n"+
			"                          UTCTime              %s\n"+
			"                          LocalTime            %s\n"+
			"                          Port                 %s\n",
			t.XMLName.Local,
			t.DeviceMacId,
			t.MeterMacId,
			gettime(t.UTCTime),
			gettime(t.LocalTime),
			t.Port)
	} else {
		return ""
	}
}
