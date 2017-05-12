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

type DeleteStorageDomainMemberSpectraS3Request struct {
    storageDomainMember string
    queryParams *url.Values
}

func NewDeleteStorageDomainMemberSpectraS3Request(storageDomainMember string) *DeleteStorageDomainMemberSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteStorageDomainMemberSpectraS3Request{
        storageDomainMember: storageDomainMember,
        queryParams: queryParams,
    }
}




func (DeleteStorageDomainMemberSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteStorageDomainMemberSpectraS3Request *DeleteStorageDomainMemberSpectraS3Request) Path() string {
    return "/_rest_/storage_domain_member/" + deleteStorageDomainMemberSpectraS3Request.storageDomainMember
}

func (deleteStorageDomainMemberSpectraS3Request *DeleteStorageDomainMemberSpectraS3Request) QueryParams() *url.Values {
    return deleteStorageDomainMemberSpectraS3Request.queryParams
}

func (DeleteStorageDomainMemberSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteStorageDomainMemberSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteStorageDomainMemberSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
