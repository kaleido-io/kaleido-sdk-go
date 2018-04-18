#!/bin/bash -eu
set -o pipefail

COMMIT_LEVEL=$(git rev-parse --short HEAD)
TARGET_REPO=jimthematrix.github.io
git clone https://github.com/jimthematrix/$TARGET_REPO.git

cp kld-* $TARGET_REPO/assets/
cd $TARGET_REPO
git add .
git commit -m "Build commit - $COMMIT_LEVEL"
git config remote.gh-pages.url https://jimthematrix1:$GITPAGE_PASSWORD@github.com/jimthematrix/$TARGET_REPO.git

# Push API docs to Target repository
git push gh-pages master
