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

type DeleteBucketAclSpectraS3Request struct {
    bucketAcl string
    queryParams *url.Values
}

func NewDeleteBucketAclSpectraS3Request(bucketAcl string) *DeleteBucketAclSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteBucketAclSpectraS3Request{
        bucketAcl: bucketAcl,
        queryParams: queryParams,
    }
}




func (DeleteBucketAclSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteBucketAclSpectraS3Request *DeleteBucketAclSpectraS3Request) Path() string {
    return "/_rest_/bucket_acl/" + deleteBucketAclSpectraS3Request.bucketAcl
}

func (deleteBucketAclSpectraS3Request *DeleteBucketAclSpectraS3Request) QueryParams() *url.Values {
    return deleteBucketAclSpectraS3Request.queryParams
}

func (DeleteBucketAclSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteBucketAclSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteBucketAclSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
