package runner

import (
    "testing"
    "gontest/model"
    "net/http"
)

type MockClient struct {
    
}

func (this *MockClient) Do (req *http.Request) (*http.Response, error) {
    var resp = new(http.Response);
    resp.ContentLength = 200
    return resp, nil
}

type MockListener struct {
    SendStartCallCtr int
    SendEndCallCtr int
    RespReceivCallCtr int
    ReceivedResponse *http.Response
}

func (this *MockListener) SendRequestStart(req *model.Request) {
    this.SendStartCallCtr += 1    
}

func (this *MockListener) SendRequestEnd(req *model.Request) {
    this.SendEndCallCtr++
}

func (this *MockListener) ResponseReceived(resp *http.Response) {
    this.RespReceivCallCtr++
    this.ReceivedResponse = resp
}

func TestRun(t *testing.T) {
    
    var req = &model.Request{
        Headers: []model.Header{
            model.Header{
                Key: "Accept-Language",
                Value: "hu_HU",
            },
        },
        Method: "GET",
        URL: "http://localhost/cyclonephp",
    }
    var mockListener = &MockListener{};
    if err := Run(req, new(MockClient), mockListener); err != nil {
        t.Errorf("error during request initialization: %s", err)
    }
    if mockListener.SendStartCallCtr != 1 {
        t.Errorf("listener.SendRequestStart() was called %d times instead of 1", mockListener.SendStartCallCtr)
    }
    if mockListener.SendEndCallCtr != 1 {
        t.Errorf("listener.SendRequestEnd() was called %d times instead of 1", mockListener.SendEndCallCtr)
    }
    if mockListener.RespReceivCallCtr != 1 {
        t.Errorf("listener.ResponseReceived was called %d times instead of 1", mockListener.RespReceivCallCtr)
    }
    if mockListener.ReceivedResponse.ContentLength != 200 {
        t.Errorf("mockListener.ReceivedResponse expected: 200, actual: %d", mockListener.ReceivedResponse.ContentLength)
    }
}

