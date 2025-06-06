// Copyright 2020 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prompb

import (
	"sync"
)

func (m Sample) T() int64   { return m.Timestamp }
func (m Sample) V() float64 { return m.Value }

func (h Histogram) IsFloatHistogram() bool {
	_, ok := h.GetCount().(*Histogram_CountFloat)
	return ok
}

func (r *ChunkedReadResponse) PooledMarshal(p *sync.Pool) ([]byte, error) {
	size := r.Size()
	data, ok := p.Get().(*[]byte)
	if ok && cap(*data) >= size {
		n, err := r.MarshalToSizedBuffer((*data)[:size])
		if err != nil {
			return nil, err
		}
		return (*data)[:n], nil
	}
	return r.Marshal()
}