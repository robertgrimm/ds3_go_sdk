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
    "strconv"
)

type GetS3TargetFailureNotificationRegistrationsSpectraS3Request struct {
    pageLength int
    pageOffset int
    pageStartMarker string
    userId string
    queryParams *url.Values
}

func NewGetS3TargetFailureNotificationRegistrationsSpectraS3Request() *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetS3TargetFailureNotificationRegistrationsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.pageLength = pageLength
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}
func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.pageOffset = pageOffset
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}
func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.pageStartMarker = pageStartMarker
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}
func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.userId = userId
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("user_id", userId)
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}


func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams.Set("last_page", "")
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (GetS3TargetFailureNotificationRegistrationsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) Path() string {
    return "/_rest_/s3_target_failure_notification_registration"
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) QueryParams() *url.Values {
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request.queryParams
}

func (GetS3TargetFailureNotificationRegistrationsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetS3TargetFailureNotificationRegistrationsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetS3TargetFailureNotificationRegistrationsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
