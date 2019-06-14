package rpc

//RPC connect configure of the omni core wallet
type ConnConfig struct {
	// The format is <host|ip:[port]>
	Host string
	// Give a username for Authentication
	User string
	// Give a password for Authentication
	Pass string
}

func NewConnConfig(host,user,pass string) *ConnConfig{
	connCfg := &ConnConfig{
		Host: host,
		User: user,
		Pass: pass,
	}
	return connCfg
}
