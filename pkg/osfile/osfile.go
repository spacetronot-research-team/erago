package osfile

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spacetronot-research-team/erago/pkg/ctxutil"
)

func AddUniqueErrCodeToErrorsJSON(ctx context.Context, uniqueErrCodes ...string) error {
	errorsJSONPath := ctxutil.GetDocsErrorsJSONFilePath(ctx)
	mapData, err := jsonFileToMap(errorsJSONPath)
	if err != nil {
		return fmt.Errorf("err json file to map: %v", err)
	}

	for _, uniqueErrCode := range uniqueErrCodes {
		mapData[uniqueErrCode] = "kalimat yang akan ditampilkan frontend"
	}

	if err := mapToJSONFile(mapData, errorsJSONPath); err != nil {
		return fmt.Errorf("err map to json file: %v", err)
	}

	return nil
}

func jsonFileToMap(jsonPath string) (map[string]interface{}, error) {
	jsonFile, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("err read JSON file: %v", err)
	}

	mapData := map[string]interface{}{}
	if err = json.Unmarshal(jsonFile, &mapData); err != nil {
		return nil, fmt.Errorf("err unmarshal JSON file: %v", err)
	}

	return mapData, nil
}

func mapToJSONFile(mapData map[string]interface{}, jsonPath string) error {
	jsonData, err := json.MarshalIndent(mapData, "", "    ")
	if err != nil {
		return fmt.Errorf("err marshal mapData: %v", err)
	}

	file, err := os.Create(jsonPath)
	if err != nil {
		return fmt.Errorf("err create json file: %v", err)
	}
	defer file.Close()

	if _, err = file.Write(jsonData); err != nil {
		return fmt.Errorf("err write JSON to file: %v", err)
	}

	return nil
}
