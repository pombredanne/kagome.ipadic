package data

import(
	"os"
	"time"
)

var _dicIpaChardefDic = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\xdc\x7f\x4f\xda\x40\x18\x00\xe0\x5e\x15\xe9\x94\x64\x9f\xe9\xc6\x98\x32\x90\x19\xd4\x3f\xcc\x3e\xce\x7e\x7c\xeb\x85\x15\x44\xb7\xa1\x0c\xb0\xad\x47\xe5\x79\x42\x29\xbd\x1e\xed\x7b\xbd\xeb\x0b\xc9\x25\xed\xcd\xbe\xe5\x21\x9f\x7d\xcf\x42\x2f\xcb\xbe\x96\xeb\xb3\xee\xc7\xc1\xa7\x78\x3b\xbe\xe9\x5c\x5f\xc5\xfe\xa0\x33\x8a\x93\xcf\xc3\x93\xeb\xbb\xcb\x0f\x5f\xc6\xdd\xc9\xed\xe5\x60\x3a\xec\x77\xe2\xf8\xea\x22\x16\x17\xc3\x69\x3c\x8f\x93\x58\x8c\xe2\x4d\x2c\xeb\xc5\xde\xa2\xf2\x43\xa5\xf3\xe9\x60\x30\x2a\xfa\x77\xd3\xe1\x78\x3c\xec\xff\x0a\xd9\xc9\x69\x56\xbe\xff\x11\x42\xc8\xfe\xde\x7e\x28\x3e\xfa\xd7\xf1\xa3\x65\x41\x67\xad\x6d\xf6\x3f\x3d\xe1\xaa\xa3\xff\x5b\x7f\x78\x1a\xb4\xb9\xdf\x58\xeb\xdd\x9e\x98\xc7\x72\x7a\xb0\xee\xfb\xe2\xf9\x52\x00\x00\x00\x00\x00\x80\xf6\x4a\x3d\x91\x98\x4e\xea\x2b\xbf\x1f\x36\x4c\xaf\xd7\xee\x78\x47\xaf\x1d\x5f\x55\xbb\xb6\xaf\xed\xed\xa5\xdd\xde\x5a\x7e\xd9\xd5\xde\x07\x48\x93\x74\x3f\x49\x35\x9d\x7f\xdb\x9e\x9f\x69\x56\xea\xf1\xc7\x61\xcb\x13\x48\x7d\xfe\xe6\x6d\x7f\xfd\x2b\xe7\x8f\x93\x66\x65\xdd\x66\xd5\x3d\x9e\xf7\xdd\x6a\xfb\xeb\xf8\x0d\x69\xb1\x83\xff\xff\x93\x3a\x53\x01\x00\x00\x00\x00\x00\x30\x57\xdf\x7c\x60\x91\xe7\xe5\xab\x5c\x8a\x06\xc2\xdc\xf1\x98\xc5\xcb\xbe\x06\x95\x54\x1f\x6f\x5b\x1d\xa1\x99\x7b\x8c\xb4\xf4\x2a\xec\x23\xf7\x25\x00\x00\x00\x00\x00\x00\x00\xbc\x0a\x53\xf4\x00\x00\x00\x00\x00\x00\x00\x00\x00\xc0\x5b\x55\xdf\xf3\x1e\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x9a\x96\xfa\xc9\xce\xed\x52\xe9\x7a\xa5\xee\x6a\xa0\xa2\xa3\x0a\x52\xc7\xfe\xac\x95\x18\x8f\x1f\x2d\x63\xee\xac\x75\x5f\x61\xd3\xfe\x85\xee\x8b\xd5\xdc\xbc\x27\x7d\xd0\x9b\xfd\xc8\x43\x3e\xfb\x99\x85\x32\x3f\xbf\x2f\xd7\x67\x65\x61\x08\x61\xbe\x84\xc5\xf6\xf2\xf3\xdc\xef\x00\x00\x00\xff\xff\x1d\xc3\x2e\x56\xa0\x00\x01\x00"

func dicIpaChardefDicBytes() ([]byte, error) {
	return bindataRead(
		_dicIpaChardefDic,
		"dic/ipa/chardef.dic",
	)
}

func dicIpaChardefDic() (*asset, error) {
	bytes, err := dicIpaChardefDicBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dic/ipa/chardef.dic", size: 65696, mode: os.FileMode(420), modTime: time.Unix(1452316192, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

