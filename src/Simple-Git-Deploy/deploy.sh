#!/bin/bash

while getopts ':k:g:d:c:' OPTION ; do
  case "$OPTION" in
    k)   KEY=$OPTARG;;
    g)   GIT_URL=$OPTARG;;
    d)   DEPLOY_DIR=$OPTARG;;
    c)	 DCONTAINER_NAME=$OPTARG;;
    *)   echo "Unknown parameter"
  esac
done

#Check for Key
if [ -z "$KEY" ]
  then
    echo "You must provide a ssh-keyfile (via -k /path/to/key)!"
    exit 1
fi

#Check for Git-URL
if [ -z "$GIT_URL" ]
  then
    echo "You must provide a git-url to clone from (via -g git_url)!"
    exit 1
fi

#Check for deploy-path
if [ -z "$DEPLOY_DIR" ]
  then
    echo "You must provide a deploydir path (via -c /path/to/deploy/dir)!"
    exit 1
fi

echo "-----------------------------"
echo "Started at:"
date
echo "Key: ${KEY}"
echo "Git-Url: ${GIT_URL}"
echo "Deploy-Dir: ${DEPLOY_DIR}"
#echo "ssh-agent bash -c 'ssh-add ${KEY}; git clone ${GIT_URL} /tmp/git-deploy'"
echo "----------------------------"

#mkdir ${DEPLOY_DIR}_dep
rm -Rf /tmp/git-deploy
ssh-agent bash -c "ssh-add ${KEY}; git clone ${GIT_URL} /tmp/git-deploy"
rm -Rf /tmp/git-deploy/.git
rm -Rf ${DEPLOY_DIR}/.* 2> /dev/null
rm -Rf ${DEPLOY_DIR}/*
#sleep 1

mv -fT /tmp/git-deploy ${DEPLOY_DIR}
#cp -urfv /tmp/git-deploy/* ${DEPLOY_DIR}/*
rm -rf /tmp/git-deploy

#echo $DCONTAINER_NAME
if [ ! -z "$DCONTAINER_NAME" ]
  then
    docker restart $DCONTAINER_NAME
fi

#[! -z "$DCONTAINER_NAME" ] && docker restart "$DCONTAINER_NAME"


echo "Finished at:"
date
echo "-----------------------------"
