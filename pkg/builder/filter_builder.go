package builder

import (
	"fmt"
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	"reflect"
	"strconv"
	"strings"
)

type IFilterBuilder interface {
	Or() IFilterBuilder
	And() IFilterBuilder
	KStringV(key string, value string) IFilterBuilder
	KBoolV(key string, value bool) IFilterBuilder
	KNumberV(key string, value float64) IFilterBuilder
	KObjectV(key string, value FilterBuilderObject) IFilterBuilder
	Build() string
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

func toFloat64(num interface{}) float64 {
	numType := reflect.TypeOf(num)
	numValue := reflect.ValueOf(num)

	switch numType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(numValue.Int())
	case reflect.Float32, reflect.Float64:
		return numValue.Float()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(numValue.Uint())
	default:
		panic(fmt.Errorf("unsupported number type %v for %v", numType, numValue))
	}
}

func (b *FilterBuilder) writeFloat64Value(value interface{}) *FilterBuilder {
	b.sb.WriteString(strconv.FormatFloat(toFloat64(value), 'f', -1, 64))
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