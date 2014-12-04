Cloudflare API
==============

This is a library for use in conjunction with the [Cloudflare API](https://www.cloudflare.com/docs/client-api.html).

It is neither complete nor very stable. It includes a very basic `cloudflare` command to add and remove dns records from a given zone.

## Parameters

Environment variables:

```
CF_TOKEN = <Your API token>
CF_EMAIL = <You Cloudflare account email>
CZ_ZONE  = <The zone to use for all operations> (optional, can be set via the -z flag)
```

## Add a record

```
./cloudflare [-z consulted.com] add -c [content] -n [name] -t [type]
```

This will also update the record to be immediately Cloudflare proxy enabled (orange cloud)

## Remove a record

```
./cloudflare [-z zone] delete -c [content]
```

Removal will ask for your confirmation, howver, you can pass the `-y` flag to skip the confirmation. All records with a given content will be deleted.
