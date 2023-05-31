#!/bin/sh
hostname "$FLY_MACHINE_ID.vm.$FLY_APP_NAME.internal"
echo "$FLY_PUBLIC_IP  $FLY_MACHINE_ID.vm.$FLY_APP_NAME.internal" >> /etc/hosts

mkdir -p /data/db /data/configdb
chown -R mongodb:mongodb /data/db /data/configdb
echo "$MONGO_KEY" > /etc/mongo.key
chown mongodb:mongodb /etc/mongo.key
chmod 400 /etc/mongo.key

cat /etc/mongo.key


start