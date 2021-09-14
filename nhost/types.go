package nhost

import (
	"sync"

	"github.com/docker/docker/api/types/container"
)

type (

	// Authentication validation response
	Response struct {
		Project           Project `json:",omitempty"`
		User              User
		Error             Error
		Email             string `json:"email"`
		VerificationToken string `json:"verificationToken"`
		VerifiedToken     string `json:"token"`
	}

	// Error structure
	Error struct {
		Code  string
		Email string
	}

	// Nhost individual team structure
	TeamData struct {
		Team Team `json:",omitempty"`
	}

	Team struct {
		Name     string    `json:",omitempty"`
		ID       string    `json:",omitempty"`
		Projects []Project `json:",omitempty"`
	}

	// Nhost project structure
	Project struct {
		ID                          string                   `json:"id" yaml:"project_id"`
		UserID                      string                   `json:"user_id"`
		Team                        Team                     `json:",omitempty"`
		TeamID                      string                   `json:"team_id,omitempty"`
		Type                        string                   `json:",omitempty"`
		Name                        string                   `json:"name"`
		HasuraGQEVersion            string                   `json:"hasura_gqe_version,omitempty"`
		BackendVersion              string                   `json:"backend_version,omitempty"`
		HasuraGQEAdminSecret        string                   `json:"hasura_gqe_admin_secret,omitempty"`
		PostgresVersion             string                   `json:"postgres_version,omitempty"`
		HasuraGQECustomEnvVariables map[string]string        `json:"hasura_gqe_custom_env_variables,omitempty"`
		BackendUserFields           string                   `json:"backend_user_fields,omitempty"`
		HBPDefaultAllowedUserRoles  string                   `json:"hbp_DEFAULT_ALLOWED_USER_ROLES,omitempty"`
		HBPRegistrationCustomFields string                   `json:"hbp_REGISTRATION_CUSTOM_FIELDS,omitempty"`
		HBPAllowedUserRoles         string                   `json:"hbp_allowed_user_roles,omitempty"`
		ProjectDomains              Domains                  `json:"project_domain"`
		ProjectEnvVars              []map[string]interface{} `json:"project_env_vars,omitempty"`
	}

	// Nhost project domains
	Domains struct {
		Hasura string `json:"hasura_domain,omitempty"`
	}

	// Nhost user structure
	User struct {
		ID       string     `json:",omitempty"`
		Projects []Project  `json:",omitempty"`
		Teams    []TeamData `json:",omitempty"`
	}

	// Session struct
	Session struct {
		Command string
		Dir     string
		Log     bool
		Browser string
	}

	// Nhost config.yaml root structure
	Configuration struct {
		MetadataDirectory string                      `yaml:"metadata_directory,omitempty"`
		Services          map[string]*Service         `yaml:",omitempty"`
		Auth              map[interface{}]interface{} `yaml:",omitempty"`
		Storage           map[interface{}]interface{} `yaml:",omitempty"`
		Version           int                         `yaml:",omitempty"`
		Sessions          map[string]Session          `yaml:",omitempty"`
		// Environment       map[string]interface{} `yaml:",omitempty"`
	}

	// Nhost config.yaml authentication structure
	Authentication struct {
		Endpoints map[string]interface{} `yaml:",omitempty"`
		Providers map[string]interface{} `yaml:",omitempty"`
		Signin    map[string]interface{} `yaml:",omitempty"`
		Signup    map[string]interface{} `yaml:",omitempty"`
		Email     map[string]interface{} `yaml:",omitempty"`
		Tokens    map[string]interface{} `yaml:",omitempty"`
		Gravatar  map[string]interface{} `yaml:",omitempty"`
	}

	// Nhost config.yaml service structure
	Service struct {
		Port           int                   `yaml:",omitempty"`
		Version        interface{}           `yaml:",omitempty"`
		Image          string                `yaml:",omitempty"`
		AdminSecret    interface{}           `yaml:"admin_secret,omitempty"`
		Name           string                `yaml:",omitempty"`
		Address        string                `yaml:",omitempty"`
		ID             string                `yaml:",omitempty"`
		Handle         string                `yaml:",omitempty"`
		Proxy          bool                  `yaml:",omitempty"`
		Config         *container.Config     `yaml:",omitempty"`
		HostConfig     *container.HostConfig `yaml:",omitempty"`
		HealthEndpoint string                `yaml:",omitempty"`

		// Channels are best thought of as queues (FIFO).
		// Therefore you can't really skip around.
		// We need a mutex to lock the service
		// before updating it's channels.
		sync.Mutex `yaml:",omitempty"`
		Active     bool `yaml:",omitempty"`

		//	HTTP Handler function.
		//	If specified, then all proxy requests will be handle with this handler.
		//	Handler func(http.ResponseWriter, *http.Request) `yaml:",omitempty"`
	}

	// .nhost/nhost.yaml information
	Information struct {
		ProjectID string `yaml:"project_id,omitempty"`
	}

	// Nhost servers structure
	Server struct {
		ID          string
		Name        string
		CountryCode string
		City        string
	}

	// Authentication credentials structure
	Credentials struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}

	// GitHub Release API reponse structure
	Release struct {
		URL         string  `json:",omitempty"`
		Name        string  `json:",omitempty"`
		TagName     string  `json:"tag_name,omitempty"`
		Prerelease  string  `json:",omitempty"`
		CreatedAt   string  `json:",omitempty"`
		PublishedAt string  `json:",omitempty"`
		Body        string  `json:",omitempty"`
		Assets      []Asset `json:",omitempty"`
	}

	// GitHub Release API Assets structure
	Asset struct {
		URL                string `json:",omitempty"`
		Name               string `json:",omitempty"`
		ID                 string `json:",omitempty"`
		Label              string `json:",omitempty"`
		BrowserDownloadURL string `json:"browser_download_url,omitempty"`
		Size               int    `json:",omitempty"`
	}
)
