package domain

type Package struct {
	Name string `json:"name"`
	//Epoch     int    `json:"epoch"`
	Version string `json:"version"`
	//Release   string `json:"release"`
	Arch string `json:"arch"`
	//Disttag   string `json:"disttag"`
	//Buildtime int    `json:"buildtime"`
	//Source    string `json:"source"`
}

type PackageMap struct {
	Arch map[string]map[string]string
}

func New() *PackageMap {
	return &PackageMap{
		Arch: make(map[string]map[string]string),
	}
}
