// Copyright 2014-2017 Spectra Logic Corporation. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use
// this file except in compliance with the License. A copy of the License is located at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file.
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// This code is auto-generated, do not modify

package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type CancelVerifyTapeSpectraS3Request struct {
    tapeId string
    queryParams *url.Values
}

func NewCancelVerifyTapeSpectraS3Request(tapeId string) *CancelVerifyTapeSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "cancel_verify")

    return &CancelVerifyTapeSpectraS3Request{
        tapeId: tapeId,
        queryParams: queryParams,
    }
}




func (CancelVerifyTapeSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (cancelVerifyTapeSpectraS3Request *CancelVerifyTapeSpectraS3Request) Path() string {
    return "/_rest_/tape/" + cancelVerifyTapeSpectraS3Request.tapeId
}

func (cancelVerifyTapeSpectraS3Request *CancelVerifyTapeSpectraS3Request) QueryParams() *url.Values {
    return cancelVerifyTapeSpectraS3Request.queryParams
}

func (CancelVerifyTapeSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (CancelVerifyTapeSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (CancelVerifyTapeSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
