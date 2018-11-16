// Copyright 2015 Light Code Labs, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package caddyhttp

import (
	// plug in the server
	_ "github.com/icasei/caddy/caddyhttp/httpserver"

	// plug in the standard directives
	_ "github.com/icasei/caddy/caddyhttp/basicauth"
	_ "github.com/icasei/caddy/caddyhttp/bind"
	_ "github.com/icasei/caddy/caddyhttp/browse"
	_ "github.com/icasei/caddy/caddyhttp/errors"
	_ "github.com/icasei/caddy/caddyhttp/expvar"
	_ "github.com/icasei/caddy/caddyhttp/extensions"
	_ "github.com/icasei/caddy/caddyhttp/fastcgi"
	_ "github.com/icasei/caddy/caddyhttp/gzip"
	_ "github.com/icasei/caddy/caddyhttp/header"
	_ "github.com/icasei/caddy/caddyhttp/index"
	_ "github.com/icasei/caddy/caddyhttp/internalsrv"
	_ "github.com/icasei/caddy/caddyhttp/limits"
	_ "github.com/icasei/caddy/caddyhttp/log"
	_ "github.com/icasei/caddy/caddyhttp/markdown"
	_ "github.com/icasei/caddy/caddyhttp/mime"
	_ "github.com/icasei/caddy/caddyhttp/pprof"
	_ "github.com/icasei/caddy/caddyhttp/proxy"
	_ "github.com/icasei/caddy/caddyhttp/push"
	_ "github.com/icasei/caddy/caddyhttp/redirect"
	_ "github.com/icasei/caddy/caddyhttp/requestid"
	_ "github.com/icasei/caddy/caddyhttp/rewrite"
	_ "github.com/icasei/caddy/caddyhttp/root"
	_ "github.com/icasei/caddy/caddyhttp/status"
	_ "github.com/icasei/caddy/caddyhttp/templates"
	_ "github.com/icasei/caddy/caddyhttp/timeouts"
	_ "github.com/icasei/caddy/caddyhttp/websocket"
	_ "github.com/icasei/caddy/onevent"
)
