package sonarcloud

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	log "github.com/sirupsen/logrus"
)

// Returns the resource represented by this file.
func resourceSonarcloudQualityProfileDefault() *schema.Resource {
	return &schema.Resource{
		Create: resourceSonarcloudQualityProfileDefaultCreate,
		Read:   resourceSonarcloudQualityProfileDefaultRead,
		Delete: resourceSonarcloudQualityProfileDefaultDelete,

		// Define the fields of this schema.
		Schema: map[string]*schema.Schema{
			"quality_profile": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
                        "language": {
                                Type:     schema.TypeString,
                                Required: true,
                                ForceNew: true,
                        },
                        "organization": {     
                                Type:     schema.TypeString,
                                Required: true,
                                ForceNew: true,
                        },
		},
	}
}

func resourceSonarcloudQualityProfileDefaultCreate(d *schema.ResourceData, m interface{}) error {
	sonarCloudURL := m.(*ProviderConfiguration).sonarCloudURL
	sonarCloudURL.Path = "api/qualityprofiles/set_default"
	sonarCloudURL.RawQuery = url.Values{
		"qualityProfile": []string{d.Get("quality_profile").(string)},
                "language": []string{d.Get("language").(string)},
                "organization": []string{d.Get("organization").(string)},
	}.Encode()

	resp, err := httpRequestHelper(
		m.(*ProviderConfiguration).httpClient,
		"POST",
		sonarCloudURL.String(),
		http.StatusNoContent,
		"resourceQualityProfileCreate",
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

 	d.SetId(strconv.FormatInt(qualityProfileResponse.ID, 10))
	return nil
}

func resourceSonarcloudQualityProfileDefaultRead(d *schema.ResourceData, m interface{}) error {
	sonarCloudURL := m.(*ProviderConfiguration).sonarCloudURL
	sonarCloudURL.Path = "api/qualityprofiles/search"
	sonarCloudURL.RawQuery = url.Values{
		"qualityProfile": []string{d.Get("quality_profile").(string)},
                "organization":   []string{d.Get("organization").(string)},
                "language":       []string{d.Get("language").(string)},
	}.Encode()

	resp, err := httpRequestHelper(
		m.(*ProviderConfiguration).httpClient,
		"GET",
		sonarCloudURL.String(),
		http.StatusOK,
		"resourceQualityProfileRead",
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response into struct
	qualityProfileReadResponse := GetQualityProfile{}
	err = json.NewDecoder(resp.Body).Decode(&qualityProfileReadResponse)
	if err != nil {
		log.WithError(err).Error("resourceQualityProfileRead: Failed to decode json into struct")
	}

	return nil
}

func resourceSonarcloudQualityProfileDefaultDelete(d *schema.ResourceData, m interface{}) error {
	sonarCloudURL := m.(*ProviderConfiguration).sonarCloudURL
	sonarCloudURL.Path = "api/qualityprofiles/set_default"
	sonarCloudURL.RawQuery = url.Values{
		"qualityProfile": []string{d.Get("quality_profile").(string)},
                "organization":   []string{d.Get("organization").(string)},
                "language":       []string{d.Get("language").(string)},
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
