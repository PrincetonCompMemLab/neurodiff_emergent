// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package patgen

import (
	"github.com/PrincetonCompMemLab/private-emergent/erand"
	"github.com/emer/etable/etable"
)

// Shuffle shuffles rows in specified columns in the table independently
func Shuffle(dt *etable.Table, rows []int, colNames []string, colIndependent bool) {
	cl := dt.Clone()
	if colIndependent { // independent
		for _, colNm := range colNames {
			sfrows := make([]int, len(rows))
			copy(sfrows, rows)
			erand.PermuteInts(sfrows)
			for i, row := range rows {
				dt.CellTensor(colNm, row).CopyFrom(cl.CellTensor(colNm, sfrows[i]))
			}
		}
	} else { // shuffle together
		sfrows := make([]int, len(rows))
		copy(sfrows, rows)
		erand.PermuteInts(sfrows)
		for _, colNm := range colNames {
			for i, row := range rows {
				dt.CellTensor(colNm, row).CopyFrom(cl.CellTensor(colNm, sfrows[i]))
			}
		}
	}

}
