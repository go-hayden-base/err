package err

const (
	// Normal error code
	ErrCodeUnknown = iota - 999999
	ErrCodeNotExist
	ErrCodeParamInvalid
	ErrCodeEmpty
	ErrCodeDuplicate

	// File System error code
	ErrCodeFileNoSuchFile
	ErrCodeFileLocking
	ErrCodeFileReadUnknown
	ErrCodeFileReadNoPermission
	ErrCodeFileReadInvalidFileName
	ErrCodeFileReadCorruptFile
	ErrCodeFileReadNoSuchFile
	ErrCodeFileReadInapplicableStringEncoding
	ErrCodeFileReadUnsupportedScheme
	ErrCodeFileReadTooLarge
	ErrCodeFileReadUnknownStringEncoding
	ErrCodeFileWriteUnknown
	ErrCodeFileWriteNoPermission
	ErrCodeFileWriteInvalidFileName
	ErrCodeFileWriteInapplicableStringEncoding
	ErrCodeFileWriteUnsupportedScheme
	ErrCodeFileWriteOutOfSpace
	ErrCodeFileWriteVolumeReadOnly
)
