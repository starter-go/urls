package entity

import "github.com/starter-go/urls/src/main/golang/lib/data/dxo"

// GetDataGroupInfo  ...
func GetDataGroupInfo() dxo.DataGroupInfo {
	return new(dgInfo)
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamePrefix string

type dgInfo struct{}

func (inst *dgInfo) Prototypes() []any {
	list := make([]any, 0)
	list = append(list, new(ShortLink))
	return list
}

func (inst *dgInfo) SetTableNamePrefix(prefix string) {
	if theTableNamePrefix == "" {
		theTableNamePrefix = prefix
	}
}

////////////////////////////////////////////////////////////////////////////////

// TableName 。。。
func (ShortLink) TableName() string {
	return theTableNamePrefix + "short_links"
}
