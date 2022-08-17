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
- Should a caching layer be added to avoid hitting the API too often (limit is 500r/m) ? Or should there be multiple weather data providers to loop through in case one hits the limit?
- Should I make a Rust (`mau-rs`) version of this just for learning?
- Add forecast for a specific time, e.g. "tomorrow"
- Add more data layers to the forecast
- Add algorithm to generate hour quality-score based on weather data
- Add tests
- Coordinates are a pain. It would be nice to add a geo-coder to convert a location to coordinates.

## Thoughts
- The `go-resty` package was a good choice for the HTTP client, it's easy to use and hides a lot of the boilerplate needed when making HTTP requests with Golang. We are now using the native `net/http` package instead for "purity".
