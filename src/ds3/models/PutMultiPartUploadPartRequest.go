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
    "strconv" //TODO import
)

type PutMultiPartUploadPartRequest struct {
    bucketName string
    objectName string
    partNumber int
    uploadId string
    content networking.ReaderWithSizeDecorator //TODO add
    queryParams *url.Values
}

//TODO add content param
func NewPutMultiPartUploadPartRequest(bucketName string, objectName string, partNumber int, uploadId string, content networking.ReaderWithSizeDecorator) *PutMultiPartUploadPartRequest {
    queryParams := &url.Values{}
    queryParams.Set("part_number", strconv.Itoa(partNumber))
    queryParams.Set("upload_id", uploadId)

    return &PutMultiPartUploadPartRequest{
        bucketName: bucketName,
        objectName: objectName,
        partNumber: partNumber,
        uploadId: uploadId,
        content: content, //TODO add
        queryParams: queryParams,
    }
}




func (PutMultiPartUploadPartRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (putMultiPartUploadPartRequest *PutMultiPartUploadPartRequest) Path() string {
    return "/" + putMultiPartUploadPartRequest.bucketName + "/" + putMultiPartUploadPartRequest.objectName
}

func (putMultiPartUploadPartRequest *PutMultiPartUploadPartRequest) QueryParams() *url.Values {
    return putMultiPartUploadPartRequest.queryParams
}

func (PutMultiPartUploadPartRequest) Header() *http.Header {
    return &http.Header{}
}

func (PutMultiPartUploadPartRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}

func (putMultiPartUploadPartRequest *PutMultiPartUploadPartRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return putMultiPartUploadPartRequest.content //TODO return content
}
