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

package documents_examples

// [START import]
import (
	"fmt"

	"github.com/nitrictech/go-sdk/api/documents"
)

// [END import]

func queryFilter() {
	// [START snippet]
	docs, _ := documents.New()

	query := docs.Collection("customers").Query().Where(
		documents.Condition("country").Eq(documents.StringValue("US")),
		documents.Condition("age").Gt(documents.NumberValue(21)),
	)

	results, err := query.Fetch()

	if err != nil {
		// handle error
	}

	for _, result := range results.Documents {
		// handle documents
		fmt.Println(result.Content())
	}
	// [END snippet]
}
