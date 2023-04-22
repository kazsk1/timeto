package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// bug 1 // go run main.go 2023-04-22T09:49:59Z00:00 JST
// bug 2 // go run main.go 2023/04/22 09:00 JS
// bug 3 // go run main.go 2023-04-21 15:49:42 PDT NZS

func param(z int) string {

	// https://www.timeanddate.com/time/zones/

	// PST, PDT, ..., SGT, JST ->  UTC-8, UTC-7, ..., UTC+8, UTC+9
	// pst, pdt, ..., SGT, JST ->  UTC-8, UTC-7, ..., UTC+8, UTC+9
	x := strings.ToUpper(os.Args[z])
	y := ""

	y = strings.ReplaceAll(x, "WEST", "UTC+1")
	x = strings.ReplaceAll(y, "WET", "UTC")
	y = strings.ReplaceAll(x, "AEDT", "UTC+11")
	x = strings.ReplaceAll(y, "AEST", "UTC+10")
	y = strings.ReplaceAll(x, "AKDT", "UTC-8")
	x = strings.ReplaceAll(y, "AKST", "UTC-9")
	y = strings.ReplaceAll(x, "AWDT", "UTC+9")
	x = strings.ReplaceAll(y, "AWST", "UTC+8")
	y = strings.ReplaceAll(x, "CAT", "UTC+2")
	x = strings.ReplaceAll(y, "CDT", "UTC-5")
	y = strings.ReplaceAll(x, "CEST", "UTC+2")
	x = strings.ReplaceAll(y, "CET", "UTC+1")
	y = strings.ReplaceAll(x, "CST", "UTC-6")
	x = strings.ReplaceAll(y, "EAT", "UTC+3")
	y = strings.ReplaceAll(x, "EDT", "UTC-4")
	x = strings.ReplaceAll(y, "EEST", "UTC+3")
	y = strings.ReplaceAll(x, "EET", "UTC+2")
	x = strings.ReplaceAll(y, "EST", "UTC-5")
	y = strings.ReplaceAll(x, "HDT", "UTC-9")
	x = strings.ReplaceAll(y, "HKT", "UTC+8")
	y = strings.ReplaceAll(x, "HST", "UTC-10")
	x = strings.ReplaceAll(y, "JST", "UTC+9")
	y = strings.ReplaceAll(x, "KST", "UTC+9")
	x = strings.ReplaceAll(y, "MDT", "UTC-6")
	y = strings.ReplaceAll(x, "MSD", "UTC+4")
	x = strings.ReplaceAll(y, "MSK", "UTC+3")
	y = strings.ReplaceAll(x, "MST", "UTC-7")
	x = strings.ReplaceAll(y, "NZDT", "UTC+13")
	y = strings.ReplaceAll(x, "NZST", "UTC+12")
	x = strings.ReplaceAll(y, "PDT", "UTC-7")
	y = strings.ReplaceAll(x, "PST", "UTC-8")
	x = strings.ReplaceAll(y, "SGT", "UTC+8")

	return x
}

func main() {

	fmt.Printf("\n")

	if len(os.Args) == 3 {

		// UTC-1, UTC, UTC+1, UTC+2,... -> -1, 0, 1, 2, ....
		inc, _ := strconv.Atoi((strings.Trim(param(2), "UTC")))

		// Split 2023-04-22T09:49:59Z to 2023 04 22T09:49:59Z
		// Split 2023-04-22T09:49:59Z to 2023-04-22T09 49 59Z
		arg1 := strings.Split(os.Args[1], "-")
		arg2 := strings.Split(os.Args[1], ":")

		// 22T09:49:59Z -> 10
		rep1 := regexp.MustCompile(`T.*`)
		str1 := rep1.ReplaceAllString(arg1[2], "")

		// 2023-04-22T09 -> 09
		rep2 := regexp.MustCompile(`.*T`)
		str2 := rep2.ReplaceAllString(arg2[0], "")

		// 59Z -> 59
		rep3 := regexp.MustCompile(`Z`)
		str3 := rep3.ReplaceAllString(arg2[2], "")

		h, _ := strconv.Atoi(str2)
		min, _ := strconv.Atoi(arg2[1])
		s, _ := strconv.Atoi(str3)

		// test command 1 // go run main.go 2023-04-22T09:49:59Z JST
		// test command 2 // go run main.go 2023-04-22T09:49:59Z PDT
		// test command 3 // go run main.go 2023-04-22T09:49:59Z Est
		// test command 4 // go run main.go 2023-04-22T09:49:59Z cest
		// test command 5 // go run main.go 2023-04-22T09:49:59Z west

		year, _ := strconv.Atoi(arg1[0])
		var m, _ = strconv.Atoi(arg1[1])
		d, _ := strconv.Atoi(str1)

		t := time.Date(year, time.Month(m), d, h, min, s, 0, time.FixedZone("UTC", 0*60*60))

		// Define the target time zone
		targetTimeZone := time.FixedZone(strings.ToUpper(os.Args[2]), inc*60*60)

		// Convert the original time to the target time zone
		targetTime := t.In(targetTimeZone).Add(time.Hour * time.Duration(0))

		// Print the original and converted times
		fmt.Println("Original time:", t.Format("2006/01/02 15:04:05 MST"))
		fmt.Println("  Target time:", targetTime.Format("2006/01/02 15:04:05 MST"))

	}

	if len(os.Args) == 4 {

		// UTC-1, UTC, UTC+1, UTC+2,... -> -1, 0, 1, 2, ....
		inc, _ := strconv.Atoi((strings.Trim(param(3), "UTC")))

		// Split 2023/04/10 to 2023 4 10
		// Split 2023-04-10 to 2023 4 10
		arg1 := strings.Split(os.Args[1], "/")
		arg2 := strings.Split(os.Args[1], "-")

		// Split 09:00 to 9 00
		arg3 := strings.Split(os.Args[2], ":")

		// Change 09:00 to 9 00 00
		if len(arg3) == 2 {
			zero := "00"
			arg3 = append(arg3, zero)
		}

		h, _ := strconv.Atoi(arg3[0])
		min, _ := strconv.Atoi(arg3[1])
		s, _ := strconv.Atoi(arg3[2])

		if len(arg1) == 3 {
			// test command 6 // go run main.go 2023/04/10 09:00 UTC-1
			// test command 7 // go run main.go 2023/04/10 09:00 UTC
			// test command 8 // go run main.go 2023/04/10 09:00 UTC+9
			// test command 9 // go run main.go 2023/04/10 08:00 UTC-2
			// test command 10 // go run main.go 2023/04/10 08:45 UTC-2
			// test command 11 // go run main.go 2023/04/10 08:45:39 UTC-2
			// test command 12 // go run main.go 2023/04/10 09:00 utc-5
			// test command 13 // go run main.go 2023/04/10 09:00 utc
			// test command 14 // go run main.go 2023/04/10 09:00 utc+9
			// test command 15 // go run main.go 2023/04/10 08:00 uTC-6
			// test command 16 // go run main.go 2023/04/10 08:45 uTC-7
			// test command 17 // go run main.go 2023/04/10 08:45:39 uTC-8
			// test command 18 // go run main.go 2023/04/22 09:00 NZST
			// test command 19 // go run main.go 2023/04/22 09:00 AWST
			// test command 20 // go run main.go 2023/04/22 09:00 JST
			// test command 21 // go run main.go 2023/04/22 09:00 HKT
			// test command 22 // go run main.go 2023/04/22 09:00 SGT
			// test command 23 // go run main.go 2023/04/22 09:00 PDt
			// test command 24 // go run main.go 2023/04/22 09:00:00 pst

			year, _ := strconv.Atoi(arg1[0])
			var m, _ = strconv.Atoi(arg1[1])
			d, _ := strconv.Atoi(arg1[2])

			t := time.Date(year, time.Month(m), d, h, min, s, 0, time.FixedZone(strings.ToUpper(os.Args[3]), inc*60*60))

			// Convert the time to UTC
			utc := t.UTC()

			if s == 0 {
				// Print the original and converted times
				fmt.Println("Original time:", t.Format("2006/01/02 15:04 MST"))
				fmt.Println("     UTC time:", utc.Format("2006/01/02 15:04 MST"))
			}

			if s != 0 {
				// Print the original and converted times
				fmt.Println("Original time:", t.Format("2006/01/02 15:04:05 MST"))
				fmt.Println("     UTC time:", utc.Format("2006/01/02 15:04:05 MST"))
			}

		} else if len(arg2) == 3 {
			// test command 25 // go run main.go 2023-04-10 09:00 UTC-1
			// test command 26 // go run main.go 2023-04-10 09:00 UTC
			// test command 27 // go run main.go 2023-04-10 09:00 UTC+9
			// test command 28 // go run main.go 2023-04-10 08:00 UTC-2
			// test command 29 // go run main.go 2023-04-10 08:45 UTC-2
			// test command 30 // go run main.go 2023-04-10 08:45:39 UTC-2
			// test command 31 // go run main.go 2023-04-10 09:00 utc-5
			// test command 32 // go run main.go 2023-04-10 09:00 utc
			// test command 33 // go run main.go 2023-04-10 09:00 utc+9
			// test command 34 // go run main.go 2023-04-10 08:00 uTC-6
			// test command 35 // go run main.go 2023-04-10 08:45 uTC-7
			// test command 36 // go run main.go 2023-04-10 08:45:39 uTC-8
			// test command 37 // go run main.go 2023-04-22 09:00 NZST
			// test command 38 // go run main.go 2023-04-22 09:00 AWST
			// test command 39 // go run main.go 2023-04-22 09:00 JST
			// test command 40 // go run main.go 2023-04-22 09:00 HKT
			// test command 41 // go run main.go 2023-04-22 09:00 SGT
			// test command 42 // go run main.go 2023-04-22 09:00 PDt
			// test command 43 // go run main.go 2023-04-22 09:00:00 pst

			year, _ := strconv.Atoi(arg2[0])
			var m, _ = strconv.Atoi(arg2[1])
			d, _ := strconv.Atoi(arg2[2])

			t := time.Date(year, time.Month(m), d, h, min, s, 0, time.FixedZone(strings.ToUpper(os.Args[3]), inc*60*60))

			// Convert the time to UTC
			utc := t.UTC()

			if s == 0 {
				// Print the original and converted times
				fmt.Println("Original time:", t.Format("2006-01-02 15:04 MST"))
				fmt.Println("     UTC time:", utc.Format("2006-01-02 15:04 MST"))
			}

			if s != 0 {
				// Print the original and converted times
				fmt.Println("Original time:", t.Format("2006-01-02 15:04:05 MST"))
				fmt.Println("     UTC time:", utc.Format("2006-01-02 15:04:05 MST"))
			}

		}
	}

	if len(os.Args) == 5 {

		// UTC-1, UTC, UTC+1, UTC+2,... -> -1, 0, 1, 2, ....
		inc1, _ := strconv.Atoi((strings.Trim(param(3), "UTC")))

		// UTC-1, UTC, UTC+1, UTC+2,... -> -1, 0, 1, 2, ....
		inc2, _ := strconv.Atoi((strings.Trim(param(4), "UTC")))

		// Split 2023/04/10 to 2023 4 10
		// Split 2023-04-10 to 2023 4 10
		arg1 := strings.Split(os.Args[1], "/")
		arg2 := strings.Split(os.Args[1], "-")

		// Split 09:00 to 9 00
		arg3 := strings.Split(os.Args[2], ":")

		// Change 09:00 to 9 00 00
		if len(arg3) == 2 {
			zero := "00"
			arg3 = append(arg3, zero)
		}

		h, _ := strconv.Atoi(arg3[0])
		min, _ := strconv.Atoi(arg3[1])
		s, _ := strconv.Atoi(arg3[2])

		if len(arg1) == 3 {
			// test command 44 // go run main.go 2023/04/09 00:00 UTC UTC+9
			// test command 45 // go run main.go 2023/04/09 00:00 UTC+1 UTC+9
			// test command 46 // go run main.go 2023/04/09 00:00 UTC+2 UTC+9
			// test command 47 // go run main.go 2023/04/09 00:00 UTC+3 UTC+9
			// test command 48 // go run main.go 2023/04/09 00:00 UTC+4 UTC+9
			// test command 49 // go run main.go 2023/04/09 00:30:01 UTC UTC+9
			// test command 50 // go run main.go 2023/04/09 00:00 utc+5 utc+9
			// test command 51 // go run main.go 2023/04/09 00:00 utc+6 utc+9
			// test command 52 // go run main.go 2023/04/09 00:00 utc+7 utc+9
			// test command 53 // go run main.go 2023/04/09 00:00 uTc+8 uTc+9
			// test command 54 // go run main.go 2023/04/09 00:00 uTc+9 uTc+9
			// test command 55 // go run main.go 2023/04/09 00:30:01 uTc+10 uTc+9
			// test command 56 // go run main.go 2023/04/21 12:59:59 UTC-7 UTC+9
			// test command 57 // go run main.go 2023/04/21 12:59:59 pdt jst
			// test command 58 // go run main.go 2023/04/21 15:49:42 PDT NZST

			year, _ := strconv.Atoi(arg1[0])
			var m, _ = strconv.Atoi(arg1[1])
			d, _ := strconv.Atoi(arg1[2])

			t := time.Date(year, time.Month(m), d, h, min, s, 0, time.FixedZone(strings.ToUpper(os.Args[3]), inc1*60*60))

			// Define the target time zone
			targetTimeZone := time.FixedZone(strings.ToUpper(os.Args[4]), inc2*60*60)

			// Convert the original time to the target time zone
			targetTime := t.In(targetTimeZone).Add(time.Hour * time.Duration(0))

			if s == 0 {
				// Print the original and target times.
				fmt.Println("Original time:", t.Format("2006/01/02 15:04 MST"))
				fmt.Println("  Target time:", targetTime.Format("2006/01/02 15:04 MST"))
			}

			if s != 0 {
				// Print the original and target times.
				fmt.Println("Original time:", t.Format("2006/01/02 15:04:05 MST"))
				fmt.Println("  Target time:", targetTime.Format("2006/01/02 15:04:05 MST"))
			}

		} else if len(arg2) == 3 {
			// test command 59 // go run main.go 2023-04-09 00:00 UTC UTC+9
			// test command 60 // go run main.go 2024-04-09 00:00 UTC+1 UTC+9
			// test command 61 // go run main.go 2023-04-09 00:00 UTC+2 UTC+9
			// test command 62 // go run main.go 2023-04-09 00:00 UTC+3 UTC+9
			// test command 63 // go run main.go 2023-04-09 00:00 UTC+4 UTC+9
			// test command 64 // go run main.go 2023-04-09 00:30:01 UTC UTC+9
			// test command 65 // go run main.go 2023-04-09 00:00 utc+5 utc+9
			// test command 66 // go run main.go 2024-04-09 00:00 utc+6 utc+9
			// test command 67 // go run main.go 2023-04-09 00:00 utc+7 utc+9
			// test command 68 // go run main.go 2023-04-09 00:00 uTc+8 uTc+9
			// test command 69 // go run main.go 2023-04-09 00:00 uTc+9 uTc+9
			// test command 70 // go run main.go 2023-04-09 00:30:01 uTc+10 uTc+9
			// test command 71 // go run main.go 2023-04-21 12:59:59 UTC-7 UTC+9
			// test command 72 // go run main.go 2023-04-21 12:59:59 pdt jst
			// test command 73 // go run main.go 2023-04-21 15:49:42 PDT NZST

			year, _ := strconv.Atoi(arg2[0])
			var m, _ = strconv.Atoi(arg2[1])
			d, _ := strconv.Atoi(arg2[2])

			t := time.Date(year, time.Month(m), d, h, min, s, 0, time.FixedZone(strings.ToUpper(os.Args[3]), inc1*60*60))

			// Define the target time zone
			targetTimeZone := time.FixedZone(strings.ToUpper(os.Args[4]), inc2*60*60)

			// Convert the original time to the target time zone
			targetTime := t.In(targetTimeZone).Add(time.Hour * time.Duration(0))

			if s == 0 {
				// Print the original and target times.
				fmt.Println("Original time:", t.Format("2006-01-02 15:04 MST"))
				fmt.Println("  Target time:", targetTime.Format("2006-01-02 15:04 MST"))
			}

			if s != 0 {
				// Print the original and target times.
				fmt.Println("Original time:", t.Format("2006-01-02 15:04:05 MST"))
				fmt.Println("  Target time:", targetTime.Format("2006-01-02 15:04:05 MST"))
			}

		}
	}

}
