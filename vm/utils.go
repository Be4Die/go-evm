// vm/utils.go
package vm

import "math"

// Float32ToUint32 converts float32 to uint32 representation
func Float32ToUint32(f float32) uint32 {
	return math.Float32bits(f)
}

// Uint32ToFloat32 converts uint32 representation to float32
func Uint32ToFloat32(u uint32) float32 {
	return math.Float32frombits(u)
}