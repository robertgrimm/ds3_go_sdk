package networking

import (
    "io"
    "net/http"
    "net/url"
    "strings"
)

type HttpRequestBuilder struct {
    reader io.Reader
    contentLength *int64
    queryParams *url.Values
    headers *http.Header
    signatureFields signatureFields
    checksumType ChecksumType
}

func NewHttpRequestBuilder() *HttpRequestBuilder {
    return &HttpRequestBuilder{
        queryParams:&url.Values{},
        headers:&http.Header{},
        checksumType:NONE,
    }
}

// Internally converts reader with size decorator to limit reader to ensure size is respected
func (builder *HttpRequestBuilder) WithReader(stream ReaderWithSizeDecorator) *HttpRequestBuilder {
    streamSize, _ := stream.Size()
    builder.reader = io.LimitReader(stream, streamSize)
    builder.contentLength = &streamSize
    return builder
}

// Internally converts reader with size decorator to limit reader to ensure size is respected
// and adds the closer functionality to the limit reader. The send network will automatically
// close the reader when finished.
func (builder *HttpRequestBuilder) WithReadCloser(stream ReadCloserWithSizeDecorator) *HttpRequestBuilder {
    streamSize, _ := stream.Size()
    builder.reader = NewLimitReadCloser(stream)
    builder.contentLength = &streamSize
    return builder
}

func (builder *HttpRequestBuilder) WithHttpVerb(verb string) *HttpRequestBuilder {
    builder.signatureFields.Verb = verb
    return builder
}

func (builder *HttpRequestBuilder) WithPath(path string) *HttpRequestBuilder {
    builder.signatureFields.Path = path
    return builder
}

func (builder *HttpRequestBuilder) WithHeader(key string, value string) *HttpRequestBuilder {
    builder.headers.Add(key, value)
    return builder
}

const ( AMZ_META_HEADER = "x-amz-meta-" )

func (builder *HttpRequestBuilder) WithAmzHeader(key string, value string) *HttpRequestBuilder {
    if strings.HasPrefix(strings.ToLower(key), AMZ_META_HEADER) {
        builder.headers.Add(strings.ToLower(key), value)
    } else {
        builder.headers.Add(strings.ToLower(AMZ_META_HEADER + key), value)
    }
    return builder
}

func (builder *HttpRequestBuilder) WithHeaders(headers map[string]string) *HttpRequestBuilder {
    for key, value := range headers {
        builder.WithAmzHeader(key, value)
    }
    return builder
}

func (builder *HttpRequestBuilder) WithQueryParam(key string, value string) *HttpRequestBuilder {
    builder.queryParams.Set(key, value)
    return builder
}

func (builder *HttpRequestBuilder) WithOptionalQueryParam(key string, value *string) *HttpRequestBuilder {
    if value != nil {
        return builder
    }
    builder.queryParams.Set(key, *value)
    return builder
}

func (builder *HttpRequestBuilder) WithOptionalVoidQueryParam(key string, value bool) *HttpRequestBuilder {
    if value {
        builder.queryParams.Set(key, "")
    }
    return builder
}

func (builder *HttpRequestBuilder) WithChecksum(checksum Checksum) *HttpRequestBuilder {
    builder.signatureFields.ContentHash = checksum.ContentHash
    builder.checksumType = checksum.Type
    return builder
}

//todo not used yet
func (builder *HttpRequestBuilder) WithContentType(contentType string) *HttpRequestBuilder {
    builder.signatureFields.ContentType = contentType
    return builder
}

func (builder *HttpRequestBuilder) Build(conn *ConnectionInfo) (*http.Request, error) {
    httpRequest, err := http.NewRequest(builder.signatureFields.Verb, builder.buildUrl(conn), builder.reader)
    if err != nil {
        return nil, err
    }

    if builder.contentLength != nil {
        httpRequest.ContentLength = *builder.contentLength
    }

    builder.signatureFields.Date = getCurrentTime()

    authHeaderVal := buildAuthHeaderValue(conn.Credentials, &builder.signatureFields)

    // Set the http request headers such as authorization and date.
    return builder.addHttpRequestHeaders(httpRequest, authHeaderVal)
}

func (builder *HttpRequestBuilder) buildUrl(conn *ConnectionInfo) string {
    httpUrl := conn.Endpoint
    httpUrl.Path = builder.signatureFields.Path
    httpUrl.RawQuery = encodeQueryParams(builder.queryParams)
    return httpUrl.String()
}

func (builder *HttpRequestBuilder) addHttpRequestHeaders(httpRequest *http.Request, authHeader string) (*http.Request, error) {
    httpRequest.Header.Add("Date", builder.signatureFields.Date)
    httpRequest.Header.Add("Authorization", authHeader)

    if builder.checksumType != NONE {
        checksumKey, err := getChecksumHeaderKey(builder.checksumType)
        if err != nil {
            return nil, err
        }
        httpRequest.Header.Add(checksumKey, builder.signatureFields.ContentHash)
    }

    // Copy the headers from the Ds3Request object.
    for key, val := range *builder.headers {
        httpRequest.Header.Add(key, val[0])
    }
    return httpRequest, nil
}

// Percent encodes query parameters and constructs encoded string.
// Spaces are percent encoded as '%20'
func encodeQueryParams(queryParams *url.Values) string {
    // url.Encode encodes spaces as plus (+), so after urlEncode we replace plus (+) signs
    // with percent encoding for spaces (%20)
    return strings.Replace(queryParams.Encode(), "+", "%20", -1)
}
