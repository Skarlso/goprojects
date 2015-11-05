package main

import "fmt"

// Factory A Factory which defines a GetFactory method to return a factory instance
type Factory interface {
	GetFactory() Factory
}

// Database describes a database which should be created. The client will choose
// the type of database they want.
type Database interface {
	RunQuery(sql string) (result string)
	AddData(sql string, data string) error
}

// Filesystem describes a filesystem with the capabilities to open and read from
// files.
type Filesystem interface {
	CreateFile(fileName string) (bool, error)
	GetFile(fileName string) (File, error)
}

// File type
type File struct {
	content  string
	fileName string
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
	files map[string]File
}

// NTFS concrete filesystem
type NTFS struct {
	files map[string]File
}

// Databases abstract database groupper which acts like it is a mongoDb or an OracleDB
type Databases struct {
	*MongoDB
	*OracleDB
}

// Filesystems abstract filesystem groupper which acts like a ZFS or an NTFS
type Filesystems struct {
	*ZFS
	*NTFS
}

// RunQuery runs a query on a OracleDB type database
func (odb OracleDB) RunQuery(sql string) (result string) {
	return odb.database[sql]
}

// RunQuery runs a query on a MongoDB type database
func (mdb MongoDB) RunQuery(sql string) (result string) {
	return mdb.database[sql]
}

// AddData add data to the database
func (odb *OracleDB) AddData(sql string, data string) error {
	odb.database[sql] = data
	return nil
}

// AddData add data to the database
func (mdb *MongoDB) AddData(sql string, data string) error {
	mdb.database[sql] = data
	return nil
}

// CreateFile creates file
func (zfs *ZFS) CreateFile(fileName string) (bool, error) {
	file := File{content: "ZFS Content", fileName: fileName}
	zfs.files[fileName] = file
	if _, ok := zfs.files[fileName]; ok {
		return true, nil
	}

	return false, fmt.Errorf("Something bad happened.")
}

// GetFile gets a file
func (zfs *ZFS) GetFile(fileName string) (File, error) {
	if f, ok := zfs.files[fileName]; ok {
		return f, nil
	}

	return File{}, fmt.Errorf("File is still there.")
}

// CreateFile creates file
func (ntfs *NTFS) CreateFile(fileName string) (bool, error) {
	file := File{content: "NTFS Content", fileName: fileName}
	ntfs.files[fileName] = file
	if _, ok := ntfs.files[fileName]; ok {
		return true, nil
	}

	return false, fmt.Errorf("Something bad happened.")
}

// GetFile gets a file
func (ntfs *NTFS) GetFile(fileName string) (File, error) {
	if f, ok := ntfs.files[fileName]; ok {
		return f, nil
	}

	return File{}, fmt.Errorf("File is still there.")
}

// GetFactory Create a Factory for the databases
func (db Databases) GetFactory() Factory {
	return Databases{&MongoDB{make(map[string]string)}, &OracleDB{make(map[string]string)}}
}

// GetFactory Create a Factory for the filesystems
func (fs Filesystems) GetFactory() Factory {
	return Filesystems{&ZFS{make(map[string]File)}, &NTFS{make(map[string]File)}}
}

// GetFactory is an abstract factory which returns factories
func GetFactory(factoryType string) Factory {
	switch factoryType {
	case "database":
		return Databases{}.GetFactory()
	case "filesystems":
		return Filesystems{}.GetFactory()
	}
	return nil
}

// GetDatabase This works like a concrete database factory. It returns a concrete
// database based on databaseType
func GetDatabase(databaseType string) Database {
	f := GetFactory("database")
	switch databaseType {
	case "mongo":
		return f.(Databases).MongoDB
	case "oracle":
		return f.(Databases).OracleDB
	}
	return nil
}

// GetFilesystems This works like a concrete filesystem factory. Returns a concrete
// filesystem
func GetFilesystems(filesystemType string) Filesystem {
	f := GetFactory("filesystems")
	switch filesystemType {
	case "zfs":
		return f.(Filesystems).ZFS
	case "ntfs":
		return f.(Filesystems).NTFS
	}
	return nil
}

func main() {
	database := GetDatabase("mongo")
	database.AddData("bla", "data bla")
	fmt.Println("database: ", database.RunQuery("bla"))

	filesystem := GetFilesystems("zfs")
	filesystem.CreateFile("bla")
	file, _ := filesystem.GetFile("bla")
	fmt.Println("file content: ", file.content)
}
