package controllers

import (
	"app/base/core"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdvisorySystemsExportJSON(t *testing.T) {
	core.SetupTest(t)
	w := CreateRequestRouterWithPath("GET", "/:advisory_id", "RH-1", "", nil, "application/json",
		AdvisorySystemsExportHandler)

	var output []SystemDBLookupExtended
	CheckResponse(t, w, http.StatusOK, &output)
	assert.Equal(t, 6, len(output))
	assert.Equal(t, output[0].ID, "00000000-0000-0000-0000-000000000001")
	assert.Equal(t, SystemTagsList{{"k1", "ns1", "val1"}, {"k2", "ns1", "val2"}}, output[0].SystemItemAttributes.Tags)
}

func TestAdvisorySystemsExportCSV(t *testing.T) {
	core.SetupTest(t)
	w := CreateRequestRouterWithPath("GET", "/:advisory_id", "RH-1", "", nil, "text/csv", AdvisorySystemsExportHandler)

	assert.Equal(t, http.StatusOK, w.Code)
	body := w.Body.String()
	lines := strings.Split(body, "\r\n")

	assert.Equal(t, 8, len(lines))
	assert.Equal(t,
		"display_name,last_upload,stale,os,rhsm,stale_timestamp,stale_warning_timestamp,culled_timestamp,created,tags,"+
			"groups,baseline_id,baseline_name,template_name,template_uuid,status,satellite_managed,built_pkgcache,id", lines[0])

	assert.Equal(t, "00000000-0000-0000-0000-000000000001,2020-09-22T16:00:00Z,false,RHEL 8.10,8.10,2018-08-26T16:00:00Z,"+
		"2018-09-02T16:00:00Z,2018-09-09T16:00:00Z,2018-08-26T16:00:00Z,\"[{'key':'k1','namespace':'ns1','value':'val1'},"+
		"{'key':'k2','namespace':'ns1','value':'val2'}]\",\"[{'id':'inventory-group-1','name':'group1'}]\","+
		"1,baseline_1-1,temp1-1,99900000-0000-0000-0000-000000000001,Installable,false,false,"+
		"00000000-0000-0000-0000-000000000001",
		lines[1])
}
