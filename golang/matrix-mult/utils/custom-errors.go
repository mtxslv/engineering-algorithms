package utils

import "errors"

var ErrMalformedMatrix = errors.New("Matrix has malformed dimensions.")
var ErrBadDimensions = errors.New("Matrices cannot be multiplied given dimensions.")