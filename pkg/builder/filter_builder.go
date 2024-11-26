package builder

import (
	"fmt"
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	"github.com/CloudNativeGame/structured-filter-go/internal/utils"
	"reflect"
	"strconv"
	"strings"
)

type IFilterBuilder interface {
	// Or And must be placed before any other method call, and only one of them can be called, at most once.
	Or() IFilterBuilder
	And() IFilterBuilder
	// KStringV KBoolV KNumberV are used to build `$eq` filters of corresponding data types.
	KStringV(key string, value string) IFilterBuilder
	KBoolV(key string, value bool) IFilterBuilder
	KNumberV(key string, value float64) IFilterBuilder
	// KObjectV is used to build filters other than `$eq`.
	KObjectV(key string, value FilterBuilderObject) IFilterBuilder
	// Build is usually placed at the end of the chaining method calls and returns the filter string. It is reentrant.
	Build() string
	// Reset is used to clean all the filters in the `IFilterBuilder` if you need to build another filter string.
	Reset()
}

type FilterBuilder struct {
	sb    strings.Builder
	stack []byte
	layer int
}

func NewFilterBuilder() *FilterBuilder {
	b := &FilterBuilder{
		sb: strings.Builder{},
	}
	return b
}

func (b *FilterBuilder) Or() IFilterBuilder {
	if b.layer != 0 {
		panic(fmt.Errorf("%s should be in layer 0", consts.OrKey))
	}
	b.sb.WriteString("{\"")
	b.appendByte('}')
	b.sb.WriteString(consts.OrKey)
	b.sb.WriteString("\":[")
	b.appendByte(']')
	b.layer++
	return b
}

func (b *FilterBuilder) And() IFilterBuilder {
	if b.layer != 0 {
		panic(fmt.Errorf("%s should be in layer 0", consts.AndKey))
	}
	b.sb.WriteString("{\"")
	b.appendByte('}')
	b.sb.WriteString(consts.AndKey)
	b.sb.WriteString("\":[")
	b.appendByte(']')
	b.layer++
	return b
}

func (b *FilterBuilder) KStringV(key string, value string) IFilterBuilder {
	return b.kv(key, value)
}

func (b *FilterBuilder) KBoolV(key string, value bool) IFilterBuilder {
	return b.kv(key, value)
}

func (b *FilterBuilder) KNumberV(key string, value float64) IFilterBuilder {
	return b.kv(key, value)
}

func (b *FilterBuilder) KObjectV(key string, value FilterBuilderObject) IFilterBuilder {
	return b.kv(key, value)
}

func (b *FilterBuilder) Build() string {
	for len(b.stack) != 0 {
		if b.topByte() == ',' {
			b.popByte()
			continue
		}
		b.sb.WriteByte(b.popByte())
	}
	return b.sb.String()
}

func (b *FilterBuilder) Reset() {
	b.sb.Reset()
	b.stack = b.stack[:0]
	b.layer = 0
}

func (b *FilterBuilder) kv(key string, value interface{}) *FilterBuilder {
	if b.topByte() == ',' {
		b.sb.WriteByte(b.popByte())
	}
	b.sb.WriteString("{\"")
	b.sb.WriteString(key)
	b.sb.WriteString("\":")
	switch value.(type) {
	case string:
		b.writeStringValue(value.(string))
	case bool:
		b.writeBoolValue(value.(bool))
	case int:
		b.writeFloat64Value(value)
	case int8:
		b.writeFloat64Value(value)
	case int16:
		b.writeFloat64Value(value)
	case int32:
		b.writeFloat64Value(value)
	case int64:
		b.writeFloat64Value(value)
	case uint:
		b.writeFloat64Value(value)
	case uint8:
		b.writeFloat64Value(value)
	case uint16:
		b.writeFloat64Value(value)
	case uint32:
		b.writeFloat64Value(value)
	case uint64:
		b.writeFloat64Value(value)
	case float32:
		b.writeFloat64Value(value)
	case float64:
		b.writeFloat64Value(value)
	case []float64:
		b.writeFloat64ArrayValue(value.([]float64))
	case []string:
		b.writeStringArrayValue(value.([]string))
	case FilterBuilderObject:
		if len(value.(FilterBuilderObject)) != 1 {
			panic(fmt.Errorf("%v should be object with 1 kv", value))
		}
		for k, v := range value.(FilterBuilderObject) {
			b.kv(k, v)
		}
		b.sb.WriteByte('}')
	default:
		panic(fmt.Errorf("%v unsupported value type %v", value, reflect.TypeOf(value)))
	}
	return b
}

func (b *FilterBuilder) writeStringValue(value string) *FilterBuilder {
	b.sb.WriteByte('"')
	b.sb.WriteString(value)
	b.sb.WriteByte('"')
	b.sb.WriteByte('}')
	b.appendByte(',')
	return b
}

func (b *FilterBuilder) writeBoolValue(value bool) *FilterBuilder {
	if value {
		b.sb.WriteString("true")
	} else {
		b.sb.WriteString("false")
	}
	b.sb.WriteByte('}')
	b.appendByte(',')

	return b
}

func (b *FilterBuilder) writeFloat64Value(value interface{}) *FilterBuilder {
	b.sb.WriteString(strconv.FormatFloat(utils.NumberToFloat64(value), 'f', -1, 64))
	b.sb.WriteByte('}')
	b.appendByte(',')
	return b
}

func (b *FilterBuilder) writeFloat64ArrayValue(value []float64) *FilterBuilder {
	b.sb.WriteByte('[')
	for i, v := range value {
		if i != 0 {
			b.sb.WriteByte(',')
		}
		b.sb.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
	}
	b.sb.WriteByte(']')
	b.sb.WriteByte('}')
	b.appendByte(',')
	return b
}

func (b *FilterBuilder) writeStringArrayValue(value []string) *FilterBuilder {
	b.sb.WriteByte('[')
	for i, v := range value {
		if i != 0 {
			b.sb.WriteByte(',')
		}
		b.sb.WriteByte('"')
		b.sb.WriteString(v)
		b.sb.WriteByte('"')
	}
	b.sb.WriteByte(']')
	b.sb.WriteByte('}')
	b.appendByte(',')
	return b
}

func (b *FilterBuilder) appendByte(c byte) {
	b.stack = append(b.stack, c)
}

func (b *FilterBuilder) popByte() byte {
	c := b.stack[len(b.stack)-1]
	b.stack = b.stack[:len(b.stack)-1]
	return c
}

func (b *FilterBuilder) topByte() byte {
	if len(b.stack) == 0 {
		return 0
	}
	return b.stack[len(b.stack)-1]
}
