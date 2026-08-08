# timeto

![Go](https://img.shields.io/badge/Go-1.25+-informational)
![License](https://img.shields.io/badge/license-MIT-blue)

timeto is a simple command-line tool for time difference calculation.

![clock](images/clock.png)

## Getting Started

```bash or powershell
go install github.com/kazsk1/timeto@main
```

## Usage
1. Convert UTC (RFC3339) to target timezone abbreviations such as JST, SGT, etc.
   ```bash or powreshell
   timeto [YYYY-MM-DDTHH:MM:SSZ] [TG Abbr]
   ```

2. Convert UTC offset to UTC.
   ```bash or powershell
   timeto [YYYY/MM/DD] [HH:MM:SS] [OG UTC offset]
   timeto [YYYY-MM-DD] [HH:MM:SS] [OG UTC offset]
   timeto [YYYY/MM/DD] [HH:MM] [OG UTC offset]
   timeto [YYYY-MM-DD] [HH:MM] [OG UTC offset]
   ```

3. Convert original UTC offset to target UTC offset.
   ```bash or powershell
   timeto [YYYY/MM/DD] [HH:MM:SS] [OG UTC offset] [TG UTC offset]
   timeto [YYYY-MM-DD] [HH:MM:SS] [OG UTC offset] [TG UTC offset]
   timeto [YYYY/MM/DD] [HH:MM] [OG UTC offset] [TG UTC offset]
   timeto [YYYY-MM-DD] [HH:MM] [OG UTC offset] [TG UTC offset]
   ```

4. Convert original timezone abbreviations to target timezone abbreviations.
   ```bash or powershell
   timeto [YYYY/MM/DD] [HH:MM:SS] [OG Abbr] [TG Abbr]
   timeto [YYYY-MM-DD] [HH:MM:SS] [OG Abbr] [TG Abbr]
   timeto [YYYY/MM/DD] [HH:MM] [OG Abbr] [TG Abbr]
   timeto [YYYY-MM-DD] [HH:MM] [OG Abbr] [TG Abbr]
   ```

5. Convert UNIX time to UTC.
   ```bash or powershell
   timeto [UNIX time] 
   ```

## Examples

1. Convert UTC (RFC3339) to JST.  
   ```bash or powershell
   timeto 2023-04-22T09:49:59Z JST
   ```
   > Conversion  
   >  Original time: 2023/04/22 09:49:59 UTC  
   >    Target time: 2023/04/22 18:49:59 JST  

2. Convert UTC+9 to UTC.  
   ```bash or powershell
   timeto 2023/04/10 09:00 UTC+9
   ```
   > Conversion  
   >  Original time: 2023/04/10 09:00 UTC+9  
   >       UTC time: 2023/04/10 00:00 UTC  

3. Convert UTC-7 to UTC+9.
   ```bash or powershell
   timeto 2023/04/21 12:59:59 UTC-7 UTC+9
   ```
   > Conversion  
   >  Original time: 2023/04/21 12:59:59 UTC-7  
   >    Target time: 2023/04/22 04:59:59 UTC+9  

4. Convert PDT to NZST.
   ```bash or powershell
   timeto 2023-04-21 15:49:42 PDT NZST
   ```
   > Conversion  
   >  Original time: 2023-04-21 15:49:42 PDT  
   >    Target time: 2023-04-22 10:49:42 NZST  

5. Convert UNIX time to UTC.
   ```bash or powershell
   timeto 1718000000123
   ```
   > Conversion
   >  UNIX time: 1718000000123
   >   UTC time: 2024-06-10 06:13:20.123 UTC


## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.
