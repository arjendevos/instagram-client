package responses

type Profile struct {
	Data struct {
		User struct {
			Biography             string        `json:"biography"`
			BioLinks              []interface{} `json:"bio_links"`
			BiographyWithEntities struct {
				RawText  string `json:"raw_text"`
				Entities []struct {
					User struct {
						Username string `json:"username"`
					} `json:"user"`
					Hashtag interface{} `json:"hashtag"`
				} `json:"entities"`
			} `json:"biography_with_entities"`
			BlockedByViewer        bool   `json:"blocked_by_viewer"`
			RestrictedByViewer     bool   `json:"restricted_by_viewer"`
			CountryBlock           bool   `json:"country_block"`
			ExternalURL            string `json:"external_url"`
			ExternalURLLinkshimmed string `json:"external_url_linkshimmed"`
			EdgeFollowedBy         struct {
				Count int `json:"count"`
			} `json:"edge_followed_by"`
			Fbid             string `json:"fbid"`
			FollowedByViewer bool   `json:"followed_by_viewer"`
			EdgeFollow       struct {
				Count int `json:"count"`
			} `json:"edge_follow"`
			FollowsViewer                       bool        `json:"follows_viewer"`
			FullName                            string      `json:"full_name"`
			HasArEffects                        bool        `json:"has_ar_effects"`
			HasClips                            bool        `json:"has_clips"`
			HasGuides                           bool        `json:"has_guides"`
			HasChannel                          bool        `json:"has_channel"`
			HasBlockedViewer                    bool        `json:"has_blocked_viewer"`
			HighlightReelCount                  int         `json:"highlight_reel_count"`
			HasRequestedViewer                  bool        `json:"has_requested_viewer"`
			HideLikeAndViewCounts               bool        `json:"hide_like_and_view_counts"`
			ID                                  string      `json:"id"`
			IsBusinessAccount                   bool        `json:"is_business_account"`
			IsEligibleToViewAccountTransparency bool        `json:"is_eligible_to_view_account_transparency"`
			IsProfessionalAccount               bool        `json:"is_professional_account"`
			IsSupervisionEnabled                bool        `json:"is_supervision_enabled"`
			IsGuardianOfViewer                  bool        `json:"is_guardian_of_viewer"`
			IsSupervisedByViewer                bool        `json:"is_supervised_by_viewer"`
			IsSupervisedUser                    bool        `json:"is_supervised_user"`
			IsEmbedsDisabled                    bool        `json:"is_embeds_disabled"`
			IsJoinedRecently                    bool        `json:"is_joined_recently"`
			GuardianID                          interface{} `json:"guardian_id"`
			BusinessAddressJSON                 interface{} `json:"business_address_json"`
			BusinessContactMethod               string      `json:"business_contact_method"`
			BusinessEmail                       string      `json:"business_email"`
			BusinessPhoneNumber                 string      `json:"business_phone_number"`
			BusinessCategoryName                string      `json:"business_category_name"`
			OverallCategoryName                 interface{} `json:"overall_category_name"`
			CategoryEnum                        interface{} `json:"category_enum"`
			CategoryName                        string      `json:"category_name"`
			IsPrivate                           bool        `json:"is_private"`
			IsVerified                          bool        `json:"is_verified"`
			EdgeMutualFollowedBy                struct {
				Count int           `json:"count"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_mutual_followed_by"`
			ProfilePicURL               string        `json:"profile_pic_url"`
			ProfilePicURLHd             string        `json:"profile_pic_url_hd"`
			RequestedByViewer           bool          `json:"requested_by_viewer"`
			ShouldShowCategory          bool          `json:"should_show_category"`
			ShouldShowPublicContacts    bool          `json:"should_show_public_contacts"`
			StateControlledMediaCountry interface{}   `json:"state_controlled_media_country"`
			TransparencyLabel           interface{}   `json:"transparency_label"`
			TransparencyProduct         string        `json:"transparency_product"`
			Username                    string        `json:"username"`
			ConnectedFbPage             interface{}   `json:"connected_fb_page"`
			Pronouns                    []interface{} `json:"pronouns"`
			EdgeFelixVideoTimeline      struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool        `json:"has_next_page"`
					EndCursor   interface{} `json:"end_cursor"`
				} `json:"page_info"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_felix_video_timeline"`
			EdgeOwnerToTimelineMedia struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool   `json:"has_next_page"`
					EndCursor   string `json:"end_cursor"`
				} `json:"page_info"`
				Edges []struct {
					Node struct {
						Typename   string `json:"__typename"`
						ID         string `json:"id"`
						Shortcode  string `json:"shortcode"`
						Dimensions struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						DisplayURL            string `json:"display_url"`
						EdgeMediaToTaggedUser struct {
							Edges []interface{} `json:"edges"`
						} `json:"edge_media_to_tagged_user"`
						FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
						FactCheckInformation   interface{} `json:"fact_check_information"`
						GatingInfo             interface{} `json:"gating_info"`
						SharingFrictionInfo    struct {
							ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
							BloksAppURL               interface{} `json:"bloks_app_url"`
						} `json:"sharing_friction_info"`
						MediaOverlayInfo interface{} `json:"media_overlay_info"`
						MediaPreview     interface{} `json:"media_preview"`
						Owner            struct {
							ID       string `json:"id"`
							Username string `json:"username"`
						} `json:"owner"`
						IsVideo              bool   `json:"is_video"`
						HasUpcomingEvent     bool   `json:"has_upcoming_event"`
						AccessibilityCaption string `json:"accessibility_caption"`
						EdgeMediaToCaption   struct {
							Edges []struct {
								Node struct {
									Text string `json:"text"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_to_caption"`
						EdgeMediaToComment struct {
							Count int `json:"count"`
						} `json:"edge_media_to_comment"`
						CommentsDisabled bool `json:"comments_disabled"`
						TakenAtTimestamp int  `json:"taken_at_timestamp"`
						EdgeLikedBy      struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
						EdgeMediaPreviewLike struct {
							Count int `json:"count"`
						} `json:"edge_media_preview_like"`
						Location struct {
							ID            string `json:"id"`
							HasPublicPage bool   `json:"has_public_page"`
							Name          string `json:"name"`
							Slug          string `json:"slug"`
						} `json:"location"`
						NftAssetInfo       interface{} `json:"nft_asset_info"`
						ThumbnailSrc       string      `json:"thumbnail_src"`
						ThumbnailResources []struct {
							Src          string `json:"src"`
							ConfigWidth  int    `json:"config_width"`
							ConfigHeight int    `json:"config_height"`
						} `json:"thumbnail_resources"`
						CoauthorProducers     []interface{} `json:"coauthor_producers"`
						PinnedForUsers        []interface{} `json:"pinned_for_users"`
						EdgeSidecarToChildren struct {
							Edges []struct {
								Node struct {
									Typename   string `json:"__typename"`
									ID         string `json:"id"`
									Shortcode  string `json:"shortcode"`
									Dimensions struct {
										Height int `json:"height"`
										Width  int `json:"width"`
									} `json:"dimensions"`
									DisplayURL            string `json:"display_url"`
									EdgeMediaToTaggedUser struct {
										Edges []interface{} `json:"edges"`
									} `json:"edge_media_to_tagged_user"`
									FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
									FactCheckInformation   interface{} `json:"fact_check_information"`
									GatingInfo             interface{} `json:"gating_info"`
									SharingFrictionInfo    struct {
										ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
										BloksAppURL               interface{} `json:"bloks_app_url"`
									} `json:"sharing_friction_info"`
									MediaOverlayInfo interface{} `json:"media_overlay_info"`
									MediaPreview     string      `json:"media_preview"`
									Owner            struct {
										ID       string `json:"id"`
										Username string `json:"username"`
									} `json:"owner"`
									IsVideo              bool   `json:"is_video"`
									HasUpcomingEvent     bool   `json:"has_upcoming_event"`
									AccessibilityCaption string `json:"accessibility_caption"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_sidecar_to_children"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
			EdgeSavedMedia struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool        `json:"has_next_page"`
					EndCursor   interface{} `json:"end_cursor"`
				} `json:"page_info"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_saved_media"`
			EdgeMediaCollections struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool        `json:"has_next_page"`
					EndCursor   interface{} `json:"end_cursor"`
				} `json:"page_info"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_collections"`
		} `json:"user"`
	} `json:"data"`
	Status string `json:"status"`
}
