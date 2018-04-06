package dateutil

import (
	"time"
	"math"
	"strings"
	"strconv"
)

var julianEaster = map[ int ]string {
	1: "5 Apr ", 2: "25 Mar ", 3: "13 Apr ", 4: "2 Apr ", 5: "22 Mar ",
	6: "10 Apr ", 7: "30 Mar ", 8: "18 Apr ", 9: "7 Apr ", 10: "27 Mar ",
	11: "15 Apr ", 12: "4 Apr ", 13: "24 Mar ", 14: "12 Apr ", 15: "1 Apr ",
	16: "21 Mar ", 17: "9 Apr ", 18: "29 Mar ", 19: "17 Apr " }

func Now( ) time.Time {
	return time.Now()
}

func Easter( whichYear int ) time.Time {
	goldenNumber := getGoldenNumber( float64( whichYear ) )
	if easterDate, OK := julianEaster[ int( goldenNumber ) ]; OK {
		stringYear := strconv.Itoa( whichYear )
		stringDate := strings.Join( []string{ easterDate, stringYear[ 2:len( stringYear ) ], " 00:00 UTC" }, "" )
		timeObj, _:= time.Parse( time.RFC822, stringDate )
		return timeObj
	}
	return time.Now()
}
func getGoldenNumber( whichYear float64 ) float64 {
	return math.Mod( float64( whichYear ), 19 ) + 1
}

func ZoneInfo() ( string, int ) {
	zoneName, offSet := time.Now().In( time.Local ).Zone()
	return zoneName, offSet
}