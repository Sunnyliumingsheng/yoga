package picture

import yangoss "github.com/Sunnyliumingsheng/yangoss"

func StartStoragePicture() {
	yangoss.Config.ServerPath = "http://localhost:80/img"
	yangoss.Config.StoragePath = "/home/yang/temp/picture"
}
