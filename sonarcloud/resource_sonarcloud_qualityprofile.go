package sonarcloud

import (
	"encoding/json"
	"net/http"
	"net/url"
        "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	log "github.com/sirupsen/logrus"
)

// Returns the resource represented by this file.
func resourceSonarcloudQualityProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSonarcloudQualityProfileCreate,
		Read:   resourceSonarcloudQualityProfileRead,
		Delete: resourceSonarcloudQualityProfileDelete,
		Importer: &schema.ResourceImporter{
			State: resourceSonarcloudQualityProfileImport,
		},

		// Define the fields of this schema.
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
                                Description: "Quality profile name",
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
                                Description: "Quality profile language",
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
		},
	}
}

func resourceSonarcloudQualityProfileCreate(d *schema.ResourceData, m interface{}) error {
	sonarCloudURL := m.(*ProviderConfiguration).sonarCloudURL
	sonarCloudURL.Path = "api/qualityprofiles/create"
	sonarCloudURL.RawQuery = url.Values{
		"name": []string{d.Get("name").(string)},
                "language": []string{d.Get("language").(string)},
                "organization": []string{d.Get("organization").(string)},
	}.Encode()

	resp, err := httpRequestHelper(
		m.(*ProviderConfiguration).httpClient,
		"POST",
		sonarCloudURL.String(),
		http.StatusOK,
		"resourceSonarcloudQualityProfileCreate",
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response into struct
	qualityProfileResponse := CreateQualityProfileResponse{}
	err = json.NewDecoder(resp.Body).Decode(&qualityProfileResponse)
	if err != nil {
		log.WithError(err).Error("resourceQualityProfileCreate: Failed to decode json into struct")
	}

        d.SetId(qualityProfileResponse.Profile.Key)
        d.Set("key", qualityProfileResponse.Profile.Key)
	return nil
}

func resourceSonarcloudQualityProfileRead(d *schema.ResourceData, m interface{}) error {
	sonarCloudURL := m.(*ProviderConfiguration).sonarCloudURL
	sonarCloudURL.Path = "api/qualityprofiles/search"
        sonarCloudURL.RawQuery = url.Values{
		"qualityProfile": []string{d.Get("name").(string)},
                "organization": []string{d.Get("organization").(string)},
                "language": []string{d.Get("language").(string)},
	}.Encode()
    
	resp, err := httpRequestHelper(
		m.(*ProviderConfiguration).httpClient,
		"GET",
		sonarCloudURL.String(),
		http.StatusOK,
		"resourceSonarcloudQualityProfileRead",
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response into struct
	qualityProfileReadResponse := GetQualityProfileList{}
	err = json.NewDecoder(resp.Body).Decode(&qualityProfileReadResponse)
	if err != nil {
		log.WithError(err).Error("resourceQualityProfileRead: Failed to decode json into struct")
	}
        for _, value := range qualityProfileReadResponse.Profiles {
		if d.Id() == value.Key {
			d.SetId(value.Key)
			d.Set("name", value.Name)
			d.Set("language", value.Language)
                        d.Set("key", value.Key)
		}
	}  
	
        return nil
}

func resourceSonarcloudQualityProfileDelete(d *schema.ResourceData, m interface{}) error {
	sonarCloudURL := m.(*ProviderConfiguration).sonarCloudURL
	sonarCloudURL.Path = "api/qualityprofiles/delete"
	sonarCloudURL.RawQuery = url.Values{
		"qualityProfile": []string{d.Get("name").(string)},
                "organization": []string{d.Get("organization").(string)},
                "language": []string{d.Get("language").(string)},
	}.Encode()

	resp, err := httpRequestHelper(
		m.(*ProviderConfiguration).httpClient,
		"POST",
		sonarCloudURL.String(),
		http.StatusNoContent,
		"resourceQualityProfileDelete",
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response into struct
	qualityProfileReadResponse := GetQualityProfile{}
	err = json.NewDecoder(resp.Body).Decode(&qualityProfileReadResponse)
	if err != nil {
		log.WithError(err).Error("resourceQualityProfileDelete: Failed to decode json into struct")
	}

	return nil
}

func resourceSonarcloudQualityProfileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if err := resourceSonarcloudQualityProfileRead(d, m); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func isValidLanguage(language string) bool {
	switch language {
	case
		"cs",
		"css",
		"flex",
		"go",
		"java",
		"js",
		"jsp",
		"kotlin",
		"php",
		"py",
		"ruby",
		"scala",
		"ts",
		"vbnet",
		"web",
		"xml":
		return true
	}
	return false
}
