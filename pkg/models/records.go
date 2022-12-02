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

type Field struct {
	Field4   string `json:"Field 4"`
	Field3   string `json:"Field 3"`
	Projects string `json:"Projects"`
	EndDate  string `json:"End date"`
}

type Record struct {
	ID          string    `json:"id"`
	CreatedTime time.Time `json:"createdTime"`
	Fields      Field     `json:"fields"`
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
			UserID struct {
				DisplayName string `json:"displayName"`
				Email       string `json:"email"`
				ID          string `json:"id"`
				Type        string `json:"type"`
			} `json:"userId"`
		} `json:"mentioned,omitempty"`
	} `json:"comments"`
	Offset interface{} `json:"offset"`
}

type BasesStruct struct {
	Bases []struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		PermissionLevel string `json:"permissionLevel"`
	} `json:"bases"`
	Offset string `json:"offset"`
}
