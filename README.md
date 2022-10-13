# Struct File Mapper
Struct File Mapper uses [Filesystem in Userspace](https://en.wikipedia.org/wiki/Filesystem_in_Userspace) feature in Linux systems to map in-memory data **(user struct)** to accessible file system.

# Install
```
go get github.com/IslamWalid/struct_file_mapper
```

# Functionality
| Function | Description |
|----------|-------------|
| `func Mount(mountPoint string, structReference any) error` | Creates a filesystem in the specified mount point. |

# Usage and Example
```
package main

import (
	"fmt"
	"os"

	sfmapper "github.com/IslamWalid/struct_file_mapper"
)

type Person struct {
	Name   string
	Age    int
}

func main() {
    p := &Person{
    	Name: "Islam",
    	Age:  22,
    }

    // Create directory person to host the filesystem
    os.MkdirAll("person", 0777)

    // Start the filesystem
    err := sfmapper.Mount("person", p)
    if err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}
```

**NOTE:** Arrays and slices are not supported.
