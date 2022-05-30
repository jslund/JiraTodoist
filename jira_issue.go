package function

type JiraIssue struct {
	Timestamp          int64  `json:"timestamp"`
	WebhookEvent       string `json:"webhookEvent"`
	IssueEventTypeName string `json:"issue_event_type_name"`
	User               struct {
		Self         string `json:"self"`
		Name         string `json:"name"`
		Key          string `json:"key"`
		AccountID    string `json:"accountId"`
		EmailAddress string `json:"emailAddress"`
		AvatarUrls   struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
		TimeZone    string `json:"timeZone"`
	} `json:"user"`
	Issue struct {
		ID     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
		Fields struct {
			Issuetype struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Description string `json:"description"`
				IconURL     string `json:"iconUrl"`
				Name        string `json:"name"`
				Subtask     bool   `json:"subtask"`
				AvatarID    int    `json:"avatarId"`
			} `json:"issuetype"`
			Timespent interface{} `json:"timespent"`
			Project   struct {
				Self           string `json:"self"`
				ID             string `json:"id"`
				Key            string `json:"key"`
				Name           string `json:"name"`
				ProjectTypeKey string `json:"projectTypeKey"`
				AvatarUrls     struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
			} `json:"project"`
			FixVersions        []interface{} `json:"fixVersions"`
			Aggregatetimespent interface{}   `json:"aggregatetimespent"`
			Resolution         struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"resolution"`
			Resolutiondate string `json:"resolutiondate"`
			Workratio      int    `json:"workratio"`
			Watches        struct {
				Self       string `json:"self"`
				WatchCount int    `json:"watchCount"`
				IsWatching bool   `json:"isWatching"`
			} `json:"watches"`
			LastViewed       string        `json:"lastViewed"`
			Created          string        `json:"created"`
			Customfield10020 interface{}   `json:"customfield_10020"`
			Customfield10021 []interface{} `json:"customfield_10021"`
			Priority         struct {
				Self    string `json:"self"`
				IconURL string `json:"iconUrl"`
				Name    string `json:"name"`
				ID      string `json:"id"`
			} `json:"priority"`
			Customfield10024              interface{}   `json:"customfield_10024"`
			Labels                        []interface{} `json:"labels"`
			Customfield10016              interface{}   `json:"customfield_10016"`
			Customfield10017              string        `json:"customfield_10017"`
			Customfield10018              interface{}   `json:"customfield_10018"`
			Customfield10019              string        `json:"customfield_10019"`
			Aggregatetimeoriginalestimate interface{}   `json:"aggregatetimeoriginalestimate"`
			Timeestimate                  interface{}   `json:"timeestimate"`
			Versions                      []interface{} `json:"versions"`
			Issuelinks                    []interface{} `json:"issuelinks"`
			Assignee                      interface{}   `json:"assignee"`
			Updated                       string        `json:"updated"`
			Status                        struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
					Name      string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
			Components           []interface{} `json:"components"`
			Timeoriginalestimate interface{}   `json:"timeoriginalestimate"`
			Description          string        `json:"description"`
			Customfield10013     interface{}   `json:"customfield_10013"`
			Customfield10014     interface{}   `json:"customfield_10014"`
			Timetracking         struct {
			} `json:"timetracking"`
			Customfield10015      interface{}   `json:"customfield_10015"`
			Customfield10005      interface{}   `json:"customfield_10005"`
			Customfield10006      interface{}   `json:"customfield_10006"`
			Customfield10007      interface{}   `json:"customfield_10007"`
			Security              interface{}   `json:"security"`
			Customfield10008      interface{}   `json:"customfield_10008"`
			Aggregatetimeestimate interface{}   `json:"aggregatetimeestimate"`
			Customfield10009      interface{}   `json:"customfield_10009"`
			Attachment            []interface{} `json:"attachment"`
			Summary               string        `json:"summary"`
			Creator               struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				AccountID    string `json:"accountId"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"creator"`
			Subtasks []interface{} `json:"subtasks"`
			Reporter struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				AccountID    string `json:"accountId"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"reporter"`
			Aggregateprogress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"aggregateprogress"`
			Customfield10000 string      `json:"customfield_10000"`
			Customfield10001 interface{} `json:"customfield_10001"`
			Customfield10002 interface{} `json:"customfield_10002"`
			Customfield10003 interface{} `json:"customfield_10003"`
			Customfield10004 interface{} `json:"customfield_10004"`
			Environment      interface{} `json:"environment"`
			Duedate          interface{} `json:"duedate"`
			Progress         struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"progress"`
			Votes struct {
				Self     string `json:"self"`
				Votes    int    `json:"votes"`
				HasVoted bool   `json:"hasVoted"`
			} `json:"votes"`
		} `json:"fields"`
	} `json:"issue"`
	Changelog struct {
		ID    string `json:"id"`
		Items []struct {
			Field      string      `json:"field"`
			Fieldtype  string      `json:"fieldtype"`
			FieldID    string      `json:"fieldId"`
			From       interface{} `json:"from"`
			FromString interface{} `json:"fromString"`
			To         string      `json:"to"`
		} `json:"items"`
	} `json:"changelog"`
}
