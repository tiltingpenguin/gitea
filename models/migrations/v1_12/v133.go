// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1_12 //nolint

import "xorm.io/xorm"

func AddEmailHashTable(x *xorm.Engine) error {
	// EmailHash represents a pre-generated hash map
	type EmailHash struct {
		Hash  string `xorm:"pk varchar(32)"`
		Email string `xorm:"UNIQUE NOT NULL"`
	}
	return x.Sync2(new(EmailHash))
}
