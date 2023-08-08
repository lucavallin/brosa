# <img src="https://github.com/lucavallin/brosa/blob/main/docs/brosa.png" style="width:40px;padding-right:10px;margin-bottom:-8px;"> brósa

A CLI tool written in Golang to retrieve information useful for astronomy.

## Setup

Build `brosa` with the following command:

```bash
go build
```

Then initializa configuration (e.g. API keys...) with `brosa init`:

```bash
brosa init

INFO  initializing brosa configuration

Enter your Tomorrow.io API key: XXXXXXXXXXXXXXXXXXXX

Enter your IPGeolocation.com API key: XXXXXXXXXXXXXXXXXXXX

SUCCESS  brosa configuration initialized
```

## Usage

- `brosa forecast`: Get the weather forecast relevant for astronomy for a set or coordinates. Example:

```bash
brosa forecast 18.955324,69.649208

┌────────────────────────────────────────────────────────────────────────────────────────┐
|             Date | Cloud Cover (%) | Humidity (%) | Temperature (ºC) | Visibility (km) |
| 2022-08-19 19:00 |              59 |           85 |               27 |              11 |
| 2022-08-19 20:00 |              20 |           86 |               27 |              11 |
| 2022-08-19 21:00 |              35 |           86 |               27 |              10 |
| 2022-08-19 22:00 |              32 |           86 |               26 |              11 |
| 2022-08-19 23:00 |              25 |           87 |               26 |              10 |
| 2022-08-20 00:00 |              86 |           87 |               26 |              10 |
| 2022-08-20 01:00 |              98 |           86 |               27 |              11 |
| 2022-08-20 02:00 |              92 |           85 |               27 |              11 |
| 2022-08-20 03:00 |             100 |           85 |               27 |              11 |
| 2022-08-20 04:00 |             100 |           85 |               27 |              10 |
| 2022-08-20 05:00 |              98 |           84 |               27 |              10 |
| 2022-08-20 06:00 |              98 |           84 |               27 |              11 |
| 2022-08-20 07:00 |              99 |           84 |               27 |              12 |
| 2022-08-20 08:00 |              98 |           85 |               27 |              12 |
| 2022-08-20 09:00 |              41 |           85 |               27 |              11 |
| 2022-08-20 10:00 |              34 |           85 |               27 |              11 |
| 2022-08-20 11:00 |              53 |           85 |               27 |              11 |
| 2022-08-20 12:00 |              43 |           86 |               27 |              11 |
| 2022-08-20 13:00 |              49 |           85 |               27 |              11 |
| 2022-08-20 14:00 |              93 |           85 |               27 |              10 |
| 2022-08-20 15:00 |              99 |           85 |               27 |              11 |
| 2022-08-20 16:00 |             100 |           84 |               27 |              10 |
| 2022-08-20 17:00 |             100 |           85 |               27 |              11 |
| 2022-08-20 18:00 |             100 |           86 |               27 |              11 |
| 2022-08-20 19:00 |             100 |           86 |               27 |              11 |
└────────────────────────────────────────────────────────────────────────────────────────┘
```

- `brosa locate`: Get the coordinates of a place. Example:

```bash
brosa locate tromso

SUCCESS  6 coordinate(s) found!
┌─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
| Name                                                                                                      | Latitude  | Longitude | Latitude, Longitude |
| Tromsø, Troms og Finnmark, Norge                                                                          | 69.651648 | 18.955819 | 69.651648,18.955819 |
| Tromso, Wenchi Municipal District, Bono Region, Ghana                                                     | 7.558820  | -2.163072 | 7.558820,-2.163072  |
| Tromsö, Haparanda kommun, Norrbottens län, Sverige                                                        | 65.712138 | 23.761817 | 65.712138,23.761817 |
| Sentrum legekontor, Tromsø, Killengreens gate, Skarpsno, Nordbyen, Tromsø, Troms og Finnmark, 9008, Norge | 69.649739 | 18.959848 | 69.649739,18.959848 |
| Utleiecompagniet AS, Tromsø, Gimlevegen, Gimle, Tromsø, Troms og Finnmark, 9019, Norge                    | 69.684069 | 18.990678 | 69.684069,18.990678 |
| Harila, Tromsø, Skattørvegen, Skattøra, Tromsø, Troms og Finnmark, 9018, Norge                            | 69.696052 | 19.013815 | 69.696052,19.013815 |
└─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
```

- `brosa iss`: Get the current position of the International Space Station. Example:

```bash
brosa iss

SUCCESS  ISS Found
┌─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
| Time             | Latitude | Longitude   | Altitude (km) | Velocity (km/h) | Visibility | Solar latitude | Solar longitude |
| 2022-08-21 16:59 | 5.319085 | -178.143555 | 419           | 27576           | eclipsed   | 11.977992      | 315.955918      |
└─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
```

- `brosa day`: Get current information about the Sun and the Moon for a given location. Example:

```bash
brosa day 45.806691,12.206316

┌────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
| Sunrise | Sunset | Day length (h) | Sun altitude (°) | Sun azimuth (°) | Moonrise | Moonset | Moon altitude (°) | Moon azimuth (°) |
|   06:42 |  19:35 |          12:53 |              -15 |             294 |    19:38 |   04:58 |                12 |              123 |
└────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
```

## Improvements
- Implement `visible` command (https://docs.astronomyapi.com)
- Add support for "best" timestep (or devise algorithm to calculate best time for observing)
- Cache weather results for a set of coordinates for one hour
- Add dashboard with all info that refreshes every hour and sends events to the user
- Add unit/integration tests
- Decouple `astro` providers like for `weather` and `geo`
