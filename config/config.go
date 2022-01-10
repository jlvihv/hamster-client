package config

const (
	// HttpGetResource get my resources
	HttpGetResource = "http://42.192.53.10:8888/api/resources/use"
	// HttpGetCode get verification code
	HttpGetCode = "https://ttchain.tntlinking.com/api/authorization/verification/code"
	// HttpLogin login
	HttpLogin = "https://ttchain.tntlinking.com/api/authorization/oauth/token"

	//Port p2p port configuration
	Port = 4001

	// p2p swarm key
	SwarmKey = "/key/swarm/psk/1.0.0/\n/base16/\n55158d9b6b7e5a8e41aa8b34dd057ff1880e38348613d27ae194ad7c5b9670d7"
)
