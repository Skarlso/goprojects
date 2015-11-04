package abstract

import "fmt"

// Factory don't know yet.
type Factory interface {
	// Database
	// Filesystem
}

// Database describes a database which should be created. The client will choose
// the type of database they want.
type Database interface {
	RunQuery(sql string) (result string)
}

// File type
type File struct {
	content  string
	fileName string
}

// Filesystem describes a filesystem with the capabilities to open and read from
// files.
type Filesystem interface {
	CreateFile(fileName string) (bool, error)
	DeleteFile(fileName string) (bool, error)
}

// MongoDB a concrete Database implementation
type MongoDB struct {
	database map[string]string
}

// OracleDB a concrete Database implementation
type OracleDB struct {
	database map[string]string
}

// ZFS concrete filesystem
type ZFS struct {
	files map[interface{}]File
}

// NTFS concrete filesystem
type NTFS struct {
	files map[interface{}]File
}

// RunQuery runs a query on a OracleDB type database
func (odb *OracleDB) RunQuery(sql string) (result string) {
	return odb.database[sql]
}

// RunQuery runs a query on a MongoDB type database
func (mdb *MongoDB) RunQuery(sql string) (result string) {
	return mdb.database[sql]
}

// CreateFile creates file
func (zfs *ZFS) CreateFile(fileName string) (bool, error) {
	file := File{content: "Content", fileName: fileName}
	zfs.files[fileName] = file
	if _, ok := zfs.files[fileName]; ok {
		return true, nil
	}

	return false, fmt.Errorf("Something bad happened.")
}

// DeleteFile deletes file
func (zfs *ZFS) DeleteFile(fileName string) (bool, error) {
	delete(zfs.files, fileName)
	if f, ok := zfs.files[fileName]; ok {
		return false, fmt.Errorf("File is still there.%v", f)
	}

	return true, nil
}

// CreateFile creates file
func (ntfs *NTFS) CreateFile(fileName string) (bool, error) {
	file := File{content: "Bla", fileName: fileName}
	ntfs.files[fileName] = file
	if _, ok := ntfs.files[fileName]; ok {
		return true, nil
	}

	return false, fmt.Errorf("Something bad happened.")
}

// DeleteFile deletes file
func (ntfs *NTFS) DeleteFile(fileName string) (bool, error) {
	delete(ntfs.files, fileName)
	if f, ok := ntfs.files[fileName]; ok {
		return false, fmt.Errorf("File is still there.%v", f)
	}

	return true, nil

}

// DatabaseFactory returns a database
func DatabaseFactory(databaseType string) Database {
	switch databaseType {
	case "mongo":
		return new(MongoDB)
	case "oracle":
		return new(OracleDB)
	}
	return nil
}

// FilesystemFactory returns a filesystem
func FilesystemFactory(filesystemType string) Filesystem {
	switch filesystemType {
	case "zfs":
		return new(ZFS)
	case "ntfs":
		return new(NTFS)
	}
	return nil
}

func main() {

}
