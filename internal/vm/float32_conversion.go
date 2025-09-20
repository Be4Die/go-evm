package vm

import "math"

// Float32ToUint32 преобразует значение float32 в его битовое представление uint32.
// Возвращает 32-битное беззнаковое целое число, которое представляет биты числа с плавающей точкой.
func Float32ToUint32(f float32) uint32 {
	return math.Float32bits(f)
}

// Uint32ToFloat32 преобразует битовое представление uint32 обратно в float32.
// Принимает 32-битное беззнаковое целое число и возвращает соответствующее число с плавающей точкой.
func Uint32ToFloat32(u uint32) float32 {
	return math.Float32frombits(u)
}