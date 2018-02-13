# Go Route Picker

I work in an industrial area, the fastest route home is backed up with traffic fairly frequently due to some road works along the way.

There is an alternative route I can take that is usually slower, but on days when traffic is backed up - it is quicker to go this way.

Leveraging the Google Maps Directions API - I can get Google to indicate the best route to take, factoring in traffic.

It has also served as a good opportunity to start toying with the Google Maps API in Go.

## Installation

1. `go get github.com/jryd/go-route-picker`
2. `cd /path-to-go/src/github.com/jryd/go-route-picker`
3. `go build`

## Running This

As this uses a `.env` file, you need to ensure that you run the executable from a directory that contains the `.env` file