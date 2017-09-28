#!/bin/bash


IFACE_NAME=$1
: ${IFACE_NAME:="eth0"}
echo $IFACE_NAME

ARG=$2
: ${ARG:="up"}
echo $ARG

 networkup () {
for id in $(seq 3 20)
do
    echo $id
    ifconfig $IFACE_NAME:$id 12.12.12.$id netmask 255.255.255.0 up
done
}

 networkdown () {
for id in $(seq 3 20)
do
    echo $id
    ifconfig $IFACE_NAME:$id  12.12.12.$id down
done
}

if [ "${ARG}" = "up" ]; then
	echo "up"
	networkup
elif [ "${ARG}" = "down" ]; then ## Clear the network
	networkdown
  	echo "down"
fi

#sudo sh -x ./ip.sh enp0s25 down
#sudo sh -x ./ip.sh enp0s25 up




