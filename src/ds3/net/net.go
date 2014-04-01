package net

import (
    "net/http"
    "net/url"
    "ds3/models"
)

type Network interface {
    Invoke(request models.Ds3Request) (*http.Response, error)
}

type ConnectionInfo struct {
    Endpoint url.URL
    Creds Credentials

    Proxy *url.URL
    RedirectRetryCount int
}

type Credentials struct {
    AccessId, Key string
}

type httpNetwork struct {
    connectionInfo *ConnectionInfo
    client *http.Client
}

func NewHttpNetwork(connectionInfo *ConnectionInfo) Network {
    return &httpNetwork{
        connectionInfo,
        &http.Client{ Transport: &http.Transport{ Proxy: http.ProxyURL(connectionInfo.Proxy) } },
    }
}

//TODO: improve error handling
func (net httpNetwork) Invoke(request models.Ds3Request) (*http.Response, error) {
    // Build the request.
    httpRequest, makeReqErr := buildHttpRequest(net.connectionInfo, request)
    if makeReqErr != nil {
        return nil, makeReqErr
    }

    // Perform the request.
    httpResponse, reqErr := net.client.Do(httpRequest)
    if reqErr != nil {
        return nil, reqErr
    }

    // Return the response.
    return httpResponse, nil
}

func buildHttpRequest(conn *ConnectionInfo, request models.Ds3Request) (*http.Request, error) {
    httpRequest, err := http.NewRequest(
        request.Verb().String(),
        buildUrl(conn, request).String(),
        request.GetContentStream(),
    )
    if err != nil {
        return nil, err
    }
    setRequestHeaders(httpRequest, conn.Creds, request)
    return httpRequest, nil
}

func buildUrl(conn *ConnectionInfo, request models.Ds3Request) *url.URL {
    url := conn.Endpoint
    url.Path = request.Path()
    url.RawQuery = request.QueryParams().Encode()
    return &url
}

