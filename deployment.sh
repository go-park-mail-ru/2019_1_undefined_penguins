#! /usr/bin/bash

ssh-keyscan -H $PRODUCTION_MACHINE_ADDRESS >> ~/.ssh/known_hosts
chmod 600 ./deployment_travis_key
ssh -i ./deployment_travis_key travis@$PRODUCTION_MACHINE_ADDRESS << EOF

export PATH=$PATH:/home/kate/.go/bin
cd /home/kate/2019_1_undefined_penguins
echo Building... && go build cmd/main.go && \
echo Restarting service... &&  systemctl restart penguin-backend.service && \
echo Enabling service ... && systemctl enable penguin-backend.service && \
echo Successfully deployed!!!
exit
EOF