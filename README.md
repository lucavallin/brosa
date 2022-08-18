# mau
Mighty Astronomical Utility. A CLI tool written in Golang to retrieve information useful for astronomy.

![](docs/mau.png)


## Usage

This utility is a work in progress!
To get the forecast for a location, run the following:

```bash
go build
./mau forecast <COORDINATES> --tomorrowio-key=<YOUR_TOMORROW_IO_API_KEY> --end-time=nowPlus24h
```

## TODOs
- Add forecast for a specific time, e.g. "tomorrow"
- Add more data layers to the forecast (precipitation, dew point, moonphase, bortle class, sun/no-sun)
- Add algorithm to generate hour quality-score based on weather data
- Add tests

## Thoughts
- The `go-resty` package was a good choice for the HTTP client, it's easy to use and hides a lot of the boilerplate needed when making HTTP requests with Golang. We are now using the native `net/http` package instead for "purity".
