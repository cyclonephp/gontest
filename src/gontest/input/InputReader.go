package input

import (
    "encoding/xml"
)

type Header struct {
    Key string `xml:"key,attr"`
    Value string `xml:"value,attr"`
}

type Request struct {
    Headers []Header `xml:"header"`
    Count int `xml:"count,attr"`
    Parallel bool `xml:"parallel,attr"`
    MaxExecTime string `xml:"maxExecTime,attr"`
}


func ReadInput(rawXML string) (req *Request, err error) {
    req = new (Request)
    err = xml.Unmarshal([]byte(rawXML), req)
    return req, err
}



