Cloudflare API
==============

This is a library for use in conjunction with the [Cloudflare API](https://www.cloudflare.com/docs/client-api.html).

It is neither complete nor very stable. It includes a very basic `cloudflare` command to add and remove dns records from a given zone.

## Add a record

```
cloudflare -c <content> -n <name> -z <name-of-zone>
```

This will also update the record to be immediately Cloudflare proxy enabled (orange cloud)

## Remove a record

```
cloudflare -n <name> -z <name-of-zone> -d
```

## Parameters

Environment variables:

```
CF_TOKEN = <Your API token>
CF_EMAIL = <You Cloudflare account email>
CZ_ZONE  = <The zone to use for all operations> (optional, can be set via the -z flag)
```

Flags:

```
-c <content> - the content for the record
-n <name> - the name of the record
-t <type> - the type for the record (A/CNAME/MX/TXT/SPF/AAAA/NS/SRV/LOC), default is "A"
-z <zone> - optional, the name of the zone
-d - delete that record
```


## TODO

- [ ] Split the commands up
