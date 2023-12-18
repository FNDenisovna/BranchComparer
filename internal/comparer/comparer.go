package comparer

import (
	"branchcomparer/internal/domain"
	"encoding/json"
	"fmt"
	"sync"
)

type Comparer struct {
	b1  string
	b2  string
	api Api
}

type Api interface {
	GetPackages(branch string) ([]domain.Package, error)
}

/*type Package struct {
	Name    string
	Version string
	Arch    string
}*/

func New(_b1 string, _b2 string, _api Api) *Comparer {
	return &Comparer{
		b1:  _b1,
		b2:  _b2,
		api: _api,
	}
}

var (
	p1map = make(map[string]map[string]string)
	p2map = make(map[string]map[string]string)
)

func (c *Comparer) Compare() (resultJson []byte, err error) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(pmap *map[string]map[string]string) {
		p := make([]domain.Package, 0, 0)
		p, err = c.api.GetPackages(c.b1)
		err = c.doPackegeMap(pmap, p)
		wg.Done()
	}(&p1map)

	go func(pmap *map[string]map[string]string) {
		p := make([]domain.Package, 0, 0)
		p, err = c.api.GetPackages(c.b2)
		err = c.doPackegeMap(pmap, p)
		wg.Done()
	}(&p2map)

	wg.Wait()
	if err != nil {
		err = fmt.Errorf("Getting branches packages is failed. Error: %w", err)
		return resultJson, err
	}

	resultStruct := &CompareResult{
		DiffPackege1:   make([]Arch, 0),
		DiffPackege2:   make([]Arch, 0),
		DiffPackegeVer: make([]Arch, 0),
	}

	wg.Add(2)
	go func() {
		resultStruct.DiffPackege1, resultStruct.DiffPackegeVer = c.getDiffs(&p1map, &p2map, true)
		wg.Done()
	}()

	go func() {
		resultStruct.DiffPackege2, _ = c.getDiffs(&p2map, &p1map, false)
		wg.Done()
	}()

	wg.Wait()
	resultJson, err = json.Marshal(resultStruct)
	return
}

func (c *Comparer) doPackegeMap(pmap *map[string]map[string]string, ps []domain.Package) (err error) {
	for _, p := range ps {
		if arch, ok := (*pmap)[p.Arch]; ok {
			if _, ok := arch[p.Name]; ok {
				continue
			} else {
				arch[p.Name] = p.Version
			}
		} else {
			(*pmap)[p.Arch] = make(map[string]string)
			(*pmap)[p.Arch][p.Name] = p.Version
		}
	}
	return nil
}

func (c *Comparer) getDiffs(pm1 *map[string]map[string]string, pm2 *map[string]map[string]string, addVersionDiff bool) (diffPackege1 []Arch, diffPackegeVer []Arch) {
	diffPackege1 = make([]Arch, 0)
	diffPackegeVer = make([]Arch, 0)

	for k, v1 := range *pm1 {
		arch := Arch{
			Name:     k,
			Packages: make([]string, 0),
		}
		archVer := Arch{
			Name:     k,
			Packages: make([]string, 0),
		}

		if v2, ok := (*pm2)[k]; ok {
			for p1, ver1 := range v1 {
				if ver2, ok := v2[p1]; !ok {
					arch.Packages = append(arch.Packages, p1)
					if addVersionDiff {
						if ver1 > ver2 {
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

type CompareResult struct {
	mu             sync.Mutex
	DiffPackege1   []Arch `json:"diffpackage1"`
	DiffPackege2   []Arch `json:"diffpackage2"`
	DiffPackegeVer []Arch `json:"diffpackagever"`
}

type Arch struct {
	Name     string   `json:"archname"`
	Packages []string `json:"packages"`
}

/*
{"ajaxId":"DDA44F72-A109-4003-A99C-229B066C9BC8","sessionId":"476413EF-72B6-4F9E-AE4C-7E029AAE1506","resSignature":"E3E8934C-235A-4B0E-825A-35A08381A191","rtl":false,"longPooling":true,"plugins":[{"name":"wsm","settingsJson":"","localizationDictionary":null},{"name":"wnt","settingsJson":"","localizationDictionary":null},{"name":"ca","settingsJson":"{\"submitHandlerEnabled\":true,\"wfdIdSelector\":\"input, form\",\"avoidTypes\":[\"checkbox\",\"radio\",\"submit\",\"button\",\"file\",\"image\",\"reset\"]}","localizationDictionary":null},{"name":"xhr_content","settingsJson":"","localizationDictionary":null}],"enableTracing":false}
*/
