# timetoUTC
timetoUTC is a simple command-line tool for time difference calculation.


## Getting Started

```
go install github.com/kazsk1/timetoUTC@master
```

## Usage
1. UTC (RFC3339) to target time on a timezone abbreviation such as "JST", "PDT", and "SGT"
```
timetoUTC [YYYY-MM-DDTHH:MM:SSZ] [TGZ Abbr]
```

2. Original UTC offset to UTC time
```
timetoUTC [YYYY/MM/DD] [HH:MM:SS] [ORG Abbr]
timetoUTC [YYYY-MM-DD] [HH:MM:SS] [ORG Abbr]
timetoUTC [YYYY/MM/DD] [HH:MM] [ORG Abbr]
timetoUTC [YYYY-MM-DD] [HH:MM] [ORG Abbr]
```

3. Original UTC offset to Target UTC offset
```
timetoUTC [YYYY/MM/DD] [HH:MM:SS] [ORG Abbr]
timetoUTC [YYYY-MM-DD] [HH:MM:SS] [ORG Abbr]
timetoUTC [YYYY/MM/DD] [HH:MM] [ORG Abbr]
timetoUTC [YYYY-MM-DD] [HH:MM] [ORG Abbr]

```

4. Original time to Target time on timezone abbreviations
```
timetoUTC [YYYY/MM/DD] [HH:MM:SS] [ORG Abbr] [TGZ Abbr]
timetoUTC [YYYY-MM-DD] [HH:MM:SS] [ORG Abbr] [TGZ Abbr]
timetoUTC [YYYY/MM/DD] [HH:MM] [ORG Abbr] [TGZ Abbr]
timetoUTC [YYYY-MM-DD] [HH:MM] [ORG Abbr] [TGZ Abbr]
```

## Examples

1. UTC (RFC3339) to target time on a timezone abbreviation such as "JST"
```
timetoUTC 2023-04-22T09:49:59Z JST

Original time: 2023/04/22 09:49:59 UTC
  Target time: 2023/04/22 18:49:59 JST
```

2. Original time UTC+9 to UTC time
```
timetoUTC 2023/04/10 09:00 UTC+9

Original time: 2023/04/10 09:00 UTC+9
     UTC time: 2023/04/10 00:00 UTC
```

3. Original time UTC-7 to Target time UTC+9
```
timetoUTC 2023/04/21 12:59:59 UTC-7 UTC+9

Original time: 2023/04/21 12:59:59 UTC-7
  Target time: 2023/04/22 04:59:59 UTC+9
```

4. Original time PDT to Target time NZST
```
timetoUTC 2023-04-21 15:49:42 PDT NZST

Original time: 2023-04-21 15:49:42 PDT
  Target time: 2023-04-22 10:49:42 NZST
```


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.
