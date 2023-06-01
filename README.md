## Fly Template Mongo

```sh
fly launch --copy-config --no-public-ips --no-deploy
fly secrets set MONGO_KEY="$(openssl rand -base64 756)" MONGO_INITDB_ROOT_USERNAME=admin MONGO_INITDB_ROOT_PASSWORD=admin
fly deploy
fly m update PRIMARY_MACHINE_ID --vm-memory 1024 --yes
fly ssh console -C "initialize-replica-set"
fly ssh console -C "create-root-user"

fly m clone PRIMARY_MACHINE_ID --region gig
fly ssh console -s -C "setup-machine NEW_MACHINE_ID"

```