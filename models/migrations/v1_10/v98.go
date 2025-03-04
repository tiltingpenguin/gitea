// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1_10 //nolint

import "xorm.io/xorm"

func AddOriginalAuthorOnMigratedReleases(x *xorm.Engine) error {
	type Release struct {
		ID               int64
		OriginalAuthor   string
		OriginalAuthorID int64 `xorm:"index"`
	}

	return x.Sync2(new(Release))
}
