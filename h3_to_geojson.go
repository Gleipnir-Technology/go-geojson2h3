package geojson2h3

import (
	"fmt"
	"strconv"

	"github.com/tidwall/geojson"
	"github.com/tidwall/geojson/geometry"
	"github.com/uber/h3-go/v4"
)

// ToFeatureCollection converts a set of hexagons to a GeoJSON `FeatureCollection`
// with the set outline(s). The feature's geometry type will be `Polygon`.
func ToFeatureCollection(indexes []h3.Cell) (*geojson.FeatureCollection, error) {
	if len(indexes) == 0 {
		return nil, fmt.Errorf("uber h3 indexes are empty")
	}
	features := make([]geojson.Object, 0, len(indexes))
	for _, index := range indexes {
		boundary, err := index.Boundary()
		if err != nil {
			return nil, err
		}
		points := make([]geometry.Point, 0, 6)
		for _, b := range boundary {
			points = append(points, geometry.Point{
				X: b.Lng,
				Y: b.Lat,
			})
		}
		points = append(points, geometry.Point{
			X: points[0].X,
			Y: points[0].Y,
		})
		polygon := geojson.NewPolygon(
			geometry.NewPoly(points, nil, &geometry.IndexOptions{
				Kind: geometry.None,
			}))
		feature := geojson.NewFeature(polygon, toH3Props(index))
		features = append(features, feature)
	}
	return geojson.NewFeatureCollection(features), nil
}

func toH3Props(index h3.Cell) string {
	res := strconv.Itoa(index.Resolution())
	return `{"h3index":"` + index.String() + `", "h3resolution": ` + res + `}`
}
