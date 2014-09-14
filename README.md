Cloudflare API
==============

This is a library for use in conjunction with the [Cloudflare API](https://www.cloudflare.com/docs/client-api.html).

## Example

```
package main

import (
    "fmt"
    cf "github.com/consulted/cloudflare/lib"
)

func main() {
    client := cf.Client{
        Email: "<Your email>",
        Token: "<Your token>",
    }
    zones, err := client.GetZoneList()
    if err != nil {
        panic(err)
    }
    fmt.Println(zones.Count)
}
```
