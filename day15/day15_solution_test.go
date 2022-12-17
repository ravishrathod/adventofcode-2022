package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findTrueCoverage_withOverlappingEndAndStart(t *testing.T) {
	last := Coverage{
		From: -2,
		To:   2,
	}
	current := Coverage{
		From: 2,
		To:   14,
	}
	beaconsCovered := make(map[int]bool)
	assert.Equal(t, 12, findTrueCoverage(last, current, nil, beaconsCovered))
}

func Test_findTrueCoverage_whenSecondRangeWithinFirst(t *testing.T) {
	last := Coverage{
		From: 2,
		To:   14,
	}
	current := Coverage{
		From: 3,
		To:   6,
	}
	beaconsCovered := make(map[int]bool)
	assert.Equal(t, 0, findTrueCoverage(last, current, nil, beaconsCovered))
}

func Test_findTrueCoverage_whenRangesOverlap(t *testing.T) {
	last := Coverage{
		From: 14,
		To:   18,
	}
	current := Coverage{
		From: 16,
		To:   24,
	}
	beaconsCovered := make(map[int]bool)
	assert.Equal(t, 6, findTrueCoverage(last, current, nil, beaconsCovered))
}

func Test_findTrueCoverage_whenRangesDoNotOverlap(t *testing.T) {
	last := Coverage{
		From: 12,
		To:   13,
	}
	current := Coverage{
		From: 14,
		To:   18,
	}
	beaconsCovered := make(map[int]bool)
	assert.Equal(t, 5, findTrueCoverage(last, current, nil, beaconsCovered))
}

func Test_findUnScannedPoint(t *testing.T) {
	coverages := []Coverage{
		{From: 11, To: 13},
		{From: 15, To: 17},
		{From: 15, To: 25},
	}

	x := findUnScannedPoint(coverages, 0, 20)
	assert.Equal(t, 14, x)
}

func Test_findUnScannedPointWhenFirstRangeSwallowsOthers(t *testing.T) {
	coverages := []Coverage{
		{From: -8, To: 12},
		{From: 6, To: 10},
		{From: 12, To: 14},
		{From: 14, To: 26},
	}

	x := findUnScannedPoint(coverages, 0, 20)
	assert.Equal(t, -1, x)
}
