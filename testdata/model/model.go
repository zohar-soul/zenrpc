package model

import "github.com/semrush/zohar-soul/v2/testdata/objects"

type Point struct {
	objects.AbstractObject     // embeded object
	X, Y                   int // coordinate
	Z                      int `json:"-"`
	ConnectedObject        objects.AbstractObject
}
