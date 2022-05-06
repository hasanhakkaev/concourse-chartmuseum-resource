package resource

type (
	// Source struct represents a source of resources.
	Source struct {
		ServerURL    string   `json:"server_url"`
		ChartName    string   `json:"chart_name"`
		VersionRange string   `json:"version_range,omitempty"`
		Username     string   `json:"username,omitempty"`
		Password     string   `json:"password,omitempty"`
		DomainCerts  []string `json:"ca_certs,omitempty"`
		HarborAPI    bool     `json:"harbor_api,omitempty"`
	}
	// CheckRequest struct represents a request to check a source.
	CheckRequest struct {
		Source  Source `json:"source"`
		Version string `json:"version"`
		Digest  string `json:"digest"`
	}
	// CheckResponse struct represents a response to check a source.
	CheckResponse struct {
		Version string `json:"version"`
		Digest  string `json:"digest"`
	}
	//InRequest struct represents a request for in step for Concourse.
	InRequest struct {
		Source  Source `json:"source"`
		Version string `json:"version"`
		Digest  string `json:"digest"`
	}
	//InResponse struct represents a response for in step for Concourse.
	InResponse struct {
		Version  string     `json:"version"`
		Digest   string     `json:"digest"`
		Metadata []Metadata `json:"metadata"`
	}
	//OutRequest struct represents a request for out step for Concourse.
	OutRequest struct {
		Source     Source       `json:"source"`
		Parameters []Parameters `json:"parameters"`
	}
	//OutResponse struct represents a response for out step for Concourse.
	OutResponse struct {
		Version  string     `json:"version"`
		Digest   string     `json:"digest"`
		Metadata []Metadata `json:"metadata"`
	}
	//Metadata struct represents a metadata for a resource.
	Metadata struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	//Parameters struct represents a parameter for a resource.
	Parameters struct {
		Chart            string `json:"chart"`
		Sign             bool   `json:"sign,omitempty"`
		KeyData          string `json:"key_data"`
		KeyFile          string `json:"key_file"`
		KeyPassphrase    string `json:"key_passphrase"`
		Version          string `json:"version,omitempty"`
		VersionFile      string `json:"version_file,omitempty"`
		Force            bool   `json:"force,omitempty"`
		DependencyUpdate bool   `json:"dependency_update"`
	}
)

/*
func (source Source) AuthOptions(repo name.Repository, scopeActions []string) error {
	if source.Username != "" && source.Password != "" {
		auth = &authn.Basic{
			Username: source.Username,
			Password: source.Password,
		}
	} else {
		auth = authn.Anonymous
	}

	tr := http.DefaultTransport.(*http.Transport)
	// a cert was provided
	if len(source.DomainCerts) > 0 {
		rootCAs, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}

		for _, cert := range source.DomainCerts {
			// append our cert to the system pool
			if ok := rootCAs.AppendCertsFromPEM([]byte(cert)); !ok {
				return nil, fmt.Errorf("failed to append registry certificate: %w", err)
			}
		}

		// trust the augmented cert pool in our client
		config := &tls.Config{
			RootCAs: rootCAs,
		}

		tr.TLSClientConfig = config
	}

	rt, err := transport.New(repo.Registry, auth, tr, scopes)
	if err != nil {
		return nil, fmt.Errorf("initialize transport: %w", err)
	}

	return []remote.Option{remote.WithAuth(auth), remote.WithTransport(rt)}, nil
}
*/
