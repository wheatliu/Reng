step-1 request
``` javascript
fetch("https://api.x.com/1.1/onboarding/task.json?flow_name=login", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-site",
    "x-client-transaction-id": "7lPti6Q/xQ/3O1RUY/qMMeVluX+C0/aidmGBm6+r07zj+oGBFQAgtqCpRHnJUVigwE/B8O2CoTQaXFPPZB6SZMXSiaEy7Q",
    "x-guest-token": "1872331613247324529",
    "x-twitter-active-user": "yes",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": "{\"input_flow_data\":{\"flow_context\":{\"debug_overrides\":{},\"start_location\":{\"location\":\"manual_link\"}}},\"subtask_versions\":{\"action_list\":2,\"alert_dialog\":1,\"app_download_cta\":1,\"check_logged_in_account\":1,\"choice_selection\":3,\"contacts_live_sync_permission_prompt\":0,\"cta\":7,\"email_verification\":2,\"end_flow\":1,\"enter_date\":1,\"enter_email\":2,\"enter_password\":5,\"enter_phone\":2,\"enter_recaptcha\":1,\"enter_text\":5,\"enter_username\":2,\"generic_urt\":3,\"in_app_notification\":1,\"interest_picker\":3,\"js_instrumentation\":1,\"menu_dialog\":1,\"notifications_permission_prompt\":2,\"open_account\":2,\"open_home_timeline\":1,\"open_link\":1,\"phone_verification\":4,\"privacy_options\":1,\"security_key\":3,\"select_avatar\":4,\"select_banner\":2,\"settings_list\":7,\"show_code\":1,\"sign_up\":2,\"sign_up_review\":4,\"tweet_selection_urt\":1,\"update_users\":1,\"upload_media\":1,\"user_recommendations_list\":4,\"user_recommendations_urt\":1,\"wait_spinner\":3,\"web_modal\":1}}",
  "method": "POST"
});
```

step-1 body
``` javascript
{"input_flow_data":{"flow_context":{"debug_overrides":{},"start_location":{"location":"manual_link"}}},"subtask_versions":{"action_list":2,"alert_dialog":1,"app_download_cta":1,"check_logged_in_account":1,"choice_selection":3,"contacts_live_sync_permission_prompt":0,"cta":7,"email_verification":2,"end_flow":1,"enter_date":1,"enter_email":2,"enter_password":5,"enter_phone":2,"enter_recaptcha":1,"enter_text":5,"enter_username":2,"generic_urt":3,"in_app_notification":1,"interest_picker":3,"js_instrumentation":1,"menu_dialog":1,"notifications_permission_prompt":2,"open_account":2,"open_home_timeline":1,"open_link":1,"phone_verification":4,"privacy_options":1,"security_key":3,"select_avatar":4,"select_banner":2,"settings_list":7,"show_code":1,"sign_up":2,"sign_up_review":4,"tweet_selection_urt":1,"update_users":1,"upload_media":1,"user_recommendations_list":4,"user_recommendations_urt":1,"wait_spinner":3,"web_modal":1}}
```

step-1 response

``` javascript
{
    "flow_token": "g;173515123746842611:-1735234321691:0P0cBS0rmV2v8zVJKzHzYt4j:0",
    "status": "success",
    "subtasks": [
        {
            "subtask_id": "LoginJsInstrumentationSubtask",
            "js_instrumentation": {
                "url": "https://twitter.com/i/js_inst?c_name=ui_metrics",
                "timeout_ms": 2000,
                "next_link": {
                    "link_type": "task",
                    "link_id": "next_link"
                }
            }
        }
    ]
}
```

step-2 request

```javascript
fetch("https://api.x.com/1.1/onboarding/task.json", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-site",
    "x-client-transaction-id": "9En3kb4l3xXtIU5OeeCWK/9/o2WYyey4bHubgbWxyab54JubDxo6rLqzXmPTS0K62lbb6vdYwuOrtZ8dn/Ooew025muj9w",
    "x-guest-token": "1872331613247324529",
    "x-twitter-active-user": "yes",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": "{\"flow_token\":\"g;:-::0\",\"subtask_inputs\":[{\"subtask_id\":\"LoginJsInstrumentationSubtask\",\"js_instrumentation\":{\"response\":\"{\\\"rf\\\":{\\\"eb42a173b7c6448ec98b1ab1e8490b89f59712c93ec96136b4e584b6aa98ba26\\\":26,\\\"a477b65c4cb67055b9086c1afaa3ba05658a1f9597878f0c1d882c5ec9cd08b3\\\":-22,\\\"\\\":85,\\\"a564f82d6f700dd173268a4f44924db1fb4e69c1c685d20cbb0f8c6159145288\\\":-43},\\\"s\\\":\\\"--RLA3fpat8-\\\"}\",\"link\":\"next_link\"}}]}",
  "method": "POST"
});
```

step-2 body

``` javascript
{"flow_token":"g;:-::0","subtask_inputs":[{"subtask_id":"LoginJsInstrumentationSubtask","js_instrumentation":{"response":"{\"rf\":{\"\":26,\"\":-22,\"\":85,\"\":-43},\"s\":\"--RLA3fpat8-\"}","link":"next_link"}}]}
```

step-2 response

```javascript
{
    "flow_token": "g;:-::1",
    "status": "success",
    "subtasks": [
        {
            "subtask_id": "LoginEnterUserIdentifierSSO",
            "settings_list": {
                "settings": [
                    {
                        "value_type": "button",
                        "value_identifier": "google_sso_button",
                        "value_data": {
                            "button": {
                                "navigation_link": {
                                    "link_type": "subtask",
                                    "link_id": "google_sso",
                                    "label": "Continue with Google",
                                    "subtask_id": "EnterIdGoogleSSOSubtask"
                                },
                                "style": "brand",
                                "icon": {
                                    "icon": "logo_google_g_color"
                                },
                                "preferred_size": "normal"
                            }
                        }
                    },
                    {
                        "value_type": "button",
                        "value_identifier": "apple_sso_button",
                        "value_data": {
                            "button": {
                                "navigation_link": {
                                    "link_type": "subtask",
                                    "link_id": "apple_id",
                                    "label": "Continue with Apple",
                                    "subtask_id": "EnterIdAppleSSOSubtask"
                                },
                                "style": "brand",
                                "icon": {
                                    "icon": "logo_apple"
                                },
                                "preferred_size": "normal"
                            }
                        }
                    },
                    {
                        "value_type": "separator",
                        "value_identifier": "separator",
                        "value_data": {
                            "separator": {
                                "label": {
                                    "text": "or",
                                    "entities": []
                                }
                            }
                        }
                    },
                    {
                        "value_type": "text_field",
                        "value_identifier": "user_identifier",
                        "value_data": {
                            "text_field": {
                                "content_type": "text",
                                "hint_text": "Phone, email, or username"
                            }
                        }
                    },
                    {
                        "value_type": "button",
                        "value_identifier": "next_button",
                        "value_data": {
                            "button": {
                                "navigation_link": {
                                    "link_type": "task",
                                    "link_id": "next_link",
                                    "label": "Next"
                                },
                                "style": "primary",
                                "preferred_size": "normal"
                            }
                        }
                    },
                    {
                        "value_type": "button",
                        "value_identifier": "forgot_password",
                        "value_data": {
                            "button": {
                                "navigation_link": {
                                    "link_type": "subtask",
                                    "link_id": "forget_password",
                                    "label": "Forgot password?",
                                    "subtask_id": "RedirectToPasswordReset"
                                },
                                "style": "secondary",
                                "preferred_size": "normal"
                            }
                        }
                    }
                ],
                "detail_text": {
                    "text": "Don't have an account? Sign up",
                    "entities": [
                        {
                            "from_index": 23,
                            "to_index": 30,
                            "navigation_link": {
                                "link_type": "deep_link_and_abort",
                                "link_id": "signup_deep_link",
                                "url": "https://twitter.com/i/flow/signup"
                            }
                        }
                    ]
                },
                "style": "step",
                "header": {
                    "primary_text": {
                        "text": "Sign in to X",
                        "entities": []
                    }
                },
                "navigation_style": "hide",
                "horizontal_style": "compact"
            },
            "subtask_back_navigation": "cancel_flow"
        },
        {
            "subtask_id": "EnterIdGoogleSSOSubtask",
            "single_sign_on": {
                "provider": "google",
                "scopes": [
                    "openid",
                    "email",
                    "profile"
                ],
                "state": "kKXWbZsLpFWH1zCPgtmOXDCRgWPO5K5WFpigOwGdPTL",
                "next_link": {
                    "link_type": "task",
                    "link_id": "next_link"
                },
                "fail_link": {
                    "link_type": "subtask",
                    "link_id": "fail_link",
                    "subtask_id": "LoginEnterUserIdentifierSSO"
                },
                "cancel_link": {
                    "link_type": "subtask",
                    "link_id": "cancel_link",
                    "subtask_id": "LoginEnterUserIdentifierSSO"
                }
            },
            "subtask_back_navigation": "cancel_flow"
        },
        {
            "subtask_id": "EnterIdAppleSSOSubtask",
            "single_sign_on": {
                "provider": "apple",
                "scopes": [
                    "email",
                    "name"
                ],
                "state": "O1cWlM9_FGQWBlEBofqBzQwJ1uwHP_ZI5AQuCLJ0KXf",
                "next_link": {
                    "link_type": "task",
                    "link_id": "next_link"
                },
                "fail_link": {
                    "link_type": "subtask",
                    "link_id": "fail_link",
                    "subtask_id": "LoginEnterUserIdentifierSSO"
                },
                "cancel_link": {
                    "link_type": "subtask",
                    "link_id": "cancel_link",
                    "subtask_id": "LoginEnterUserIdentifierSSO"
                }
            },
            "subtask_back_navigation": "cancel_flow"
        },
        {
            "subtask_id": "RedirectToPasswordReset",
            "open_link": {
                "link": {
                    "link_type": "deep_link_and_abort",
                    "link_id": "password_reset_deep_link",
                    "url": "https://twitter.com/i/flow/password_reset?input_flow_data=%7B%22requested_variant%22%3A%22eyJwbGF0Zm9ybSI6IlJ3ZWIifQ%3D%3D%22%7D"
                }
            }
        }
    ]
}

```

step-3 request

``` javascript
fetch("https://api.x.com/1.1/onboarding/task.json", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-site",
    "x-client-transaction-id": "AL0DZUrRK+EZ1bq6jRRi3wuLV5FsPRhMmI9vdUFFPVINFG9v++7OWE5Hqpcnv7ZOLg4yHgM8BlqvEv7nlHYIczInGVMcAw",
    "x-guest-token": "1872331613247324529",
    "x-twitter-active-user": "yes",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": "{\"flow_token\":\"g;:-::1\",\"subtask_inputs\":[{\"subtask_id\":\"LoginEnterUserIdentifierSSO\",\"settings_list\":{\"setting_responses\":[{\"key\":\"user_identifier\",\"response_data\":{\"text_data\":{\"result\":\"\"}}}],\"link\":\"next_link\"}}]}",
  "method": "POST"
});
```

step-3 payload

```javascript
{"flow_token":"g;:-::1","subtask_inputs":[{"subtask_id":"LoginEnterUserIdentifierSSO","settings_list":{"setting_responses":[{"key":"user_identifier","response_data":{"text_data":{"result":""}}}],"link":"next_link"}}]}
```

step-3 response

```javascript
{
    "flow_token": "g;:-::7",
    "status": "success",
    "subtasks": [
        {
            "subtask_id": "LoginEnterPassword",
            "enter_password": {
                "primary_text": {
                    "text": "Enter your password",
                    "entities": []
                },
                "next_link": {
                    "link_type": "task",
                    "link_id": "next_link",
                    "label": "Log in"
                },
                "hint": "Password",
                "username": "xxxxxxxx",
                "user_identifier_display_type": "username",
                "skip_password_validation": true,
                "os_content_type": "password",
                "footer": {
                    "style": "fixed",
                    "footnote_text": {
                        "text": "Don't have an account? Sign up",
                        "entities": [
                            {
                                "from_index": 23,
                                "to_index": 30,
                                "navigation_link": {
                                    "link_type": "deep_link_and_abort",
                                    "link_id": "signup_deep_link",
                                    "url": "https://twitter.com/i/flow/signup"
                                }
                            }
                        ]
                    }
                },
                "password_field": {
                    "content_type": "password",
                    "hint_text": "Password",
                    "detail_text": {
                        "text": "Forgot password?\t",
                        "entities": [
                            {
                                "from_index": 0,
                                "to_index": 16,
                                "navigation_link": {
                                    "link_type": "subtask",
                                    "link_id": "forget_password",
                                    "label": "Forgot password?",
                                    "subtask_id": "RedirectToPasswordReset"
                                }
                            }
                        ]
                    }
                }
            },
            "subtask_back_navigation": "cancel_flow"
        },
        {
            "subtask_id": "RedirectToPasswordReset",
            "open_link": {
                "link": {
                    "link_type": "deep_link_and_abort",
                    "link_id": "password_reset_deep_link",
                    "url": "https://twitter.com/i/flow/password_reset?input_flow_data=%7B%22requested_variant%22%3A%22eyJ1c2VyX2lkZW50aWZpZXIiOiJwaGlsZW1lbnc2ODYyMiIsInBsYXRmb3JtIjoiUndlYiJ9%22%7D"
                }
            }
        }
    ]
}
```

step-4 request

```javascript
fetch("https://api.x.com/1.1/onboarding/task.json", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-site",
    "x-client-transaction-id": "KJUrTWL5A8kx/ZKSpTxK9yOjf7lEFTBksKdHXWltFXolPEdH08bmcGZvgr8Pl55mBt8aNitdbsDgFYukPccl0akUE4RuKw",
    "x-guest-token": "1872331613247324529",
    "x-twitter-active-user": "yes",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": "{\"flow_token\":\"g;:-::7\",\"subtask_inputs\":[{\"subtask_id\":\"LoginEnterPassword\",\"enter_password\":{\"password\":\"\",\"link\":\"next_link\"}}]}",
  "method": "POST"
});
```

step-4 payload

```javascript
{"flow_token":"g;:-1735234321691::7","subtask_inputs":[{"subtask_id":"LoginEnterPassword","enter_password":{"password":"","link":"next_link"}}]}
```

step-4 response

```javascript
{
    "flow_token": "g;:-::9",
    "status": "success",
    "subtasks": [
        {
            "subtask_id": "LoginTwoFactorAuthChallenge",
            "enter_text": {
                "header": {
                    "primary_text": {
                        "text": "Enter your verification code",
                        "entities": []
                    },
                    "secondary_text": {
                        "text": "Use your code generator app to generate a code and enter it below.",
                        "entities": []
                    },
                    "user": {
                        "id": 1722,
                        "id_str": "1722",
                        "name": "xxxxxxx",
                        "screen_name": "xxxxxxx",
                        "location": "CA",
                        "url": null,
                        "description": "cxxxxxxxx",
                        "protected": false,
                        "followers_count": -1,
                        "friends_count": -1,
                        "listed_count": null,
                        "created_at": "Fri Nov 10 14:05:48 +0000 2093",
                        "favourites_count": -1,
                        "utc_offset": null,
                        "time_zone": null,
                        "geo_enabled": false,
                        "verified": false,
                        "statuses_count": -1,
                        "lang": null,
                        "contributors_enabled": false,
                        "is_translator": false,
                        "is_translation_enabled": false,
                        "profile_background_color": "",
                        "profile_background_image_url": null,
                        "profile_background_image_url_https": null,
                        "profile_background_tile": false,
                        "profile_image_url": "http:\/\/pbs.twimg.com\/profile_images\/xxxxx\/xxxxx.png",
                        "profile_image_url_https": "https:\/\/pbs.twimg.com\/profile_images\/xxxx\/xxxx.png",
                        "profile_link_color": "",
                        "profile_sidebar_border_color": "",
                        "profile_sidebar_fill_color": "",
                        "profile_text_color": "",
                        "profile_use_background_image": false,
                        "default_profile": true,
                        "default_profile_image": false,
                        "can_media_tag": null,
                        "following": null,
                        "follow_request_sent": null,
                        "notifications": null,
                        "muting": null,
                        "blocking": null,
                        "blocked_by": null,
                        "translator_type": "none",
                        "withheld_in_countries": [],
                        "followed_by": null
                    }
                },
                "detail_text": {
                    "text": "Choose a different verification method\n\nContact X Support",
                    "entities": [
                        {
                            "from_index": 0,
                            "to_index": 38,
                            "navigation_link": {
                                "link_type": "subtask",
                                "link_id": "choose_different_method",
                                "label": "choose_methods_link",
                                "subtask_id": "LoginTwoFactorAuthChooseMethod"
                            }
                        },
                        {
                            "from_index": 40,
                            "to_index": 57,
                            "navigation_link": {
                                "link_type": "chromeless_web_link",
                                "link_id": "login_issues_link",
                                "url": "https://help.twitter.com/forms/account-access/regain-access"
                            }
                        }
                    ]
                },
                "hint_text": "Enter code",
                "multiline": false,
                "auto_capitalization_type": "none",
                "auto_correction_enabled": false,
                "keyboard_type": "number",
                "next_link": {
                    "link_type": "task",
                    "link_id": "next_link",
                    "label": "Next"
                }
            },
            "subtask_back_navigation": "cancel_flow"
        },
        {
            "subtask_id": "LoginTwoFactorAuthChooseMethod",
            "choice_selection": {
                "selection_type": "single_select",
                "choices": [
                    {
                        "id": "0",
                        "text": {
                            "text": "Use a code generator app",
                            "entities": []
                        }
                    },
                    {
                        "id": "1",
                        "text": {
                            "text": "Use a backup code",
                            "entities": []
                        }
                    }
                ],
                "primary_text": {
                    "text": "Choose a different verification method",
                    "entities": []
                },
                "next_link": {
                    "link_type": "task",
                    "link_id": "next_link",
                    "label": "Next"
                },
                "style": "default",
                "sections": []
            },
            "subtask_back_navigation": "cancel_flow"
        },
        {
            "subtask_id": "login_security_key_not_supported_cta",
            "cta": {
                "header": {
                    "primary_text": {
                        "text": "Why can’t you log in using a security key?",
                        "entities": []
                    },
                    "secondary_text": {
                        "text": "You can log in to twitter.com with a security key only when using a compatible web browser. Currently, the X mobile app doesn't support the use of security keys. ",
                        "entities": []
                    }
                },
                "primary_action_link": {
                    "link_type": "subtask",
                    "link_id": "choose_different_method",
                    "label": "Next",
                    "subtask_id": "LoginTwoFactorAuthChooseMethod"
                },
                "text_alignment": "left",
                "primary_action_style": "primary",
                "secondary_action_style": "secondary",
                "style": "scrollable"
            }
        }
    ]
}
```

step-5 request

```javascript
fetch("https://api.x.com/1.1/onboarding/task.json", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-site",
    "x-client-transaction-id": "3GHfuZYN9z3FCWZmUci+A9dXi02w4cSQRFOzqZ2Z4Y7RyLOzJzIShJKbdkv7Y2qS8jLowt80HsOu7ALnEBbOz599k7Zp3w",
    "x-guest-token": "",
    "x-twitter-active-user": "yes",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": "{\"flow_token\":\"g;:-::9\",\"subtask_inputs\":[{\"subtask_id\":\"LoginTwoFactorAuthChallenge\",\"enter_text\":{\"text\":\"\",\"link\":\"next_link\"}}]}",
  "method": "POST"
});
```

step-5 payload

```javascript
{"flow_token":"g;:-::9","subtask_inputs":[{"subtask_id":"LoginTwoFactorAuthChallenge","enter_text":{"text":"","link":"next_link"}}]}
```