package request

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//

import (
	"encoding/json"
	"fmt"
	"github.com/cortezaproject/corteza/server/pkg/payload"
	"github.com/go-chi/chi/v5"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// dummy vars to prevent
// unused imports complain
var (
	_ = chi.URLParam
	_ = multipart.ErrMessageTooLarge
	_ = payload.ParseUint64s
	_ = strings.ToLower
	_ = io.EOF
	_ = fmt.Errorf
	_ = json.NewEncoder
)

type (
	// Internal API interface
	DalSchemaAlterationList struct {
		// AlterationID GET parameter
		//
		// Filter by alteration ID
		AlterationID []string

		// BatchID GET parameter
		//
		// Filter by batch ID
		BatchID []string

		// Resource GET parameter
		//
		// Search by resource
		Resource []string

		// ResourceType GET parameter
		//
		// Search by resource type
		ResourceType string

		// Kind GET parameter
		//
		// Search by kind
		Kind string

		// Deleted GET parameter
		//
		// Exclude (0, default), include (1) or return only (2) deleted alterations
		Deleted uint

		// Completed GET parameter
		//
		// Exclude (0, default), include (1) or return only (2) completed alterations
		Completed uint

		// Dismissed GET parameter
		//
		// Exclude (0, default), include (1) or return only (2) dismissed alterations
		Dismissed uint

		// IncTotal GET parameter
		//
		// Include total counter
		IncTotal bool
	}

	DalSchemaAlterationRead struct {
		// AlterationID PATH parameter
		//
		// Alteration ID
		AlterationID uint64 `json:",string"`
	}

	DalSchemaAlterationApply struct {
		// AlterationID GET parameter
		//
		// Alteration ID
		AlterationID []uint64
	}

	DalSchemaAlterationDismiss struct {
		// AlterationID GET parameter
		//
		// Alteration ID
		AlterationID []uint64
	}
)

// NewDalSchemaAlterationList request
func NewDalSchemaAlterationList() *DalSchemaAlterationList {
	return &DalSchemaAlterationList{}
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationList) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"alterationID": r.AlterationID,
		"batchID":      r.BatchID,
		"resource":     r.Resource,
		"resourceType": r.ResourceType,
		"kind":         r.Kind,
		"deleted":      r.Deleted,
		"completed":    r.Completed,
		"dismissed":    r.Dismissed,
		"incTotal":     r.IncTotal,
	}
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationList) GetAlterationID() []string {
	return r.AlterationID
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationList) GetBatchID() []string {
	return r.BatchID
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationList) GetResource() []string {
	return r.Resource
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationList) GetResourceType() string {
	return r.ResourceType
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationList) GetKind() string {
	return r.Kind
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationList) GetDeleted() uint {
	return r.Deleted
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationList) GetCompleted() uint {
	return r.Completed
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationList) GetDismissed() uint {
	return r.Dismissed
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationList) GetIncTotal() bool {
	return r.IncTotal
}

// Fill processes request and fills internal variables
func (r *DalSchemaAlterationList) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["alterationID[]"]; ok {
			r.AlterationID, err = val, nil
			if err != nil {
				return err
			}
		} else if val, ok := tmp["alterationID"]; ok {
			r.AlterationID, err = val, nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["batchID[]"]; ok {
			r.BatchID, err = val, nil
			if err != nil {
				return err
			}
		} else if val, ok := tmp["batchID"]; ok {
			r.BatchID, err = val, nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["resource[]"]; ok {
			r.Resource, err = val, nil
			if err != nil {
				return err
			}
		} else if val, ok := tmp["resource"]; ok {
			r.Resource, err = val, nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["resourceType"]; ok && len(val) > 0 {
			r.ResourceType, err = val[0], nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["kind"]; ok && len(val) > 0 {
			r.Kind, err = val[0], nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["deleted"]; ok && len(val) > 0 {
			r.Deleted, err = payload.ParseUint(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["completed"]; ok && len(val) > 0 {
			r.Completed, err = payload.ParseUint(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["dismissed"]; ok && len(val) > 0 {
			r.Dismissed, err = payload.ParseUint(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["incTotal"]; ok && len(val) > 0 {
			r.IncTotal, err = payload.ParseBool(val[0]), nil
			if err != nil {
				return err
			}
		}
	}

	return err
}

// NewDalSchemaAlterationRead request
func NewDalSchemaAlterationRead() *DalSchemaAlterationRead {
	return &DalSchemaAlterationRead{}
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationRead) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"alterationID": r.AlterationID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationRead) GetAlterationID() uint64 {
	return r.AlterationID
}

// Fill processes request and fills internal variables
func (r *DalSchemaAlterationRead) Fill(req *http.Request) (err error) {

	{
		var val string
		// path params

		val = chi.URLParam(req, "alterationID")
		r.AlterationID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewDalSchemaAlterationApply request
func NewDalSchemaAlterationApply() *DalSchemaAlterationApply {
	return &DalSchemaAlterationApply{}
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationApply) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"alterationID": r.AlterationID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationApply) GetAlterationID() []uint64 {
	return r.AlterationID
}

// Fill processes request and fills internal variables
func (r *DalSchemaAlterationApply) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["alterationID[]"]; ok {
			r.AlterationID, err = payload.ParseUint64s(val), nil
			if err != nil {
				return err
			}
		} else if val, ok := tmp["alterationID"]; ok {
			r.AlterationID, err = payload.ParseUint64s(val), nil
			if err != nil {
				return err
			}
		}
	}

	return err
}

// NewDalSchemaAlterationDismiss request
func NewDalSchemaAlterationDismiss() *DalSchemaAlterationDismiss {
	return &DalSchemaAlterationDismiss{}
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationDismiss) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"alterationID": r.AlterationID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r DalSchemaAlterationDismiss) GetAlterationID() []uint64 {
	return r.AlterationID
}

// Fill processes request and fills internal variables
func (r *DalSchemaAlterationDismiss) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["alterationID[]"]; ok {
			r.AlterationID, err = payload.ParseUint64s(val), nil
			if err != nil {
				return err
			}
		} else if val, ok := tmp["alterationID"]; ok {
			r.AlterationID, err = payload.ParseUint64s(val), nil
			if err != nil {
				return err
			}
		}
	}

	return err
}
