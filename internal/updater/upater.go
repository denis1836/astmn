package updater

import (
	"fmt"
	"os"

	"astmn/internal/config"
	"astmn/internal/db"
	"astmn/internal/log"
	"astmn/internal/manifest"
)

type UpdateInfo struct {
	UpdatablePackagesAmount int64
	UpdatablePackagesNames  []string
}

func LookForUpdates(c *config.Config) (UpdateInfo, error) {
	var upinf UpdateInfo
	var manifests []manifest.Manifest
	log.Infof("fetching manifests from: %s", c.AssetsRegistryDir)
	os.Stat(c.AssetsRegistryDir)
	//TODO: fetch all the current manifests from db
	db.GetPackages() //TODO: packages objects array getter from db
	//TODO: check manifests hashes if they match

	//TODO: fill update info object with info
	//TODO: display user the info
	//TODO: user selects what to update
	//TODO: download packages with downloader.DownloadFile

	fmt.Println()
	return upinf, nil
}
