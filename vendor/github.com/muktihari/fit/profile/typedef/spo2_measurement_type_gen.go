// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Spo2MeasurementType byte

const (
	Spo2MeasurementTypeOffWrist        Spo2MeasurementType = 0
	Spo2MeasurementTypeSpotCheck       Spo2MeasurementType = 1
	Spo2MeasurementTypeContinuousCheck Spo2MeasurementType = 2
	Spo2MeasurementTypePeriodic        Spo2MeasurementType = 3
	Spo2MeasurementTypeInvalid         Spo2MeasurementType = 0xFF
)

func (s Spo2MeasurementType) Byte() byte { return byte(s) }

func (s Spo2MeasurementType) String() string {
	switch s {
	case Spo2MeasurementTypeOffWrist:
		return "off_wrist"
	case Spo2MeasurementTypeSpotCheck:
		return "spot_check"
	case Spo2MeasurementTypeContinuousCheck:
		return "continuous_check"
	case Spo2MeasurementTypePeriodic:
		return "periodic"
	default:
		return "Spo2MeasurementTypeInvalid(" + strconv.Itoa(int(s)) + ")"
	}
}

// FromString parse string into Spo2MeasurementType constant it's represent, return Spo2MeasurementTypeInvalid if not found.
func Spo2MeasurementTypeFromString(s string) Spo2MeasurementType {
	switch s {
	case "off_wrist":
		return Spo2MeasurementTypeOffWrist
	case "spot_check":
		return Spo2MeasurementTypeSpotCheck
	case "continuous_check":
		return Spo2MeasurementTypeContinuousCheck
	case "periodic":
		return Spo2MeasurementTypePeriodic
	default:
		return Spo2MeasurementTypeInvalid
	}
}

// List returns all constants.
func ListSpo2MeasurementType() []Spo2MeasurementType {
	return []Spo2MeasurementType{
		Spo2MeasurementTypeOffWrist,
		Spo2MeasurementTypeSpotCheck,
		Spo2MeasurementTypeContinuousCheck,
		Spo2MeasurementTypePeriodic,
	}
}