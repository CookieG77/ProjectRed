package PPR

import (
	"encoding/json"
	"os"
)

func OpenJson(res *map[string]map[string]interface{}, filepath string) bool {
	//Retourne 0 si aucune erreur n'a été rencontré lors du dépactage du json. Sinon retourne 1
	file, err1 := os.ReadFile(filepath)
	if err1 != nil || json.Unmarshal(file, res) != nil {
		return true
	}
	return false
}

func InitItemList(res *map[string]map[string]interface{}, filepath string) bool {
	if OpenJson(res, filepath) {
		return true
	}
	for k, v := range *res {
		ninter := make(map[string]interface{})
		for kk, vv := range v {
			tmp, ok := vv.(float64)
			if ok {
				ninter[kk] = int(tmp)
			} else {
				ninter[kk] = vv
			}
		}
		(*res)[k] = ninter
	}
	return false
}

func InitInventory() map[string]int {
	return make(map[string]int)
}

func InitPlayer() map[string]interface{} {
	var res map[string]interface{} = {
		"name" : ""
		"pv" : 0
		"pvmax" : 0
	}
	return res
}
