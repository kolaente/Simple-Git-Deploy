# Simple Git Deploy

This simple application provides automatic update of git repositories on your server.
Triggered by a webhook (like the one from github), it can be very useful when you don't have (or don't want) a fully-fleged CI-Server to update your website.

**Note: ** The application currently runs only under linux-systems as it uses a bash-script to update.

# Usage

Usage is pretty simple. Simply configure the application via config.ini:

```
ssh_key=/path/to/deploy-key # Path to the Deploy-SSH
git_url=ssh://git@your-server.tld/gitproject.git # SSH-Path to your Git-Project
deploy_dir=/path/to/deploy/dir # The folder to which the updated version of your site will be deployed. You need write permissions.
container_name= # If you run your site within a docker-container, the application can automatically restart the container for any changes to take effect.
secret= # The Secret sent by the webhook - to prevent abuse
interface=":8080" # The interface the server should listen on
branch_name=refs/heads/master # The name of the branch. If an update to this branch is triggered, the application will run the update.
```

# Troubleshooting

### Not Cloning - the script runs pretty fast but wont clone

Check the permissions of your keyfile. They should be 600, otherwise the ssh-agent will ignore the keyfile and won't clone:

```
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@         WARNING: UNPROTECTED PRIVATE KEY FILE!          @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
Permissions 0664 for '/path/to/deploy-key' are too open.
It is required that your private key files are NOT accessible by others.
This private key will be ignored.
```