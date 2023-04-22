# timetoUTC
timetoUTC is a simple command-line tool for time difference calculation.


## Getting Started

```
go install github.com/kazsk1/timetoUTC@latest
```


## Instructions
1. UTC (RFC3339) to local time on a timezone abbreviation such as "JST", "PDT", and "SGT"
```
./timetoUTC 2023-04-22T09:49:59Z JST

Original time: 2023/04/22 09:49:59 UTC
  Target time: 2023/04/22 18:49:59 JST
```

2. UTC offset to UTC
```
./timetoUTC 2023/04/10 09:00 UTC+9

Original time: 2023/04/10 09:00 UTC+9
     UTC time: 2023/04/10 00:00 UTC
```

3. UTC offset to UTC offset
```
./timetoUTC 2023/04/21 12:59:59 UTC-7 UTC+9

Original time: 2023/04/21 12:59:59 UTC-7
  Target time: 2023/04/22 04:59:59 UTC+9
```

4. Local time on a timezone abbreviation to local time on another timezone abbreviation
```
./timetoUTC 2023-04-21 15:49:42 PDT NZST

Original time: 2023-04-21 15:49:42 PDT
  Target time: 2023-04-22 10:49:42 NZST
```


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
