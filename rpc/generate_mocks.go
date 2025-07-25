// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package rpc

//go:generate protoc --go_out=paths=source_relative:.  --go-grpc_out=paths=source_relative:. rpc.proto
//go:generate go run github.com/golang/mock/mockgen -destination mocks/rpc_mocks.go -copyright_file ../scripts/copyright.txt . CNIBackendClient,NPBackendClient,ConfigServerBackendClient
