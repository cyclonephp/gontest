package runner

import (
    "gontest/model"
    "net/http"
//    "fmt"
)

func Run(req *model.Request, client Client, listener Listener) (err error) {
    var httpReq *http.Request
    if httpReq, err = http.NewRequest(req.Method, req.URL, nil); err != nil {
        return err
    }
    
    var resp *http.Response
    listener.SendRequestStart(req)
    resp, err = client.Do(httpReq)
    listener.SendRequestEnd(req)
    
    if err != nil {
        return err
    }
    listener.ResponseReceived(resp)
    var _ = resp.Body
    
    return 
}

