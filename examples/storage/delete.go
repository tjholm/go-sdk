// Copyright 2021 Nitric Technologies Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage_examples

// [START import]
import (
	"github.com/nitrictech/go-sdk/api/storage"
)

// [END import]

func deleteFile() {
	// [START snippet]
	// Create a new storage client
	sc, err := storage.New()

	if err != nil {
		// handle client creation error...
	}

	err = sc.Bucket("my-bucket").File("path/to/file").Delete()

	if err != nil {
		// handle bucket file delete error
	}
	// [END snippet]
}
