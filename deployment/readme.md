# What is This?
- to run the server in the background, we run it in the background as a linux process
- the process is defined in `byteserver.service`
- it is copied over to `/etc/systemd/system/` which is where all services are stored
- `restart-service` refreshes the `systemctl` daemon when the application is updated
- `update-service` refreshes the `systemctl` daemon when the service configuration is changed
- `nginx.conf` is the configuration for nginx which allows our server to serve HTTP and HTTPS without running the application in root