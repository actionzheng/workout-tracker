// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ChopExerciseName uint16

const (
	ChopExerciseNameCablePullThrough                   ChopExerciseName = 0
	ChopExerciseNameCableRotationalLift                ChopExerciseName = 1
	ChopExerciseNameCableWoodchop                      ChopExerciseName = 2
	ChopExerciseNameCrossChopToKnee                    ChopExerciseName = 3
	ChopExerciseNameWeightedCrossChopToKnee            ChopExerciseName = 4
	ChopExerciseNameDumbbellChop                       ChopExerciseName = 5
	ChopExerciseNameHalfKneelingRotation               ChopExerciseName = 6
	ChopExerciseNameWeightedHalfKneelingRotation       ChopExerciseName = 7
	ChopExerciseNameHalfKneelingRotationalChop         ChopExerciseName = 8
	ChopExerciseNameHalfKneelingRotationalReverseChop  ChopExerciseName = 9
	ChopExerciseNameHalfKneelingStabilityChop          ChopExerciseName = 10
	ChopExerciseNameHalfKneelingStabilityReverseChop   ChopExerciseName = 11
	ChopExerciseNameKneelingRotationalChop             ChopExerciseName = 12
	ChopExerciseNameKneelingRotationalReverseChop      ChopExerciseName = 13
	ChopExerciseNameKneelingStabilityChop              ChopExerciseName = 14
	ChopExerciseNameKneelingWoodchopper                ChopExerciseName = 15
	ChopExerciseNameMedicineBallWoodChops              ChopExerciseName = 16
	ChopExerciseNamePowerSquatChops                    ChopExerciseName = 17
	ChopExerciseNameWeightedPowerSquatChops            ChopExerciseName = 18
	ChopExerciseNameStandingRotationalChop             ChopExerciseName = 19
	ChopExerciseNameStandingSplitRotationalChop        ChopExerciseName = 20
	ChopExerciseNameStandingSplitRotationalReverseChop ChopExerciseName = 21
	ChopExerciseNameStandingStabilityReverseChop       ChopExerciseName = 22
	ChopExerciseNameInvalid                            ChopExerciseName = 0xFFFF
)

func (c ChopExerciseName) Uint16() uint16 { return uint16(c) }

func (c ChopExerciseName) String() string {
	switch c {
	case ChopExerciseNameCablePullThrough:
		return "cable_pull_through"
	case ChopExerciseNameCableRotationalLift:
		return "cable_rotational_lift"
	case ChopExerciseNameCableWoodchop:
		return "cable_woodchop"
	case ChopExerciseNameCrossChopToKnee:
		return "cross_chop_to_knee"
	case ChopExerciseNameWeightedCrossChopToKnee:
		return "weighted_cross_chop_to_knee"
	case ChopExerciseNameDumbbellChop:
		return "dumbbell_chop"
	case ChopExerciseNameHalfKneelingRotation:
		return "half_kneeling_rotation"
	case ChopExerciseNameWeightedHalfKneelingRotation:
		return "weighted_half_kneeling_rotation"
	case ChopExerciseNameHalfKneelingRotationalChop:
		return "half_kneeling_rotational_chop"
	case ChopExerciseNameHalfKneelingRotationalReverseChop:
		return "half_kneeling_rotational_reverse_chop"
	case ChopExerciseNameHalfKneelingStabilityChop:
		return "half_kneeling_stability_chop"
	case ChopExerciseNameHalfKneelingStabilityReverseChop:
		return "half_kneeling_stability_reverse_chop"
	case ChopExerciseNameKneelingRotationalChop:
		return "kneeling_rotational_chop"
	case ChopExerciseNameKneelingRotationalReverseChop:
		return "kneeling_rotational_reverse_chop"
	case ChopExerciseNameKneelingStabilityChop:
		return "kneeling_stability_chop"
	case ChopExerciseNameKneelingWoodchopper:
		return "kneeling_woodchopper"
	case ChopExerciseNameMedicineBallWoodChops:
		return "medicine_ball_wood_chops"
	case ChopExerciseNamePowerSquatChops:
		return "power_squat_chops"
	case ChopExerciseNameWeightedPowerSquatChops:
		return "weighted_power_squat_chops"
	case ChopExerciseNameStandingRotationalChop:
		return "standing_rotational_chop"
	case ChopExerciseNameStandingSplitRotationalChop:
		return "standing_split_rotational_chop"
	case ChopExerciseNameStandingSplitRotationalReverseChop:
		return "standing_split_rotational_reverse_chop"
	case ChopExerciseNameStandingStabilityReverseChop:
		return "standing_stability_reverse_chop"
	default:
		return "ChopExerciseNameInvalid(" + strconv.FormatUint(uint64(c), 10) + ")"
	}
}

// FromString parse string into ChopExerciseName constant it's represent, return ChopExerciseNameInvalid if not found.
func ChopExerciseNameFromString(s string) ChopExerciseName {
	switch s {
	case "cable_pull_through":
		return ChopExerciseNameCablePullThrough
	case "cable_rotational_lift":
		return ChopExerciseNameCableRotationalLift
	case "cable_woodchop":
		return ChopExerciseNameCableWoodchop
	case "cross_chop_to_knee":
		return ChopExerciseNameCrossChopToKnee
	case "weighted_cross_chop_to_knee":
		return ChopExerciseNameWeightedCrossChopToKnee
	case "dumbbell_chop":
		return ChopExerciseNameDumbbellChop
	case "half_kneeling_rotation":
		return ChopExerciseNameHalfKneelingRotation
	case "weighted_half_kneeling_rotation":
		return ChopExerciseNameWeightedHalfKneelingRotation
	case "half_kneeling_rotational_chop":
		return ChopExerciseNameHalfKneelingRotationalChop
	case "half_kneeling_rotational_reverse_chop":
		return ChopExerciseNameHalfKneelingRotationalReverseChop
	case "half_kneeling_stability_chop":
		return ChopExerciseNameHalfKneelingStabilityChop
	case "half_kneeling_stability_reverse_chop":
		return ChopExerciseNameHalfKneelingStabilityReverseChop
	case "kneeling_rotational_chop":
		return ChopExerciseNameKneelingRotationalChop
	case "kneeling_rotational_reverse_chop":
		return ChopExerciseNameKneelingRotationalReverseChop
	case "kneeling_stability_chop":
		return ChopExerciseNameKneelingStabilityChop
	case "kneeling_woodchopper":
		return ChopExerciseNameKneelingWoodchopper
	case "medicine_ball_wood_chops":
		return ChopExerciseNameMedicineBallWoodChops
	case "power_squat_chops":
		return ChopExerciseNamePowerSquatChops
	case "weighted_power_squat_chops":
		return ChopExerciseNameWeightedPowerSquatChops
	case "standing_rotational_chop":
		return ChopExerciseNameStandingRotationalChop
	case "standing_split_rotational_chop":
		return ChopExerciseNameStandingSplitRotationalChop
	case "standing_split_rotational_reverse_chop":
		return ChopExerciseNameStandingSplitRotationalReverseChop
	case "standing_stability_reverse_chop":
		return ChopExerciseNameStandingStabilityReverseChop
	default:
		return ChopExerciseNameInvalid
	}
}

// List returns all constants.
func ListChopExerciseName() []ChopExerciseName {
	return []ChopExerciseName{
		ChopExerciseNameCablePullThrough,
		ChopExerciseNameCableRotationalLift,
		ChopExerciseNameCableWoodchop,
		ChopExerciseNameCrossChopToKnee,
		ChopExerciseNameWeightedCrossChopToKnee,
		ChopExerciseNameDumbbellChop,
		ChopExerciseNameHalfKneelingRotation,
		ChopExerciseNameWeightedHalfKneelingRotation,
		ChopExerciseNameHalfKneelingRotationalChop,
		ChopExerciseNameHalfKneelingRotationalReverseChop,
		ChopExerciseNameHalfKneelingStabilityChop,
		ChopExerciseNameHalfKneelingStabilityReverseChop,
		ChopExerciseNameKneelingRotationalChop,
		ChopExerciseNameKneelingRotationalReverseChop,
		ChopExerciseNameKneelingStabilityChop,
		ChopExerciseNameKneelingWoodchopper,
		ChopExerciseNameMedicineBallWoodChops,
		ChopExerciseNamePowerSquatChops,
		ChopExerciseNameWeightedPowerSquatChops,
		ChopExerciseNameStandingRotationalChop,
		ChopExerciseNameStandingSplitRotationalChop,
		ChopExerciseNameStandingSplitRotationalReverseChop,
		ChopExerciseNameStandingStabilityReverseChop,
	}
}