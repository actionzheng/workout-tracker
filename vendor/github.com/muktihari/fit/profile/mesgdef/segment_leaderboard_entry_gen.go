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

// SegmentLeaderboardEntry is a SegmentLeaderboardEntry message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type SegmentLeaderboardEntry struct {
	Name             string // Friendly name assigned to leader
	ActivityIdString string // String version of the activity_id. 21 characters long, express in decimal
	GroupPrimaryKey  uint32 // Primary user ID of this leader
	ActivityId       uint32 // ID of the activity associated with this leader time
	SegmentTime      uint32 // Scale: 1000; Units: s; Segment Time (includes pauses)
	MessageIndex     typedef.MessageIndex
	Type             typedef.SegmentLeaderboardType // Leader classification

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewSegmentLeaderboardEntry creates new SegmentLeaderboardEntry struct based on given mesg.
// If mesg is nil, it will return SegmentLeaderboardEntry with all fields being set to its corresponding invalid value.
func NewSegmentLeaderboardEntry(mesg *proto.Message) *SegmentLeaderboardEntry {
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

	return &SegmentLeaderboardEntry{
		MessageIndex:     typedef.MessageIndex(vals[254].Uint16()),
		Name:             vals[0].String(),
		Type:             typedef.SegmentLeaderboardType(vals[1].Uint8()),
		GroupPrimaryKey:  vals[2].Uint32(),
		ActivityId:       vals[3].Uint32(),
		SegmentTime:      vals[4].Uint32(),
		ActivityIdString: vals[5].String(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts SegmentLeaderboardEntry into proto.Message. If options is nil, default options will be used.
func (m *SegmentLeaderboardEntry) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumSegmentLeaderboardEntry}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.Name != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.String(m.Name)
		fields = append(fields, field)
	}
	if m.Type != typedef.SegmentLeaderboardTypeInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(byte(m.Type))
		fields = append(fields, field)
	}
	if m.GroupPrimaryKey != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.GroupPrimaryKey)
		fields = append(fields, field)
	}
	if m.ActivityId != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.ActivityId)
		fields = append(fields, field)
	}
	if m.SegmentTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.SegmentTime)
		fields = append(fields, field)
	}
	if m.ActivityIdString != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.String(m.ActivityIdString)
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

// SegmentTimeScaled return SegmentTime in its scaled value.
// If SegmentTime value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: s; Segment Time (includes pauses)
func (m *SegmentLeaderboardEntry) SegmentTimeScaled() float64 {
	if m.SegmentTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.SegmentTime)/1000 - 0
}

// SetMessageIndex sets MessageIndex value.
func (m *SegmentLeaderboardEntry) SetMessageIndex(v typedef.MessageIndex) *SegmentLeaderboardEntry {
	m.MessageIndex = v
	return m
}

// SetName sets Name value.
//
// Friendly name assigned to leader
func (m *SegmentLeaderboardEntry) SetName(v string) *SegmentLeaderboardEntry {
	m.Name = v
	return m
}

// SetType sets Type value.
//
// Leader classification
func (m *SegmentLeaderboardEntry) SetType(v typedef.SegmentLeaderboardType) *SegmentLeaderboardEntry {
	m.Type = v
	return m
}

// SetGroupPrimaryKey sets GroupPrimaryKey value.
//
// Primary user ID of this leader
func (m *SegmentLeaderboardEntry) SetGroupPrimaryKey(v uint32) *SegmentLeaderboardEntry {
	m.GroupPrimaryKey = v
	return m
}

// SetActivityId sets ActivityId value.
//
// ID of the activity associated with this leader time
func (m *SegmentLeaderboardEntry) SetActivityId(v uint32) *SegmentLeaderboardEntry {
	m.ActivityId = v
	return m
}

// SetSegmentTime sets SegmentTime value.
//
// Scale: 1000; Units: s; Segment Time (includes pauses)
func (m *SegmentLeaderboardEntry) SetSegmentTime(v uint32) *SegmentLeaderboardEntry {
	m.SegmentTime = v
	return m
}

// SetSegmentTimeScaled is similar to SetSegmentTime except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: s; Segment Time (includes pauses)
func (m *SegmentLeaderboardEntry) SetSegmentTimeScaled(v float64) *SegmentLeaderboardEntry {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.SegmentTime = uint32(basetype.Uint32Invalid)
		return m
	}
	m.SegmentTime = uint32(unscaled)
	return m
}

// SetActivityIdString sets ActivityIdString value.
//
// String version of the activity_id. 21 characters long, express in decimal
func (m *SegmentLeaderboardEntry) SetActivityIdString(v string) *SegmentLeaderboardEntry {
	m.ActivityIdString = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *SegmentLeaderboardEntry) SetUnknownFields(unknownFields ...proto.Field) *SegmentLeaderboardEntry {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *SegmentLeaderboardEntry) SetDeveloperFields(developerFields ...proto.DeveloperField) *SegmentLeaderboardEntry {
	m.DeveloperFields = developerFields
	return m
}