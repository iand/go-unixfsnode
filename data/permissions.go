package data

const FilePermissionsDefault = 0o0644
const DirectorPerimissionsDefault = 0o0755
const HAMTShardPerimissionsDefault = 0o0755

func (u UnixFSData) Permissions() int {
	if u.FieldMode().Exists() {
		return int(u.FieldMode().Must().Int() & 0xFFF)
	}
	return DefaultPermissions(u)
}

func DefaultPermissions(u UnixFSData) int {
	if u.FieldDataType().Int() == Data_File {
		return FilePermissionsDefault
	}
	if u.FieldDataType().Int() == Data_Directory {
		return DirectorPerimissionsDefault
	}
	if u.FieldDataType().Int() == Data_HAMTShard {
		return HAMTShardPerimissionsDefault
	}
	return 0
}
