package interfacepractice

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type economy struct {
	SystemObjectID string `json:"system_object_id,omitempty"`
	Code           string `json:"Kod,omitempty"`
	IsDeleted      int    `json:"is_deleted"`
	SignatureData  string `json:"signature_date"`
	Description    string `json:"Nomdescr,omitempty"`
	GlobalID       int64  `json:"global_id"`
	IDx            string `json:"Idx"`
	Unit           string `json:"Razdel"`
	Name           string `json:"Name"`
}

func DataParse() {

	file, err := os.Open("K:/Projects/go_study/interfacePractice/data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var e []economy
	if err := json.Unmarshal(data, &e); err != nil {
		fmt.Println(err)
		return
	}
	var globalSum int64 = 0
	for _, eStruct := range e {
		globalSum += eStruct.GlobalID
	}
	fmt.Println(globalSum)

}
