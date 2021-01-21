package fields

// Columns which the report will have
type ReportColumns struct {
	// operator name
	Operator string
	// timestamp at which csv was created
	CreatedAt string
	// csv provider
	Company string
	// version of SDK from stamps
	SDKVersion string
	// operatortype from stamps
	OperatorType string
}

// Inputs obtained from index image label
type Inputs struct {
	Path    string
	Source  string
	Version string
}

// Type of operator - go/ansible/helm
type OperatorType struct {
	Go      int
	Ansible int
	Helm    int
}

type SDKVersion struct {
	// refers to all versions <1.0
	PreMajorRel int
	// refers to all versions >1.0
	PostMajorel int
}

// Operator data to be collected after finding the dump
type OperatorData struct {
	CommunityOperators   map[string]ReportColumns
	CertifiedOperators   (map[string]ReportColumns)
	MarketplaceOperators (map[string]ReportColumns)
	OperatorHub          (map[string]ReportColumns)
	RedHatOperators      (map[string]ReportColumns)
	ProdOperators        (map[string]ReportColumns)
	SDKVersionCount      SDKVersion
	OperatorTypeCount    OperatorType
}
