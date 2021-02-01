package collector

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/operator-framework/api/pkg/operators/v1alpha1"
	"github.com/varshaprasad96/operator-sdk-data-collector/pkg/fields"
	"golang.org/x/mod/semver"
)

var (
	builder = "operators.operatorframework.io/builder"
	layout  = "operators.operatorframework.io/project_layout"

	communityOperators   = make(map[string]fields.ReportColumns)
	certifiedOperators   = make(map[string]fields.ReportColumns)
	marketplaceOperators = make(map[string]fields.ReportColumns)
	operatorHub          = make(map[string]fields.ReportColumns)
	redHatOperators      = make(map[string]fields.ReportColumns)
	prodOperators        = make(map[string]fields.ReportColumns)
	allOperators         = make(map[string]fields.ReportColumns)
	layoutStamp          = make(map[string]int)
	versionStamp         = make(map[string]int)
	// count of pre 1.0 and post 1.0 SDK versions
	preSDKMajorRel = 0
	postSDKMajRel  = 0
	// count of operator types
	goOp      = 0
	ansibleOp = 0
	helmOp    = 0
)

const (
	source_redhat      = "redhat"
	source_community   = "community"
	source_marketplace = "marketplace"
	source_certified   = "certified"
	source_operatorhub = "operatorhub"
	source_prod        = "prod"
)

// Execupte SQL query and extract relvant data from database
func CollectDump(inputList []fields.Inputs) fields.OperatorData {

	for i := 0; i < len(inputList); i++ {
		db, err := sql.Open("sqlite3", inputList[i].Path)
		if err != nil {
			panic(err)
		}
		dump(db, inputList[i].Source)
	}

	calculateCounts(allOperators)

	return fields.OperatorData{
		CommunityOperators:   communityOperators,
		CertifiedOperators:   certifiedOperators,
		MarketplaceOperators: marketplaceOperators,
		OperatorHub:          operatorHub,
		RedHatOperators:      redHatOperators,
		ProdOperators:        prodOperators,
		AllOperators:         allOperators,
		LayoutData:           layoutStamp,
		VersionData:          versionStamp,
		SDKVersionCount: fields.SDKVersion{
			PreMajorRel: preSDKMajorRel,
			PostMajorel: postSDKMajRel,
		},
		OperatorTypeCount: fields.OperatorType{
			Go:      goOp,
			Ansible: ansibleOp,
			Helm:    helmOp,
		},
	}

}

func calculateCounts(overallData map[string]fields.ReportColumns) {
	for _, data := range overallData {
		getOperatorType(data.OperatorType)
		getReleaseCount(data.SDKVersion)
	}

	for _, v := range overallData {
		val, ok := layoutStamp[v.OperatorType]
		if ok {
			layoutStamp[v.OperatorType] = val + 1
		} else {
			layoutStamp[v.OperatorType] = 1
		}
	}

	for _, v := range overallData {
		val, ok := versionStamp[v.SDKVersion]
		if ok {
			versionStamp[v.SDKVersion] = val + 1
		} else {
			versionStamp[v.SDKVersion] = 1
		}
	}

}

func dump(db *sql.DB, sourceDescription string) {

	// execute db query
	row, err := db.Query("SELECT name, csv, bundlepath FROM operatorbundle where csv is not null  order by name")
	if err != nil {
		panic(err)
	}

	defer row.Close()

	for row.Next() {
		var name string
		var csv string
		var bundlepath string
		var operatorType string
		var sdkVersion string
		var csvStruct v1alpha1.ClusterServiceVersion

		row.Scan(&name, &csv, &bundlepath)
		err := json.Unmarshal([]byte(csv), &csvStruct)
		if err != nil {
			fmt.Printf("error unmarshalling csv %s\n", err.Error())
		}

		createdAt := csvStruct.ObjectMeta.Annotations["createdAt"]
		companyName := csvStruct.Spec.Provider.Name

		annotations := csvStruct.GetAnnotations()

		_, ok := annotations[builder]
		if ok {
			sdkVersion, operatorType = annotations[builder], annotations[layout]
		}

		op := fields.ReportColumns{
			Operator:     getName(name),
			CreatedAt:    createdAt,
			Company:      companyName,
			SDKVersion:   sdkVersion,
			OperatorType: operatorType,
			CSVName:      csvStruct.Name,
		}

		allOperators[op.Operator] = op

		switch sourceDescription {
		case source_redhat:
			redHatOperators[op.Operator] = op
		case source_community:
			communityOperators[op.Operator] = op
		case source_marketplace:
			marketplaceOperators[op.Operator] = op
		case source_prod:
			prodOperators[op.Operator] = op
		case source_certified:
			certifiedOperators[op.Operator] = op
		case source_operatorhub:
			operatorHub[op.Operator] = op
		}
	}

}

func getName(operatorName string) string {
	if operatorName == "" {
		fmt.Printf("error in the operator name")
	}

	return strings.Split(operatorName, ".")[0]
}

func getOperatorType(operatorType string) {
	if operatorType == "" {
		return
	}
	opType := strings.Split(operatorType, ".")[0]

	if opType == "go" {
		goOp++
	} else if opType == "ansible" {
		ansibleOp++
	} else if opType == "helm" {
		helmOp++
	}
}

func getReleaseCount(sdkVersion string) {
	if sdkVersion == "" {
		return
	}
	ver := strings.Replace(sdkVersion, "operator-sdk-", "", -1)
	c := semver.Compare(ver, "v1.0.0")
	if c >= 0 {
		// post release
		postSDKMajRel++
	} else {
		// pre release
		preSDKMajorRel++
	}
}
