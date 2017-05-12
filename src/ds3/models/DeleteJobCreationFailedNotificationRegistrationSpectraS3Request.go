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

type DeleteJobCreationFailedNotificationRegistrationSpectraS3Request struct {
    notificationId string
    queryParams *url.Values
}

func NewDeleteJobCreationFailedNotificationRegistrationSpectraS3Request(notificationId string) *DeleteJobCreationFailedNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteJobCreationFailedNotificationRegistrationSpectraS3Request{
        notificationId: notificationId,
        queryParams: queryParams,
    }
}




func (DeleteJobCreationFailedNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteJobCreationFailedNotificationRegistrationSpectraS3Request *DeleteJobCreationFailedNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/job_creation_failed_notification_registration/" + deleteJobCreationFailedNotificationRegistrationSpectraS3Request.notificationId
}

func (deleteJobCreationFailedNotificationRegistrationSpectraS3Request *DeleteJobCreationFailedNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return deleteJobCreationFailedNotificationRegistrationSpectraS3Request.queryParams
}

func (DeleteJobCreationFailedNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteJobCreationFailedNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteJobCreationFailedNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
