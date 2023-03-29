// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package weights

import (
	"encoding/json"
	"io"
	// "io/ioutil"
	"log"
	// "fmt"
)

// Prec is the precision for weight output in text formats -- default is aggressive
// for Leabra models -- may need to increase for other models.
var Prec = 4

// NetReadJSON reads weights for entire network in a JSON format into Network structure
func NetReadJSON(r io.Reader) (*Network, error) {
	nw := &Network{}
	dec := json.NewDecoder(r)
	err := dec.Decode(nw) // this is way to do it on reader instead of bytes
	// fmt.Println("nw %+v", nw.Layers)
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
	}
	return nw, nil
}

// LayReadJSON reads weights for layer in a JSON format into Layer structure
func LayReadJSON(r io.Reader) (*Layer, error) {
	lw := &Layer{}
	dec := json.NewDecoder(r)
	err := dec.Decode(lw) // this is way to do it on reader instead of bytes
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
	}
	return lw, nil
}

// PrjnReadJSON reads weights for prjn in a JSON format into Prjn structure
func PrjnReadJSON(r io.Reader) (*Prjn, error) {
	pw := &Prjn{}
	dec := json.NewDecoder(r)
	err := dec.Decode(pw) // this is way to do it on reader instead of bytes
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
	}
	return pw, nil
}
