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

type VerifyPhysicalPlacementForObjectsSpectraS3Request struct {
    BucketName string
    Objects []Ds3GetObject
    StorageDomain *string
}

func NewVerifyPhysicalPlacementForObjectsSpectraS3Request(bucketName string, objectNames []string) *VerifyPhysicalPlacementForObjectsSpectraS3Request {

    return &VerifyPhysicalPlacementForObjectsSpectraS3Request{
        BucketName: bucketName,
        Objects: buildDs3GetObjectSliceFromNames(objectNames),
    }
}

func NewVerifyPhysicalPlacementForObjectsSpectraS3RequestWithPartialObjects(bucketName string, objects []Ds3GetObject) *VerifyPhysicalPlacementForObjectsSpectraS3Request {

    return &VerifyPhysicalPlacementForObjectsSpectraS3Request{
        BucketName: bucketName,
        Objects: objects,
    }
}

func (verifyPhysicalPlacementForObjectsSpectraS3Request *VerifyPhysicalPlacementForObjectsSpectraS3Request) WithStorageDomain(storageDomain string) *VerifyPhysicalPlacementForObjectsSpectraS3Request {
    verifyPhysicalPlacementForObjectsSpectraS3Request.StorageDomain = &storageDomain
    return verifyPhysicalPlacementForObjectsSpectraS3Request
}

