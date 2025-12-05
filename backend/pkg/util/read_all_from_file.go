package util

import "io/fs"

// Helper function to read all data from file (compatibility for older Go versions)
func ReadAllFromFile(file fs.File) ([]byte, error) {
	// Try to get file info to allocate correct buffer size
	var data []byte

	if fileInfo, err := file.Stat(); err == nil {
		// If we can get file size, allocate exact buffer
		size := fileInfo.Size()
		if size > 0 && size < 50*1024*1024 { // Max 50MB for safety
			data = make([]byte, size)
			n, err := file.Read(data)
			if err != nil && err.Error() != "EOF" {
				return nil, err
			}
			return data[:n], nil
		}
	}

	// Fallback: read in chunks
	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			data = append(data, buffer[:n]...)
		}
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
	}
	return data, nil
}
