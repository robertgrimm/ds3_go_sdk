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

func (ds3DataReplicationRuleList *Ds3DataReplicationRuleList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Ds3DataReplicationRule":
            var model Ds3DataReplicationRule
            model.parse(&child, aggErr)
            ds3DataReplicationRuleList.Ds3DataReplicationRules = append(ds3DataReplicationRuleList.Ds3DataReplicationRules, model)
        }
    }
}