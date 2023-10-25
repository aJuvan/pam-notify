# PAM Notify

## Build

```sh
go build
```

## Configure

Copy example config to  `/etc/pam-notify.yml` and edit it, then add the
following lines to `/etc/pam.d/<service>` for each service you'd like to
monitor:

```
session optional pam_exec.so seteuid /path/to/pam-notify
```
