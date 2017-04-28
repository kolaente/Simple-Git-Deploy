# Simple Git Deploy

# Troubleshooting

### Not Cloning - the script runs pretty fast but wont clone

Check the permissions of your keyfile. They should be 600, otherwise the ssh-agent will ignore the keyfile and won't clone.

```
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@         WARNING: UNPROTECTED PRIVATE KEY FILE!          @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
Permissions 0664 for '/home/konrad/Schreibtisch/life-wob/life-deploy' are too open.
It is required that your private key files are NOT accessible by others.
This private key will be ignored.
```