# timeto
timeto is a simple command-line tool for time difference calculation.


## Getting Started

```
go install github.com/kazsk1/timeto@main
```

## Usage
1) Change UTC (RFC3339) into target time using abbreviation such as JST, PDT and SGT
```powershell or go
timeto [YYYY-MM-DDTHH:MM:SSZ] [TG Abbr]
```

2) Change original UTC offset into UTC time
```powershell or go
timeto [YYYY/MM/DD] [HH:MM:SS] [OG UTC offset]
timeto [YYYY-MM-DD] [HH:MM:SS] [OG UTC offset]
timeto [YYYY/MM/DD] [HH:MM] [OG UTC offset]
timeto [YYYY-MM-DD] [HH:MM] [OG UTC offset]
```

3) Change original UTC offset into target UTC offset
```powershell or go
timeto [YYYY/MM/DD] [HH:MM:SS] [OG UTC offset] [TZ UTC offset]
timeto [YYYY-MM-DD] [HH:MM:SS] [OG UTC offset] [TZ UTC offset]
timeto [YYYY/MM/DD] [HH:MM] [OG UTC offset] [TZ UTC offset]
timeto [YYYY-MM-DD] [HH:MM] [OG UTC offset] [TZ UTC offset]
```

4) Change original time into target time on timezone abbreviations
```powershell or go
timeto [YYYY/MM/DD] [HH:MM:SS] [OG UTC offset] [TZ UTC offset]
timeto [YYYY/MM/DD] [HH:MM:SS] [OG Abbr] [TZ Abbr]
timeto [YYYY-MM-DD] [HH:MM:SS] [OG Abbr] [TZ Abbr]
timeto [YYYY/MM/DD] [HH:MM] [OG Abbr] [TZ Abbr]
timeto [YYYY-MM-DD] [HH:MM] [OG Abbr] [TZ Abbr]
```

## Examples

1) Just convert UTC (RFC3339) to JST
```
timeto 2023-04-22T09:49:59Z JST
```
> Conversion  
>  &ensp; Original time: 2023/04/22 09:49:59 UTC  
>  &ensp;&ensp;&ensp; Target time: 2023/04/22 18:49:59 JST  

2) Convert Original time UTC+9 to Target time UTC
```
timeto 2023/04/10 09:00 UTC+9
```
> Conversion  
>  &ensp; Original time: 2023/04/10 09:00 UTC+9  
>  &ensp;&ensp;&ensp;&ensp;&ensp; UTC time: 2023/04/10 00:00 UTC  

3) Convert Original time UTC-7 to Target time UTC+9
```
timeto 2023/04/21 12:59:59 UTC-7 UTC+9
```
> Conversion  
>  &ensp; Original time: 2023/04/21 12:59:59 UTC-7  
>  &ensp;&ensp;&ensp; Target time: 2023/04/22 04:59:59 UTC+9  

4) Convert Original time PDT to Target time NZST
```
timeto 2023-04-21 15:49:42 PDT NZST
```
> Conversion  
>  &ensp; Original time: 2023-04-21 15:49:42 PDT  
>  &ensp;&ensp;&ensp; Target time: 2023-04-22 10:49:42 NZST  


## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.
