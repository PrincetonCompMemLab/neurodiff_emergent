// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package erand

import "math/rand"

// PChoose32 chooses an index in given slice of float32's at random according
// to the probilities of each item (must be normalized to sum to 1)
func PChoose32(ps []float32) int {
	pv := rand.Float32()
	sum := float32(0)
	for i, p := range ps {
		sum += p
		if pv < sum { // note: lower values already excluded
			return i
		}
	}
	return len(ps) - 1
}

// PChoose64 chooses an index in given slice of float64's at random according
// to the probilities of each item (must be normalized to sum to 1)
func PChoose64(ps []float64) int {
	pv := rand.Float64()
	sum := float64(0)
	for i, p := range ps {
		sum += p
		if pv < sum { // note: lower values already excluded
			return i
		}
	}
	return len(ps) - 1
}
