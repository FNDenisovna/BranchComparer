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
