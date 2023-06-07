## Fly Template Mongo

This is an example on how to run a MongoDB Replica Set Cluster. Use this for educational or curiosity purposes, it's not meant to be ran in production.

The first step is to deploy the first machine. Clone this repo and inside it run these commands.

```sh
fly launch --copy-config --no-public-ips --no-deploy
fly secrets set MONGO_KEY="$(openssl rand -base64 756)" MONGO_INITDB_ROOT_USERNAME=admin MONGO_INITDB_ROOT_PASSWORD=admin
fly deploy
fly m update PRIMARY_MACHINE_ID --vm-memory 1024 --yes
```

By now you'll have a uninitialized single machine Mongo server. The default user will have `MONGO_INITDB_ROOT_USERNAME` and `MONGO_INITDB_ROOT_PASSWORD` as credentials.

```sh
fly ssh console -C "initialize-replica-set"
fly ssh console -C "create-root-user"
```

## Horizontal scale

All you need to do is clone the primary then run a `setup-command` inside your primary node.

```sh
fly m clone PRIMARY_MACHINE_ID --region gig
fly ssh console -s -C "setup-machine NEW_MACHINE_ID" # run this on your primary
```
