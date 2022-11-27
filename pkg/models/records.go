package models

import "time"

type RecordsStruct struct {
	Records []struct {
		ID          string    `json:"id"`
		CreatedTime time.Time `json:"createdTime"`
		Fields      struct {
			Field4   string `json:"Field 4"`
			Field3   string `json:"Field 3"`
			Projects string `json:"Projects"`
			EndDate  string `json:"End date"`
		} `json:"fields"`
	} `json:"records"`
}

type Record struct {
	ID          string    `json:"id"`
	CreatedTime time.Time `json:"createdTime"`
	Fields      struct {
		Field4   string `json:"Field 4"`
		Field3   string `json:"Field 3"`
		Projects string `json:"Projects"`
		EndDate  string `json:"End date"`
	} `json:"fields"`
}

type CommentStruct struct {
	Comments []struct {
		Author struct {
			Email string `json:"email"`
			ID    string `json:"id"`
			Name  string `json:"name"`
		} `json:"author"`
		CreatedTime     time.Time   `json:"createdTime"`
		ID              string      `json:"id"`
		LastUpdatedTime interface{} `json:"lastUpdatedTime"`
		Text            string      `json:"text"`
		Mentioned       struct {
			UsrL2PNC5O3H4LBEi struct {
				DisplayName string `json:"displayName"`
				Email       string `json:"email"`
				ID          string `json:"id"`
				Type        string `json:"type"`
			} `json:"usrL2PNC5o3H4lBEi"`
		} `json:"mentioned,omitempty"`
	} `json:"comments"`
	Offset interface{} `json:"offset"`
}

type TableStruct struct {
	Description string `json:"description"`
	Fields      []struct {
		Description string `json:"description,omitempty"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"fields"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	PrimaryFieldID string `json:"primaryFieldId"`
	Views          []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"views"`
}

type BasesStruct struct {
	Bases []struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		PermissionLevel string `json:"permissionLevel"`
	} `json:"bases"`
	Offset string `json:"offset"`
}
