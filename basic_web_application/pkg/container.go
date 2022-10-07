package pkg

import "github.com/alexedwards/scs/v2"

type Container struct {
	AppConfig AppConfig
	Session   *scs.SessionManager
}
