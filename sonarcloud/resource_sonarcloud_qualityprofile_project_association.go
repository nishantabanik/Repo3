package sonarcloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Returns the resource represented by this file.
func resourceSonarcloudQualityProfileProjectAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSonarcloudQualityProfileProjectAssociationCreate,
		Read:   resourceSonarcloudQualityProfileProjectAssociationRead,
		Delete: resourceSonarcloudQualityProfileProjectAssociationDelete,
     
                Importer: &schema.ResourceImporter{
			State: resourceSonarcloudQualityProfileProjectAssociationImport,
		},

		// Define the fields of this schema.
		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
                                ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					l := len(val.(string))
					if l > 100 {
						errs = append(errs, fmt.Errorf("%q must not be longer than 100 characters", key))
					}
					return
				},
			},
                        "language": {
                                Type:     schema.TypeString,
                                Required: true,
                                ForceNew: true,
                                ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if !isValidLanguage(v) {
						errs = append(errs, fmt.Errorf("%q must be a supported language, got: %v", key, v))
					}
					return
				},
                        },
			"organization": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
                                
			},
                        "quality_profile": {
                                Type:     schema.TypeString,
                                Required: true,
                                ForceNew: true,
                                ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					l := len(val.(string))
					if l > 100 {
						errs = append(errs, fmt.Errorf("%q must not be longer than 100 characters", key))
					}
					return
				},
                        },
		},
	}
}

func resourceSonarcloudQualityProfileProjectAssociationCreate(d *schema.ResourceData, m interface{}) error {
	sonarCloudURL := m.(*ProviderConfiguration).sonarCloudURL
	sonarCloudURL.Path = "api/qualityprofiles/add_project"
	sonarCloudURL.RawQuery = url.Values{
		"project":        []string{d.Get("project").(string)},
                "organization":   []string{d.Get("organization").(string)},
                "qualityProfile": []string{d.Get("quality_profile").(string)},
                "language":       []string{d.Get("language").(string)},
	}.Encode()

	resp, err := httpRequestHelper(
		m.(*ProviderConfiguration).httpClient,
		"POST",
		sonarCloudURL.String(),
		http.StatusNoContent,
		"resourceSonarcloudQualityProfileProjectAssociationCreate",
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

        id := fmt.Sprintf("%v/%v", d.Get("quality_profile").(string), d.Get("project").(string))
	d.SetId(id)
	return nil
}

func resourceSonarcloudQualityProfileProjectAssociationRead(d *schema.ResourceData, m interface{}) error {
	var language string
	var qualityProfile string

	// Id is composed of qualityProfile name and project name
	idSlice := strings.Split(d.Id(), "/")

	// Call api/qualityprofiles/search to return the qualityProfileID
	sonarCloudURL := m.(*ProviderConfiguration).sonarCloudURL
	sonarCloudURL.Path = "api/qualityprofiles/search"
	sonarCloudURL.RawQuery = url.Values{
		"qualityProfile": []string{idSlice[0]},
                "organization":   []string{d.Get("organization").(string)},
	}.Encode()

	resp, err := httpRequestHelper(
		m.(*ProviderConfiguration).httpClient,
		"GET",
		sonarCloudURL.String(),
		http.StatusOK,
		"resourceSonarcloudQualityProfileProjectAssociationRead",
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response into struct
	getQualityProfileResponse := GetQualityProfileList{}
	err = json.NewDecoder(resp.Body).Decode(&getQualityProfileResponse)
	if err != nil {
		return fmt.Errorf("resourceSonarcloudQualityProfileProjectAssociationRead: Failed to decode json into struct: %+v", err)
	}

	var qualityProfileID string
	for _, value := range getQualityProfileResponse.Profiles {
		qualityProfileID = value.Key
		language = value.Language
		qualityProfile = value.Name
	}

	// With the qualityProfileID we can check if the project name is associated
	sonarCloudURL.Path = "api/qualityprofiles/projects"
	sonarCloudURL.RawQuery = url.Values{
		"key": []string{qualityProfileID},
	}.Encode()

	resp, err = httpRequestHelper(
		m.(*ProviderConfiguration).httpClient,
		"GET",
		sonarCloudURL.String(),
		http.StatusOK,
		"resourceSonarcloudQualityProfileProjectAssociationRead",
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response into struct
	getQualityProfileProjectResponse := GetQualityProfileProjectAssociation{}
	err = json.NewDecoder(resp.Body).Decode(&getQualityProfileProjectResponse)
	if err != nil {
		return fmt.Errorf("resourceSonarcloudQualityProfileProjectAssociationRead: Failed to decode json into struct: %+v", err)
	}

	for _, value := range getQualityProfileProjectResponse.Results {
		if idSlice[1] == value.Name {
			d.Set("project", value.Name)
			d.Set("quality_profile", qualityProfile)
			d.Set("language", language)
		}
	}

	return nil

}

func resourceSonarcloudQualityProfileProjectAssociationDelete(d *schema.ResourceData, m interface{}) error {
	sonarCloudURL := m.(*ProviderConfiguration).sonarCloudURL
	sonarCloudURL.Path = "api/qualityprofiles/remove_project"
	sonarCloudURL.RawQuery = url.Values{
                "project":        []string{d.Get("project").(string)},
                "organization":   []string{d.Get("organization").(string)},
                "qualityProfile": []string{d.Get("quality_profile").(string)},
                "language":       []string{d.Get("language").(string)},
	}.Encode()

	resp, err := httpRequestHelper(
		m.(*ProviderConfiguration).httpClient,
		"POST",
		sonarCloudURL.String(),
		http.StatusNoContent,
		"resourceSonarcloudQualityProfileProjectAssociationDelete",
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func resourceSonarcloudQualityProfileProjectAssociationImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if err := resourceSonarcloudQualityProfileRead(d, m); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
