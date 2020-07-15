package china_division

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

type Code int

const (
	UnknownCode Code = iota
	ProvinceCode
	CityCode
	AreaCode
)

// 所有省
var provinces []Row
var provinceJson []byte
var provinceMap = make(map[string]string, 40)

// 所有市，key为省code
var cities = make(map[string][]Row, 40)
var cityJson = make(map[string][]byte, 40)
var cityMap = make(map[string]string, 300)

// 所有区县，key为市code
var areas = make(map[string][]Row, 300)
var areaJson = make(map[string][]byte, 300)
var areaMap = make(map[string]string, 1000)

type Row struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	dataFile := filepath.Join(filepath.Dir(currentFile), `data`, `data.json`)
	content, err := ioutil.ReadFile(dataFile)
	if err != nil {
		panic(err)
	}
	rows := make([]Row, 3000)

	if err := json.Unmarshal(content, &rows); err != nil {
		panic(err)
	}

	for _, row := range rows {
		provinceCode := row.Code[:2]
		cityCode := row.Code[2:4]
		areaCode := row.Code[4:]
		provinceCityCode := row.Code[:4]

		if cityCode == "00" { // 省
			if areaCode != "00" {
				log.Println("Can't parse code: ", row.Code)
			}

			provinces = append(provinces, row)
			provinceMap[row.Code] = row.Name
		} else if areaCode == "00" { // 市
			cities[provinceCode] = append(cities[provinceCode], row)
			cityMap[row.Code] = row.Name
		} else { // 区
			areas[provinceCityCode] = append(areas[provinceCityCode], row)
			areaMap[row.Code] = row.Name
		}
	}

	provinceJson, err = json.Marshal(provinces)
	if err != nil {
		panic(err)
	}

	for provinceCode, v := range cities {
		cityJson[provinceCode], err = json.Marshal(v)
		if err != nil {
			panic(err)
		}
	}

	for provinceCityCode, v := range areas {
		areaJson[provinceCityCode], err = json.Marshal(v)
		if err != nil {
			panic(err)
		}
	}
}

// 返回 code 对应的省市区
func GetFullName(code string) (province, city, area string) {
	province = provinceMap[code[:2]+`0000`]
	city = cityMap[code[:4]+`00`]
	area = areaMap[code[:6]]
	return
}

// 获得 json 格式的下一级
func GetJsonChildren(code string) []byte {
	if code == `` {
		return GetJsonProvinces()
	}

	switch CodeType(code) {
	case ProvinceCode:
		return GetJsonCities(code)
	case CityCode:
		return GetJsonAreas(code)
	default:
		return []byte("{}")
	}
}

// 获得 slice 格式的下一级
func GetChildren(code string) []Row {
	if code == `` {
		return GetProvinces()
	}

	switch CodeType(code) {
	case ProvinceCode:
		return GetCities(code)
	case CityCode:
		return GetAreas(code)
	default:
		return nil
	}
}

// 获得 json 格式的所有省名及其代码
func GetJsonProvinces() []byte {
	return provinceJson
}

// 获得 slice 格式的所有省名及其代码
func GetProvinces() []Row {
	return provinces
}

// 获得 json 格式的所有市名及其代码
func GetJsonCities(code string) []byte {
	if CodeType(code) != ProvinceCode {
		return []byte("{}")
	}
	return cityJson[code[:2]]
}

// 获得 slice 格式的所有市名及其代码
func GetCities(code string) []Row {
	if CodeType(code) != ProvinceCode {
		return nil
	}
	return cities[code[:2]]
}

// 获得 json 格式的所有区县名及其代码
func GetJsonAreas(code string) []byte {
	if CodeType(code) != CityCode {
		return []byte("{}")
	}
	return areaJson[code[:4]]
}

// 获得 slice 格式的所有区县名及其代码
func GetAreas(code string) []Row {
	if CodeType(code) != CityCode {
		return nil
	}
	return areas[code[:4]]
}

// 返回当前代码的类型
func CodeType(code string) Code {
	if len(code) != 6 {
		return UnknownCode
	}

	if code[2:4] == `00` {
		if _, ok := provinceMap[code]; ok {
			return ProvinceCode
		}
	} else if code[4:6] == `00` {
		if _, ok := cityMap[code]; ok {
			return CityCode
		}
	} else {
		if _, ok := areaMap[code]; ok {
			return AreaCode
		}
	}
	return UnknownCode
}
