// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
)

// SpeedZone is a SpeedZone message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type SpeedZone struct {
	Name         string
	MessageIndex typedef.MessageIndex
	HighValue    uint16 // Scale: 1000; Units: m/s

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewSpeedZone creates new SpeedZone struct based on given mesg.
// If mesg is nil, it will return SpeedZone with all fields being set to its corresponding invalid value.
func NewSpeedZone(mesg *proto.Message) *SpeedZone {
	vals := [255]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &SpeedZone{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		HighValue:    vals[0].Uint16(),
		Name:         vals[1].String(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts SpeedZone into proto.Message. If options is nil, default options will be used.
func (m *SpeedZone) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumSpeedZone}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.HighValue != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.HighValue)
		fields = append(fields, field)
	}
	if m.Name != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.String(m.Name)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	*arr = [poolsize]proto.Field{}
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// HighValueScaled return HighValue in its scaled value.
// If HighValue value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *SpeedZone) HighValueScaled() float64 {
	if m.HighValue == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.HighValue)/1000 - 0
}

// SetMessageIndex sets MessageIndex value.
func (m *SpeedZone) SetMessageIndex(v typedef.MessageIndex) *SpeedZone {
	m.MessageIndex = v
	return m
}

// SetHighValue sets HighValue value.
//
// Scale: 1000; Units: m/s
func (m *SpeedZone) SetHighValue(v uint16) *SpeedZone {
	m.HighValue = v
	return m
}

// SetHighValueScaled is similar to SetHighValue except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *SpeedZone) SetHighValueScaled(v float64) *SpeedZone {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.HighValue = uint16(basetype.Uint16Invalid)
		return m
	}
	m.HighValue = uint16(unscaled)
	return m
}

// SetName sets Name value.
func (m *SpeedZone) SetName(v string) *SpeedZone {
	m.Name = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *SpeedZone) SetUnknownFields(unknownFields ...proto.Field) *SpeedZone {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *SpeedZone) SetDeveloperFields(developerFields ...proto.DeveloperField) *SpeedZone {
	m.DeveloperFields = developerFields
	return m
}