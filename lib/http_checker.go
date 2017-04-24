package lib

import (
    "net/url"
    "net/http"
    "io"
    "strconv"
)

type HttpChecker struct {}

func (hc HttpChecker) GetHealth(url *url.URL, opts map[string]interface{}) HealthStatus {
    health_meta := make(map[string]string)

    req, err := http.NewRequest("GET", url.String(), nil)
    if err != nil {
        panic(err)
    }

    if val, ok := opts["method"].(string); ok {
        req.Method = val
    }

    if val, ok := opts["body"].(io.ReadCloser); ok {
        req.Body = val
    }

    expect_status_code := "200"
    if val, ok := opts["expect_status_code"].(string); ok {
        expect_status_code = val
    }

    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        return HealthStatus{
            DOWN,
            HealthReason{
                "http.response",
                HIGH,
                err.Error(),
                health_meta,
           }}
    }

    health_meta["proto"] = resp.Proto
    health_meta["status_code"] = strconv.Itoa(resp.StatusCode)

    if expect_status_code != health_meta["status_code"] {
        return HealthStatus{
            DOWN,
            HealthReason{
                "http.response",
                HIGH,
                "Server did not respond with correct status code",
                health_meta,
           }}
    }

    return HealthStatus{
        UP,
	HealthReason{
	    "http.response",
	    INFO,
	    "The server answered correctly within the given timeout",
            health_meta,
       }}
}
