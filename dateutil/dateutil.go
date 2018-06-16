package dateutil

import (
	"time"
	"math"
	"strings"
	"strconv"
	"fmt"
)

var julianEaster = map[ int ]string {
	1: "5 Apr ", 2: "25 Mar ", 3: "13 Apr ", 4: "2 Apr ", 5: "22 Mar ",
	6: "10 Apr ", 7: "30 Mar ", 8: "18 Apr ", 9: "7 Apr ", 10: "27 Mar ",
	11: "15 Apr ", 12: "4 Apr ", 13: "24 Mar ", 14: "12 Apr ", 15: "1 Apr ",
	16: "21 Mar ", 17: "9 Apr ", 18: "29 Mar ", 19: "17 Apr " }

var paschalFullMoon = map[ int ]string {
	1: "12 Apr ", 2: "11 Apr ", 3: "10 Apr ", 4: "9 Apr ", 5: "8 Apr ", 6: "7 Apr ",
	7: "6 Apr ", 8: "5 Apr ", 9: "4 Apr ", 10: "3 Apr ", 11: "2 Apr ", 12: "1 Apr ",
	13: "31 Mar ", 14: "30 Mar ", 15: "29 Mar ", 16: "28 Mar ", 17: "27 Mar ",
	18: "26 Mar ", 19: "25 Mar ", 20: "24 Mar ", 21: "23 Mar ", 22: "22 Mar ",
	23: "21 Mar ", 24: "18 Apr ", 25: "18 Apr ", 26: "17 Apr ", 27: "16 Apr ",
	28: "15 Apr ", 29: "14 Apr ", 30: "13 Apr ", 31: "17 Apr " }

var SUNDAY = time.Sunday
var MONDAY = time.Monday
var TUESDAY = time.Tuesday
var WEDNESDAY = time.Wednesday
var THURSDAY = time.Thursday
var FRIDAY = time.Friday
var SATURDAY = time.Saturday

func Now( ) time.Time {
	return time.Now()
}
func NextDay( startDate string, whichDay time.Weekday ) time.Time {
	fullStringDate := strings.Join( []string { startDate, " 00:00 UTC" }, "" )
	timeObj, _:= time.Parse( time.RFC822, fullStringDate )
	daysBetween := 7 - int( timeObj.Weekday() )

	return timeObj.AddDate( 0, 0, daysBetween )
}
func JulianEaster( whichYear int ) time.Time {
	goldenNumber := getGoldenNumber( float64( whichYear ) )
	if easterDate, OK := julianEaster[ int( goldenNumber ) ]; OK {
		stringYear := strconv.Itoa( whichYear )
		stringDate := strings.Join( []string{ easterDate, stringYear[ 2:len( stringYear ) ], " 00:00 UTC" }, "" )
		timeObj, _:= time.Parse( time.RFC822, stringDate )

		return timeObj
	}
	return Now()
}

func GregorianEaster( whichYear int ) time.Time {
	goldenNumber := getGoldenNumber( float64( whichYear ) ) -1
	julianEpact := int( math.Mod( 11 * goldenNumber, 30 ) )
	stringCentury := strconv.Itoa( whichYear )[ :2 ]
	intCentury, _ := strconv.Atoi( stringCentury )
	solarEquation := ( 3 * intCentury ) / 4
	lunarEquation := ( 8 * intCentury + 5 ) / 25
	gregorianEpact := julianEpact - solarEquation + lunarEquation + 8
	if gregorianEpact == 25 && goldenNumber > 11 {
		gregorianEpact = 31
	}
	if fullMoon, OK := paschalFullMoon[ int( gregorianEpact )]; OK {
		stringYear := strconv.Itoa( whichYear )
		stringDate := strings.Join( []string{ fullMoon, stringYear[ 2:len( stringYear ) ] }, "" )
		fmt.Println( stringDate )

		return NextDay( stringDate, SUNDAY )
	}

	return Now()
}

func getGoldenNumber( whichYear float64 ) float64 {
	return math.Mod( float64( whichYear ), 19 ) + 1
}

func ZoneInfo() ( string, int ) {
	zoneName, offSet := time.Now().In( time.Local ).Zone()
	return zoneName, offSet
}