# timetoUTC
timetoUTC is a simple command-line tool for time difference calculation.


## Getting Started

```
go install github.com/kazsk1/timetoUTC@master
```

## Usage
1. Change UTC (RFC3339) into target time on a timezone abbreviation such as "JST", "PDT", and "SGT"
```
timetoUTC [YYYY-MM-DDTHH:MM:SSZ] [TG Abbr]
```

2. Change original UTC offset into UTC time
```
timetoUTC [YYYY/MM/DD] [HH:MM:SS] [OG UTC offset]
timetoUTC [YYYY-MM-DD] [HH:MM:SS] [OG UTC offset]
timetoUTC [YYYY/MM/DD] [HH:MM] [OG UTC offset]
timetoUTC [YYYY-MM-DD] [HH:MM] [OG UTC offset]
```

3. Change original UTC offset into target UTC offset
```
timetoUTC [YYYY/MM/DD] [HH:MM:SS] [OG UTC offset] [TZ UTC offset]
timetoUTC [YYYY-MM-DD] [HH:MM:SS] [OG UTC offset] [TZ UTC offset]
timetoUTC [YYYY/MM/DD] [HH:MM] [OG UTC offset] [TZ UTC offset]
timetoUTC [YYYY-MM-DD] [HH:MM] [OG UTC offset] [TZ UTC offset]

```

4. Change original time into target time on timezone abbreviations
```
timetoUTC [YYYY/MM/DD] [HH:MM:SS] [OG Abbr] [TZ Abbr]
timetoUTC [YYYY-MM-DD] [HH:MM:SS] [OG Abbr] [TZ Abbr]
timetoUTC [YYYY/MM/DD] [HH:MM] [OG Abbr] [TZ Abbr]
timetoUTC [YYYY-MM-DD] [HH:MM] [OG Abbr] [TZ Abbr]
```

## Examples

1. UTC (RFC3339) >> Target time on a timezone abbreviation such as "JST"
```
timetoUTC 2023-04-22T09:49:59Z JST

Original time: 2023/04/22 09:49:59 UTC
  Target time: 2023/04/22 18:49:59 JST
```

2. Original time UTC+9 >> UTC time
```
timetoUTC 2023/04/10 09:00 UTC+9

Original time: 2023/04/10 09:00 UTC+9
     UTC time: 2023/04/10 00:00 UTC
```

3. Original time UTC-7 >> Target time UTC+9
```
timetoUTC 2023/04/21 12:59:59 UTC-7 UTC+9

Original time: 2023/04/21 12:59:59 UTC-7
  Target time: 2023/04/22 04:59:59 UTC+9
```

4. Original time PDT >> Target time NZST
```
timetoUTC 2023-04-21 15:49:42 PDT NZST

Original time: 2023-04-21 15:49:42 PDT
  Target time: 2023-04-22 10:49:42 NZST
```


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.
