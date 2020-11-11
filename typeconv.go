package typeconv

import (
	"strconv"
)

//Intc converts input to int
func Intc(a interface{}) (i int) {
	switch b := a.(type) {
	case string:
		c, _ := strconv.ParseInt(b, 10, 64)
		i = int(c)
	case bool:
		if b == true {
			i = 1
		} else {
			i = 0
		}
	case int8:
		i = int(b)
	case int16:
		i = int(b)
	case int32:
		i = int(b)
	case int64:
		i = int(b)
	case int:
		i = b
	default:
		i = 0
	}
	return
}

//Uintc converts input to uint
func Uintc(a interface{}) (i uint) {
	switch b := a.(type) {
	case string:
		c, _ := strconv.ParseUint(b, 10, 64)
		i = uint(c)
	case bool:
		if b == true {
			i = 1
		} else {
			i = 0
		}
	case uint8:
		i = uint(b)
	case uint16:
		i = uint(b)
	case uint32:
		i = uint(b)
	case uint64:
		i = uint(b)
	case uint:
		i = b
	default:
		i = 0
	}
	return
}

//Stringc converts input to string
func Stringc(a interface{}) (s string) {
	switch b := a.(type) {
	case string:
		s = b
	case bool:
		if b == true {
			s = "true"
		} else {
			s = "false"
		}
	case int8:
		s = strconv.FormatInt(int64(b), 10)
	case int16:
		s = strconv.FormatInt(int64(b), 10)
	case int32:
		s = strconv.FormatInt(int64(b), 10)
	case int64:
		s = strconv.FormatInt(b, 10)
	case int:
		s = strconv.FormatInt(int64(b), 10)
	case uint8:
		s = strconv.FormatUint(uint64(b), 10)
	case uint16:
		s = strconv.FormatUint(uint64(b), 10)
	case uint32:
		s = strconv.FormatUint(uint64(b), 10)
	case uint64:
		s = strconv.FormatUint(b, 10)
	case uint:
		s = strconv.FormatUint(uint64(b), 10)
	case float32:
		s = strconv.FormatFloat(float64(b), 'f', -1, 32)
	case float64:
		s = strconv.FormatFloat(b, 'f', -1, 64)
	case complex64:
		s = strconv.FormatComplex(complex128(b), 'f', -1, 64)
	case complex128:
		s = strconv.FormatComplex(b, 'f', -1, 128)
	default:
		s = ""
	}
	return
}

//Boolc converts input to bool
func Boolc(a interface{}) (b bool) {
	switch c := a.(type) {
	case string:
		if c != "" {
			b = true
		} else {
			b = false
		}
	case bool:
		b = c
	case int8:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case int16:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case int32:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case int64:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case int:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case uint8:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case uint16:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case uint32:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case uint64:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case uint:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case float32:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case float64:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case complex64:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	case complex128:
		if c == 0 {
			b = false
		} else {
			b = true
		}
	default:
		b = false
	}
	return
}

//Float64c converts input to float64
func Float64c(a interface{}) (f float64) {
	switch b := a.(type) {
	case string:
		f, _ = strconv.ParseFloat(b, 64)
	case bool:
		if b == true {
			f = 1
		} else {
			f = 0
		}
	case int8:
		f = float64(b)
	case int16:
		f = float64(b)
	case int32:
		f = float64(b)
	case int64:
		f = float64(b)
	case int:
		f = float64(b)
	case uint8:
		f = float64(b)
	case uint16:
		f = float64(b)
	case uint32:
		f = float64(b)
	case uint64:
		f = float64(b)
	case uint:
		f = float64(b)
	case float32:
		f = float64(b)
	case float64:
		f = b
	default:
		f = 0
	}
	return
}

//Float32c converts input to float32
func Float32c(a interface{}) (f float32) {
	switch b := a.(type) {
	case string:
		c, _ := strconv.ParseFloat(b, 32)
		f = float32(c)
	case bool:
		if b == true {
			f = 1
		} else {
			f = 0
		}
	case int8:
		f = float32(b)
	case int16:
		f = float32(b)
	case int32:
		f = float32(b)
	case int64:
		f = float32(b)
	case int:
		f = float32(b)
	case uint8:
		f = float32(b)
	case uint16:
		f = float32(b)
	case uint32:
		f = float32(b)
	case uint64:
		f = float32(b)
	case uint:
		f = float32(b)
	case float32:
		f = b
	default:
		f = 0
	}
	return
}

//Complex64c converts input to complex64
func Complex64c(a interface{}) (c complex64) {
	switch b := a.(type) {
	case string:
		d, _ := strconv.ParseComplex(b, 64)
		c = complex64(d)
	case bool:
		if b == true {
			c = 1
		} else {
			c = 0
		}
	case int8:
		c = complex(float32(b), 0)
	case int16:
		c = complex(float32(b), 0)
	case int32:
		c = complex(float32(b), 0)
	case int64:
		c = complex(float32(b), 0)
	case int:
		c = complex(float32(b), 0)
	case uint8:
		c = complex(float32(b), 0)
	case uint16:
		c = complex(float32(b), 0)
	case uint32:
		c = complex(float32(b), 0)
	case uint64:
		c = complex(float32(b), 0)
	case uint:
		c = complex(float32(b), 0)
	case float32:
		c = complex(b, 0)
	case complex64:
		c = b
	default:
		c = 0
	}
	return
}

//Complex128c converts input to complex128
func Complex128c(a interface{}) (c complex128) {
	switch b := a.(type) {
	case string:
		d, _ := strconv.ParseComplex(b, 128)
		c = d
	case bool:
		if b == true {
			c = 1
		} else {
			c = 0
		}
	case int8:
		c = complex(float64(b), 0)
	case int16:
		c = complex(float64(b), 0)
	case int32:
		c = complex(float64(b), 0)
	case int64:
		c = complex(float64(b), 0)
	case int:
		c = complex(float64(b), 0)
	case uint8:
		c = complex(float64(b), 0)
	case uint16:
		c = complex(float64(b), 0)
	case uint32:
		c = complex(float64(b), 0)
	case uint64:
		c = complex(float64(b), 0)
	case uint:
		c = complex(float64(b), 0)
	case float32:
		c = complex(float64(b), 0)
	case complex64:
		c = complex128(b)
	case complex128:
		c = b
	default:
		c = 0
	}
	return
}

//Int8c converts input to int8
func Int8c(a interface{}) (i int8) {
	switch b := a.(type) {
	case string:
		c, _ := strconv.ParseInt(b, 10, 8)
		i = int8(c)
	case bool:
		if b == true {
			i = 1
		} else {
			i = 0
		}
	case int8:
		i = b
	default:
		i = 0
	}
	return
}

//Int16c converts input to int16
func Int16c(a interface{}) (i int16) {
	switch b := a.(type) {
	case string:
		c, _ := strconv.ParseInt(b, 10, 16)
		i = int16(c)
	case bool:
		if b == true {
			i = 1
		} else {
			i = 0
		}
	case int8:
		i = int16(b)
	case int16:
		i = b
	default:
		i = 0
	}
	return
}

//Int32c converts input to int32
func Int32c(a interface{}) (i int32) {
	switch b := a.(type) {
	case string:
		c, _ := strconv.ParseInt(b, 10, 32)
		i = int32(c)
	case bool:
		if b == true {
			i = 1
		} else {
			i = 0
		}
	case int8:
		i = int32(b)
	case int16:
		i = int32(b)
	case int32:
		i = b
	default:
		i = 0
	}
	return
}

//Int64c converts input to int64
func Int64c(a interface{}) (i int64) {
	switch b := a.(type) {
	case string:
		c, _ := strconv.ParseInt(b, 10, 64)
		i = c
	case bool:
		if b == true {
			i = 1
		} else {
			i = 0
		}
	case int8:
		i = int64(b)
	case int16:
		i = int64(b)
	case int32:
		i = int64(b)
	case int64:
		i = b
	case int:
		i = int64(b)
	default:
		i = 0
	}
	return
}

//Uint8c converts input to uint8
func Uint8c(a interface{}) (i uint8) {
	switch b := a.(type) {
	case string:
		c, _ := strconv.ParseUint(b, 10, 8)
		i = uint8(c)
	case bool:
		if b == true {
			i = 1
		} else {
			i = 0
		}
	case uint8:
		i = b
	default:
		i = 0
	}
	return
}

//Uint16c converts input to uint16
func Uint16c(a interface{}) (i uint16) {
	switch b := a.(type) {
	case string:
		c, _ := strconv.ParseUint(b, 10, 16)
		i = uint16(c)
	case bool:
		if b == true {
			i = 1
		} else {
			i = 0
		}
	case uint8:
		i = uint16(b)
	case uint16:
		i = b
	default:
		i = 0
	}
	return
}

//Uint32c converts input to uint32
func Uint32c(a interface{}) (i uint32) {
	switch b := a.(type) {
	case string:
		c, _ := strconv.ParseUint(b, 10, 32)
		i = uint32(c)
	case bool:
		if b == true {
			i = 1
		} else {
			i = 0
		}
	case uint8:
		i = uint32(b)
	case uint16:
		i = uint32(b)
	case uint32:
		i = b
	default:
		i = 0
	}
	return
}

//Uint64c converts input to int64
func Uint64c(a interface{}) (i uint64) {
	switch b := a.(type) {
	case string:
		c, _ := strconv.ParseUint(b, 10, 64)
		i = c
	case bool:
		if b == true {
			i = 1
		} else {
			i = 0
		}
	case uint8:
		i = uint64(b)
	case uint16:
		i = uint64(b)
	case uint32:
		i = uint64(b)
	case uint64:
		i = b
	case uint:
		i = uint64(b)
	default:
		i = 0
	}
	return
}

//Scontains return if there is in slice
func Scontains(a interface{}, b interface{}) (i bool) {
	switch c := a.(type) {
	case []string:
		d := b.(string)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []int:
		d := b.(int)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []bool:
		d := b.(bool)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []int8:
		d := b.(int8)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []int16:
		d := b.(int16)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []int32:
		d := b.(int32)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []int64:
		d := b.(int64)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []uint8:
		d := b.(uint8)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []uint16:
		d := b.(uint16)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []uint32:
		d := b.(uint32)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []uint64:
		d := b.(uint64)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []float32:
		d := b.(float32)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []float64:
		d := b.(float64)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []complex64:
		d := b.(complex64)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	case []complex128:
		d := b.(complex128)
		for _, x := range c {
			if x == d {
				i = true
				break
			}
		}
	default:
		i = false
	}
	return
}

//Sindex returns index of passed value
func Sindex(a interface{}, b interface{}) (i int) {
	//-1 return means not found
	//the function will panic if a and b whose type are different.
	i = -1
	switch c := a.(type) {
	case []string:
		d := b.(string)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []int:
		d := b.(int)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []bool:
		d := b.(bool)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []int8:
		d := b.(int8)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []int16:
		d := b.(int16)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []int32:
		d := b.(int32)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []int64:
		d := b.(int64)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []uint:
		d := b.(uint)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []uint8:
		d := b.(uint8)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []uint16:
		d := b.(uint16)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []uint32:
		d := b.(uint32)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []uint64:
		d := b.(uint64)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []float32:
		d := b.(float32)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []float64:
		d := b.(float64)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []complex64:
		d := b.(complex64)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	case []complex128:
		d := b.(complex128)
		for y, x := range c {
			if x == d {
				i = y
				break
			}
		}
	default:
		i = -1
	}
	return
}

//Slastindex returns index of passed value
func Slastindex(a interface{}, b interface{}) (i int) {
	//-1 return means not found
	//the function will panic if a and b whose type are different.
	i = -1
	switch c := a.(type) {
	case []string:
		d := b.(string)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []int:
		d := b.(int)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []bool:
		d := b.(bool)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []int8:
		d := b.(int8)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []int16:
		d := b.(int16)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []int32:
		d := b.(int32)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []int64:
		d := b.(int64)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []uint:
		d := b.(uint)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []uint8:
		d := b.(uint8)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []uint16:
		d := b.(uint16)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []uint32:
		d := b.(uint32)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []uint64:
		d := b.(uint64)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []float32:
		d := b.(float32)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []float64:
		d := b.(float64)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []complex64:
		d := b.(complex64)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	case []complex128:
		d := b.(complex128)
		index := len(c) - 1
		for index >= 0 {
			if c[index] == d {
				i = index
				break
			}
			index = index - 1
		}
	default:
		i = -1
	}
	return
}

//Scount returns how many passed value
func Scount(a interface{}, b interface{}) (i int) {
	switch c := a.(type) {
	case []string:
		d := b.(string)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []int:
		d := b.(int)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []bool:
		d := b.(bool)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []int8:
		d := b.(int8)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []int16:
		d := b.(int16)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []int32:
		d := b.(int32)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []int64:
		d := b.(int64)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []uint:
		d := b.(uint)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []uint8:
		d := b.(uint8)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []uint16:
		d := b.(uint16)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []uint32:
		d := b.(uint32)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []uint64:
		d := b.(uint64)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []float32:
		d := b.(float32)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []float64:
		d := b.(float64)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []complex64:
		d := b.(complex64)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	case []complex128:
		d := b.(complex128)
		for _, x := range c {
			if x == d {
				i = i + 1
			}
		}
	default:
		i = 0
	}
	return
}
