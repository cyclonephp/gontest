package input

import (
    "testing"
    "fmt"
)

func TestReadInput(t *testing.T) {
    req, err := ReadInput(`
<request count="20" parallel="true" maxExecTime="1sec" method="get">
    <header key="accept-language" value="hu_HU"/>
    <nextRequests>
        <request count="5" parallel="false" maxExecTime="3sec">
            <header key="mykey" value="myvalue"/>
        </request>
    </nextRequests>
</request>
    `)
    fmt.Println(err)
    if req.Count != 20 {
        t.Errorf("req.Count: expected: 20, actual: %d\n", req.Count)
    }
    if req.Parallel != true {
        t.Errorf("req.Parallel: expected: true, actual: false\n");
    }
    if req.MaxExecTime != "1sec" {
        t.Errorf("req.MaxExecTime: expected: 1sec, actual: %s\n", req.MaxExecTime)
    }
    if req.Method != "get" {
        t.Errorf("req.Method: expected: get, found: %s\n", req.Method);
    }
    
    if len(req.Headers) != 1 {
        t.Errorf("expected 1 header, found %d\n", len(req.Headers))
    }
    header := req.Headers[0]
    if header.Key != "accept-language" || header.Value != "hu_HU" {
        t.Errorf("expected header: 'accept-language: hu_HU', found '%s: %s'\n",
            header.Key, header.Value) 
    }
    if nextLen := len(req.NextRequests); nextLen != 1 {
        t.Errorf("len(nextRequests): expected: 1; actual: %d\n", nextLen)
        return
    }
    req = &req.NextRequests[0];
    
    if req.Method != "" {
        t.Errorf("failed to parse empty method")
    }
    
}

