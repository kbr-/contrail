package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRoutingInstance makes RoutingInstance
// nolint
func MakeRoutingInstance() *RoutingInstance {
	return &RoutingInstance{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		ConfigurationVersion: 0,
	}
}

// MakeRoutingInstance makes RoutingInstance
// nolint
func InterfaceToRoutingInstance(i interface{}) *RoutingInstance {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RoutingInstance{
		//TODO(nati): Apply default
		UUID:                 common.InterfaceToString(m["uuid"]),
		ParentUUID:           common.InterfaceToString(m["parent_uuid"]),
		ParentType:           common.InterfaceToString(m["parent_type"]),
		FQName:               common.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          common.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion: common.InterfaceToInt64(m["configuration_version"]),
	}
}

// MakeRoutingInstanceSlice() makes a slice of RoutingInstance
// nolint
func MakeRoutingInstanceSlice() []*RoutingInstance {
	return []*RoutingInstance{}
}

// InterfaceToRoutingInstanceSlice() makes a slice of RoutingInstance
// nolint
func InterfaceToRoutingInstanceSlice(i interface{}) []*RoutingInstance {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RoutingInstance{}
	for _, item := range list {
		result = append(result, InterfaceToRoutingInstance(item))
	}
	return result
}