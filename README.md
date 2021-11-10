# GetCountryFromIP

Simple go executable to task an ip and get the results from ip-api
https://ip-api.com/
https://ip-api.com/docs/api:json

This is a fantastic site and I hope the continue to provide this

Allows flag -p to replace the properties output

## Usage Example

GetCountryFromIP.exe -ip 8.8.8.8

Output:

United States

GetCountryFromIP.exe -ip 8.8.8.8 -p country,isp

Output:

United States,Google LLC
