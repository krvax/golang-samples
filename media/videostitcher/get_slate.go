// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package videostitcher

// [START video_stitcher_get_slate]
import (
	"context"
	"fmt"
	"io"

	stitcher "cloud.google.com/go/video/stitcher/apiv1"
	stitcherstreampb "google.golang.org/genproto/googleapis/cloud/video/stitcher/v1"
)

// getSlate gets a previously-created slate.
func getSlate(w io.Writer, projectID, slateID string) error {
	// projectID := "my-project-id"
	// slateID := "my-slate-id"
	location := "us-central1"
	ctx := context.Background()
	client, err := stitcher.NewVideoStitcherClient(ctx)
	if err != nil {
		return fmt.Errorf("stitcher.NewVideoStitcherClient: %v", err)
	}
	defer client.Close()

	req := &stitcherstreampb.GetSlateRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/slates/%s", projectID, location, slateID),
	}

	response, err := client.GetSlate(ctx, req)
	if err != nil {
		return fmt.Errorf("client.GetSlate: %v", err)
	}

	fmt.Fprintf(w, "Slate: %+v", response)
	return nil
}

// [END video_stitcher_get_slate]
