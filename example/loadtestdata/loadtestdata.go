package loadtestdata

import (
	"bytes"
	"fmt"
	"io"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad/nquads"
)

func LoadTestData(store *cayley.Handle) error {
	dec := nquads.NewReader(bytes.NewReader([]byte(testdata_nq)), false)
	for {
		q, err := dec.ReadQuad()
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("Failed to read documentRaw: %v", err)
			}
			break
		}
		if !q.IsValid() {
			return fmt.Errorf("Unexpected quad, got:%v", q)
		}
		store.AddQuad(q)
	}
	dec.Close()
	return nil
}

const testdata_nq = `<alice> <follows> <bob> .
<bob> <follows> <fred> .
<bob> <status> "cool_person" .
<charlie> <follows> <bob> .
<charlie> <follows> <dani> .
<dani> <follows> <bob> .
<dani> <follows> <greg> .
<dani> <status> "cool_person" .
<emily> <follows> <fred> .
<fred> <follows> <greg> .
<greg> <status> "cool_person" .
<predicates> <are> <follows> .
<predicates> <are> <status> .
<emily> <status> "smart_person" <smart_graph> .
<greg> <status> "smart_person" <smart_graph> .
`
