/*
Copyright 2013 Tamás Gulácsi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package camliWeed

import (
	"log"

	"camlistore.org/pkg/blob"
	"camlistore.org/pkg/blobserver"
)

var _ blobserver.MaxEnumerateConfig = (*weedStorage)(nil)

func (sto *weedStorage) MaxEnumerate() int { return 1000 }

func (sto *weedStorage) EnumerateBlobs(dest chan<- blob.SizedRef, after string, limit int) error {
	defer close(dest)
	objs, err := sto.weedClient.List(after, limit)
	if err != nil {
		log.Printf("weed List: %v", err)
		return err
	}
	var (
		ok bool
		br blob.Ref
	)
	for _, obj := range objs {
		br, ok = blob.Parse(obj.Key)
		if !ok {
			continue
		}
		dest <- blob.SizedRef{Ref: br, Size: obj.Size}
	}
	return nil
}
