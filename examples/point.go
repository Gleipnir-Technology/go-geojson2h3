package main

import (
	"github.com/Gleipnir-Technology/go-geojson2h3/v2"
	"github.com/tidwall/geojson"
	"github.com/uber/h3-go/v4"
)

func pointToH3(res int) ([]h3.Cell, error) {
	o, err := geojson.Parse(`{
    "type": "Point",
    "coordinates": [
        -105.01621,
        39.57422
    ]
}`, nil)
	checkError(err)
	return geojson2h3.ToH3(res, o)
}

func multiPointToH3(res int) ([]h3.Cell, error) {
	o, err := geojson.Parse(`{
    "type": "MultiPoint",
    "coordinates": [
        [
            -105.01621,
            39.57422
        ],
        [
            -80.666513,
            35.053994
        ]
    ]
}`, nil)
	checkError(err)
	return geojson2h3.ToH3(res, o)
}
