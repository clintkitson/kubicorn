// Copyright 2017, Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

// Package debugger is an experimental, auto-generated package for the
// Stackdriver Debugger API.
//
// Examines the call stack and variables of a running application
// without stopping or slowing it down.
//
// Use the client at cloud.google.com/go/cmd/go-cloud-debug-agent in preference to this.
package debugger // import "cloud.google.com/go/debugger/apiv2"

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func insertXGoog(ctx context.Context, val []string) context.Context {
	md, _ := metadata.FromOutgoingContext(ctx)
	md = md.Copy()
	md["x-goog-api-client"] = val
	return metadata.NewOutgoingContext(ctx, md)
}

// DefaultAuthScopes reports the authentication scopes required
// by this package.
func DefaultAuthScopes() []string {
	return []string{
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/cloud_debugger",
	}
}
