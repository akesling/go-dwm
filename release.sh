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

function determine_old_version() {
    cat ./dwm/config.h | \
        grep '#define VERSION' | \
        cut -d' ' -f3 | \
        sed -e 's/"//g'
}

function prompt_for_input() {
    local NEW_VERSION
    read -p'> ' NEW_VERSION
    printf "${NEW_VERSION}"
}

function generate_commit_message() {
    local REVISION_NOTES=$1
    local OLD_VERSION=$2
    local NEW_VERSION=$3
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
        exit 1
    fi
}

function write_new_version_to_config_file() {
    local NEW_VERSION="$1"
    local CONFIG_FILE="$2"
    local TMP_CONFIG="$3"
    cat ${CONFIG_FILE} |
        sed -e "s/\(\#define VERSION \"\)[.0-9]*\(\"\)/\1${NEW_VERSION}\2/" > ${TMP_CONFIG}
}

function exit_if_new_config_file_doesnt_look_right() {
    local EXPECTED_CONFIG_DIFF=$(cat <<HERE_DOC
3c3
< #define VERSION "${OLD_VERSION}"
---
> #define VERSION "${NEW_VERSION}"
HERE_DOC
)
    local ACTUAL_CONFIG_DIFF=$(diff ${CONFIG_FILE} ${TMP_CONFIG})
    if [ "${EXPECTED_CONFIG_DIFF}" != "${ACTUAL_CONFIG_DIFF}" ]; then
        echo "Generated configuration did not match expectation"
        echo "Expected:"
        echo "${EXPECTED_CONFIG_DIFF}"
        echo "Got:"
        echo "${ACTUAL_CONFIG_DIFF}"
        exit 1
    fi
}

################################################################################
## Script Body #################################################################
################################################################################

function release_script_main() {
    if [ -z "$(gather_revision_notes)" ]; then
        echo "Version up to date with dev branch."
        exit 0
    fi

    echo "Gathering revision notes."
    REVISION_NOTES="$(gather_revision_notes)"
    change_branch_to_master

    echo "Determining old version."
    OLD_VERSION="$(determine_old_version)"

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

    CONFIG_FILE=./dwm/config.h
    echo "Rewriting version from ${OLD_VERSION} to ${NEW_VERSION} in ${CONFIG_FILE}."
    TMP_CONFIG=$(mktemp)
    write_new_version_to_config_file "${NEW_VERSION}" "${CONFIG_FILE}" "${TMP_CONFIG}"

    echo "Verifying new version of ${CONFIG_FILE} looks as expected."
    exit_if_new_config_file_doesnt_look_right
    mv ${TMP_CONFIG} ${CONFIG_FILE}

    echo "Rolling up dev changes into commit."
    roll_up_dev_changes_into_master_commit

    echo "Committing new versioned config."
    git add ${CONFIG_FILE}
    git commit -m "Update version to ${NEW_VERSION}"

    echo "Creating tag for version ${NEW_VERSION}"
    git tag -a "${NEW_VERSION}" -m "Update version to ${NEW_VERSION}"

    echo "Version updated to ${NEW_VERSION}, rolled up from dev branch, and tagged."
}

release_script_main
