// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type BatteryStatus uint8

const (
	BatteryStatusNew      BatteryStatus = 1
	BatteryStatusGood     BatteryStatus = 2
	BatteryStatusOk       BatteryStatus = 3
	BatteryStatusLow      BatteryStatus = 4
	BatteryStatusCritical BatteryStatus = 5
	BatteryStatusCharging BatteryStatus = 6
	BatteryStatusUnknown  BatteryStatus = 7
	BatteryStatusInvalid  BatteryStatus = 0xFF
)

func (b BatteryStatus) Uint8() uint8 { return uint8(b) }

func (b BatteryStatus) String() string {
	switch b {
	case BatteryStatusNew:
		return "new"
	case BatteryStatusGood:
		return "good"
	case BatteryStatusOk:
		return "ok"
	case BatteryStatusLow:
		return "low"
	case BatteryStatusCritical:
		return "critical"
	case BatteryStatusCharging:
		return "charging"
	case BatteryStatusUnknown:
		return "unknown"
	default:
		return "BatteryStatusInvalid(" + strconv.FormatUint(uint64(b), 10) + ")"
	}
}

// FromString parse string into BatteryStatus constant it's represent, return BatteryStatusInvalid if not found.
func BatteryStatusFromString(s string) BatteryStatus {
	switch s {
	case "new":
		return BatteryStatusNew
	case "good":
		return BatteryStatusGood
	case "ok":
		return BatteryStatusOk
	case "low":
		return BatteryStatusLow
	case "critical":
		return BatteryStatusCritical
	case "charging":
		return BatteryStatusCharging
	case "unknown":
		return BatteryStatusUnknown
	default:
		return BatteryStatusInvalid
	}
}

// List returns all constants.
func ListBatteryStatus() []BatteryStatus {
	return []BatteryStatus{
		BatteryStatusNew,
		BatteryStatusGood,
		BatteryStatusOk,
		BatteryStatusLow,
		BatteryStatusCritical,
		BatteryStatusCharging,
		BatteryStatusUnknown,
	}
}