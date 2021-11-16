# lookupip

Simple go executable to task an ip and get the results from ip-api
https://ip-api.com/

This is a fantastic site and I hope the continue to provide this

Allows flag -p to replace the properties output

## Usage

> **Default (Country Name)**
> lookupip.exe -ip 8.8.8.8
> **Output:**
> United States

> **Custom return (Country Name and ISP)**
> lookupip.exe -ip 8.8.8.8 -p country,isp
> **Output:**
> United States,Google LLC

## Custom Properties

https://ip-api.com/docs/api:json

| name          | description                                | example             | type   |
| ------------- | ------------------------------------------ | ------------------- | ------ |
| continent     | Continent name                             | North America       | string |
| continentCode | Two-letter continent code                  | NA                  | string |
| country       | Country name                               | United States       | string |
| countryCode   | Two-letter country code ISO 3166-1 alpha-2 | US                  | string |
| region        | Region/state short code (FIPS or ISO)      | CA or 10            | string |
| regionName    | Region/state                               | California          | string |
| city          | City                                       | Mountain View       | string |
| district      | District (subdivision of city)             | Old Farm District   | string |
| zip           | Zip code                                   | 94043               | string |
| lat           | Latitude                                   | 37.4192             | float  |
| lon           | Longitude                                  | -122.0574           | float  |
| timezone      | Timezone (tz)                              | America/Los_Angeles | string |
| offset        | Timezone UTC DST offset in seconds         | -25200              | int    |
| currency      | National currency                          | USD                 | string |
| isp           | ISP name                                   | Google              | string |
| org           | Organization name                          | Google              | string |
| as            | AS number and organization (RIR).          | AS15169 Google Inc. | string |
| asname        | AS name (RIR).                             | GOOGLE              | string |
| reverse       | Reverse DNS of the IP (can delay response) | wi-in-f94.1e100.net | string |
| hosting       | Hosting, colocated or data center          | true                | bool   |
| query         | IP used for the query                      | 173.194.67.94       | string |
