#!/usr/bin/env bash

REVISION_NOTES=$(git log master...dev --pretty=format:'*) %s')

if [ -z ${REVISION_NOTES} ]; then
    echo "Version up to date with dev branch."
    exit 0
fi

if [ $(git rev-parse --abbrev-ref HEAD) != 'master' ]; then
    git checkout master
fi

OLD_VERSION=$(cat ./dwm/config.h | \
              grep '#define VERSION' | \
              cut -d' ' -f3 | \
              sed -e 's/"//g')

echo "Please provide new version number.  Old version was ${OLD_VERSION}."
read -p'> ' NEW_VERSION

COMMIT_MESSAGE=$(cat <<HERE_DOC
Merge for version ${NEW_VERSION} update

Revision notes:
${REVISION_NOTES}
HERE_DOC
)

cat <<HERE_DOC
Commit message for version merge from dev branch is as follows:
"""
${COMMIT_MESSAGE}
"""

HERE_DOC

read -p'Is this correct? [Y/n] ' AFFIRM

if [ ${AFFIRM:-Y} = 'n' ]; then
    printf "Commit message rejected, aborting"
    exit 1
fi

echo "Rolling up dev changes into commit."

git merge --no-ff dev -m ${COMMIT_MESSAGE}

CONFIG_FILE=./dwm/config.h
#cat ${CONFIG_FILE} | sed -e 's/(#define VERSION ")[.\d]*(")/'

git add ${CONFIG_FILE}
git commit -m "Update version to ${NEW_VERSION}"
git tag -a ${NEW_VERSION} -m "Update version to ${NEW_VERSION}"

echo "Version updated to ${NEW_VERSION}, rolled up from dev branch, and tagged."
