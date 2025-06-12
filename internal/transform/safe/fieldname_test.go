package safe

import (
	"testing"

	"github.com/ufukty/gonfique/internal/files/config"
)

func TestFieldnameFromUppercase(t *testing.T) {
	tcs := []struct {
		input string
		want  config.Fieldname
	}{
		{"TITLE", "Title"},
		{"TOKEN", "Token"},
		{"TRANSITIONS", "Transitions"},
		{"TUTORIALS", "Tutorials"},
		{"TYPOGRAPHY", "Typography"},
		{"UDP", "Udp"},
		{"USER", "User"},
		{"USERNAME", "Username"},
		{"VERBOSE", "Verbose"},
		{"VERSION", "Version"},
		{"VIDEOS", "Videos"},
		{"VPN", "Vpn"},
		{"WARNINGS", "Warnings"},
		{"WIDGET", "Widget"},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			if got := FieldName(tc.input); tc.want != got {
				t.Errorf("titleCase(%q) got %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestFieldnameFromLowercase(t *testing.T) {
	tcs := []struct {
		input string
		want  config.Fieldname
	}{
		{"address", "Address"},
		{"age", "Age"},
		{"alerts", "Alerts"},
		{"algorithm", "Algorithm"},
		{"animations", "Animations"},
		{"audios", "Audios"},
		{"avatar", "Avatar"},
		{"bio", "Bio"},
		{"birthdate", "Birthdate"},
		{"branding", "Branding"},
		{"cache", "Cache"},
		{"cdn", "Cdn"},
		{"certificate", "Certificate"},
		{"city", "City"},
		{"component", "Component"},
		{"compression", "Compression"},
		{"confirmations", "Confirmations"},
		{"contact", "Contact"},
		{"cookie", "Cookie"},
		{"cors", "Cors"},
		{"country", "Country"},
		{"currency", "Currency"},
		{"database", "Database"},
		{"debug", "Debug"},
		{"description", "Description"},
		{"directory", "Directory"},
		{"disabled", "Disabled"},
		{"dns", "Dns"},
		{"effects", "Effects"},
		{"email", "Email"},
		{"enabled", "Enabled"},
		{"encoding", "Encoding"},
		{"encryption", "Encryption"},
		{"errors", "Errors"},
		{"extension", "Extension"},
		{"faq", "Faq"},
		{"favicon", "Favicon"},
		{"feature", "Feature"},
		{"feedback", "Feedback"},
		{"firewall", "Firewall"},
		{"fonts", "Fonts"},
		{"footer", "Footer"},
		{"format", "Format"},
		{"game", "Game"},
		{"gateway", "Gateway"},
		{"gender", "Gender"},
		{"graphql", "Graphql"},
		{"group", "Group"},
		{"guides", "Guides"},
		{"header", "Header"},
		{"help", "Help"},
		{"host", "Host"},
		{"http", "Http"},
		{"https", "Https"},
		{"icons", "Icons"},
		{"images", "Images"},
		{"instructions", "Instructions"},
		{"ip", "Ip"},
		{"kerberos", "Kerberos"},
		{"key", "Key"},
		{"keywords", "Keywords"},
		{"language", "Language"},
		{"layout", "Layout"},
		{"ldap", "Ldap"},
		{"locale", "Locale"},
		{"logo", "Logo"},
		{"messages", "Messages"},
		{"migration", "Migration"},
		{"module", "Module"},
		{"nationality", "Nationality"},
		{"navbar", "Navbar"},
		{"nonce", "Nonce"},
		{"notification", "Notification"},
		{"notifications", "Notifications"},
		{"oauth", "Oauth"},
		{"password", "Password"},
		{"permission", "Permission"},
		{"phone", "Phone"},
		{"plugin", "Plugin"},
		{"poll", "Poll"},
		{"port", "Port"},
		{"privacy", "Privacy"},
		{"profile", "Profile"},
		{"prompt", "Prompt"},
		{"protocol", "Protocol"},
		{"proxy", "Proxy"},
		{"quiet", "Quiet"},
		{"quiz", "Quiz"},
		{"radius", "Radius"},
		{"region", "Region"},
		{"rest", "Rest"},
		{"retry", "Retry"},
		{"role", "Role"},
		{"routes", "Routes"},
		{"rpc", "Rpc"},
		{"saml", "Saml"},
		{"schema", "Schema"},
		{"scope", "Scope"},
		{"security", "Security"},
		{"server", "Server"},
		{"session", "Session"},
		{"sidebar", "Sidebar"},
		{"signature", "Signature"},
		{"skin", "Skin"},
		{"soap", "Soap"},
		{"sounds", "Sounds"},
		{"ssl", "Ssl"},
		{"sslmode", "Sslmode"},
		{"state", "State"},
		{"subnet", "Subnet"},
		{"support", "Support"},
		{"survey", "Survey"},
		{"tcp", "Tcp"},
		{"template", "Template"},
		{"theme", "Theme"},
		{"throttling", "Throttling"},
		{"timeout", "Timeout"},
		{"timezone", "Timezone"},
		{"tips", "Tips"},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			if got := FieldName(tc.input); tc.want != got {
				t.Errorf("titleCase(%q) got %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestFieldnameFromCamelCase(t *testing.T) {
	tcs := []struct {
		input string
		want  config.Fieldname
	}{
		{"webSocket", "WebSocket"},
		{"accessToken", "AccessToken"},
		{"apiGateway", "ApiGateway"},
		{"apiKey", "ApiKey"},
		{"appEnvironment", "AppEnvironment"},
		{"appName", "AppName"},
		{"appUrl", "AppUrl"},
		{"appVersion", "AppVersion"},
		{"clientId", "ClientId"},
		{"clientSecret", "ClientSecret"},
		{"colorScheme", "ColorScheme"},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			if got := FieldName(tc.input); tc.want != got {
				t.Errorf("titleCase(%q) got %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestFieldnameFromSnakeCase(t *testing.T) {
	tcs := []struct {
		input string
		want  config.Fieldname
	}{
		{"conn_max_lifetime", "ConnMaxLifetime"},
		{"db_driver", "DbDriver"},
		{"db_host", "DbHost"},
		{"db_name", "DbName"},
		{"db_Password", "DbPassword"},
		{"db_port", "DbPort"},
		{"db_url", "DbUrl"},
		{"db_User", "DbUser"},
		{"file_path", "FilePath"},
		{"grant_type", "GrantType"},
		{"load_Balancer", "LoadBalancer"},
		{"log_level", "LogLevel"},
		{"max_connections", "MaxConnections"},
		{"max_idle_conns", "MaxIdleConns"},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			if got := FieldName(tc.input); tc.want != got {
				t.Errorf("titleCase(%q) got %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestFieldnameFromKebabCase(t *testing.T) {
	tcs := []struct {
		input string
		want  config.Fieldname
	}{
		{"max-open-conns", "MaxOpenConns"},
		{"password-policy", "PasswordPolicy"},
		{"pool-size", "PoolSize"},
		{"query-logging", "QueryLogging"},
		{"rate-limit", "RateLimit"},
		{"redirect-uri", "RedirectUri"},
		{"refresh-token", "RefreshToken"},
		{"response-type", "ResponseType"},
		{"table-prefix", "TablePrefix"},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			if got := FieldName(tc.input); tc.want != got {
				t.Errorf("titleCase(%q) got %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestFieldnameFromKubernetesConfig(t *testing.T) {
	tcs := []struct {
		input string
		want  config.Fieldname
	}{
		{"apiVersion", "ApiVersion"},
		{"app", "App"},
		{"backend", "Backend"},
		{"configMapRef", "ConfigMapRef"},
		{"containerPort", "ContainerPort"},
		{"containers", "Containers"},
		{"data", "Data"},
		{"envFrom", "EnvFrom"},
		{"host", "Host"},
		{"http", "Http"},
		{"image", "Image"},
		{"kind", "Kind"},
		{"labels", "Labels"},
		{"matchLabels", "MatchLabels"},
		{"metadata", "Metadata"},
		{"my-key", "MyKey"},
		{"name", "Name"},
		{"namespace", "Namespace"},
		{"number", "Number"},
		{"password", "Password"},
		{"path", "Path"},
		{"paths", "Paths"},
		{"pathType", "PathType"},
		{"port", "Port"},
		{"ports", "Ports"},
		{"protocol", "Protocol"},
		{"replicas", "Replicas"},
		{"rules", "Rules"},
		{"secretRef", "SecretRef"},
		{"selector", "Selector"},
		{"service", "Service"},
		{"spec", "Spec"},
		{"targetPort", "TargetPort"},
		{"template", "Template"},
		{"type", "Type"},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			if got := FieldName(tc.input); tc.want != got {
				t.Errorf("titleCase(%q) got %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
