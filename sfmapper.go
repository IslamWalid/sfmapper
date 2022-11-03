package sfmapper

import (
	"fmt"
	"reflect"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/IslamWalid/sfmapper/internal/fsnode"
	"github.com/fatih/structs"
)

// FS represents file system type.
type FS struct {
	// Reference to the struct that the user need to represent in the file system.
	UserStructRef any
}

// newFS Creates new file system object
func newFS(userStruct any) *FS {
	return &FS{
		UserStructRef: userStruct,
	}
}

// Mount mounts the file system to the given mount point and starts the file system server.
func Mount(mountPointPath string, userStruct any) error {
	conn, err := fuse.Mount(mountPointPath)
	if err != nil {
		return err
	}

	err = fs.Serve(conn, newFS(userStruct))
	if err != nil {
		return err
	}

	err = conn.Close()
	if err != nil {
		return err
	}

	return nil
}

func UnMount(mountPointPath string) error {
	err := fuse.Unmount(mountPointPath)
	if err != nil {
		return err
	}

	return nil
}

// Root initialize the root directory.
func (f *FS) Root() (fs.Node, error) {
	dir := fsnode.NewDir()
	structMap := structs.Map(f.UserStructRef)
	dir.Entries = f.createEntries(structMap, []string{})
	return dir, nil
}

// createEntries creates a map of directories and files a directory have.
func (f *FS) createEntries(structMap map[string]any, currentPath []string) map[string]any {
	entries := map[string]any{}

	for key, val := range structMap {
		if reflect.TypeOf(val).Kind() == reflect.Map {
			dir := fsnode.NewDir()
			dir.Entries = f.createEntries(val.(map[string]any), append(currentPath, key))
			entries[key] = dir
		} else {
			filePath := make([]string, len(currentPath))
			copy(filePath, currentPath)
			content := []byte(fmt.Sprintln(reflect.ValueOf(val)))
			file := fsnode.NewFile(key, filePath, len(content), f.UserStructRef)
			entries[key] = file
		}
	}

	return entries
}
