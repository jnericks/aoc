#!/bin/bash

# this runs each time the container starts

echo "post-start start"
echo "$(date +'%Y-%m-%d %H:%M:%S')    post-start start" >> "$HOME/status"

# put things in here like below you need to run after startup
# sudo tailscale up --accept-routes --advertise-exit-node --auth-key $TS_AUTH_KEY
echo "post-start complete"
echo "$(date +'%Y-%m-%d %H:%M:%S')    post-start complete" >> "$HOME/status"
