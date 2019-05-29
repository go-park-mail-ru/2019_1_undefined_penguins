#! /usr/bin/bash

ssh-keyscan -H $PRODUCTION_MACHINE_ADDRESS >> ~/.ssh/known_hosts
chmod 600 ./deployment_travis_key
ssh -i ./deployment_travis_key travis@$PRODUCTION_MACHINE_ADDRESS << EOF
cd front
echo Aquiring fresh version of repo... && git checkout deployment && \
echo Pulling changes... && git pull && \
echo Building... && go build cmd/main.go && \
echo Restarting service... &&  systemctl restart penguin-backend.service && \
echo Enabling service ... && systemctl enable penguin-backend.service && \
echo Successfully deployed!!!
exit
EOF