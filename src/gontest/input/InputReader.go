package input

import (
    "encoding/xml"
    "gontest/model"
)

func ReadInput(rawXML string) (req *model.Request, err error) {
    return ReadRawInput( []byte(rawXML));
}

func ReadRawInput(rawXML []byte) (req *model.Request, err error) {
    req = new (model.Request)
    err = xml.Unmarshal(rawXML, req)
    return req, err
}



