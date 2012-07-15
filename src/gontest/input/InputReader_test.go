package input

import (
    "testing"
    "fmt"
)

func TestReadInput(t *testing.T) {
    req, err := ReadInput(`
<request count="20" parallel="true" maxExecTime="1sec">
    <header key="accept-language" value="hu_HU"/>
</request>
    `)
    if req.Count != 20 {
        t.Errorf("req.Count: expected: 20, actual: %d\n", req.Count)
    }
    if req.Parallel != true {
        t.Errorf("req.Parallel: expected: true, actual: false\n");
    }
    if req.MaxExecTime != "1sec" {
        t.Errorf("req.MaxExecTime: expected: 1sec, actual: %s\n", req.MaxExecTime)
    }
    if len(req.Headers) != 1 {
        t.Errorf("expected 1 header, found %d\n", len(req.Headers))
    }
    header := req.Headers[0]
    if header.Key != "accept-language" || header.Value != "hu_HU" {
        t.Errorf("expected header: 'accept-language: hu_HU', found '%s: %s'\n",
            header.Key, header.Value) 
    }
    fmt.Println(err)
}

