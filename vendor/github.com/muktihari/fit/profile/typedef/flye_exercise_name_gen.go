// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type FlyeExerciseName uint16

const (
	FlyeExerciseNameCableCrossover                    FlyeExerciseName = 0
	FlyeExerciseNameDeclineDumbbellFlye               FlyeExerciseName = 1
	FlyeExerciseNameDumbbellFlye                      FlyeExerciseName = 2
	FlyeExerciseNameInclineDumbbellFlye               FlyeExerciseName = 3
	FlyeExerciseNameKettlebellFlye                    FlyeExerciseName = 4
	FlyeExerciseNameKneelingRearFlye                  FlyeExerciseName = 5
	FlyeExerciseNameSingleArmStandingCableReverseFlye FlyeExerciseName = 6
	FlyeExerciseNameSwissBallDumbbellFlye             FlyeExerciseName = 7
	FlyeExerciseNameArmRotations                      FlyeExerciseName = 8
	FlyeExerciseNameHugATree                          FlyeExerciseName = 9
	FlyeExerciseNameInvalid                           FlyeExerciseName = 0xFFFF
)

func (f FlyeExerciseName) Uint16() uint16 { return uint16(f) }

func (f FlyeExerciseName) String() string {
	switch f {
	case FlyeExerciseNameCableCrossover:
		return "cable_crossover"
	case FlyeExerciseNameDeclineDumbbellFlye:
		return "decline_dumbbell_flye"
	case FlyeExerciseNameDumbbellFlye:
		return "dumbbell_flye"
	case FlyeExerciseNameInclineDumbbellFlye:
		return "incline_dumbbell_flye"
	case FlyeExerciseNameKettlebellFlye:
		return "kettlebell_flye"
	case FlyeExerciseNameKneelingRearFlye:
		return "kneeling_rear_flye"
	case FlyeExerciseNameSingleArmStandingCableReverseFlye:
		return "single_arm_standing_cable_reverse_flye"
	case FlyeExerciseNameSwissBallDumbbellFlye:
		return "swiss_ball_dumbbell_flye"
	case FlyeExerciseNameArmRotations:
		return "arm_rotations"
	case FlyeExerciseNameHugATree:
		return "hug_a_tree"
	default:
		return "FlyeExerciseNameInvalid(" + strconv.FormatUint(uint64(f), 10) + ")"
	}
}

// FromString parse string into FlyeExerciseName constant it's represent, return FlyeExerciseNameInvalid if not found.
func FlyeExerciseNameFromString(s string) FlyeExerciseName {
	switch s {
	case "cable_crossover":
		return FlyeExerciseNameCableCrossover
	case "decline_dumbbell_flye":
		return FlyeExerciseNameDeclineDumbbellFlye
	case "dumbbell_flye":
		return FlyeExerciseNameDumbbellFlye
	case "incline_dumbbell_flye":
		return FlyeExerciseNameInclineDumbbellFlye
	case "kettlebell_flye":
		return FlyeExerciseNameKettlebellFlye
	case "kneeling_rear_flye":
		return FlyeExerciseNameKneelingRearFlye
	case "single_arm_standing_cable_reverse_flye":
		return FlyeExerciseNameSingleArmStandingCableReverseFlye
	case "swiss_ball_dumbbell_flye":
		return FlyeExerciseNameSwissBallDumbbellFlye
	case "arm_rotations":
		return FlyeExerciseNameArmRotations
	case "hug_a_tree":
		return FlyeExerciseNameHugATree
	default:
		return FlyeExerciseNameInvalid
	}
}

// List returns all constants.
func ListFlyeExerciseName() []FlyeExerciseName {
	return []FlyeExerciseName{
		FlyeExerciseNameCableCrossover,
		FlyeExerciseNameDeclineDumbbellFlye,
		FlyeExerciseNameDumbbellFlye,
		FlyeExerciseNameInclineDumbbellFlye,
		FlyeExerciseNameKettlebellFlye,
		FlyeExerciseNameKneelingRearFlye,
		FlyeExerciseNameSingleArmStandingCableReverseFlye,
		FlyeExerciseNameSwissBallDumbbellFlye,
		FlyeExerciseNameArmRotations,
		FlyeExerciseNameHugATree,
	}
}