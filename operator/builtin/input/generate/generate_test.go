// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generate

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-log-collection/entry"
	"github.com/open-telemetry/opentelemetry-log-collection/operator"
	"github.com/open-telemetry/opentelemetry-log-collection/testutil"
	"github.com/stretchr/testify/require"
)

func TestInputGenerate(t *testing.T) {
	cfg := NewGenerateInputConfig("test_operator_id")
	cfg.OutputIDs = []string{"fake"}
	cfg.Count = 5
	cfg.Entry = entry.Entry{
		Record: "test message",
	}

	ops, err := cfg.Build(testutil.NewBuildContext(t))
	require.NoError(t, err)
	op := ops[0]

	fake := testutil.NewFakeOutput(t)
	err = op.SetOutputs([]operator.Operator{fake})
	require.NoError(t, err)

	require.NoError(t, op.Start())
	defer op.Stop()

	for i := 0; i < 5; i++ {
		fake.ExpectRecord(t, "test message")
	}
}
