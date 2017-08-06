// Copyright 2017-present The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

import (
	"github.com/gohugoio/hugo/config"
	"github.com/spf13/cast"
)

func GetContentDirsNoFallbackToContentDir(cfg config.Provider) []string {
	contentDirs := cast.ToStringSlice(cfg.Get("contentDirs"))
	return contentDirs
}

func GetContentDirs(cfg config.Provider) []string {
	contentDirs := GetContentDirsNoFallbackToContentDir(cfg)
	if len(contentDirs) == 0 && cfg.IsSet("contentDir") {
		contentDirs = append(contentDirs, cfg.GetString("contentDir"))
	}
	return contentDirs
}

func GetContentDirsAbsolutePaths(cfg config.Provider, ps *PathSpec) []string {
	contentDirs := GetContentDirs(cfg)
	for i, contentDir := range contentDirs {
		contentDirs[i] = ps.AbsPathify(contentDir)
	}
	return contentDirs
}
