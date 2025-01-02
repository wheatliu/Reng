package iface

// Twitter User Info
type UserInfo struct {
	ID                  int64   `json:"id"`
	IDStr               string  `json:"id_str"`
	Name                string  `json:"name"`
	ScreenName          string  `json:"screen_name"`
	Location            string  `json:"location"`
	URL                 *string `json:"url"`
	Description         string  `json:"description"`
	Protected           bool    `json:"protected"`
	DefaultProfile      bool    `json:"default_profile"`
	DefaultProfileImage bool    `json:"default_profile_image"`
}

// Start Login Request / Response Data
type (
	StartLoginRequest struct {
		InputFlowData struct {
			FlowContext struct {
				DebugOverrides map[string]string `json:"debug_overrides"`
				StartLocation  struct {
					Location string `json:"location"`
				} `json:"start_location"`
			} `json:"flow_context"`
		} `json:"input_flow_data"`
		SubtaskVersions map[string]int `json:"subtask_versions"`
	}

	StartLoginResponse struct {
		FlowToken string `json:"flow_token"`
		Status    string `json:"status"`
		Subtasks  []struct {
			SubtaskId         string `json:"subtask_id"`
			JsInstrumentation struct {
				Url       string `json:"url"`
				TimeoutMs int    `json:"timeout_ms"`
				NextLink  struct {
					LinkType string `json:"link_type"`
					LinkId   string `json:"link_id"`
				} `json:"next_link"`
			} `json:"js_instrumentation"`
		} `json:"subtasks"`
	}
)

// Js Instrumentation Request / Response Data
type (
	Subtask struct {
		SubtaskID             string        `json:"subtask_id"`
		SettingsList          *SettingsList `json:"settings_list,omitempty"`
		SubtaskBackNavigation string        `json:"subtask_back_navigation"`
	}

	SettingsList struct {
		Settings []Setting `json:"settings"`
	}

	Setting struct {
		ValueType       string    `json:"value_type"`
		ValueIdentifier string    `json:"value_identifier"`
		ValueData       ValueData `json:"value_data"`
	}

	ValueData struct {
		Button *Button `json:"button,omitempty"`
	}

	Button struct {
		NavigationLink NavigationLink `json:"navigation_link"`
	}

	NavigationLink struct {
		LinkType  string `json:"link_type"`
		LinkID    string `json:"link_id"`
		Label     string `json:"label,omitempty"`
		SubtaskID string `json:"subtask_id,omitempty"`
		URL       string `json:"url,omitempty"`
	}

	JsInstrumentation struct {
		Response string `json:"response"`
		Link     string `json:"link"`
	}

	JsInstrumentationSubtaskInput struct {
		SubtaskId         string             `json:"subtask_id"`
		JsInstrumentation *JsInstrumentation `json:"js_instrumentation"`
	}

	JsInstrumentationRequest struct {
		FlowToken     string                           `json:"flow_token"`
		SubtaskInputs []*JsInstrumentationSubtaskInput `json:"subtask_inputs"`
	}

	JsInstrumentationResponse struct {
		FlowToken string    `json:"flow_token"`
		Status    string    `json:"status"`
		Subtasks  []Subtask `json:"subtasks"`
	}
)

// Enter User Identifier Request / Response Data
type (
	TextData struct {
		Result string `json:"result"`
	}
	SettingResponseData struct {
		TextData *TextData `json:"text_data"`
	}
	SettingResponse struct {
		Key          string               `json:"key"`
		ResponseData *SettingResponseData `json:"response_data"`
	}

	EnterUserIdentifierSettingsList struct {
		SettingResponses []SettingResponse `json:"setting_responses"`
		Link             string            `json:"link"`
	}

	EnterUserIdentifierSubtaskInput struct {
		SubtaskId    string                           `json:"subtask_id"`
		SettingsList *EnterUserIdentifierSettingsList `json:"settings_list"`
	}

	EnterUserIdentifierRequest struct {
		FlowToken     string                            `json:"flow_token"`
		SubtaskInputs []EnterUserIdentifierSubtaskInput `json:"subtask_inputs"`
	}

	EnterUserIdentifierSubTask struct {
		SubtaskId string `json:"subtask_id"`
	}

	EnterUserIdentifierResponse struct {
		FlowToken string                        `json:"flow_token"`
		Status    string                        `json:"status"`
		Subtasks  []*EnterUserIdentifierSubTask `json:"subtasks"`
	}
)

// Enter Password Request / Response Data
type (
	EnterPassword struct {
		Password string `json:"password"`
		Link     string `json:"link"`
	}

	EnterPasswordSubtaskInput struct {
		SubtaskId     string         `json:"subtask_id"`
		EnterPassword *EnterPassword `json:"enter_password"`
	}

	EnterPasswordRequest struct {
		// {"flow_token":"g;173530831161593371:-1735311452571:uriVXkTkwRhIMO50ZlGr8XMf:7","subtask_inputs":[{"subtask_id":"LoginEnterPassword","enter_password":{"password":"Cassava2023@#","link":"next_link"}}]}
		FlowToken     string                       `json:"flow_token"`
		SubtaskInputs []*EnterPasswordSubtaskInput `json:"subtask_inputs"`
	}

	EnterPasswordSubTask struct {
		SubtaskId string    `json:"subtask_id"`
		User      *UserInfo `json:"user,omitempty"`
	}
	EnterPasswordResponse struct {
		FlowToken string                  `json:"flow_token"`
		Status    string                  `json:"status"`
		Subtasks  []*EnterPasswordSubTask `json:"subtasks"`
	}
)

// Two Factor Auth Challenge Request / Response Data
type (
	EnterText struct {
		Text string `json:"text"`
		Link string `json:"link"`
	}

	TwoFactorAuthChallengeSubtaskInput struct {
		SubtaskId string     `json:"subtask_id"`
		EnterText *EnterText `json:"enter_text"`
	}

	TwoFactorAuthChallengeRequest struct {
		// {"flow_token":"g;173530831161593371:-1735311452571:uriVXkTkwRhIMO50ZlGr8XMf:9","subtask_inputs":[{"subtask_id":"LoginTwoFactorAuthChallenge","enter_text":{"text":"776342","link":"next_link"}}]}
		FlowToken     string                                `json:"flow_token"`
		SubtaskInputs []*TwoFactorAuthChallengeSubtaskInput `json:"subtask_inputs"`
	}
)

func (s1 *StartLoginRequest) WithDefaultValue() {
	s1.InputFlowData.FlowContext.DebugOverrides = make(map[string]string)
	s1.InputFlowData.FlowContext.StartLocation.Location = "manual_link"
	s1.SubtaskVersions = map[string]int{"action_list": 2, "alert_dialog": 1, "app_download_cta": 1, "check_logged_in_account": 1, "choice_selection": 3, "contacts_live_sync_permission_prompt": 0, "cta": 7, "email_verification": 2, "end_flow": 1, "enter_date": 1, "enter_email": 2, "enter_password": 5, "enter_phone": 2, "enter_recaptcha": 1, "enter_text": 5, "enter_username": 2, "generic_urt": 3, "in_app_notification": 1, "interest_picker": 3, "js_instrumentation": 1, "menu_dialog": 1, "notifications_permission_prompt": 2, "open_account": 2, "open_home_timeline": 1, "open_link": 1, "phone_verification": 4, "privacy_options": 1, "security_key": 3, "select_avatar": 4, "select_banner": 2, "settings_list": 7, "show_code": 1, "sign_up": 2, "sign_up_review": 4, "tweet_selection_urt": 1, "update_users": 1, "upload_media": 1, "user_recommendations_list": 4, "user_recommendations_urt": 1, "wait_spinner": 3, "web_modal": 1}
}
