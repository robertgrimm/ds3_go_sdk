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

type GetDs3TargetDataPoliciesSpectraS3Request struct {
    ds3TargetDataPolicies string
    queryParams *url.Values
}

func NewGetDs3TargetDataPoliciesSpectraS3Request(ds3TargetDataPolicies string) *GetDs3TargetDataPoliciesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDs3TargetDataPoliciesSpectraS3Request{
        ds3TargetDataPolicies: ds3TargetDataPolicies,
        queryParams: queryParams,
    }
}




func (GetDs3TargetDataPoliciesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDs3TargetDataPoliciesSpectraS3Request *GetDs3TargetDataPoliciesSpectraS3Request) Path() string {
    return "/_rest_/ds3_target_data_policies/" + getDs3TargetDataPoliciesSpectraS3Request.ds3TargetDataPolicies
}

func (getDs3TargetDataPoliciesSpectraS3Request *GetDs3TargetDataPoliciesSpectraS3Request) QueryParams() *url.Values {
    return getDs3TargetDataPoliciesSpectraS3Request.queryParams
}

func (GetDs3TargetDataPoliciesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDs3TargetDataPoliciesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDs3TargetDataPoliciesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
