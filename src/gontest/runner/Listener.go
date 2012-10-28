package runner

import (
    "gontest/model"
    "net/http"
)

type Listener interface {

    SendRequestStart(req *model.Request)
    
    SendRequestEnd(req *model.Request)
    
    ResponseReceived(resp *http.Response) 

}

