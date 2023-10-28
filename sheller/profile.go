package sheller

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

// Profile represents an Akeyless CLI profile.
type Profile struct {
	Name       string `toml:"name"`
	AccessID   string `toml:"access_id"`
	AccessType string `toml:"access_type"`
	// ... other properties as needed
}

// GetProfile loads the specified profile from the .akeyless/profiles directory.
func GetProfile(name string, config *Config) (*Profile, error) {
	profilePath := filepath.Join(config.AkeylessPath, "profiles", fmt.Sprintf("%s.toml", name))
	profileData, err := os.ReadFile(profilePath)
	if err != nil {
		return nil, err
	}

	profile := &Profile{}
	err = toml.Unmarshal(profileData, profile)
	if err != nil {
		return nil, err
	}
	profile.Name = name // Setting the profile name from the file name

	return profile, nil
}

// ListProfiles lists all profiles in the .akeyless/profiles directory.
func ListProfiles(config *Config) ([]Profile, error) {
	profilesDir := filepath.Join(config.AkeylessPath, "profiles")
	files, err := os.ReadDir(profilesDir) // Updated to use os.ReadDir
	if err != nil {
		return nil, err
	}

	var profiles []Profile
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".toml" {
			profileName := file.Name()[0 : len(file.Name())-len(filepath.Ext(file.Name()))]
			profile, err := GetProfile(profileName, config)
			if err != nil {
				return nil, err
			}
			profiles = append(profiles, *profile)
		}
	}

	return profiles, nil
}
