# exceltool

Usage of exceltool.exe:

  -axis string
  -file string
  -sheet string
  -type string
  -value string


exceltool.bat
```
@echo off
exceltool.exe -file="test.xlsx" -sheet="Sheet1" -axis="A1" -type="text" -value="text"
exceltool.exe -file="test.xlsx" -sheet="Sheet1" -axis="A2" -type="number" -value="11"
exceltool.exe -file="test.xlsx" -sheet="Sheet1" -axis="A3" -type="number" -value="22"
exceltool.exe -file="test.xlsx" -sheet="Sheet1" -axis="A4" -type="formula" -value="=a2+a3"
```
