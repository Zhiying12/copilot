serverPort=$1

bin/server -maddr=172.31.22.119 -addr=172.31.18.42 -port=${serverPort} -e=false -p=8 -copilot=true -exec=true -dreply=true

