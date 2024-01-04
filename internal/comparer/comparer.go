package comparer

import (
	"branchcomparer/internal/domain"
	versioncomparer "branchcomparer/internal/versionComparer/v2"
	"encoding/json"
	"fmt"
	"sync"
)

type Comparer struct {
	b1    string
	b2    string
	api   Api
	mu    sync.RWMutex
	p1map *domain.PackageMap
	p2map *domain.PackageMap
}

type Api interface {
	GetPackages(branch string) ([]domain.Package, error)
}

func New(_b1 string, _b2 string, _api Api) *Comparer {
	return &Comparer{
		b1:    _b1,
		b2:    _b2,
		api:   _api,
		p1map: domain.New(),
		p2map: domain.New(),
	}
}

type CompareResult struct {
	DiffPackege1   []Arch `json:"diffpackage1"`
	DiffPackege2   []Arch `json:"diffpackage2"`
	DiffPackegeVer []Arch `json:"diffpackagever"`
}

type Arch struct {
	Name     string   `json:"archname"`
	Packages []string `json:"packages"`
}

func (c *Comparer) Compare() (resultJson []byte, err error) {
	var wg sync.WaitGroup
	wg.Add(2)
	var err1, err2 error
	go func(pmap *domain.PackageMap) {
		defer wg.Done()
		var p []domain.Package
		p, err1 = c.api.GetPackages(c.b1)
		if err1 != nil {
			return
		}
		err1 = c.doPackegeMap(pmap, p)
	}(c.p1map)

	go func(pmap *domain.PackageMap) {
		defer wg.Done()
		var p []domain.Package
		p, err2 = c.api.GetPackages(c.b2)
		if err2 != nil {
			return
		}
		err2 = c.doPackegeMap(pmap, p)
	}(c.p2map)

	wg.Wait()
	if err1 != nil {
		err = fmt.Errorf("Getting branches packages is failed. Error: \n%w", err1)
		return resultJson, err
	} else if err2 != nil {
		err = fmt.Errorf("Getting branches packages is failed. Error: \n%w", err2)
		return resultJson, err
	}

	resultStruct := &CompareResult{
		DiffPackege1:   make([]Arch, 0),
		DiffPackege2:   make([]Arch, 0),
		DiffPackegeVer: make([]Arch, 0),
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		resultStruct.DiffPackege1, resultStruct.DiffPackegeVer = c.getDiffs(c.p1map, c.p2map, true)
	}()

	go func() {
		defer wg.Done()
		resultStruct.DiffPackege2, _ = c.getDiffs(c.p2map, c.p1map, false)
	}()

	wg.Wait()
	resultJson, err = json.Marshal(resultStruct)
	return
}

func (c *Comparer) doPackegeMap(pmap *domain.PackageMap, ps []domain.Package) (err error) {
	for _, p := range ps {
		if arch, ok := (pmap.Arch)[p.Arch]; ok {
			if _, ok := arch[p.Name]; ok {
				continue
			} else {
				arch[p.Name] = fmt.Sprintf("%d:%s-%s", p.Epoch, p.Version, p.Release)
			}
		} else {
			(pmap.Arch)[p.Arch] = make(map[string]string)
			(pmap.Arch)[p.Arch][p.Name] = p.Version
		}
	}
	return nil
}

func (c *Comparer) getDiffs(pm1 *domain.PackageMap, pm2 *domain.PackageMap, addVersionDiff bool) (diffPackege1 []Arch, diffPackegeVer []Arch) {
	diffPackege1 = make([]Arch, 0)
	diffPackegeVer = make([]Arch, 0)

	for k, v1 := range pm1.Arch {
		arch := Arch{
			Name:     k,
			Packages: make([]string, 0),
		}
		archVer := Arch{
			Name:     k,
			Packages: make([]string, 0),
		}

		if v2, ok := (pm2.Arch)[k]; ok {
			for p1, ver1 := range v1 {
				if ver2, ok := v2[p1]; !ok {
					arch.Packages = append(arch.Packages, p1)
				} else {
					if addVersionDiff {
						compRes, err := versioncomparer.Compare(ver1, ver2)
						if err != nil {
							fmt.Printf("Can't compare versions: %s, %s. Arch: %v, package: %v. Error: \n%v\n", ver1, ver2, k, p1, err)
							continue
						}
						if compRes > 0 {
							archVer.Packages = append(arch.Packages, p1)
						}
					}
				}
			}
		}
		if len(arch.Packages) > 0 {
			diffPackege1 = append(diffPackege1, arch)
		}
		if len(archVer.Packages) > 0 {
			diffPackegeVer = append(diffPackegeVer, archVer)
		}
	}

	return
}

/*
{"ajaxId":"DDA44F72-A109-4003-A99C-229B066C9BC8","sessionId":"476413EF-72B6-4F9E-AE4C-7E029AAE1506","resSignature":"E3E8934C-235A-4B0E-825A-35A08381A191","rtl":false,"longPooling":true,"plugins":[{"name":"wsm","settingsJson":"","localizationDictionary":null},{"name":"wnt","settingsJson":"","localizationDictionary":null},{"name":"ca","settingsJson":"{\"submitHandlerEnabled\":true,\"wfdIdSelector\":\"input, form\",\"avoidTypes\":[\"checkbox\",\"radio\",\"submit\",\"button\",\"file\",\"image\",\"reset\"]}","localizationDictionary":null},{"name":"xhr_content","settingsJson":"","localizationDictionary":null}],"enableTracing":false}
*/
