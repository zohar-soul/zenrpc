package subarithservice

import (
	"github.com/zohar-soul/zenrpc/v2/testdata/objects"
	"time"
)

type Point struct {
	objects.AbstractObject
	A, B int        // coordinate
	C    int        `json:"-"`
	When *time.Time `json:"when"` // when it happened
}
