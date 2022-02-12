package config

import (
    "encoding/json"
    "os"
)

// ReadJSON reads a JSON config and writes it to “Conf”.
func ReadJSON(configPath string) error {
    configContents, err := os.ReadFile(configPath)
    if err != nil {
        return err
    }

    err = json.Unmarshal(configContents, &Conf)
    if err != nil {
        return err
    }

    return nil
}