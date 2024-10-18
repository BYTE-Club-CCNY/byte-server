#!/usr/bin/env bash

# our server routes port 80 to 3000 to get around admin perms to use port 80 
                                      # (doing that would mess up our CI/CD)

# sudo iptables -t nat -D PREROUTING -p tcp --dport 80 -j REDIRECT --to-port 3000

sudo certbot certonly --standalone -n --agree-tos --renew-by-default -d "test.byteccny.com"
# this may ask for inputs
# check certbot -h for details

# sudo iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-port 3000
