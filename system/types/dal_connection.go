package types

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/cortezaproject/corteza-server/pkg/sql"
	"time"

	"github.com/cortezaproject/corteza-server/pkg/dal"
	"github.com/cortezaproject/corteza-server/pkg/dal/capabilities"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/pkg/geolocation"
)

type (
	DalConnection struct {
		ID     uint64 `json:"connectionID,string"`
		Name   string `json:"name"`
		Handle string `json:"handle"`

		Type string `json:"type"`

		Location         geolocation.Full `json:"location"`
		Ownership        string           `json:"ownership"`
		SensitivityLevel uint64           `json:"sensitivityLevel,string,omitempty"`

		Issues []string `json:"issues,omitempty" db:"-"`

		Config       ConnectionConfig       `json:"config"`
		Capabilities ConnectionCapabilities `json:"capabilities"`

		Labels map[string]string `json:"labels,omitempty"`

		CreatedAt time.Time  `json:"createdAt,omitempty"`
		CreatedBy uint64     `json:"createdBy,string" `
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		UpdatedBy uint64     `json:"updatedBy,string,omitempty" `
		DeletedAt *time.Time `json:"deletedAt,omitempty"`
		DeletedBy uint64     `json:"deletedBy,string,omitempty" `
	}

	ConnectionCapabilities struct {
		Enforced    capabilities.Set `json:"enforced"`
		Supported   capabilities.Set `json:"supported"`
		Unsupported capabilities.Set `json:"unsupported"`
		Enabled     capabilities.Set `json:"enabled"`
	}

	ConnectionConfig struct {
		DefaultModelIdent     string `json:"defaultModelIdent"`
		DefaultAttributeIdent string `json:"defaultAttributeIdent"`

		DefaultPartitionFormat string `json:"defaultPartitionFormat"`

		PartitionFormatValidator string `json:"partitionFormatValidator"`

		Connection dal.ConnectionParams `json:"connection"`
	}

	DalConnectionFilter struct {
		ConnectionID []uint64 `json:"connectionID,string"`
		Handle       string   `json:"handle"`
		Type         string   `json:"type"`

		Deleted filter.State `json:"deleted"`

		// Check fn is called by store backend for each resource found function can
		// modify the resource and return false if store should not return it
		//
		// Store then loads additional resources to satisfy the paging parameters
		Check func(*DalConnection) (bool, error) `json:"-"`

		// Standard helpers for paging and sorting
		filter.Paging
	}
)

var (
	// Used to identify the primary DAL connection instead of an extra flag
	DalPrimaryConnectionResourceType = "corteza::system:primary-dal-connection"
	DalPrimaryConnectionHandle       = "primary-database"
)

func (c DalConnection) ActiveCapabilities() capabilities.Set {
	return c.Capabilities.Supported.
		Union(c.Capabilities.Enforced).
		Union(c.Capabilities.Enabled)
}

func (c DalConnection) HasIssues() bool {
	return len(c.Issues) > 0
}

func ParseConnectionConfig(ss []string) (m ConnectionConfig, err error) {
	if len(ss) == 0 {
		return
	}

	err = json.Unmarshal([]byte(ss[0]), &m)
	return
}

func ParseConnectionCapabilities(ss []string) (m ConnectionCapabilities, err error) {
	if len(ss) == 0 {
		return
	}

	err = json.Unmarshal([]byte(ss[0]), &m)
	return
}

func (nm *ConnectionConfig) Scan(src any) error          { return sql.ParseJSON(src, nm) }
func (nm ConnectionConfig) Value() (driver.Value, error) { return json.Marshal(nm) }

func (nm *ConnectionCapabilities) Scan(src any) error          { return sql.ParseJSON(src, nm) }
func (nm ConnectionCapabilities) Value() (driver.Value, error) { return json.Marshal(nm) }