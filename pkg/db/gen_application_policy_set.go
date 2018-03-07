// nolint
package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

// ApplicationPolicySetFields is db columns for ApplicationPolicySet
var ApplicationPolicySetFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"user_visible",
	"permissions_owner_access",
	"permissions_owner",
	"other_access",
	"group_access",
	"group",
	"last_modified",
	"enable",
	"description",
	"creator",
	"created",
	"fq_name",
	"display_name",
	"configuration_version",
	"key_value_pair",
	"all_applications",
}

// ApplicationPolicySetRefFields is db reference fields for ApplicationPolicySet
var ApplicationPolicySetRefFields = map[string][]string{

	"firewall_policy": []string{
		// <schema.Schema Value>
		"sequence",
	},

	"global_vrouter_config": []string{
	// <schema.Schema Value>

	},
}

// ApplicationPolicySetBackRefFields is db back reference fields for ApplicationPolicySet
var ApplicationPolicySetBackRefFields = map[string][]string{}

// ApplicationPolicySetParentTypes is possible parents for ApplicationPolicySet
var ApplicationPolicySetParents = []string{

	"project",

	"policy_management",
}

// CreateApplicationPolicySet inserts ApplicationPolicySet to DB
// nolint
func (db *DB) createApplicationPolicySet(
	ctx context.Context,
	request *models.CreateApplicationPolicySetRequest) error {
	qb := db.queryBuilders["application_policy_set"]
	tx := GetTransaction(ctx)
	model := request.ApplicationPolicySet
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		bool(model.GetIDPerms().GetUserVisible()),
		int(model.GetIDPerms().GetPermissions().GetOwnerAccess()),
		string(model.GetIDPerms().GetPermissions().GetOwner()),
		int(model.GetIDPerms().GetPermissions().GetOtherAccess()),
		int(model.GetIDPerms().GetPermissions().GetGroupAccess()),
		string(model.GetIDPerms().GetPermissions().GetGroup()),
		string(model.GetIDPerms().GetLastModified()),
		bool(model.GetIDPerms().GetEnable()),
		string(model.GetIDPerms().GetDescription()),
		string(model.GetIDPerms().GetCreator()),
		string(model.GetIDPerms().GetCreated()),
		common.MustJSON(model.GetFQName()),
		string(model.GetDisplayName()),
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()),
		bool(model.GetAllApplications()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.FirewallPolicyRefs {

		if ref.Attr == nil {
			ref.Attr = &models.FirewallSequence{}
		}

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("firewall_policy"), model.UUID, ref.UUID, string(ref.Attr.GetSequence()))
		if err != nil {
			return errors.Wrap(err, "FirewallPolicyRefs create failed")
		}
	}

	for _, ref := range model.GlobalVrouterConfigRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("global_vrouter_config"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "GlobalVrouterConfigRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "application_policy_set",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "application_policy_set", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanApplicationPolicySet(values map[string]interface{}) (*models.ApplicationPolicySet, error) {
	m := models.MakeApplicationPolicySet()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = common.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = common.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = common.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = common.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = common.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = common.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = common.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = common.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = common.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["all_applications"]; ok {

		m.AllApplications = common.InterfaceToBool(value)

	}

	if value, ok := values["ref_firewall_policy"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ApplicationPolicySetFirewallPolicyRef{}
			referenceModel.UUID = uuid
			m.FirewallPolicyRefs = append(m.FirewallPolicyRefs, referenceModel)

			attr := models.MakeFirewallSequence()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_global_vrouter_config"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ApplicationPolicySetGlobalVrouterConfigRef{}
			referenceModel.UUID = uuid
			m.GlobalVrouterConfigRefs = append(m.GlobalVrouterConfigRefs, referenceModel)

		}
	}

	return m, nil
}

// ListApplicationPolicySet lists ApplicationPolicySet with list spec.
func (db *DB) listApplicationPolicySet(ctx context.Context, request *models.ListApplicationPolicySetRequest) (response *models.ListApplicationPolicySetResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["application_policy_set"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.ApplicationPolicySet{}

	if spec.ParentFQName != nil {
		parentMetaData, err := db.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}
	query, columns, values := qb.ListQuery(auth, spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "row error")
	}

	for rows.Next() {
		valuesMap := map[string]interface{}{}
		values := make([]interface{}, len(columns))
		valuesPointers := make([]interface{}, len(columns))
		for _, index := range columns {
			valuesPointers[index] = &values[index]
		}
		if err := rows.Scan(valuesPointers...); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		for column, index := range columns {
			val := valuesPointers[index].(*interface{})
			valuesMap[column] = *val
		}
		m, err := scanApplicationPolicySet(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListApplicationPolicySetResponse{
		ApplicationPolicySets: result,
	}
	return response, nil
}

// UpdateApplicationPolicySet updates a resource
func (db *DB) updateApplicationPolicySet(
	ctx context.Context,
	request *models.UpdateApplicationPolicySetRequest,
) error {
	//TODO
	return nil
}

// DeleteApplicationPolicySet deletes a resource
func (db *DB) deleteApplicationPolicySet(
	ctx context.Context,
	request *models.DeleteApplicationPolicySetRequest) error {
	qb := db.queryBuilders["application_policy_set"]

	selectQuery := qb.SelectForDeleteQuery()
	deleteQuery := qb.DeleteQuery()

	var err error
	var count int
	uuid := request.ID
	tx := GetTransaction(ctx)
	auth := common.GetAuthCTX(ctx)
	if auth.IsAdmin() {
		row := tx.QueryRowContext(ctx, selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	err = db.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateApplicationPolicySet handle a Create API
// nolint
func (db *DB) CreateApplicationPolicySet(
	ctx context.Context,
	request *models.CreateApplicationPolicySetRequest) (*models.CreateApplicationPolicySetResponse, error) {
	model := request.ApplicationPolicySet
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createApplicationPolicySet(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "application_policy_set",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateApplicationPolicySetResponse{
		ApplicationPolicySet: request.ApplicationPolicySet,
	}, nil
}

//UpdateApplicationPolicySet handles a Update request.
func (db *DB) UpdateApplicationPolicySet(
	ctx context.Context,
	request *models.UpdateApplicationPolicySetRequest) (*models.UpdateApplicationPolicySetResponse, error) {
	model := request.ApplicationPolicySet
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateApplicationPolicySet(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "application_policy_set",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateApplicationPolicySetResponse{
		ApplicationPolicySet: model,
	}, nil
}

//DeleteApplicationPolicySet delete a resource.
func (db *DB) DeleteApplicationPolicySet(ctx context.Context, request *models.DeleteApplicationPolicySetRequest) (*models.DeleteApplicationPolicySetResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteApplicationPolicySet(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteApplicationPolicySetResponse{
		ID: request.ID,
	}, nil
}

//GetApplicationPolicySet a Get request.
func (db *DB) GetApplicationPolicySet(ctx context.Context, request *models.GetApplicationPolicySetRequest) (response *models.GetApplicationPolicySetResponse, err error) {
	spec := &models.ListSpec{
		Limit:  1,
		Detail: true,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListApplicationPolicySetRequest{
		Spec: spec,
	}
	var result *models.ListApplicationPolicySetResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listApplicationPolicySet(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ApplicationPolicySets) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetApplicationPolicySetResponse{
		ApplicationPolicySet: result.ApplicationPolicySets[0],
	}
	return response, nil
}

//ListApplicationPolicySet handles a List service Request.
// nolint
func (db *DB) ListApplicationPolicySet(
	ctx context.Context,
	request *models.ListApplicationPolicySetRequest) (response *models.ListApplicationPolicySetResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listApplicationPolicySet(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}