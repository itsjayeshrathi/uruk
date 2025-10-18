package util

import (
	"errors"
	"fmt"
)

// file or folder errors
var (
	ErrFileNotFound     = errors.New("file not found")
	ErrInvalidPath      = errors.New("invalid file path")
	ErrPermissionDenied = errors.New("permission denied")
	ErrFileEmpty        = errors.New("file is empty")
	ErrDirectoryEmpty   = errors.New("directory is empty")
	ErrUnsupportedType  = errors.New("unsupported file type")
)

// custom error types for more detailed error info
type FileError struct {
	Op   string
	Path string
	Err  error
}

func (e *FileError) Error() string {
	if e.Path == "" {
		return fmt.Sprintf("%s: %v:", e.Op, e.Err)
	}
	return fmt.Sprintf("%s %s: %v", e.Op, e.Path, e.Err)
}

func (e *FileError) UnWrap() error {
	return e.Err
}

func NewFileError(op, path string, err error) *FileError {
	return &FileError{
		Op:   op,
		Path: path,
		Err:  err,
	}
}

// compression error wraps compression related errors

type CompressionError struct {
	Method string
	Err    error
}

func (e *CompressionError) Error() string {
	return fmt.Sprintf("compression error (%s): %v", e.Method, e.Err)
}

func (e *CompressionError) UnWrap() error {
	return e.Err
}

func NewCompressionError(method string, err error) *CompressionError {
	return &CompressionError{
		Method: method,
		Err:    err,
	}
}

// validation eror represents validation failures

type ValidationError struct {
	Field   string
	Value   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s='%s': %s", e.Field, e.Value, e.Message)
}

func NewValidationError(field, value, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
	}
}


