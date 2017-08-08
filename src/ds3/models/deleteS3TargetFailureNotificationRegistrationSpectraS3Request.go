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

type DeleteS3TargetFailureNotificationRegistrationSpectraS3Request struct {
    s3TargetFailureNotificationRegistration string
    queryParams *url.Values
}

func NewDeleteS3TargetFailureNotificationRegistrationSpectraS3Request(s3TargetFailureNotificationRegistration string) *DeleteS3TargetFailureNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteS3TargetFailureNotificationRegistrationSpectraS3Request{
        s3TargetFailureNotificationRegistration: s3TargetFailureNotificationRegistration,
        queryParams: queryParams,
    }
}




func (DeleteS3TargetFailureNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteS3TargetFailureNotificationRegistrationSpectraS3Request *DeleteS3TargetFailureNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/s3_target_failure_notification_registration/" + deleteS3TargetFailureNotificationRegistrationSpectraS3Request.s3TargetFailureNotificationRegistration
}

func (deleteS3TargetFailureNotificationRegistrationSpectraS3Request *DeleteS3TargetFailureNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return deleteS3TargetFailureNotificationRegistrationSpectraS3Request.queryParams
}

func (DeleteS3TargetFailureNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteS3TargetFailureNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteS3TargetFailureNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}