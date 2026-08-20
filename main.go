package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var tzOffsets = map[string]int{
	"AEDT": 11, "AEST": 10, "AKDT": -8, "AKST": -9,
	"AWDT": 9, "AWST": 8, "CAT": 2, "CDT": -5,
	"CEST": 2, "CET": 1, "CST": -6, "EAT": 3,
	"EDT": -4, "EEST": 3, "EET": 2, "EST": -5,
	"HDT": -9, "HKT": 8, "HST": -10, "JST": 9,
	"KST": 9, "MDT": -6, "MSD": 4, "MSK": 3,
	"MST": -7, "NZDT": 13, "NZST": 12, "PDT": -7,
	"PST": -8, "SGT": 8, "WEST": 1, "WET": 0,
	"UTC": 0,
}

func parseLocation(tzStr string) (*time.Location, error) {
	upper := strings.ToUpper(strings.TrimSpace(tzStr))
	if offsetHours, ok := tzOffsets[upper]; ok {
		return time.FixedZone(upper, offsetHours*3600), nil
	}
	if strings.HasPrefix(upper, "UTC") {
		offsetPart := strings.TrimPrefix(upper, "UTC")
		if offsetPart == "" {
			return time.UTC, nil
		}
		offsetHours, err := strconv.Atoi(offsetPart)
		if err == nil && offsetHours >= -12 && offsetHours <= 14 {
			return time.FixedZone(upper, offsetHours*3600), nil
		}
	}
	return nil, fmt.Errorf("\033[33mERROR:\033[0m unknown timezone/offset: %q", tzStr)
}

func parseDateTime(dateStr, timeStr string) (time.Time, string, error) {
	combined := dateStr + " " + timeStr
	layouts := []struct {
		layout    string
		outFormat string
	}{
		{"2006-01-02 15:04:05", "2006-01-02 15:04:05 MST"},
		{"2006/01/02 15:04:05", "2006/01/02 15:04:05 MST"},
		{"2006-01-02 15:04", "2006-01-02 15:04 MST"},
		{"2006/01/02 15:04", "2006/01/02 15:04 MST"},
	}
	for _, l := range layouts {
		if t, err := time.ParseInLocation(l.layout, combined, time.UTC); err == nil {
			return t, l.outFormat, nil
		}
	}
	return time.Time{}, "", fmt.Errorf("invalid date/time format: %q %q", dateStr, timeStr)
}

func printConversion(origLabel, origVal, targetLabel, targetVal string) {
	width := max(len(origLabel), len(targetLabel))
	fmt.Printf("\033[36mConversion\033[0m\n")
	fmt.Printf(" %*s: %s\n", width, origLabel, origVal)
	fmt.Printf(" %*s: %s\n", width, targetLabel, targetVal)
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  timeto [UNIX timestamp]")
	fmt.Println("  timeto [YYYY-MM-DDTHH:MM:SSZ] [Target TZ]")
	fmt.Println("  timeto [Date] [Time] [Source TZ]")
	fmt.Println("  timeto [Date] [Time] [Source TZ] [Target TZ]")
}

func main() {
	args := os.Args[1:]

	switch len(args) {
	case 1: // UNIX timestamp
		val, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid UNIX timestamp: %v\n", err)
			os.Exit(1)
		}
		var t time.Time
		switch {
		case len(args[0]) >= 18: // Nanoseconds
			t = time.Unix(0, val).UTC()
		case len(args[0]) >= 15: // Microseconds
			t = time.UnixMicro(val).UTC()
		case len(args[0]) >= 12: // Milliseconds
			t = time.UnixMilli(val).UTC()
		default: // Seconds
			t = time.Unix(val, 0).UTC()
		}
		printConversion("UNIX time", args[0], "UTC time", t.Format("2006-01-02 15:04:05.999999999 UTC"))

	case 2: // RFC3339 -> Target TZ
		t, err := time.Parse(time.RFC3339, args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid RFC3339 time format: %v\n", err)
			os.Exit(1)
		}
		loc, err := parseLocation(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		printConversion("Original time", t.UTC().Format("2006/01/02 15:04:05 MST"),
			"Target time", t.In(loc).Format("2006/01/02 15:04:05 MST"))

	case 3, 4: // Date Time SrcTZ [DstTZ]
		srcLoc, err := parseLocation(args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		dstLoc := time.UTC
		dstLabel := "UTC time"
		if len(args) == 4 {
			dstLoc, err = parseLocation(args[3])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			dstLabel = "Target time"
		}

		t, outFormat, err := parseDateTime(args[0], args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Reconstruct time in source location
		origTime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, srcLoc)
		printConversion("Original time", origTime.Format(outFormat), dstLabel, origTime.In(dstLoc).Format(outFormat))

	default:
		printUsage()
		os.Exit(0)
	}
}
