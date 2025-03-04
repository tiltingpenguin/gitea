// Copyright 2022 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db_test

import (
	"context"
	"testing"

	"code.gitea.io/gitea/models/db"
	repo_model "code.gitea.io/gitea/models/repo"
	"code.gitea.io/gitea/models/unittest"

	"github.com/stretchr/testify/assert"
)

func TestIterate(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	xe := unittest.GetXORMEngine()
	assert.NoError(t, xe.Sync(&repo_model.RepoUnit{}))

	var repoCnt int
	err := db.Iterate(db.DefaultContext, nil, func(ctx context.Context, repo *repo_model.RepoUnit) error {
		repoCnt++
		return nil
	})
	assert.NoError(t, err)
	assert.EqualValues(t, 79, repoCnt)

	err = db.Iterate(db.DefaultContext, nil, func(ctx context.Context, repoUnit *repo_model.RepoUnit) error {
		reopUnit2 := repo_model.RepoUnit{ID: repoUnit.ID}
		has, err := db.GetByBean(ctx, &reopUnit2)
		if err != nil {
			return err
		} else if !has {
			return db.ErrNotExist{Resource: "repo_unit", ID: repoUnit.ID}
		}
		assert.EqualValues(t, repoUnit.RepoID, repoUnit.RepoID)
		assert.EqualValues(t, repoUnit.CreatedUnix, repoUnit.CreatedUnix)
		return nil
	})
	assert.NoError(t, err)
}
