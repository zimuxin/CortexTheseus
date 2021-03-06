// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package inference

import (
	"encoding/json"
	"errors"
)

// MarshalJSON marshals as JSON.
func (i IHWork) MarshalJSON() ([]byte, error) {
	type IHWork struct {
		Type         InferType `json:"type"  gencodec:"required"`
		Model        string    `json:"model" gencodec:"required"`
		Input        string    `json:"input" gencodec:"required"`
		ModelSize    uint64    `json:"modelSize"`
		InputSize    uint64    `json:"inputSize"`
		CvmVersion   int       `json:"cvm_version"`
		CvmNetworkId int64     `json:"cvm_networkid"`
	}
	var enc IHWork
	enc.Type = i.Type
	enc.Model = i.Model
	enc.Input = i.Input
	enc.ModelSize = i.ModelSize
	enc.InputSize = i.InputSize
	enc.CvmVersion = i.CvmVersion
	enc.CvmNetworkId = i.CvmNetworkId
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (i *IHWork) UnmarshalJSON(input []byte) error {
	type IHWork struct {
		Type         *InferType `json:"type"  gencodec:"required"`
		Model        *string    `json:"model" gencodec:"required"`
		Input        *string    `json:"input" gencodec:"required"`
		ModelSize    *uint64    `json:"modelSize"`
		InputSize    *uint64    `json:"inputSize"`
		CvmVersion   *int       `json:"cvm_version"`
		CvmNetworkId *int64     `json:"cvm_networkid"`
	}
	var dec IHWork
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Type == nil {
		return errors.New("missing required field 'type' for IHWork")
	}
	i.Type = *dec.Type
	if dec.Model == nil {
		return errors.New("missing required field 'model' for IHWork")
	}
	i.Model = *dec.Model
	if dec.Input == nil {
		return errors.New("missing required field 'input' for IHWork")
	}
	i.Input = *dec.Input
	if dec.ModelSize != nil {
		i.ModelSize = *dec.ModelSize
	}
	if dec.InputSize != nil {
		i.InputSize = *dec.InputSize
	}
	if dec.CvmVersion != nil {
		i.CvmVersion = *dec.CvmVersion
	}
	if dec.CvmNetworkId != nil {
		i.CvmNetworkId = *dec.CvmNetworkId
	}
	return nil
}
