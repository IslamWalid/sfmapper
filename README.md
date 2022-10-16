# Struct File Mapper
Struct File Mapper uses [Filesystem in Userspace](https://en.wikipedia.org/wiki/Filesystem_in_Userspace) feature in Linux systems to map in-memory data **(user struct)** to accessible file system.

# Install
```
go get github.com/IslamWalid/struct_file_mapper@latest
```

# Functionality
| Function | Description |
|----------|-------------|
| `func Mount(mountPointPath string, structReference any) error` | Creates a filesystem and mounts it to the given mount point path. |
| `func Unmount(mountPointPath string) error` | Unmount the filesystem in the given mount point path.

**NOTE**: Unmount the filesystem after using it using the `Umount` function, or through the command:
```
fusermount -u <mount_point_path>
```

# Usage and Example
```go
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

    // Create signal to recieve ctrl+c
    sigs := make(chan os.Signal)
    signal.Notify(sigs, syscall.SIGINT)

    go func() {
        <- sigs
        sfmapper.UnMount("person")
    }()

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
