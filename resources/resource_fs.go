// Code generated by vfsgen; DO NOT EDIT.

package resources

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FileSystem statically implements the virtual filesystem provided to vfsgen.
var FileSystem = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2021, 1, 6, 9, 42, 37, 242023452, time.UTC),
		},
		"/migrations": &vfsgen۰DirInfo{
			name:    "migrations",
			modTime: time.Date(2021, 1, 28, 6, 21, 54, 899353173, time.UTC),
		},
		"/migrations/000001_create_project_table.down.sql": &vfsgen۰FileInfo{
			name:    "000001_create_project_table.down.sql",
			modTime: time.Date(2020, 12, 18, 6, 45, 40, 917000000, time.UTC),
			content: []byte("\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x70\x72\x6f\x6a\x65\x63\x74\x3b"),
		},
		"/migrations/000001_create_project_table.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_project_table.up.sql",
			modTime:          time.Date(2021, 1, 1, 9, 30, 39, 157449214, time.UTC),
			uncompressedSize: 326,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x8c\xc1\x4a\xc3\x40\x14\x45\xf7\xf9\x8a\x4b\x57\x09\x28\x54\x70\xd7\xd5\xb4\x7d\xa5\xa3\xc9\xa4\xce\xbc\xd1\xd6\x4d\x08\x9d\x67\x89\x68\x12\xd2\xc4\xef\x97\x8e\xe2\xc2\x8d\xb8\xbc\x97\x73\xce\xca\x92\x62\x02\xed\x99\x8c\xd3\xa5\x81\xde\xc0\x94\x0c\xda\x6b\xc7\x0e\xb3\x69\x6a\xc2\x75\x77\x3e\xf7\xb3\x45\xf2\xcd\xb2\x5a\xe6\xf4\x8b\xeb\x87\xee\x55\x8e\x23\xd2\x04\x40\x13\xe0\xbd\x5e\x63\x67\x75\xa1\xec\x01\xf7\x74\xc0\x9a\x36\xca\xe7\x8c\x4b\xaf\x3a\x49\x2b\x43\x3d\x4a\xf5\x71\x9b\x66\x57\x17\xa5\xad\xdf\x05\x8f\xca\xae\xb6\xca\xa6\x37\xf3\x79\x16\xe3\xc6\xe7\x39\xbc\xd1\x0f\x9e\x22\x75\xec\xda\x97\xe6\x84\x3b\x57\x9a\xe5\xd7\x31\x48\x3d\x4a\xa8\xea\x11\xac\x0b\x72\xac\x8a\x1d\x9e\x34\x6f\xe3\xc4\x73\x69\xe8\x27\x14\x85\xa9\x0f\xff\x13\x82\xbc\xc9\x1f\x42\x92\x2d\x3e\x03\x00\x00\xff\xff\x0f\x1c\xea\x55\x46\x01\x00\x00"),
		},
		"/migrations/000002_create_job_table.down.sql": &vfsgen۰FileInfo{
			name:    "000002_create_job_table.down.sql",
			modTime: time.Date(2020, 12, 19, 15, 17, 27, 530370581, time.UTC),
			content: []byte("\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x6a\x6f\x62\x3b"),
		},
		"/migrations/000002_create_job_table.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000002_create_job_table.up.sql",
			modTime:          time.Date(2021, 1, 1, 10, 57, 24, 779310759, time.UTC),
			uncompressedSize: 822,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x92\x4f\x73\xd3\x30\x10\xc5\xef\xf9\x14\x7b\xb4\x67\x72\x08\x1d\x38\xf5\xa4\x24\x4a\x2b\x70\xe4\x22\xcb\xd0\x70\xd1\x08\x6b\x53\x54\x8a\xe4\xb1\xe4\x84\xe1\xd3\x33\x91\x9b\x3f\x06\x66\x32\x1c\xfd\x9e\x7e\x4f\xda\xf5\x5b\x08\x4a\x24\x05\x49\xe6\x05\x05\xb6\x02\x5e\x4a\xa0\x8f\xac\x92\x15\x3c\xfb\xaf\x90\x4d\x00\xc0\x1a\xa8\x6b\xb6\x84\x07\xc1\xd6\x44\x6c\xe0\x03\xdd\xc0\x92\xae\x48\x5d\x48\xe8\x7b\x6b\xd4\x13\x3a\xec\x74\x44\xb5\x7b\x9b\xe5\xd3\x03\xd2\x76\xfe\x19\x9b\xa8\x8e\xe8\x21\x96\xd7\x45\x01\x82\xae\xa8\xa0\x7c\x41\xab\xe3\x19\xc8\xac\x19\xa0\x1d\x76\xc1\x7a\x07\x8c\x4b\x7a\x47\x45\xd2\x9c\xfe\x81\xf0\x89\x88\xc5\x3d\x11\xd9\xcd\xcd\x2c\x3f\x45\x25\xdb\xef\x1d\x76\x27\xff\xcd\x6c\x36\x24\x85\xa8\xbb\xa8\x8c\x8e\x08\x92\xad\x69\x25\xc9\xfa\x61\x0c\xa2\x33\x7f\xf8\x49\xb6\x2e\x62\xb7\xd3\x2f\xa7\xc8\x77\xaf\x89\x06\x5b\x74\x26\x28\xef\x54\xab\x43\x84\x79\x59\x16\x94\xf0\xe4\x35\x3a\x36\xdf\x54\xdf\x8e\xc4\x01\x40\xd7\x58\x0c\xe3\x01\x88\x10\x64\x33\x9d\x1c\x0e\x45\x1d\xbe\xab\xf1\x88\xc7\x11\x92\xd5\x78\xb7\xb5\x4f\xf0\xbe\x2a\xf9\x3c\xa9\x7b\xeb\x8c\xdf\xab\x60\x7f\x21\xcc\xd9\x1d\xe3\xf2\x52\xf6\xdb\x6d\xc0\xf8\x0f\x23\x76\xbd\x6b\x0e\x3f\x28\xfa\x8b\x65\xe5\xc3\x23\x74\x08\x18\xc3\xf1\x92\x34\x50\x87\x3a\xa2\x51\x3a\x5e\xac\xef\x33\x93\xf7\xe9\x13\xbe\x94\x9c\x8e\xb7\xd9\xb7\xe6\xff\x00\x83\x2f\x78\x05\x18\x9e\x52\x73\xf6\xb1\xa6\x90\x9d\x0b\x35\x4d\x9d\xc8\x27\xf9\xed\xe4\xb5\xbc\x8c\x2f\xe9\xe3\xdf\xe5\x4d\x8b\x55\xd6\xfc\x84\x92\x0f\x65\x4e\xe0\x55\xec\x7c\xd5\x08\x3e\xcb\xf9\xed\xef\x00\x00\x00\xff\xff\xde\xb0\xe3\x06\x36\x03\x00\x00"),
		},
		"/migrations/000003_create_instance_table.down.sql": &vfsgen۰FileInfo{
			name:    "000003_create_instance_table.down.sql",
			modTime: time.Date(2021, 1, 21, 7, 45, 56, 645719521, time.UTC),
			content: []byte("\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x69\x6e\x73\x74\x61\x6e\x63\x65\x3b"),
		},
		"/migrations/000003_create_instance_table.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000003_create_instance_table.up.sql",
			modTime:          time.Date(2021, 1, 28, 6, 21, 54, 899491203, time.UTC),
			uncompressedSize: 349,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\xcc\x31\x4f\xc3\x30\x10\x05\xe0\x3d\xbf\xe2\x8d\x89\xd4\x01\x09\x36\x26\x37\xbd\xa8\x86\xc4\xa9\x6c\x07\x28\x4b\x64\x72\x27\x08\xaa\x02\x6a\x1d\x7e\x3f\x32\x11\x48\xb0\xa0\x8e\xef\xee\x7d\xaf\xb4\xa4\x3c\xc1\xab\x75\x4d\xd0\x15\x4c\xeb\x41\x0f\xda\x79\x87\x71\x3a\xc5\x30\x0d\x82\x3c\x03\x30\x32\xba\x4e\x6f\xb0\xb3\xba\x51\x76\x8f\x5b\xda\x63\x43\x95\xea\x6a\x8f\x79\x1e\xb9\x7f\x96\x49\x8e\x21\x4a\xff\x71\x95\x17\xab\x44\x5e\xdf\x9e\xfa\x6f\x96\x76\x4d\x57\xd7\xb0\x54\x91\x25\x53\x92\x4b\x7f\xe4\x23\x2f\xe5\xd3\xf0\x22\x3c\x1f\x84\xfb\x10\xe1\x75\x43\xce\xab\x66\xf7\xe3\x96\x4e\x0c\x51\x70\xa7\x6c\xb9\x55\x36\xbf\xbc\x28\x56\x59\x3a\x73\x88\x01\x37\xae\x35\xeb\x25\x0f\x47\x09\xf1\xef\xd0\xbd\xf6\xdb\xaf\x88\xc7\xd6\xd0\xef\xdd\xf9\x9d\xcf\x03\x2c\x07\xf9\x07\x64\xc5\x75\xf6\x19\x00\x00\xff\xff\xd7\x6a\xbb\x33\x5d\x01\x00\x00"),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Date(2021, 1, 13, 7, 0, 37, 690273544, time.UTC),
		},
		"/templates/scheduler": &vfsgen۰DirInfo{
			name:    "scheduler",
			modTime: time.Date(2021, 1, 18, 4, 3, 26, 505614050, time.UTC),
		},
		"/templates/scheduler/airflow_1": &vfsgen۰DirInfo{
			name:    "airflow_1",
			modTime: time.Date(2021, 1, 28, 6, 21, 54, 899687720, time.UTC),
		},
		"/templates/scheduler/airflow_1/__lib.py": &vfsgen۰CompressedFileInfo{
			name:             "__lib.py",
			modTime:          time.Date(2021, 1, 21, 7, 55, 56, 458097872, time.UTC),
			uncompressedSize: 16006,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xc4\x7b\xdb\x8e\xe3\x38\x92\xe8\xbb\xbf\x22\xda\xfd\x60\xbb\x8e\x4b\xa7\xba\xd1\x98\x07\x63\xd5\x98\xdc\xaa\x9c\x9e\x42\x5f\x51\x99\x33\xbd\x83\xda\x84\x9a\x96\x42\x36\xc7\x32\xa9\x21\xa9\x74\xba\x13\x09\xec\xb7\xec\xa7\xed\x97\x2c\x82\x17\x49\x94\xe4\xec\xac\x9e\x05\x26\x1f\x9c\x96\xc8\xb8\x30\x6e\x8c\x08\xd2\xfc\x58\x4b\x65\x40\xea\x99\xff\xf6\x77\x2d\x45\xf8\x9e\xb3\x0a\x45\xc1\x54\x78\x56\x38\x2b\x95\x3c\x82\x39\xd7\x5c\xec\xc0\xbf\xbd\x12\xe7\x35\xbc\x65\x55\xc5\xb6\x15\xae\xe1\x1d\xcf\xcd\x1a\x7e\xac\x0d\x97\x82\x55\x0e\xa0\x60\x06\x0d\x3f\x62\x00\x09\xcf\x6b\xa0\xcf\x02\x2b\xc3\xdc\xd7\x5f\xa5\xf0\x34\xb4\x51\x3d\x1a\xb7\x78\xac\x2b\x66\xfc\x58\xae\xa4\xe0\x06\x55\x18\x0d\xcf\x33\x37\xcc\xb8\x2a\x2b\x79\x4a\x0e\xcd\x16\x95\x40\x83\x3a\x4c\xa4\x37\x59\x5e\x71\x14\x66\x0d\xb5\x2c\xb2\x8a\x35\x22\xdf\xa3\x8a\x01\x8f\xb2\xc0\xaa\x05\x7a\x77\xf5\xcd\x1a\xfe\xca\x14\xf7\xeb\x63\xbb\x0f\x8d\xb0\xff\xbf\xa7\x79\x6b\xb8\x65\xfa\xf0\x5e\x68\xc3\x44\x8e\x6b\xf8\x77\xa6\xf1\xc7\x1a\x15\x33\x52\xad\xe1\x3f\xde\xca\x23\x7d\xfe\xf8\x7d\xf6\xe1\xfa\xf6\x2f\x1f\x7e\xc8\xbe\xbd\xfe\x5b\x4c\x6d\x2f\xe5\x41\x27\x7b\x63\xea\x8c\xbe\x06\xb2\x7f\x36\xa6\xfe\xb3\x94\x87\x78\xb2\xf4\x98\x75\x52\x9f\xcd\x5e\x8a\x2c\xbc\x08\x60\x3f\xd9\xd7\x81\x81\x18\x58\xa3\xd0\x04\xba\x65\x1a\x33\xf7\x30\x82\x27\xf6\x6f\xec\xd0\x34\x8e\x5c\x0a\xa3\xf8\xb6\xc7\x48\x27\xe6\x8c\x44\x3a\x44\xf8\x6d\x3b\xfc\x93\x2c\xa6\x71\x76\x18\x12\x8d\xb9\x42\x13\x60\x6f\xec\xd3\xf4\x22\xf0\xc1\xa0\x12\xac\xca\x0c\xd3\x07\xbf\x9a\x00\x77\xed\xc7\x48\x33\x6e\x31\x31\x0e\x7c\xc8\xd1\x1a\x68\xab\xe4\x2b\x37\x72\x1d\x06\xe2\xf9\x8d\xe1\x95\x4e\x0a\xcc\xa5\x5b\x73\x80\x62\x75\x5d\x9d\xb3\x02\x4b\xd6\x54\x46\x4f\xc2\x6c\xc3\xdc\x5a\xc9\x7b\x5e\x90\xdc\xb5\x1e\x11\xc8\xa5\x28\xf9\xae\x51\x8c\x68\xb7\x66\x2d\x45\x39\x85\x53\x1b\x66\x5a\x57\xba\x31\xad\x5b\xc4\x06\x65\x95\xdc\x37\x28\xd2\xac\x35\xa8\xd9\x2c\xaf\x98\xd6\x70\xd3\xd4\xa8\x26\xd5\xb3\x9c\x7c\xbb\xda\xcc\x00\x00\xe6\xf3\xb9\xfd\xff\xea\x15\xdc\x5c\x7d\x7f\x0d\x57\x37\xd3\x3a\xde\xc0\xf5\x03\xe6\x8d\x41\x60\x40\x3a\x02\x2e\x80\xf5\xa6\xc2\x4f\xb2\x80\x57\xaf\x2c\xae\x9f\x15\xab\x6b\x54\x60\x24\xd4\x8d\xde\xc3\x43\x4e\xeb\xd1\xc0\x40\xa1\x69\x94\x80\x7b\x56\x35\x08\x07\x3c\x03\xde\xa3\x00\x5e\x92\x74\x0c\xe3\x02\x15\xe4\xf2\x58\x57\x16\xe3\x89\x9b\x3d\x08\x29\x40\x37\x79\x8e\x5a\x03\x89\xaa\xd1\x33\x4b\x23\x49\x40\x48\x83\x1b\x38\x20\xd6\x60\xf6\x5c\x43\x53\x13\xc5\xc2\x8a\xb3\x04\xb3\x47\x85\xc0\x35\x30\x71\x86\x7c\xcf\xc4\x0e\x89\xe7\xc9\xc5\x01\xfa\xb5\x1d\xd1\xec\x65\x11\x09\xc6\xf8\x60\x95\x95\x1c\xab\x42\x43\x0a\xcb\x05\x3f\xb2\x1d\x2e\xd6\xb0\xc8\x8f\x85\xa6\xff\x4c\xed\x9a\x23\x0a\x63\x1f\x50\xdc\x67\xf7\x4c\xd9\xef\xce\x14\xb2\x92\x57\x76\x3e\xb9\x54\x0f\x61\x85\x8b\x95\x5b\xce\x1f\x07\xc6\x47\xef\x0a\x2c\x21\xcb\xb8\xe0\x26\xcb\x96\x1a\xab\x72\x6d\x5f\xf7\xff\x5e\x31\xb5\xd3\x13\xaf\x5f\x1d\x4e\x34\xe2\x75\x4c\x7f\x9a\xcc\x63\x79\xd9\x48\xd6\x40\x14\x56\x49\x4b\xd0\xa1\xee\x50\xcd\x3a\x54\x58\x95\x49\x21\x33\x52\x6b\x66\x15\x9c\x82\x9b\x94\xec\xd0\x2c\x17\xfd\xa1\xc5\x2a\x06\x13\xec\x88\xba\x66\x39\x0e\x60\xda\xf7\x43\x00\x2e\xb2\xbc\x6a\x34\xed\x0f\x31\x44\x37\x30\x04\xf1\xaf\x33\xb2\x29\x7c\x30\x03\xb8\xc1\xe8\x10\x58\x21\x33\x86\xe5\xfb\x4c\x8a\x4c\xa1\x36\x4c\x0d\x11\x4c\xcc\x18\x71\xd0\x69\x7d\x48\xbd\x67\x0f\x5e\xa4\xa4\x66\x6f\x7f\x4e\xcb\xe0\x59\xeb\x69\xcf\xa8\xf3\x26\xd2\x32\x2f\x47\xe2\xe1\x9a\x3c\x02\x7e\x90\x02\x37\x23\x8b\x70\xdb\x24\x31\xd3\x6d\x9a\xc4\x51\xd6\x7b\x5e\x76\xc8\xd2\x01\xf2\xb1\x89\xbd\xe8\x6f\x20\xeb\x74\x4a\x3d\xbf\x17\x75\x27\xc8\x74\x28\xf3\x55\x84\x12\x2b\xfd\x7b\x05\xf2\x2f\xe3\x3f\xc2\xf9\x39\x5c\x15\x05\xc5\xc5\x2d\x17\x6e\x4b\x91\x25\x54\x6c\x4b\x59\x8d\x91\xd0\x08\xfe\x8f\x06\xab\x33\xf0\x02\x85\xe1\xe5\x99\xe2\x6c\x23\x04\x25\x5c\xb5\x0f\x66\xe1\xcf\x43\xa5\xde\x4c\xc9\x94\x31\x73\x2f\xb3\x52\x2a\xda\xf2\x97\xc1\xfa\x66\x63\xc8\x4c\x63\x85\x39\x05\x4c\x8f\x21\x23\x89\x51\x50\x0b\xb4\xb9\xd8\x65\x7e\xaa\x4d\xf9\x96\x0e\xf9\x00\x99\x4d\xd6\xb8\x26\xf1\x7b\xc9\xd3\x53\xd6\x06\x81\xc2\x32\x12\x07\x8c\xf5\x80\x85\x34\x7e\x1c\x50\xe0\x25\x54\x28\x96\x81\x50\xc2\x0d\x1e\xf5\x0a\xbe\x86\x2f\x80\x89\xe2\xa2\xab\x8f\xed\x44\x31\xae\x71\x94\x50\x2c\x27\x95\xbe\xf8\x5e\x2a\x04\xb3\x67\x02\xa4\x40\x5a\x65\xab\x08\xbb\x97\x39\x59\x6c\x60\x31\x0d\xfd\x18\xaf\xe8\x69\x91\x94\x52\x1d\x99\x59\x3e\xbf\xf0\x91\xa2\x5c\x12\x0c\x69\x94\x13\x27\x3f\xc9\xe2\x3b\xff\x7d\xd9\xb3\xf1\x34\x24\xd0\xf8\x60\x14\xcb\x8d\x0d\xdd\xe9\x28\xc4\xbf\x4c\xbc\x69\x0a\x5f\x8c\x45\x68\xd4\x39\x13\xcd\x71\x8b\x4a\x67\x47\x66\xf2\x7d\x6b\x3d\xa3\x91\x60\x7c\x6b\x88\x51\x7f\x7c\x73\xb7\x1a\xe1\x2d\x39\x25\x8c\x36\x7f\x5a\x83\x42\xdd\x54\x26\x60\xde\x33\x51\x54\xe8\x32\xd8\x7b\x54\x15\xab\xbd\x21\xae\xc7\xdc\xac\x5b\x91\x75\x54\x5f\x12\x41\x22\xf2\xd9\x90\x03\xef\x5b\x02\x4f\x96\x8b\xb2\x97\x9c\xb7\xac\x04\xba\x63\xe1\xf6\x70\xc3\x67\xa9\x4b\x0b\x93\x9b\xbf\xbc\x7d\x7b\x7d\x73\x33\xe6\xe4\xf3\x5e\x9a\xe5\x52\xab\x90\x56\x91\x09\x96\x8c\x57\x7a\x36\x11\x7f\x48\xd0\x6e\x57\xb2\x49\x37\xf7\x45\xcf\x62\x95\xb4\x6a\x5f\x1e\xf0\x9c\x0e\xea\x9d\xb5\xa3\x91\xba\xf5\xae\xfe\x39\x8f\xa1\xac\xd1\x25\x85\x58\x00\xb3\xbc\x36\x0a\x37\xf0\x68\xd7\xde\xb9\x80\x7d\x4c\x7b\x62\x59\xc5\x84\x7d\x62\xe9\x78\x6a\x47\x5c\x65\x30\xe2\x85\x72\x51\x7c\x88\xe5\x78\x81\x6b\xcb\xa0\xf3\x1b\x72\x63\xe2\x0f\x8b\x0d\x3c\xa2\x52\x7d\x0f\xb5\x8f\x29\x3e\x90\x33\xf6\xf3\xf1\x71\xe9\xb2\x1c\x97\x64\x83\x4c\xfc\x67\xc6\x8d\x86\x52\x2a\x60\x50\xf0\xb2\x44\x45\x1b\xd5\xbb\xab\x6f\xc0\xbe\xea\x92\xef\x78\xd0\xc8\x36\x75\x76\xc0\x16\x99\xae\x31\xe7\x25\xcf\x7d\x86\x41\x8b\x3f\x71\x51\xc8\x93\x33\xb9\x4d\xcd\x14\x3b\x42\x5b\x7d\x15\x6c\x97\xf1\x62\x03\xb7\x7b\x04\xf7\x9d\xc2\x99\x09\xe9\xb9\xa6\xac\xda\x71\x70\x96\x0d\x9c\x98\x30\x60\x64\x2b\xc7\x13\xe3\x86\x68\x3b\xd4\xe6\x5c\xe3\x18\xb3\x36\xaa\x4f\x99\x55\x95\x3c\x61\xe1\x74\xaa\x37\x60\x77\x05\x59\x86\xf7\xe0\xde\xaf\xc1\x27\xc6\x94\xe4\xfc\xf2\xcb\xc7\x85\xaf\x08\x16\x77\xbf\xfc\xd2\x23\x36\x85\xac\x4f\xcc\xad\x3c\xd3\xfc\x57\xdc\x00\x7d\x12\x25\x5a\x91\x1b\x20\xa9\xee\x65\xa3\xec\xae\x5a\x51\xad\x45\x72\xf4\xa4\xca\xa6\x82\x76\xa1\xaa\xa1\x5a\x53\x40\x53\x6b\xa3\x90\x1d\x49\x56\x09\x5c\x27\xbb\x35\xcc\xbf\xfc\x6a\x0e\x27\x5e\x55\x90\xef\x31\x77\x28\x2a\xa6\x0d\x7c\xf9\x95\x47\x4e\xd5\x5d\x8b\x29\x6f\x94\xd5\x60\xa7\x1e\x5b\xbf\x58\xbe\xb8\x76\x78\xdf\x1b\x87\xcb\x19\x85\x8b\x5d\x34\xa3\xe3\xac\x45\xc7\x0d\xba\xaa\x53\xd3\x84\x3e\x7b\xc4\xae\x2f\x5a\x0b\xbf\xde\x04\xae\xaa\xca\x8b\xe0\x08\x02\xb1\xb0\x2b\xdf\x62\xaf\x68\x68\xd7\x4e\xa4\x2d\x4b\xbe\x36\xef\x99\x5b\x02\xef\x7c\xd9\x42\x6f\x19\x14\xec\x4c\x58\x1d\x91\xe5\x97\x5f\xad\x7a\x2a\x8a\x54\xc0\x85\x69\x2d\xff\x13\xaa\x20\x18\x85\x92\x81\x95\x8d\x53\xb2\x21\xd9\x8b\x33\x64\x59\x6a\x34\xcf\xcf\x31\xaa\x11\x39\x05\xf7\xa6\x36\xd2\x5a\xf4\xa5\xc2\x0c\x9e\xab\xcc\x66\x5d\xf8\x76\xf1\x60\xa1\x61\xcf\xee\x11\xcc\x49\xc2\x51\x16\xd6\x0a\xdc\x9e\x41\x0a\xdd\x50\xf5\x78\xc0\x85\xcd\x5f\x16\x0a\x75\xbe\xc7\xa2\xa9\x70\x91\xf8\x81\x1e\x3e\x0b\xcd\x35\x54\xfc\x80\x84\x93\x42\x17\x83\xd3\x9e\xea\x91\x4a\xca\x3a\x81\xd3\x1e\x85\x35\xfd\x80\x46\x39\xa3\x0e\xfe\xbd\xa6\x6f\x3d\x8c\x5e\xeb\x54\x6c\x6b\x67\x8d\x36\x1c\x4a\x05\xb5\xc2\x82\x93\x38\x28\x48\x14\xdc\xda\x70\x23\x0c\xaf\x80\x1b\xd8\x62\x2e\x8f\xa8\xc1\xa8\x06\x13\xb8\xdd\x73\xdd\xc3\xb9\x67\x8e\x1e\x96\x25\xe6\xc6\x05\x1a\x29\x72\x04\x16\xc8\xd9\x9c\x4c\xaf\x09\x93\xa3\x6c\x98\xa5\xab\x50\xcb\x46\xe5\xa8\x3d\xa5\x1e\x52\x6e\x2c\x30\xd1\x24\x7c\x84\xbf\xe3\x90\x28\x6e\x11\x05\xd5\xf9\x5e\x06\x1a\x29\x84\x45\x02\x25\x1c\xf8\xc0\x4d\x9f\x57\x7e\x3c\x62\xc1\x99\xb1\x79\x76\x39\x40\xcb\x35\x94\xac\xd2\x68\x55\x43\x2e\xe2\x31\x15\xc0\x0c\x30\xa0\x62\x5f\xd9\x86\x68\x12\x49\x14\x2d\x9a\x42\xe6\xb6\x75\xe0\x32\x7b\x92\xe8\x78\x87\xb0\xaf\x8f\x94\x5c\x72\xe1\x76\x1c\x2e\x45\x8b\xcb\x59\xd4\xc7\x05\x69\x7d\x71\x37\x28\x38\xed\xcb\x75\xb4\xbe\x6e\xe7\x8c\xab\xd6\x10\x2e\xc8\x8b\x20\x1d\x3a\x55\x3c\xb7\xe7\x50\x90\xf6\xdd\x6b\x72\x9a\xf3\xaa\x6e\xa2\x7b\x9e\x9c\x1a\x39\x57\x07\x11\xbd\x8e\x01\x43\xd4\x6f\xd9\x77\xe1\x1f\x52\xf8\x18\xe5\x4d\x77\xb3\xa9\x66\xc8\x78\x87\x7e\x61\x27\xe4\x8f\xc3\x16\x60\x08\x56\xe4\x8c\x71\x21\x4f\x18\xed\x9c\x94\x8a\xf3\xbe\xeb\x93\x68\x8d\xcc\xec\xc6\x99\x86\x59\xc9\x3f\x1a\x54\xe7\x65\x68\x49\xaf\x92\x92\x57\x06\x55\x9c\x41\x85\xd1\xc4\x6f\xd3\x69\x3a\x56\x62\x0b\x40\x28\x94\x36\xcb\x55\x3f\xe8\xb8\xfd\x89\x97\x94\xce\xf1\x22\xda\x2c\x66\xbd\x54\x54\x48\xd3\x67\xf3\x65\x59\x13\xe5\x0f\xc1\x7e\x6c\x76\xb2\x78\x51\x95\xbc\x78\x7c\x82\x42\xa2\x6b\x64\xe0\x03\xa5\xff\x5d\x06\x38\x5c\x5d\x2f\x03\x1c\x27\xe8\x9e\x71\xa9\x93\x9a\x99\x7d\x62\x71\xe9\x65\x6f\x1d\x24\x54\xac\x64\xbe\x7a\x71\xc9\xf7\x7b\x17\x15\x16\x76\x62\x1a\x0a\xa4\x0d\xb3\x78\x7e\x55\x7d\x1d\xb1\x2a\x6f\x28\x7e\x78\x3f\xd0\xbd\xd4\xd6\x27\x0c\x99\x4d\x18\xd2\x60\x6b\x1f\x17\xf1\xc8\xe2\x6e\x36\xdc\x04\x29\xa8\xae\xc3\x13\x8a\x22\x94\x2c\x3b\x14\x14\x6e\x30\xf3\x3b\x77\x8c\x68\x3d\x72\xfc\xf5\x84\x8f\xaf\x2f\x3b\xf3\xa0\x47\x56\xc9\x5d\x42\xd1\x6c\x39\xcf\xa5\xd0\xcd\x91\xe2\x7a\x6b\x83\x3e\x21\xdb\xa2\x39\x21\x8a\x0d\x3c\x3e\xc1\x6b\x78\x7c\x9a\x07\xb1\xf5\x97\x92\x70\x2d\xfd\xeb\x55\x7f\x59\xfd\xf7\xb1\x54\x4b\x2e\x8a\xb6\xa9\xdc\x4b\x99\x4e\x68\xb3\x20\x4a\x95\x7c\x92\xdc\x49\xbb\xc6\xdc\xf4\x83\x4c\x2b\x1c\x1b\x68\xee\x22\x8f\xce\x95\x14\xa4\x11\x7f\x90\x15\x99\x5d\x88\xc2\x19\x17\x06\xd5\x3d\xab\xd6\x91\x5e\x12\x85\x75\xc5\x72\x5c\x9a\x5f\x49\x38\x2e\x60\x74\x92\x73\x5b\xf8\xad\x6a\x06\xe6\x2e\xf0\xc1\x64\xaa\x21\xaa\x81\x01\xdb\xcc\xa2\xf7\xcb\x70\x3e\xb7\x1a\x79\x48\x80\xfa\xba\x2f\xb6\x29\x0e\xc6\x4e\xb2\x55\xc8\x0e\x71\x95\xfc\x8c\x8c\x12\x56\xd7\x28\x8a\x65\xa0\x78\xd1\x16\x02\x92\xce\x14\x7a\x82\x5e\x3e\x3e\xad\x36\x7d\x33\xa8\x50\x2c\x9f\x23\xbb\x5a\x3f\xcb\x55\x6c\x15\x51\xb6\x6c\x93\x21\x6f\x7e\xc0\x45\xdd\x98\x60\x93\xb6\x93\xd3\x3f\x91\xe8\xb6\x3c\x96\x9b\x86\x55\x97\x6c\x04\x54\x32\x70\x5a\xda\xd5\x15\x59\xdb\x28\xf0\x7f\x68\xc4\x60\x72\xac\xbc\xc9\x2d\xc1\x6f\x0b\x04\xfb\xcc\xa6\xb0\xbe\x04\x33\x60\xee\xeb\x38\x5e\xbc\x10\xea\xdf\xd2\x9e\x29\x3d\x03\x14\x4e\xfb\x14\xdf\xed\x50\x11\xa3\x7f\xa2\x0c\xea\x22\x84\x15\x74\xc2\x85\x2b\x01\x2e\xed\xf9\xb1\x90\x56\x89\x54\x05\xaa\x6c\x7b\x49\xa4\x84\x66\xb9\x82\xbb\x4b\xc6\xe8\xf4\xf9\x72\x53\xbc\xac\xff\xd5\xfa\x19\xe3\xe8\x9b\xe1\x91\x6b\xcd\xc5\xee\x82\x0d\x69\x34\xcf\xdb\x3b\xbc\xb6\x73\x9e\xa1\x35\x1b\xb4\xf0\x9e\x21\xb8\x82\xaf\xe1\x4d\xec\xf9\x03\x01\x79\xe0\x49\x09\x45\xd2\x79\x8e\xca\x6a\x9a\xc2\x89\x29\xc1\xc5\x6e\x6c\xe2\xf3\x46\xb0\x6d\x85\x94\xb5\xdb\x40\x8e\x42\x36\xbb\xbd\x37\x14\x08\x5d\x2c\x57\x2c\xb7\x8c\x2d\x1e\x9f\x16\xb6\xb4\x2e\x5a\xaf\x7e\x7c\xb2\xf9\xfa\xe3\xd3\x92\x8b\xbc\x6a\x34\xbf\xc7\x95\xed\xe2\xb9\x08\x4d\xeb\x72\x65\xc8\x7c\x72\x9b\xbf\xb8\x83\x0f\x02\xfa\xcb\x76\xa7\x89\x5e\x96\x75\x88\xce\x32\xfc\x5b\x8a\xfd\xdd\xb9\xd1\x70\xcb\x76\x69\x27\x8a\x22\x73\x17\x32\xa2\xdd\x7a\xb0\x51\x4f\xee\xd1\x9d\xba\xcb\x4a\x32\x43\x3a\x73\x19\x42\xc0\xd9\x8f\x98\xb6\x5a\x07\x8f\xa0\x5f\x91\xf0\x12\xa6\xd3\xf9\x14\xe6\xa7\xf9\x66\x70\xba\xa1\xf0\x28\xa9\xec\xe5\x47\x04\x3b\xad\x60\x67\x6d\x75\x63\xf5\x2b\x90\x29\xd4\x06\x4e\x88\x87\x01\xe4\x0e\xdd\x6b\xdb\xe9\x71\x1a\x0f\x0d\x95\xa3\x14\x66\x1f\x4d\xa7\x89\xd9\x91\x19\xc5\x1f\xb2\x1a\x55\x66\x67\xd0\x36\xed\xef\xc5\x24\x6f\xfd\x97\xe5\x2a\xb1\x63\x64\x2e\x3a\x8c\x2e\x83\x00\x92\x33\x32\xd5\xc9\xd8\x4d\x5d\x0d\x18\xb3\x8c\x5b\xce\x4e\xf6\x18\x3a\x70\x55\xb0\x33\x54\x1c\xe3\x76\xac\x1f\xcc\xec\xfc\xd4\x1e\xe0\x45\xe3\xb4\x2c\x3b\x46\x79\xc9\xd4\x22\x26\x7a\xd3\x52\x59\x5a\x1e\x62\x33\x69\xc1\xbc\xb4\x73\xd2\x4e\xb9\x09\x2d\x79\xb9\xda\x5c\xcc\x6b\x07\xac\x5a\x95\xc4\xbc\xc6\x56\x13\x52\x8f\xc4\x9d\x5e\xe1\xb2\x8f\xe0\xe3\x1f\xee\x7a\x72\xa4\x0f\xf3\xeb\x72\xe0\x0b\x03\x7c\xfd\xc7\x61\xa2\x12\x6e\x1c\x25\x8d\xc9\x7b\x11\x75\x0c\x08\x9f\xff\xbf\xee\xa6\xd2\x52\x63\x2e\x45\xa1\xd3\xb8\x54\x7d\x05\x7f\x78\x63\x3f\x7a\x99\x8a\x3f\x0b\x26\x0c\xaf\x2f\x23\xb0\x25\xf1\x04\xb8\xf7\x5f\x9f\x7d\xa3\x28\xba\x4e\x71\xc5\xf2\xc3\xcf\xb8\xdd\x4b\x79\x68\xef\x6c\xf4\x2f\x1e\x0d\x1a\xc4\xb7\x7b\xae\xa1\xbd\x98\x63\xb7\x42\x6d\xfb\xb1\x46\x42\x2d\xb5\x81\x23\x6a\xcd\x76\x68\xdb\x71\x16\x39\x34\x36\x4c\x73\x91\x4b\x9b\x67\x9f\x1c\x31\xed\xfa\x11\xb7\xec\x80\x1a\xb6\xd2\xec\xfd\x6c\x3f\x0c\x46\x1e\x50\x40\xc1\x15\xe6\xa6\x3a\x5b\x8f\xcc\xa5\x10\x98\xdb\x56\x85\xed\xaf\xec\x99\x9e\x02\x72\x88\xdf\x97\x0e\xab\x6e\xea\xba\xe2\x58\xac\xc1\xde\x8b\x22\x1c\x94\x9d\xd8\x2e\xe9\x16\xa1\xd1\x58\x00\xd3\x60\xef\xb8\x34\xaa\x72\x49\x00\xb3\x1e\x64\x71\x66\x8e\x91\x30\xdf\x30\x7a\x62\x9a\x84\x58\x4b\x2e\x8c\x6d\x56\x81\xc2\x8a\x19\x7e\x8f\x40\x15\x5f\xe8\xec\x36\xaa\x72\xbc\x5c\xb3\x7c\x7a\x79\x39\x13\x84\xb4\x56\xf8\x3a\xdc\xdb\xc1\xc2\x9e\xab\x6a\xdb\x89\x0a\x0d\xf4\x7c\xcf\x84\xc0\x6a\x4d\xef\x95\x60\x47\xdb\xed\xb1\xa8\x79\x2e\x45\x02\x7f\x93\x8d\xc5\x25\xef\x51\x29\x5e\xd8\xe6\x8e\xc6\xd0\xb8\xb6\xed\x62\xdb\x3f\x25\xd2\x49\xbf\x2b\xdd\x97\xc9\xe6\xa5\x02\x76\xe8\xd0\x9d\xda\x81\xbd\x9a\xd2\xeb\xb2\xc6\x28\x07\x2d\xf7\x48\xa8\x9b\x29\xe4\xfd\x7e\x6d\x3c\x79\x80\xca\x1b\x9a\x3b\x2f\xf0\x0f\xfd\x93\x01\xda\x40\x0b\x90\xc2\x11\xe9\xa1\x6d\x01\x07\x08\xb7\x95\xcc\x0f\xda\xe1\x73\xdf\x47\x58\x12\xb8\xd9\xcb\xa6\xa2\x7d\x1c\x58\x38\x2f\xe8\xca\x2f\x6e\x65\xc7\x14\x47\x0d\x0a\x6b\x85\x1a\x05\x39\xbe\x5f\xa7\x43\x9a\xf4\x58\x09\x24\xdb\xc3\x82\xb6\x1b\x3d\xbe\xf9\xf3\x71\x11\xc9\x63\xb1\x86\x85\x5f\x09\x7d\x75\x98\x16\x77\xff\xd4\x8d\x9e\x48\x77\xb6\xf8\x9a\x98\x13\x71\x71\x69\x92\xe7\x2c\x9d\xcf\x27\x06\x1d\xaf\x97\x40\x2f\x5c\x2b\x7a\xe6\x5e\xd1\x44\x08\xfb\xd4\x1b\x45\x51\x6c\x48\x23\x39\x0c\x3a\x85\x51\x54\xe8\xdf\x44\xb0\x6f\x96\xd1\x78\x1c\x73\x06\x75\x67\x30\xd9\x34\x08\x2b\x1e\xf6\x16\x98\x7a\x69\x0d\xd8\x25\x77\xf1\xbb\x75\xa7\xdc\x8e\x0b\x97\x8a\x4d\xf1\xd0\x09\x2f\x44\x74\xfa\xfb\x86\xdf\xa3\x00\xe4\x66\x8f\x0a\x18\x1c\x99\x68\x58\x55\x9d\x7d\x83\x9a\x56\x6a\x8f\xff\x3c\x92\x75\xd8\x4f\xec\xf1\x55\x24\x10\x17\xba\xba\x56\xb3\x77\x2d\xef\xc1\xd6\x53\x03\xee\xf6\x40\xa8\x73\xfb\xce\x31\x06\x2e\x7f\x31\x62\x11\xc6\xa0\xb6\x80\x70\x80\x6a\x3a\x1e\xd9\x71\xb7\x8c\xcd\xe4\x1a\x7a\xb3\x08\x4d\x0c\xd9\x97\x1d\x2f\x3d\xb7\x53\x79\x74\xbc\x36\xac\x78\x19\xf3\x33\x1b\x1c\x8b\x8b\xae\x1f\x66\xb2\x2e\x1c\x2f\xa7\xed\xc8\x1f\x3f\x29\xe6\x5a\x70\xb6\xb6\x55\x2c\x2b\xd0\x5e\xbd\x9e\x60\xc7\x8e\xbb\xde\xfc\x28\x98\x2c\x9e\x6b\x69\x5e\x6a\x4b\xbe\x65\x42\x48\x63\x73\x61\xaf\xb2\x1f\xa4\xef\xe9\xba\x98\xd7\xad\xa1\x7f\xe5\x2c\xdb\x36\xbc\x2a\x32\x4d\x53\x32\x6f\xfe\xd6\x68\x2f\xd8\xe7\x5b\x29\xb4\x51\x4d\xee\xce\x54\x1c\x66\x0f\xe6\x0e\x12\x15\x56\x78\x4f\x71\xdf\x1a\x09\x1a\x54\x1a\x18\xe5\xbe\x2e\xf5\x2b\xc0\x66\xc2\xf6\x84\xd0\x72\xd7\x62\x76\xb8\x48\x5e\x2d\xc2\xb1\x7d\x44\x04\xc3\xa6\xf0\x62\x03\xc9\x8f\x45\xfa\xf8\x34\x1b\xde\x9e\xf3\xb1\x3f\xb6\x80\x63\xf1\xb1\x0d\xe5\x69\x6f\xda\x2c\x9a\x61\xef\x0e\xfa\xf1\x61\xf0\xf0\x9a\xa6\x15\x25\x45\x73\xac\xf5\x32\x3f\x16\x9f\x70\xdb\x2f\xd2\x89\x23\x31\xa9\xad\x8b\x5d\x34\x92\x8d\x3d\xf3\xab\x50\xb9\x7d\x98\xe0\xe6\x97\xe6\x47\x58\x47\x61\x59\xca\x43\x1a\xae\xaf\x47\x4e\x90\x8e\xc2\x76\x3f\xdb\xd5\xb5\x14\xda\x73\x6f\x73\x1e\xd5\x0c\x2e\x81\x8c\x83\x79\xbc\xe5\x14\xcc\xb0\x34\xe2\x2d\x1e\xdf\x23\x2b\x50\xe9\xf4\x71\xf1\x96\x44\x28\xcc\x6b\x32\x82\xc5\x06\x16\xb4\xf5\x72\x57\x93\xfe\x7f\xd2\xc2\xa2\x53\x7d\x6f\x79\xbc\x6c\xd9\x4c\xdc\xd5\xe2\x2c\x97\x05\x52\x41\xf4\xe5\x9b\x37\x9b\x0b\xf7\x4b\xdc\x7c\xd2\xd8\xec\x37\x3c\x73\xee\x6e\x8b\xb4\x19\x8c\x5d\x8a\xd3\x49\xd4\x1c\x89\x90\xda\x1b\x24\x64\x24\x76\x5e\xe6\x50\x64\x46\x3a\xd5\x2f\x63\x63\xb9\xf9\xee\xea\xed\xb7\xd9\xdb\x1f\x7f\xf8\x21\x7b\xff\x0e\x52\x98\x3b\x69\x59\x50\x5f\x2b\x5c\xdd\x7c\xfb\xa7\xab\xf7\xdf\x65\x57\xdf\x5d\x7f\xb8\x85\x14\xb8\x30\xcb\xf0\xf3\x07\x1b\x85\xe6\x86\xe9\x03\x91\xf1\x60\xed\x5d\x8b\xec\x9e\xa9\xf4\x8b\x55\x3f\x64\xd8\x7b\x41\xf6\xf6\x4f\xb6\x67\x3a\xb3\xb7\x5e\x96\xf6\xed\x0a\x5e\x7f\x0d\x5b\x29\xab\xcd\xd0\x0b\xec\x70\x72\x40\x5b\x68\x0e\xae\x12\xf9\x03\xd3\xd0\xb3\x71\xa8\x12\x4b\x60\x6d\x53\xb9\x95\x3b\xed\xb6\x84\x16\x94\xf3\xf6\x66\xd8\xa1\xde\xf3\x47\x3f\xed\x0e\x3e\xeb\x6f\xcb\xbc\x1c\x0a\xe1\xb3\xe8\x96\x9a\x67\x73\x4e\x45\x8a\x42\x4d\xb5\x88\xbf\x85\x04\x5e\x8e\xb3\xce\x27\x43\xca\x11\xee\x86\xf7\xaf\xe5\xc6\x33\xc2\xfd\xfc\xe1\x2e\x12\xa9\x6c\x95\xd4\x4c\xeb\x93\x54\x2e\x92\xb9\xeb\x4a\x1d\xc6\x5a\x91\xb6\xe6\xc2\xfb\x6f\xbf\x36\xb8\xf7\x2a\xa4\xf4\x60\x3e\x2a\x32\xe7\x8f\xdd\xd4\xa7\x49\x30\x7b\xaa\x87\x25\xc5\xe5\x35\x74\x0d\x35\x6b\xa7\x76\xd9\xba\x35\xcf\x0e\x3c\x8d\x99\x9f\xb5\x47\xc8\xa1\xaa\x0f\x6d\xe7\xe7\x6f\x97\xf5\xce\x93\x03\xa0\x9b\xf2\x9b\x90\x7e\x5a\x04\x7a\xe9\x20\xcc\x61\x18\x1c\x86\x79\x63\xfe\xbc\x55\x71\xd8\x50\xea\x46\xef\xb1\x80\xed\x19\x82\xd3\x32\xed\x43\xbe\x9f\x9a\xb5\x55\x75\x7b\xde\x53\x4a\xe5\x2e\xdd\x71\x61\x7f\x1a\x64\x75\x7d\x64\xe2\x3c\x6e\x5a\x4e\x73\x3b\x4e\xb1\x0f\x78\xbe\x90\x94\xfb\xb5\x5f\xca\xd9\x9d\x50\x75\x1a\xab\x62\x3c\xcf\xf6\x3a\x0b\xcc\x6a\xc5\xa5\xb2\x4c\xe8\xf4\x42\xeb\xbd\xe2\x47\x6e\xd2\x2f\xde\xf4\x76\x26\x5e\x42\xdf\x9d\xbd\xc7\xc5\xa1\x72\x28\xaf\x70\xf6\xd3\x79\x6a\xd4\x84\x9e\x0e\x2a\x36\xa6\x7c\x2a\xde\x36\x02\xac\xa6\x14\x47\x6e\xbb\x86\x79\xf2\x77\xc9\xc5\x72\x88\xcc\x41\x78\x87\xf3\x83\x3a\x8a\xd2\x03\x88\x10\x17\xfd\x63\xb6\x95\xc5\x99\x28\xfc\xa7\xf0\x14\x3e\x76\x89\xc8\xff\xfc\xd7\x7f\xc3\xab\x77\x57\xdf\xbc\x8a\x10\xc6\x9a\x5a\xad\x07\xf3\x6f\x99\x3e\x4c\x03\x78\x4b\x18\x41\x5c\xb7\x77\xc8\x6e\xf9\x11\xa7\x61\x07\x27\x20\x43\x14\x1f\x1a\x01\xef\xdf\x0d\x40\xfb\xfe\xa4\x1a\xda\xe7\x17\xbe\x4b\x77\x37\x10\x42\x29\xa5\xfb\xa1\xc6\xa4\x18\x36\xdb\x4a\x6e\x5f\x97\x2c\xc7\x9a\x55\xc7\x0d\xfc\x78\x12\xa8\x2e\x93\x2a\xd8\x6e\xb1\x4a\x24\x4d\xea\xf3\xb9\xd9\xcb\x46\xed\x2a\xa6\xf5\x06\xde\xf9\xdf\x58\xd9\x03\x62\x8d\xf9\x34\xa2\x51\xfc\xf1\x50\x11\xd6\x23\x1e\xe5\x06\xde\xa1\x61\xbc\x7a\x5e\xf1\xd1\xca\xdb\x12\xb1\x5b\xe7\x63\x64\xb5\x73\xca\x49\xe6\x1b\x98\x6b\x17\x43\x07\xc5\xf8\x9c\xd8\x9c\x6f\x06\x40\x11\xe0\x51\x1d\x8a\x93\x98\x28\xe2\x03\xec\x9c\x4c\x25\xc4\xad\x4d\xc9\x15\x6e\xe2\x23\x8b\x2e\xfb\x79\x5a\xff\xeb\xd8\xec\x7b\xca\xef\x66\x8f\x59\xee\xf4\x90\x3d\xac\xd0\xfe\xfc\x6a\xbe\xe9\x29\x62\x1a\xd5\x08\xe5\xb6\x31\x46\x4e\x31\x6e\x67\x69\x73\xae\xec\xb4\x82\x89\x1d\xaa\x4b\xd3\x2e\xca\x67\x44\xae\xae\x18\x17\x99\x9d\xbf\x7e\x66\xb6\x57\xed\x5f\x39\x9e\xa0\x92\x3b\xd8\xf8\xdf\x03\x6e\x2e\x40\x3d\x5d\x60\xac\x51\xd5\x7c\xf3\xfc\xae\x5a\xc9\x5d\xd7\x76\x1d\x21\x70\x12\xcf\x78\x41\xec\xdc\x73\x3c\x65\x95\xdc\x4d\x30\xf1\x14\xbd\xb9\xfb\x04\x9d\x16\xfc\x9e\x17\xd8\x3b\x66\xfb\x6d\x10\xbf\x9e\xff\x4b\x33\xb8\x68\xbf\x53\x36\xec\x02\xdd\xec\x37\x94\x30\x12\xc2\x5d\xbb\x2f\x61\xe1\xd2\x6d\x48\xa7\xdb\xff\xb3\xc1\xb6\x9f\x2e\x5c\x7a\xd9\x87\x5d\x74\xd4\xa2\xb2\x2c\xca\xd2\xba\x39\x71\xd3\xb0\x97\xad\x76\x53\x7c\x57\xd0\xfd\x73\xaf\x5d\xb8\xf3\x79\x65\x9f\x7a\x12\xaa\xd8\xf0\xe3\xa8\xf6\x77\x43\xff\x1b\x00\x00\xff\xff\xb9\xae\xf3\x2f\x86\x3e\x00\x00"),
		},
		"/templates/scheduler/airflow_1/base_dag.py": &vfsgen۰CompressedFileInfo{
			name:             "base_dag.py",
			modTime:          time.Date(2021, 1, 28, 6, 21, 54, 899853208, time.UTC),
			uncompressedSize: 4398,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x56\x6d\x6f\xdb\x38\x12\xfe\xae\x5f\x31\xd0\xa6\x70\x72\x90\x95\xec\x2e\x6e\x71\x30\xa0\x02\x6e\xa2\x36\x6e\x5e\x6c\xd8\x4a\x7b\xc5\xde\x81\x60\xa4\xb1\xcd\x35\x4d\xea\x48\x2a\x2f\xf5\xea\xbf\x1f\x48\xbd\x58\x8e\x9d\xdb\x00\xbd\x7c\x70\x44\xce\xcc\x33\xe4\xf0\x99\x87\x64\xeb\x5c\x2a\x03\x52\x7b\xf5\xd7\x1f\x5a\x8a\xe6\x3b\xa5\x1c\x45\x46\x55\x33\x56\xe8\xcd\x95\x5c\x83\x79\xce\x99\x58\x40\x3d\x3b\x14\xcf\x01\x9c\x53\xce\xe9\x3d\xc7\x00\x2e\x58\x6a\x02\x18\xe7\x86\x49\x41\x79\x15\x90\x51\x83\x86\xad\xb1\x09\x69\xc6\x01\xd8\xdf\x0c\xb9\xa1\xd5\xe7\x77\x29\xea\x1c\xda\xa8\x4e\x8e\x04\xd7\x39\xa7\xa6\xb6\xa5\x4a\x0a\x66\x50\x35\xd6\x66\xec\x55\x66\xca\xd4\x9c\xcb\xc7\x70\x55\xdc\xa3\x12\x68\x50\x37\x8e\x76\x86\xa4\x9c\xa1\x30\x01\xe4\x32\x23\x9c\x16\x22\x5d\xa2\xda\x0d\x5c\xcb\x0c\x79\x1b\x74\x31\xfc\x14\xc0\x17\xaa\x58\xbd\x3f\xba\x98\x16\xc2\xfd\xbf\xb1\x7e\x01\x24\x54\xaf\x46\x42\x1b\x2a\x52\x0c\xe0\x03\xd5\x38\xce\x51\x51\x23\x55\x00\xff\x3c\x97\x6b\xfb\x3b\xbe\x21\xd3\x38\xb9\x9b\xde\x92\xab\xf8\xdb\x6e\xb6\xa5\x94\x2b\x1d\x2e\x8d\xc9\x89\xfd\x6c\xd2\x5e\x1a\x93\x5f\x4a\xb9\xda\x75\x96\x35\xb2\x0e\xf3\x67\xb3\x94\x82\x34\x13\x4d\xd8\xc4\x4d\x37\x0b\xd8\x0d\xd6\x28\xb4\x0d\xbd\xa7\x1a\x49\x35\xd8\x8b\xb7\xcb\x9f\x39\xd3\x61\x8c\x54\x0a\xa3\xd8\x7d\x67\x21\xdb\x32\x13\x5b\xd2\x97\x80\x57\xad\x79\x22\xb3\xc3\x98\x5b\x84\x50\x63\xaa\xd0\x34\xb1\x33\x37\x3a\xbc\x09\x7c\x32\xa8\x04\xe5\xc4\x50\xbd\xaa\x77\xd3\xc4\xc5\xb5\xcd\x9e\x4c\xb5\x99\x5d\x0c\x7c\x4a\xd1\x11\xb4\x3d\xe4\x61\x65\x89\x1b\xc3\xae\x7f\x61\x18\xd7\x61\x86\xa9\xac\xf6\xdc\x44\xd1\x3c\xe7\xcf\x24\xc3\x39\x2d\xb8\xd1\x07\x63\xee\x1b\xdf\x5c\xc9\x07\x96\xd9\xba\x6b\xbd\x97\x20\x95\x62\xce\x16\x85\xa2\x36\x77\x4b\x6b\x29\xe6\x87\x30\xb5\xa1\xa6\x6d\xa5\x99\x69\xdb\x62\x97\x50\xee\x90\xbb\x84\xb2\x27\xbb\x4f\xa8\x0a\xf2\x11\xd9\x62\x69\x88\x2a\x78\x0b\xfc\xd5\x4d\x4d\x0b\x8e\x9e\xd7\x69\x9f\xfa\x9c\xaa\x26\x6a\x3a\x85\x6a\x58\xfd\x43\xbf\x7a\xa8\x0f\x92\x17\xdb\xe6\xff\xe2\x46\x7f\xe1\x4c\xd6\xb2\x10\x66\x37\xe4\xc6\x4e\xd5\x3d\x4e\x08\x67\x6d\x69\x29\x47\x65\xc8\x9c\x32\x8e\x19\x31\x92\x68\x4e\xd3\x55\x00\xb3\x22\x47\x75\x90\x7e\xb5\x6d\x9f\x26\x01\xfc\xcb\x03\x00\x98\x59\x84\xaf\x78\x6f\xeb\xd7\x72\xd6\xf3\x66\xf1\xf9\x34\x4e\xc8\xed\xf0\x26\x86\xa8\xd5\x84\x70\x81\xe6\xd8\xaf\x98\x4b\x04\x5d\xa3\x1f\x80\x2f\x73\xc3\xd6\x85\xee\x2f\xa4\x5c\x70\xec\xa7\x0a\x33\x14\x86\x51\xae\xfd\x93\x06\xe7\x2a\xfe\xf6\x1a\xcc\x0a\x9f\x2d\x0a\x2d\xcc\x32\xb4\x92\xbc\x0d\xfa\x32\xbe\xbe\xbb\x89\xc9\x64\x98\x5c\x42\x04\xbd\x53\x99\x9b\xd3\x3a\xd9\x69\x15\xac\x4f\x7b\xde\x2c\xbe\x9d\x8d\xa7\xe4\x22\xfe\x38\xbc\xbb\x4e\xc8\x64\x7c\x15\x93\xd1\x6d\x12\x4f\xbf\x0c\xaf\xc9\xe8\x96\xcc\xe2\xf3\x19\x44\xc0\x84\x39\x7e\x99\xdf\xc9\x42\x2e\x57\x48\x98\x30\xa8\x1e\x28\x27\x4c\x10\x8d\xa9\xf6\x03\xa8\xb9\x4e\x1e\xa8\x8a\x7e\xfe\x3b\xfc\x0d\x7e\x3b\x3b\x39\x79\x99\x2d\x19\xdd\xc4\xe3\xbb\xe4\x2d\x79\xac\xe8\xcb\xc2\xfc\xef\x0c\x4d\x1a\x6f\x91\x72\x59\x64\xa4\x53\x4c\x92\x53\xb3\xb4\x75\xd8\x94\x9b\xb2\x17\xce\xa5\x5a\x53\x73\xbc\x5f\xaa\x00\xb6\x35\x3f\x69\x70\x6a\xb1\x89\x6a\x9d\x39\x76\x47\xdf\xab\x08\xd8\x0b\xa0\x62\xc2\x3e\x54\x77\xde\x52\x61\xd7\xd3\x25\xf0\xbc\x66\x1b\x54\x2d\x34\x44\xb0\x71\x2e\xbe\x7c\x14\xa8\xfc\x01\xf8\x9b\x4d\xf8\x59\xde\x87\x63\x3b\x2e\x4b\xbf\xc2\xf4\x33\xcc\x51\x64\x9a\x48\x41\x72\xaa\x8d\x3f\x80\xcd\xa6\x0f\x6c\x0e\xce\xf9\x03\x2e\xe9\x03\x93\x2a\xbc\xa8\xdc\xc6\x62\x42\xb5\x81\xb2\x84\x44\x15\x08\x9b\x0d\x20\xd7\x68\xc7\x1f\xa9\xfd\xb0\x13\x22\x83\x7e\x59\xd6\xf0\x0a\x8d\x62\xa8\xfd\x01\xfc\xda\x99\xb1\x02\xc6\xe9\xb3\x3f\xd8\x5e\xc6\xc7\x1a\x53\x29\x32\x1d\xfd\x7a\x76\x76\x52\xbb\x6a\x43\x95\x21\xf6\xe6\xf6\x07\xed\x05\x1e\x6a\xa3\x2c\xf7\xf0\x78\xb3\xa9\x16\x39\x4b\x97\x98\x15\x1c\xc3\x99\xf5\xbf\xa0\x06\xc3\x8f\xee\x50\xc0\xff\xe5\xec\xec\xb7\xfe\xd9\xcf\xfd\xb3\x5f\x7c\xf8\x13\xfe\x53\x48\x63\x17\x1b\x80\xff\xee\x5b\xff\xdd\xba\xff\x2e\xf3\x9b\x5c\x52\xb8\x6e\x2e\x14\x92\x94\x72\x7e\x4f\xd3\x95\x3f\x78\xa5\xcd\xab\x88\x5c\x31\xa9\x98\x79\x26\x95\x92\xb9\xca\xb9\xf5\xd8\xde\x0e\x27\xb5\xb5\xad\x44\x47\xef\xfc\x41\x47\xea\xc2\xe1\x87\xd9\xf8\xfa\x2e\x89\xbd\xd2\xf3\x32\xba\x80\xc8\x5e\xff\x15\x31\x32\xba\x20\x2c\x8b\x9a\xa3\xbb\xa5\x6b\x6c\x4f\xae\x7b\xda\x51\x77\x50\x99\x75\x5d\x94\xb6\xa3\x5a\x94\xb6\x5c\xa3\xda\xd2\x42\xa6\xd4\xa4\xcb\x22\xb7\xdc\xd9\xec\x53\xe0\xdc\x5a\xef\xf2\xbf\x3c\xfc\xb2\xf4\x4e\x3c\xcf\x28\x2a\x74\xd5\x1b\x4c\x0a\xd2\x2d\xcd\x9d\x60\x26\xfc\x84\xc6\x6e\x07\xfe\x04\x85\x39\xa7\x29\x82\xdf\xf7\xc1\x27\x24\xa3\x7a\x49\x88\xdf\x35\x84\x95\x41\x1a\x42\xfc\xb2\xb4\xcd\xf3\xaa\xcc\x56\x75\x63\x6b\xba\x40\x92\x17\x9c\x93\x5c\x72\x96\x3e\x47\xfe\x90\x3f\xd2\x67\x5d\x6f\xd4\x2a\xa6\xce\x2d\x76\xe4\xae\x3c\xa7\x0f\xbd\xed\x8d\xd0\x0b\xa0\xd7\xfa\xf4\x02\x98\xd7\x8c\x88\xfc\xba\xd0\x0d\x6d\x5c\x22\x88\xc0\xdf\x94\x7e\xa3\x04\xfe\xa1\xbd\x8e\xac\x63\x59\x36\x71\xe9\x3a\xd3\xd1\xef\xff\xde\xae\x26\xf2\xdf\x52\x21\xe2\xbb\x2a\xb5\x07\xe6\xde\x21\x1d\x82\xec\x05\xb7\x9e\x0b\x34\x84\xcb\x85\x8e\xec\xd1\x05\x0d\xb9\xa2\x8c\x2e\xea\x8d\x08\x92\xf2\x42\x1b\x54\x1d\x0f\xa6\x6d\x9f\xa2\xc1\xf6\x85\x65\x9f\x5b\x5d\x08\x49\x9e\x52\xb9\x26\x79\xa1\x97\x91\x23\x41\x4d\xbe\xea\x52\x88\x7e\xdf\x51\xbd\x7a\xbb\x28\x1e\xac\xd2\xea\xa8\x12\x28\xd7\x1c\x9f\xc6\xe3\x4f\xd7\x31\x19\x4e\x26\xd7\xa3\xf3\x61\x32\x1a\xdf\x92\xf3\x69\x7c\x11\xdf\x26\xa3\xe1\xf5\xcc\x1f\xc0\x2b\x32\x1c\x6c\x21\x3e\x8f\x3f\x38\x69\xf4\x07\xbd\x9d\x76\xe9\x05\xe0\x8f\x27\xc9\xe8\xe6\x6e\x46\x2e\xc7\xb3\xa4\xf2\x01\xeb\x74\x29\xb5\x11\xb5\xd3\x2e\xd0\xc5\x68\xea\x0f\x7a\xa7\x19\x35\xd4\xc6\x4f\xa6\xe3\xcf\xf1\x79\x52\x41\x4f\x94\xfc\x03\x53\xd3\x81\x4f\x86\xb3\x2b\x92\x7c\x9b\xd8\xdc\xf6\x0d\xd4\x45\x9b\x9d\x5f\xc6\x17\x77\xd7\xf1\x05\x19\x56\xf1\x56\x88\x41\xe0\x93\x21\xf8\x84\x69\xe1\xba\xc3\x6a\x9b\x3d\x55\x68\x57\x52\xab\x86\x42\x6a\x0c\x4d\x97\x56\x9e\x15\x3a\x3d\xac\xcb\x7f\xe2\x79\x3f\x81\x7b\x76\x01\x97\x32\x07\x67\xf3\x36\x1b\x50\x54\x2c\x10\x8e\x48\x00\x47\x06\x06\x51\xd5\xc4\x97\xce\xb1\x2c\x3d\x2b\xee\xb5\x42\xef\xc6\x5b\x75\xb7\x90\xa9\x4d\x89\x50\xe4\xda\x28\xa4\x6b\xa8\xdf\xbe\x2e\xb0\x86\x66\x0d\xf4\x91\xc3\xae\xae\x06\x14\x29\x43\x5d\x96\xde\x23\x65\x86\x6c\x36\x47\xa6\x3d\x83\x1f\xe8\xf1\xfd\xe7\x52\xd5\xe0\xed\x4b\xbc\x52\x48\xd7\x81\xdd\x94\x2d\xef\x1f\x99\xc8\xe4\x23\xd1\xec\x3b\x3a\x59\xab\x7d\x5c\xa3\x7c\x75\xb6\x70\xc6\xbe\x63\x78\x29\x0b\x65\x0b\xb4\x13\x25\xe7\x73\xed\x2e\xeb\x83\x71\x63\x67\x3d\x1c\x69\x54\x21\x52\x6a\x90\x14\xb9\x91\x3b\xab\xeb\x22\x24\xb5\x57\x22\x5f\x76\xb4\x0d\xb1\x85\xec\xbf\xd8\x55\x7f\x17\xe7\x60\xab\xef\xbc\xa3\x6c\x21\xdf\xf0\x2e\xab\x93\x57\x2f\xa3\xfd\x98\x17\xaf\xab\x1d\xfd\xf0\x4e\x1a\x52\x95\xa5\x25\x10\x55\x15\x4d\xdc\x0a\xdc\x7e\x20\xeb\x30\xc4\xfb\xe9\x0d\x7f\x16\xe7\x25\x03\xa1\xff\x1e\x6c\x73\xc1\xee\xcd\xe2\x32\xbc\x91\x9d\xb0\xa5\x27\xfc\x08\x3f\xed\x95\xf7\xfe\x3d\xec\x5d\x71\x47\xff\xb7\x3b\x6e\x5b\xd2\xff\x06\x00\x00\xff\xff\x5c\xe0\xb9\x5b\x2e\x11\x00\x00"),
		},
		"/templates/scheduler/airflow_2": &vfsgen۰DirInfo{
			name:    "airflow_2",
			modTime: time.Date(2021, 1, 11, 7, 4, 36, 499681802, time.UTC),
		},
		"/templates/task": &vfsgen۰DirInfo{
			name:    "task",
			modTime: time.Date(2021, 1, 12, 11, 35, 43, 639604734, time.UTC),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/migrations"].(os.FileInfo),
		fs["/templates"].(os.FileInfo),
	}
	fs["/migrations"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/migrations/000001_create_project_table.down.sql"].(os.FileInfo),
		fs["/migrations/000001_create_project_table.up.sql"].(os.FileInfo),
		fs["/migrations/000002_create_job_table.down.sql"].(os.FileInfo),
		fs["/migrations/000002_create_job_table.up.sql"].(os.FileInfo),
		fs["/migrations/000003_create_instance_table.down.sql"].(os.FileInfo),
		fs["/migrations/000003_create_instance_table.up.sql"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/scheduler"].(os.FileInfo),
		fs["/templates/task"].(os.FileInfo),
	}
	fs["/templates/scheduler"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/scheduler/airflow_1"].(os.FileInfo),
		fs["/templates/scheduler/airflow_2"].(os.FileInfo),
	}
	fs["/templates/scheduler/airflow_1"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/scheduler/airflow_1/__lib.py"].(os.FileInfo),
		fs["/templates/scheduler/airflow_1/base_dag.py"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}