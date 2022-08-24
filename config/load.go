package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"sigs.k8s.io/yaml"

	v1alpha12 "github.com/uor-framework/uor-client-go/api/v1alpha1"
)

// ReadDataSetConfig reads the specified config into a DataSetConfiguration type.
func ReadDataSetConfig(configPath string) (v1alpha12.DataSetConfiguration, error) {
	var configuration v1alpha12.DataSetConfiguration
	data, err := readInConfig(configPath, v1alpha12.DataSetConfigurationKind)
	if err != nil {
		return configuration, err
	}

	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&configuration); err != nil {
		return configuration, err
	}

	return configuration, nil
}

// ReadAttributeQuery reads the specified config into a AttributeQuery type.
func ReadAttributeQuery(configPath string) (v1alpha12.AttributeQuery, error) {
	var configuration v1alpha12.AttributeQuery
	data, err := readInConfig(configPath, v1alpha12.AttributeQueryKind)
	if err != nil {
		return configuration, err
	}

	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&configuration); err != nil {
		return configuration, err
	}

	return configuration, nil
}

// readInConfig reads in the file from the path and checks the type.
func readInConfig(configPath, kind string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath.Clean(configPath))
	if err != nil {
		return nil, err
	}

	if data, err = yaml.YAMLToJSON(data); err != nil {
		return nil, err
	}

	typeMeta, err := getTypeMeta(data)
	if err != nil {
		return nil, err
	}
	if typeMeta.Kind != kind {
		return nil, fmt.Errorf("config kind %s, does not match expected %s", typeMeta.Kind, kind)
	}
	return data, nil
}

// getTypeMeta retrieves TypeMeta information from the input.
func getTypeMeta(data []byte) (typeMeta v1alpha12.TypeMeta, err error) {
	if err := json.Unmarshal(data, &typeMeta); err != nil {
		return typeMeta, fmt.Errorf("get type meta: %v", err)
	}
	return typeMeta, nil
}
