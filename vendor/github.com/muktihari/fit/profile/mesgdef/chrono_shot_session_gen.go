// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// ChronoShotSession is a ChronoShotSession message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type ChronoShotSession struct {
	Timestamp         time.Time
	MinSpeed          uint32 // Scale: 1000; Units: m/s
	MaxSpeed          uint32 // Scale: 1000; Units: m/s
	AvgSpeed          uint32 // Scale: 1000; Units: m/s
	GrainWeight       uint32 // Scale: 10; Units: gr
	StandardDeviation uint32 // Scale: 1000; Units: m/s
	ShotCount         uint16
	ProjectileType    typedef.ProjectileType

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewChronoShotSession creates new ChronoShotSession struct based on given mesg.
// If mesg is nil, it will return ChronoShotSession with all fields being set to its corresponding invalid value.
func NewChronoShotSession(mesg *proto.Message) *ChronoShotSession {
	vals := [254]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 || mesg.Fields[i].Name == factory.NameUnknown {
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

	return &ChronoShotSession{
		Timestamp:         datetime.ToTime(vals[253].Uint32()),
		MinSpeed:          vals[0].Uint32(),
		MaxSpeed:          vals[1].Uint32(),
		AvgSpeed:          vals[2].Uint32(),
		ShotCount:         vals[3].Uint16(),
		ProjectileType:    typedef.ProjectileType(vals[4].Uint8()),
		GrainWeight:       vals[5].Uint32(),
		StandardDeviation: vals[6].Uint32(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts ChronoShotSession into proto.Message. If options is nil, default options will be used.
func (m *ChronoShotSession) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumChronoShotSession}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.MinSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(m.MinSpeed)
		fields = append(fields, field)
	}
	if m.MaxSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(m.MaxSpeed)
		fields = append(fields, field)
	}
	if m.AvgSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.AvgSpeed)
		fields = append(fields, field)
	}
	if m.ShotCount != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.ShotCount)
		fields = append(fields, field)
	}
	if m.ProjectileType != typedef.ProjectileTypeInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(byte(m.ProjectileType))
		fields = append(fields, field)
	}
	if m.GrainWeight != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint32(m.GrainWeight)
		fields = append(fields, field)
	}
	if m.StandardDeviation != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(m.StandardDeviation)
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

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *ChronoShotSession) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// MinSpeedScaled return MinSpeed in its scaled value.
// If MinSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) MinSpeedScaled() float64 {
	if m.MinSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.MinSpeed)/1000 - 0
}

// MaxSpeedScaled return MaxSpeed in its scaled value.
// If MaxSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) MaxSpeedScaled() float64 {
	if m.MaxSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.MaxSpeed)/1000 - 0
}

// AvgSpeedScaled return AvgSpeed in its scaled value.
// If AvgSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) AvgSpeedScaled() float64 {
	if m.AvgSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.AvgSpeed)/1000 - 0
}

// GrainWeightScaled return GrainWeight in its scaled value.
// If GrainWeight value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Units: gr
func (m *ChronoShotSession) GrainWeightScaled() float64 {
	if m.GrainWeight == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.GrainWeight)/10 - 0
}

// StandardDeviationScaled return StandardDeviation in its scaled value.
// If StandardDeviation value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) StandardDeviationScaled() float64 {
	if m.StandardDeviation == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.StandardDeviation)/1000 - 0
}

// SetTimestamp sets Timestamp value.
func (m *ChronoShotSession) SetTimestamp(v time.Time) *ChronoShotSession {
	m.Timestamp = v
	return m
}

// SetMinSpeed sets MinSpeed value.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) SetMinSpeed(v uint32) *ChronoShotSession {
	m.MinSpeed = v
	return m
}

// SetMinSpeedScaled is similar to SetMinSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) SetMinSpeedScaled(v float64) *ChronoShotSession {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.MinSpeed = uint32(basetype.Uint32Invalid)
		return m
	}
	m.MinSpeed = uint32(unscaled)
	return m
}

// SetMaxSpeed sets MaxSpeed value.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) SetMaxSpeed(v uint32) *ChronoShotSession {
	m.MaxSpeed = v
	return m
}

// SetMaxSpeedScaled is similar to SetMaxSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) SetMaxSpeedScaled(v float64) *ChronoShotSession {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.MaxSpeed = uint32(basetype.Uint32Invalid)
		return m
	}
	m.MaxSpeed = uint32(unscaled)
	return m
}

// SetAvgSpeed sets AvgSpeed value.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) SetAvgSpeed(v uint32) *ChronoShotSession {
	m.AvgSpeed = v
	return m
}

// SetAvgSpeedScaled is similar to SetAvgSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) SetAvgSpeedScaled(v float64) *ChronoShotSession {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.AvgSpeed = uint32(basetype.Uint32Invalid)
		return m
	}
	m.AvgSpeed = uint32(unscaled)
	return m
}

// SetShotCount sets ShotCount value.
func (m *ChronoShotSession) SetShotCount(v uint16) *ChronoShotSession {
	m.ShotCount = v
	return m
}

// SetProjectileType sets ProjectileType value.
func (m *ChronoShotSession) SetProjectileType(v typedef.ProjectileType) *ChronoShotSession {
	m.ProjectileType = v
	return m
}

// SetGrainWeight sets GrainWeight value.
//
// Scale: 10; Units: gr
func (m *ChronoShotSession) SetGrainWeight(v uint32) *ChronoShotSession {
	m.GrainWeight = v
	return m
}

// SetGrainWeightScaled is similar to SetGrainWeight except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 10; Units: gr
func (m *ChronoShotSession) SetGrainWeightScaled(v float64) *ChronoShotSession {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.GrainWeight = uint32(basetype.Uint32Invalid)
		return m
	}
	m.GrainWeight = uint32(unscaled)
	return m
}

// SetStandardDeviation sets StandardDeviation value.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) SetStandardDeviation(v uint32) *ChronoShotSession {
	m.StandardDeviation = v
	return m
}

// SetStandardDeviationScaled is similar to SetStandardDeviation except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) SetStandardDeviationScaled(v float64) *ChronoShotSession {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.StandardDeviation = uint32(basetype.Uint32Invalid)
		return m
	}
	m.StandardDeviation = uint32(unscaled)
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *ChronoShotSession) SetUnknownFields(unknownFields ...proto.Field) *ChronoShotSession {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *ChronoShotSession) SetDeveloperFields(developerFields ...proto.DeveloperField) *ChronoShotSession {
	m.DeveloperFields = developerFields
	return m
}