/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"encoding/json"
	"fmt"
)

// ItemFields Used to control the data that gets attached to DtoBaseItems.
type ItemFields string

// List of ItemFields
const (
	ITEMFIELDS_AIR_TIME ItemFields = "AirTime"
	ITEMFIELDS_CAN_DELETE ItemFields = "CanDelete"
	ITEMFIELDS_CAN_DOWNLOAD ItemFields = "CanDownload"
	ITEMFIELDS_CHANNEL_INFO ItemFields = "ChannelInfo"
	ITEMFIELDS_CHAPTERS ItemFields = "Chapters"
	ITEMFIELDS_TRICKPLAY ItemFields = "Trickplay"
	ITEMFIELDS_CHILD_COUNT ItemFields = "ChildCount"
	ITEMFIELDS_CUMULATIVE_RUN_TIME_TICKS ItemFields = "CumulativeRunTimeTicks"
	ITEMFIELDS_CUSTOM_RATING ItemFields = "CustomRating"
	ITEMFIELDS_DATE_CREATED ItemFields = "DateCreated"
	ITEMFIELDS_DATE_LAST_MEDIA_ADDED ItemFields = "DateLastMediaAdded"
	ITEMFIELDS_DISPLAY_PREFERENCES_ID ItemFields = "DisplayPreferencesId"
	ITEMFIELDS_ETAG ItemFields = "Etag"
	ITEMFIELDS_EXTERNAL_URLS ItemFields = "ExternalUrls"
	ITEMFIELDS_GENRES ItemFields = "Genres"
	ITEMFIELDS_HOME_PAGE_URL ItemFields = "HomePageUrl"
	ITEMFIELDS_ITEM_COUNTS ItemFields = "ItemCounts"
	ITEMFIELDS_MEDIA_SOURCE_COUNT ItemFields = "MediaSourceCount"
	ITEMFIELDS_MEDIA_SOURCES ItemFields = "MediaSources"
	ITEMFIELDS_ORIGINAL_TITLE ItemFields = "OriginalTitle"
	ITEMFIELDS_OVERVIEW ItemFields = "Overview"
	ITEMFIELDS_PARENT_ID ItemFields = "ParentId"
	ITEMFIELDS_PATH ItemFields = "Path"
	ITEMFIELDS_PEOPLE ItemFields = "People"
	ITEMFIELDS_PLAY_ACCESS ItemFields = "PlayAccess"
	ITEMFIELDS_PRODUCTION_LOCATIONS ItemFields = "ProductionLocations"
	ITEMFIELDS_PROVIDER_IDS ItemFields = "ProviderIds"
	ITEMFIELDS_PRIMARY_IMAGE_ASPECT_RATIO ItemFields = "PrimaryImageAspectRatio"
	ITEMFIELDS_RECURSIVE_ITEM_COUNT ItemFields = "RecursiveItemCount"
	ITEMFIELDS_SETTINGS ItemFields = "Settings"
	ITEMFIELDS_SCREENSHOT_IMAGE_TAGS ItemFields = "ScreenshotImageTags"
	ITEMFIELDS_SERIES_PRIMARY_IMAGE ItemFields = "SeriesPrimaryImage"
	ITEMFIELDS_SERIES_STUDIO ItemFields = "SeriesStudio"
	ITEMFIELDS_SORT_NAME ItemFields = "SortName"
	ITEMFIELDS_SPECIAL_EPISODE_NUMBERS ItemFields = "SpecialEpisodeNumbers"
	ITEMFIELDS_STUDIOS ItemFields = "Studios"
	ITEMFIELDS_TAGLINES ItemFields = "Taglines"
	ITEMFIELDS_TAGS ItemFields = "Tags"
	ITEMFIELDS_REMOTE_TRAILERS ItemFields = "RemoteTrailers"
	ITEMFIELDS_MEDIA_STREAMS ItemFields = "MediaStreams"
	ITEMFIELDS_SEASON_USER_DATA ItemFields = "SeasonUserData"
	ITEMFIELDS_SERVICE_NAME ItemFields = "ServiceName"
	ITEMFIELDS_THEME_SONG_IDS ItemFields = "ThemeSongIds"
	ITEMFIELDS_THEME_VIDEO_IDS ItemFields = "ThemeVideoIds"
	ITEMFIELDS_EXTERNAL_ETAG ItemFields = "ExternalEtag"
	ITEMFIELDS_PRESENTATION_UNIQUE_KEY ItemFields = "PresentationUniqueKey"
	ITEMFIELDS_INHERITED_PARENTAL_RATING_VALUE ItemFields = "InheritedParentalRatingValue"
	ITEMFIELDS_EXTERNAL_SERIES_ID ItemFields = "ExternalSeriesId"
	ITEMFIELDS_SERIES_PRESENTATION_UNIQUE_KEY ItemFields = "SeriesPresentationUniqueKey"
	ITEMFIELDS_DATE_LAST_REFRESHED ItemFields = "DateLastRefreshed"
	ITEMFIELDS_DATE_LAST_SAVED ItemFields = "DateLastSaved"
	ITEMFIELDS_REFRESH_STATE ItemFields = "RefreshState"
	ITEMFIELDS_CHANNEL_IMAGE ItemFields = "ChannelImage"
	ITEMFIELDS_ENABLE_MEDIA_SOURCE_DISPLAY ItemFields = "EnableMediaSourceDisplay"
	ITEMFIELDS_WIDTH ItemFields = "Width"
	ITEMFIELDS_HEIGHT ItemFields = "Height"
	ITEMFIELDS_EXTRA_IDS ItemFields = "ExtraIds"
	ITEMFIELDS_LOCAL_TRAILER_COUNT ItemFields = "LocalTrailerCount"
	ITEMFIELDS_IS_HD ItemFields = "IsHD"
	ITEMFIELDS_SPECIAL_FEATURE_COUNT ItemFields = "SpecialFeatureCount"
)

// All allowed values of ItemFields enum
var AllowedItemFieldsEnumValues = []ItemFields{
	"AirTime",
	"CanDelete",
	"CanDownload",
	"ChannelInfo",
	"Chapters",
	"Trickplay",
	"ChildCount",
	"CumulativeRunTimeTicks",
	"CustomRating",
	"DateCreated",
	"DateLastMediaAdded",
	"DisplayPreferencesId",
	"Etag",
	"ExternalUrls",
	"Genres",
	"HomePageUrl",
	"ItemCounts",
	"MediaSourceCount",
	"MediaSources",
	"OriginalTitle",
	"Overview",
	"ParentId",
	"Path",
	"People",
	"PlayAccess",
	"ProductionLocations",
	"ProviderIds",
	"PrimaryImageAspectRatio",
	"RecursiveItemCount",
	"Settings",
	"ScreenshotImageTags",
	"SeriesPrimaryImage",
	"SeriesStudio",
	"SortName",
	"SpecialEpisodeNumbers",
	"Studios",
	"Taglines",
	"Tags",
	"RemoteTrailers",
	"MediaStreams",
	"SeasonUserData",
	"ServiceName",
	"ThemeSongIds",
	"ThemeVideoIds",
	"ExternalEtag",
	"PresentationUniqueKey",
	"InheritedParentalRatingValue",
	"ExternalSeriesId",
	"SeriesPresentationUniqueKey",
	"DateLastRefreshed",
	"DateLastSaved",
	"RefreshState",
	"ChannelImage",
	"EnableMediaSourceDisplay",
	"Width",
	"Height",
	"ExtraIds",
	"LocalTrailerCount",
	"IsHD",
	"SpecialFeatureCount",
}

func (v *ItemFields) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ItemFields(value)
	for _, existing := range AllowedItemFieldsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ItemFields", value)
}

// NewItemFieldsFromValue returns a pointer to a valid ItemFields
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewItemFieldsFromValue(v string) (*ItemFields, error) {
	ev := ItemFields(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ItemFields: valid values are %v", v, AllowedItemFieldsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ItemFields) IsValid() bool {
	for _, existing := range AllowedItemFieldsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ItemFields value
func (v ItemFields) Ptr() *ItemFields {
	return &v
}

type NullableItemFields struct {
	value *ItemFields
	isSet bool
}

func (v NullableItemFields) Get() *ItemFields {
	return v.value
}

func (v *NullableItemFields) Set(val *ItemFields) {
	v.value = val
	v.isSet = true
}

func (v NullableItemFields) IsSet() bool {
	return v.isSet
}

func (v *NullableItemFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemFields(val *ItemFields) *NullableItemFields {
	return &NullableItemFields{value: val, isSet: true}
}

func (v NullableItemFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

