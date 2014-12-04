package client

import (
	"encoding/binary"
)


func (patient *Patient) String(field string) string {
	return string(patient.Data[field])
}

func (patient *Patient) SetString(field string, value string) {
	patient.Data[field] = []byte(value)
}

func (patient *Patient) SetUint32(field string, value uint32) {
	patient.Data[field] = make([]byte, 4)
	binary.BigEndian.PutUint32(patient.Data[field], value)
}

func (patient *Patient) Uint32(field string) uint32 {
	return binary.BigEndian.Uint32(patient.Data[field])
}
