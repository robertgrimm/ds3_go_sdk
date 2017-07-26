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

type ActiveJob struct {
    Aggregating bool
    BucketId string
    CachedSizeInBytes int64
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    CompletedSizeInBytes int64
    CreatedAt string
    DeadJobCleanupAllowed bool
    ErrorMessage *string
    Id string
    ImplicitJobIdResolution bool
    MinimizeSpanningAcrossMedia bool
    Naked bool
    Name *string
    OriginalSizeInBytes int64
    Priority Priority
    Rechunked *string
    Replicating bool
    RequestType JobRequestType
    Truncated bool
    TruncatedDueToTimeout bool
    UserId string
    VerifyAfterWrite bool
}