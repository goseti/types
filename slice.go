package types

import (
    "math/rand"
)

type Slice []Any

func (s Slice)Rand() Any {
    return s[rand.Intn(len(s))]
}
