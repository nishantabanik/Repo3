package sonarcloud

/*
 * Make sure the fields are public (First letter Uppercase),
 * otherwise the JSON serialization will fail.
 */

// Version for unmarshalling
type Version struct {
	PageIndex int64 `json:"pageIndex"`
}

// ErrorResponse struct
type ErrorResponse struct {
	Errors []ErrorMessage `json:"errors,omitempty"`
}

// ErrorMessage struct
type ErrorMessage struct {
	Message string `json:"msg,omitempty"`
}

// GetQualityGate for unmarshalling response body of quality gate get
type GetQualityGate struct {
	ID         int64                                `json:"id"`
	Name       string                               `json:"name"`
	Conditions []CreateQualityGateConditionResponse `json:"conditions"`
	IsBuiltIn  bool                                 `json:"isBuiltIn"`
	Actions    QualityGateActions                   `json:"actions"`
}

// QualityGateActions used in GetQualityGate
type QualityGateActions struct {
	Rename            bool `json:"rename"`
	SetAsDefault      bool `json:"setAsDefault"`
	Copy              bool `json:"copy"`
	AssociateProjects bool `json:"associateProjects"`
	Delete            bool `json:"delete"`
	ManageConditions  bool `json:"manageConditions"`
}

// CreateQualityGateResponse for unmarshalling response body of quality gate creation
type CreateQualityGateResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// CreateQualityGateConditionResponse for unmarshalling response body of condition creation
type CreateQualityGateConditionResponse struct {
	ID      int64  `json:"id"`
	Metric  string `json:"metric"`
	OP      string `json:"op"`
	Error   string `json:"error"`
	Warning string `json:"warning"`
}

// CreateProjectResponse for unmarshalling response body of project creation
type CreateProjectResponse struct {
	Project Project `json:"project"`
}

// Project used in CreateProjectResponse
type Project struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Qualifier string `json:"qualifier"`
}

// GetProject for unmarshalling response body from getting project details
type GetProject struct {
	Paging     Paging              `json:"paging"`
	Components []ProjectComponents `json:"components"`
}

// User struct
type User struct {
	Login       string   `json:"login,omitempty"`
	Name        string   `json:"name,omitempty"`
	Email       string   `json:"email,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	IsActive    bool     `json:"active,omitempty"`
	IsLocal     bool     `json:"local,omitempty"`
}

// GetUser for unmarshalling response body where users are retured
type GetUser struct {
	Paging Paging `json:"paging"`
	Users  []User `json:"users"`
}

// CreateUserResponse struct
type CreateUserResponse struct {
	User User `json:"user"`
}

// GetTokens represents the structure
type GetTokens struct {
	Login  string  `json:"login,omitempty"`
	Tokens []Token `json:"userTokens,omitempty"`
}

// Token struct
type Token struct {
	Login string `json:"login,omitempty"`
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
}

// CreateGroupResponse for unmarshalling response body of group creation
type CreateGroupResponse struct {
	Group Group `json:"group"`
}

// Group struct
type Group struct {
	ID           int      `json:"id,omitempty"`
	Organization string   `json:"organization,omitempty"`
	Name         string   `json:"name,omitempty"`
	Description  string   `json:"description,omitempty"`
	MembersCount int      `json:"membersCount,omitempty"`
	IsDefault    bool     `json:"default,omitempty"`
	Permissions  []string `json:"permissions,omitempty"`
}

// GroupPermission struct
type GroupPermission struct {
	Name        string   `json:"name,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// GetGroup for unmarshalling response body from getting group details
type GetGroup struct {
	Paging Paging  `json:"paging"`
	Groups []Group `json:"groups"`
}

// GetGroupPermissions struct
type GetGroupPermissions struct {
	Paging Paging            `json:"paging"`
	Groups []GroupPermission `json:"groups"`
}

// CreatePermissionTemplateResponse struct
type CreatePermissionTemplateResponse struct {
	PermissionTemplate PermissionTemplate `json:"permissionTemplate"`
}

// GetPermissionTemplates struct
type GetPermissionTemplates struct {
	Paging              Paging               `json:"paging"`
	PermissionTemplates []PermissionTemplate `json:"permissionTemplates"`
}

// PermissionTemplate struct
type PermissionTemplate struct {
	ID                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	ProjectKeyPattern string `json:"projectKeyPattern,omitempty"`
}

// Paging used in /search API endpoints
type Paging struct {
	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
	Total     int64 `json:"total"`
}

// ProjectComponents used in GetProject
type ProjectComponents struct {
	Organization     string `json:"organization"`
	Key              string `json:"key"`
	Name             string `json:"name"`
	Qualifier        string `json:"qualifier"`
	Visibility       string `json:"visibility"`
	LastAnalysisDate string `json:"lastAnalysisDate"`
	Revision         string `json:"revision"`
}

// GetQualityGateAssociation for unmarshalling response body from getting quality gate association
type GetQualityGateAssociation struct {
	Paging  Paging                              `json:"paging"`
	Results []GetQualityGateAssociationProjects `json:"results"`
}

// GetQualityGateAssociationProjects used in GetQualityGateAssociation
type GetQualityGateAssociationProjects struct {
	Name     string `json:"name"`
	Key      string `json:"key"`
	Selected bool   `json:"selected"`
}

// QualityProfile struct
type QualityProfile struct {
	IsDefault    bool   `json:"isDefault"`
	IsInherited  bool   `json:"isInherited"`
	Language     string `json:"language"`
	LanguageName string `json:"languageName"`
	Name         string `json:"name"`
	Key          string `json:"key"`
}

// CreateQualityProfileResponse for unmarshalling response body from creating quality profiles
type CreateQualityProfileResponse struct {
        ID       int64          `json:"id"`
	Profile  QualityProfile `json:"profile"`
	Warnings []string       `json:"warnings"`
}

// GetQualityProfileList for unmarshalling response body of quality gate get
type GetQualityProfileList struct {
	Profiles []GetQualityProfile `json:"profiles"`
}

// GetQualityProfile for unmarshalling response body of quality gate get
type GetQualityProfile struct {
	Key                       string                   `json:"key"`
	Name                      string                   `json:"name"`
	Language                  string                   `json:"language"`
	LanguageName              string                   `json:"languageName"`
	IsInherited               bool                     `json:"isInherited"`
	IsBuiltIn                 bool                     `json:"isBuiltIn"`
	ActiveRuleCount           int                      `json:"activeRuleCount"`
	ActiveDeprecatedRuleCount int                      `json:"activeDeprecatedRuleCount"`
	IsDefault                 bool                     `json:"isDefault"`
	RuleUpdatedAt             string                   `json:"ruleUpdatedAt"`
	LastUsed                  string                   `json:"lastUsed"`
	Actions                   GetQualityProfileActions `json:"actions"`
}

// GetQualityProfileActions for unmarshalling response body of quality gate get
type GetQualityProfileActions struct {
	Edit              bool `json:"edit"`
	SetAsDefault      bool `json:"setAsDefault"`
	Copy              bool `json:"copy"`
	Delete            bool `json:"delete"`
	AssociateProjects bool `json:"associateProjects"`
}

// GetQualityProfileProjectAssociation for unmarshalling response body from getting quality profile association
type GetQualityProfileProjectAssociation struct {
	Paging  Paging                                       `json:"paging"`
	Results []GetQualityProfileProjectAssociationResults `json:"results"`
}

// GetQualityProfileProjectAssociationResults used in GetQualityProfileProjectAssociation
type GetQualityProfileProjectAssociationResults struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Key      string `json:"key"`
	Selected bool   `json:"selected"`
}
