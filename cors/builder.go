package cors

import (
	"net/http"
	"strconv"
	"strings"
)

// OptionsBuilder facilitates building CORS options.
type OptionsBuilder struct {
	options Options
	origins []string
}

// Configure inits CORS OptionsBuilder.
func Configure() *OptionsBuilder {
	return &OptionsBuilder{
		origins: make([]string, 0),
		options: Options{SuccessStatusCode: http.StatusNoContent}}
}

// SuccessStatusCode sets a custom status code returned on CORS Options request.
// If none is specified, the default status code is http.StatusNoContent.
func (b *OptionsBuilder) SuccessStatusCode(code int) *OptionsBuilder {
	b.options.SuccessStatusCode = code
	return b
}

// MaxAge sets CORS max age.
func (b *OptionsBuilder) MaxAge(maxAge int) *OptionsBuilder {
	b.options.MaxAge = maxAge
	return b
}

// AllowOrigin sets allowed origins.
func (b *OptionsBuilder) AllowOrigin(origin ...string) *OptionsBuilder {
	b.origins = append(b.origins, origin...)
	return b
}

// AllowCredentials sets "Access-Control-Allow-Credentials" header.
func (b *OptionsBuilder) AllowCredentials(allow bool) *OptionsBuilder {
	b.options.AllowCredentials = strconv.FormatBool(allow)
	return b
}

// ExposeHeaders sets "Access-Control-Expose-Header" header.
func (b *OptionsBuilder) ExposeHeaders(headers ...string) *OptionsBuilder {
	b.options.ExposeHeaders = strings.Join(headers, ",")
	return b
}

// AllowedHeaders sets allowed headers.
// It will set the header "Access-Control-Allow-Header".
func (b *OptionsBuilder) AllowedHeaders(headers ...string) *OptionsBuilder {
	b.options.AllowedHeaders = strings.Join(headers, ",")
	return b
}

// AllowMethods sets the allowed HTTP methods.
// The header "Access-Control-Allow-Methods" will be used.
func (b *OptionsBuilder) AllowMethods(methods ...string) *OptionsBuilder {
	b.options.AllowedMethods = strings.Join(methods, ",")
	return b
}

// Build returns an Option with previously configured values.
func (b *OptionsBuilder) Build() Options {
	if len(b.origins) > 0 {
		if len(b.origins) == 1 {
			b.options.AllowedOrigin = b.origins[0]
		} else {
			b.options.AllowedOrigin = strings.Join(b.origins, ",")
		}
	}

	return b.options
}
