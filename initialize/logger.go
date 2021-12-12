package initialize

import "log"

// TODO: log to multiple writer
func InitLogger() {
	log.SetPrefix("[BLOG] ")
	log.SetFlags(log.Lmsgprefix)

}
