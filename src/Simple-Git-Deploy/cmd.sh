#!/bin/bash

while getopts ':k:g:d:' OPTION ; do
  case "$OPTION" in
    k)   KEY=$OPTARG;;
    g)   GIT_URL=$OPTARG;;
    d)   DEPLOY_DIR=$OPTARG;;
    *)   echo "Unknown parameter"
  esac
done

echo "-----------------------------"
echo "Key: ${KEY}"
echo "Git-Url: ${GIT_URL}"
echo "Deploy-Dir: ${DEPLOY_DIR}"
echo "----------------------------"

rm -Rf /tmp/git-deploy
ssh-agent bash -c "ssh-add ${KEY}; git clone ${GIT_URL} /tmp/git-deploy"

rm -Rf /tmp/git-deploy/.git # Remove the Git Folder
# Empty the deploy-directory without deleting itself
rm -Rf ${DEPLOY_DIR}/.* 2> /dev/null
rm -Rf ${DEPLOY_DIR}/*
# Move the newly downloaded files to the deploy-folder
mv -fT /tmp/git-deploy ${DEPLOY_DIR}