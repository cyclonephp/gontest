package input

import (
    "testing"
    "fmt"
)

func TestReadInput(t *testing.T) {
    req, err := ReadInput(`
<request count="20" parallel="true" maxExecTime="1sec" method="get">
    <header key="accept-language" value="hu_HU"/>
    
    <headerPattern key="Content-Type" expectedValue="text/html;charset=utf8"/>
    <cssPattern query="#cnt > p.myclass > a[rel=lightbox]" minCount="5" maxCount="10"/>
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
    
    if headersLen := len(req.Headers); headersLen != 1 {
        t.Errorf("expected 1 header, found %d\n", headersLen)
        return
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
    
    if cssLen := len(req.CssPatterns); cssLen != 1 {
        t.Errorf("len(req.CssPatterns): expected: 1, actual: %d\n", cssLen);
        return
    }
    
    cssPattern := req.CssPatterns[0];
    if cssPattern.Query != "#cnt > p.myclass > a[rel=lightbox]" ||
        cssPattern.MinCount != 5 ||
        cssPattern.MaxCount != 10 ||
        cssPattern.Count != 0 {
        t.Errorf("failed to parse cssPattern tag, actual: %v\n", cssPattern);
    }
    
    if hpLen := len(req.HeaderPatterns); hpLen != 1 {
        t.Errorf("len(req.HeaderPatterns): expected: 1, actual: %d\n", hpLen)
        return
    }
    
    hp := req.HeaderPatterns[0]
    if hp.Key != "Content-Type" ||
        hp.ExpectedValue != "text/html;charset=utf8" {
        t.Errorf("failed to parse headerPattern tag, actual: %v\n", hp)   
    }
    
    
    req = &req.NextRequests[0];
    if req.Method != "" {
        t.Errorf("failed to parse empty method")
    }
    
}

