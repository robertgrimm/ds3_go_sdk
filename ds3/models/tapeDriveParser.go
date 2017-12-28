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

import "log"

func (tapeDrive *TapeDrive) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CleaningRequired":
            tapeDrive.CleaningRequired = parseBool(child.Content, aggErr)
        case "ErrorMessage":
            tapeDrive.ErrorMessage = parseNullableString(child.Content)
        case "ForceTapeRemoval":
            tapeDrive.ForceTapeRemoval = parseBool(child.Content, aggErr)
        case "Id":
            tapeDrive.Id = parseString(child.Content)
        case "LastCleaned":
            tapeDrive.LastCleaned = parseNullableString(child.Content)
        case "MfgSerialNumber":
            tapeDrive.MfgSerialNumber = parseNullableString(child.Content)
        case "PartitionId":
            tapeDrive.PartitionId = parseString(child.Content)
        case "Quiesced":
            parseEnum(child.Content, &tapeDrive.Quiesced, aggErr)
        case "SerialNumber":
            tapeDrive.SerialNumber = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &tapeDrive.State, aggErr)
        case "TapeId":
            tapeDrive.TapeId = parseNullableString(child.Content)
        case "Type":
            parseEnum(child.Content, &tapeDrive.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeDrive.", child.XMLName.Local)
        }
    }
}