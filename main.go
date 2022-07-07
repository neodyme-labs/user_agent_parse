package user_agent_parse

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"

	ua "github.com/mileusna/useragent"
)

func init() {
	caddy.RegisterModule(UserAgentParse{})
}

type UserAgentParse struct {
	logger *zap.Logger
}

// CaddyModule returns the Caddy module information.
func (UserAgentParse) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.user_agent_parse",
		New: func() caddy.Module { return new(UserAgentParse) },
	}
}

func (l *UserAgentParse) Provision(ctx caddy.Context) error {
	l.logger = ctx.Logger(l)

	return nil
}

func (m *UserAgentParse) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	agent := r.Header.Get("User-Agent")
	ua := ua.Parse(agent)

	repl := r.Context().Value(caddy.ReplacerCtxKey).(*caddy.Replacer)
	repl.Set("user_agent.name", ua.Name)
	repl.Set("user_agent.version", ua.Version)
	repl.Set("user_agent.os", ua.OS)
	repl.Set("user_agent.os_version", ua.OSVersion)
	repl.Set("user_agent.device", ua.Device)
	repl.Set("user_agent.mobile", ua.Mobile)
	repl.Set("user_agent.tablet", ua.Tablet)
	repl.Set("user_agent.desktop", ua.Desktop)
	repl.Set("user_agent.bot", ua.Bot)
	repl.Set("user_agent.url", ua.URL)

	return next.ServeHTTP(w, r)
}

// Interface guards.
var (
	_ caddy.Provisioner = (*UserAgentParse)(nil)
	_ caddy.Module      = (*UserAgentParse)(nil)
)
