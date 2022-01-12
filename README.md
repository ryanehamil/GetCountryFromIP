# lookupip

Simple go executable to task an ip and get the results from ip-api
https://ip-api.com/

This is a fantastic site and I hope the continue to provide this

Allows flag -p to replace the properties output

## Build

I build this on Windows with

    go build -o bin/

## Usage

> **Default (Country Name)**
>
> lookupip.exe -ip 8.8.8.8
>
> **Output:**
>
> United States

> **Custom return (Country Name and ISP)**
>
> lookupip.exe -ip 8.8.8.8 -p country,isp
>
> **Output:**
>
> United States,Google LLC

## Custom Properties

https://ip-api.com/docs/api:json

| name          | description                                | example             | type   |
| ------------- | ------------------------------------------ | ------------------- | ------ |
| Continent     | Continent name                             | North America       | string |
| ContinentCode | Two-letter continent code                  | NA                  | string |
| Country       | Country name                               | United States       | string |
| CountryCode   | Two-letter country code ISO 3166-1 alpha-2 | US                  | string |
| Region        | Region/state short code (FIPS or ISO)      | CA or 10            | string |
| RegionName    | Region/state                               | California          | string |
| City          | City                                       | Mountain View       | string |
| District      | District (subdivision of city)             | Old Farm District   | string |
| Zip           | Zip code                                   | 94043               | string |
| Lat           | Latitude                                   | 37.4192             | float  |
| Lon           | Longitude                                  | -122.0574           | float  |
| Timezone      | Timezone (tz)                              | America/Los_Angeles | string |
| Offset        | Timezone UTC DST offset in seconds         | -25200              | int    |
| Currency      | National currency                          | USD                 | string |
| ISP           | ISP name                                   | Google              | string |
| Org           | Organization name                          | Google              | string |
| AS            | AS number and organization (RIR).          | AS15169 Google Inc. | string |
| ASName        | AS name (RIR).                             | GOOGLE              | string |
| Reverse       | Reverse DNS of the IP (can delay response) | wi-in-f94.1e100.net | string |
| Hosting       | Hosting, colocated or data center          | true                | bool   |
| Query         | IP used for the query                      | 173.194.67.94       | string |
