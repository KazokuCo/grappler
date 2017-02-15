package cloudbuild

type PayloadLink struct {
	Method string `json:"method"`
	Href   string `json:"href"`
}

type Payload struct {
	ProjectName     string `json:"projectName"`
	BuildTargetName string `json:"buildTargetName"`
	ProjectGUID     string `json:"projectGuid"`
	OrgForeignKey   string `json:"orgForeignKey"`
	BuildNumber     int    `json:"buildNumber"`
	BuildStatus     string `json:"buildStatus"`
	StartedBy       string `json:"startedBy"`
	Platform        string `json:"platform"`
	Links           struct {
		APISelf          PayloadLink `json:"api_self"`
		DashboardURL     PayloadLink `json:"dashboard_url"`
		DashboardProject PayloadLink `json:"dashboard_project"`
		DashboardSummary PayloadLink `json:"dashboard_summary"`
		DashboardLog     PayloadLink `json:"dashboard_log"`
	} `json:"links"`
}
