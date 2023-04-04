# tmpchk
A minimal CLI weather application, which connects to the [Open-Meteo.com](https://open-meteo.com/) 3rd-party Weather API.
This is an application inteded for learning purposes for beginners with Go.

## Installation
```
$ go get github.com/chrisdaly3/tmpchk
& cd $GOPATH/src/github.com/chrisdaly3/tmpchk && go install
```

## Use
`$ tmpchk -l (or --location) <CITY>`
* Results will appear as so:
```
Showing Results for Buffalo, Erie County, New York, United States
Latitude: 42.886612, Longitude: -78.878174
Timezone: America/New_York
Current Weather - Temp: 54.2 Fahrenheit
Current Weather - Windspeed: 11.1 mph
```


