// Code generated by generate_types.go; DO NOT EDIT.

package greenpass

import (
	"fmt"
)

type DiseaseAgentTargeted string

func (dat DiseaseAgentTargeted) String() string {
	name, ok := diseaseAgentTargetedNames[dat]
	if !ok {
		return fmt.Sprintf("unknown('%s')", dat)
	}
	return name
}

var diseaseAgentTargetedNames = map[DiseaseAgentTargeted]string{
	"840539006": "COVID-19",
}
