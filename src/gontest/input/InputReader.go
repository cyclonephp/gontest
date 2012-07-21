package input

import (
    "encoding/xml"
)

type HeaderPattern struct {
    Key string `xml:"key,attr"`
    ExpectedValue string `xml:"expectedValue,attr"`
    ExpectedValuePattern string `xml:"expectedValuePattern,attr"`
}

type BodyPattern struct {
    Query string `xml:"query,attr"`
    Count int `xml:"count,attr"`
    MinCount int `xml:"minCount,attr"`
    MaxCount int `xml:"maxCount,attr"`
}

type Header struct {
    Key string `xml:"key,attr"`
    Value string `xml:"value,attr"`
}

type Request struct {
    Headers []Header `xml:"header"`
    Method string `xml:"method,attr"`
    Count int `xml:"count,attr"`
    Parallel bool `xml:"parallel,attr"`
    MaxExecTime string `xml:"maxExecTime,attr"`
    NextRequests []Request `xml:"nextRequests>request"`
    HeaderPatterns []HeaderPattern `xml:"headerPattern"`
    RegexPatterns []BodyPattern `xml:"regexPattern"`
    XpathPatterns []BodyPattern `xml:"xpathPattern"`
    CssPatterns []BodyPattern `xml:"cssPattern"`
}


func ReadInput(rawXML string) (req *Request, err error) {
    req = new (Request)
    err = xml.Unmarshal([]byte(rawXML), req)
    return req, err
}



