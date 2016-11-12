#!/usr/bin/env bash

################################################################################
## Component Functions #########################################################
################################################################################

trap "exit 1" TERM
export SCRIPT_CONTEXT=$$

function exit_script() {
    kill -s TERM ${SCRIPT_CONTEXT}
}

function gather_revision_notes() {
    git log master...dev --pretty=format:'*) %s'
}

function change_branch_to_master() {
    if [ $(git rev-parse --abbrev-ref HEAD) != 'master' ]; then
        git checkout master
    fi
}

function prompt_for_input() {
    local NEW_VERSION
    read -p'> ' NEW_VERSION
    printf "${NEW_VERSION}"
}

function generate_commit_message() {
    local REVISION_NOTES="$1"
    local OLD_VERSION="$2"
    local NEW_VERSION="$3"
    cat <<HERE_DOC
Merge for version ${NEW_VERSION} update

Previous version was ${OLD_VERSION}

Revision notes:
${REVISION_NOTES}
HERE_DOC
}

function roll_up_dev_changes_into_master_commit() {
    change_branch_to_master
    git merge --no-ff dev -m "${COMMIT_MESSAGE}"
}

function ask_user_whether_to_continue_or_exit() {
    local AFFIRM
    read -p'Is this correct? [Y/n] ' AFFIRM

    if [ ${AFFIRM:-Y} = 'n' ]; then
        printf "Commit message rejected, aborting"
        exit_script
    fi
}

################################################################################
## Script Body #################################################################
################################################################################

function release_script_main() {
    local VERSION_CONFIG=VERSION

    if [ -z "$(gather_revision_notes)" ]; then
        echo "Version up to date with dev branch."
        exit 0
    fi

    echo "Gathering revision notes."
    REVISION_NOTES="$(gather_revision_notes)"
    change_branch_to_master

    echo "Determining old version."
    OLD_VERSION="$(cat ${VERSION_CONFIG})"

    echo "Please provide new version number.  Old version was ${OLD_VERSION}."
    NEW_VERSION="$(prompt_for_input)"

    echo "Generating commit message from revision notes, old version, and new version."
    COMMIT_MESSAGE=$(generate_commit_message "${REVISION_NOTES}" "${OLD_VERSION}" "${NEW_VERSION}")

    echo "Commit message for version merge from dev branch is as follows:"
    echo "\"\"\""
    echo "${COMMIT_MESSAGE}"
    echo "\"\"\""
    echo ""

    ask_user_whether_to_continue_or_exit

    echo "Rolling up dev changes into commit."
    roll_up_dev_changes_into_master_commit

    echo "Writing new version to the VERSION file."
    echo "${NEW_VERSION}" > "${VERSION_CONFIG}"

    echo "Running go generate to update files for the new release."
    go generate ./...

    echo "Committing new versioned config."
    git add "${VERSION_CONFIG}"
    git commit -m "Update version to ${NEW_VERSION}"

    echo "Creating tag for version ${NEW_VERSION}"
    git tag -a "${NEW_VERSION}" -m "Update version to ${NEW_VERSION}"

    echo "Version updated to ${NEW_VERSION}, rolled up from dev branch, and tagged."
}

release_script_main
