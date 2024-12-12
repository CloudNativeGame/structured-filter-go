package builder

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/builder"
)

type PlayerFilterBuilder struct {
	filterBuilder builder.IFilterBuilder
}

func NewPlayerFilterBuilder() *PlayerFilterBuilder {
	return &PlayerFilterBuilder{
		filterBuilder: builder.NewFilterBuilder(),
	}
}

func (p *PlayerFilterBuilder) Or() *PlayerFilterBuilder {
	p.filterBuilder.Or()
	return p
}

func (p *PlayerFilterBuilder) And() *PlayerFilterBuilder {
	p.filterBuilder.And()
	return p
}

func (p *PlayerFilterBuilder) KStringV(key string, value string) *PlayerFilterBuilder {
	p.filterBuilder.KStringV(key, value)
	return p
}

func (p *PlayerFilterBuilder) KBoolV(key string, value bool) *PlayerFilterBuilder {
	p.filterBuilder.KBoolV(key, value)
	return p
}

func (p *PlayerFilterBuilder) KNumberV(key string, value float64) *PlayerFilterBuilder {
	p.filterBuilder.KNumberV(key, value)
	return p
}

func (p *PlayerFilterBuilder) KObjectV(key string, value builder.FilterBuilderObject) *PlayerFilterBuilder {
	p.filterBuilder.KObjectV(key, value)
	return p
}

func (p *PlayerFilterBuilder) Build() string {
	return p.filterBuilder.Build()
}

func (p *PlayerFilterBuilder) Reset() {
	p.filterBuilder.Reset()
}

func (p *PlayerFilterBuilder) IsMale(isMale bool) *PlayerFilterBuilder {
	p.filterBuilder.KBoolV("isMale", isMale)
	return p
}

func (p *PlayerFilterBuilder) IsMaleObject(obj builder.FilterBuilderObject) *PlayerFilterBuilder {
	p.filterBuilder.KObjectV("isMale", obj)
	return p
}

func (p *PlayerFilterBuilder) UserName(userName string) *PlayerFilterBuilder {
	p.filterBuilder.KStringV("userName", userName)
	return p
}

func (p *PlayerFilterBuilder) UserNameObject(obj builder.FilterBuilderObject) *PlayerFilterBuilder {
	p.filterBuilder.KObjectV("userName", obj)
	return p
}

func (p *PlayerFilterBuilder) Level(level int) *PlayerFilterBuilder {
	p.filterBuilder.KNumberV("level", float64(level))
	return p
}

func (p *PlayerFilterBuilder) LevelObject(obj builder.FilterBuilderObject) *PlayerFilterBuilder {
	p.filterBuilder.KObjectV("level", obj)
	return p
}

func (p *PlayerFilterBuilder) PlayerIds(ids []int) *PlayerFilterBuilder {
	floatArr := make([]float64, 0, len(ids))
	for _, id := range ids {
		floatArr = append(floatArr, float64(id))
	}
	p.filterBuilder.KNumberArrayV("playerIds", floatArr)
	return p
}

func (p *PlayerFilterBuilder) PlayerIdsObject(obj builder.FilterBuilderObject) *PlayerFilterBuilder {
	p.filterBuilder.KObjectV("playerIds", obj)
	return p
}

func (p *PlayerFilterBuilder) PlayerLabels(labels []string) *PlayerFilterBuilder {
	p.filterBuilder.KStringArrayV("playerLabels", labels)
	return p
}

func (p *PlayerFilterBuilder) PlayerLabelsObject(obj builder.FilterBuilderObject) *PlayerFilterBuilder {
	p.filterBuilder.KObjectV("playerLabels", obj)
	return p
}
