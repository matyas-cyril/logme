package logme

import (
	"log/syslog"
	"sync"
)

type LogMe struct {
	msgID  messageID
	writer *syslog.Writer
	opt    options
	mu     sync.Mutex
}

type messageID struct {
	length uint8 // La valeur max ne doit pas depasser 32
	enable bool  // Même si défini le messageID est-il activé
}

type options struct {
	tag      string // Identifier le programme dans syslog
	facility LogFacility
	print    LogPrint // Afficher le log
}

type LogPrint uint
type LogFacility syslog.Priority
type LogPriority syslog.Priority
type MsgID string

const (
	MSGID_EMPTY MsgID = ""

	LOGME_NO     LogPrint = 0 // Log désactivé
	LOGME_TERM   LogPrint = 1 // Log en sortie standard
	LOGME_SYSLOG LogPrint = 2 // Log dans syslog
	LOGME_BOTH   LogPrint = 3 // Log en sortie standard et syslog

	// Niveau du priorité du log (classement par ordre croissant)
	LOGME_P_DEBUG   LogPriority = LogPriority(syslog.LOG_DEBUG)
	LOGME_P_INFO    LogPriority = LogPriority(syslog.LOG_INFO)
	LOGME_P_NOTICE  LogPriority = LogPriority(syslog.LOG_NOTICE)
	LOGME_P_WARNING LogPriority = LogPriority(syslog.LOG_WARNING)
	LOGME_P_ERR     LogPriority = LogPriority(syslog.LOG_ERR)
	LOGME_P_CRIT    LogPriority = LogPriority(syslog.LOG_CRIT)
	LOGME_P_ALERT   LogPriority = LogPriority(syslog.LOG_ALERT)
	LOGME_P_EMERG   LogPriority = LogPriority(syslog.LOG_EMERG)

	// Sous système applicatif dont les logs sont associés
	LOGME_F_AUTH   LogFacility = LogFacility(syslog.LOG_AUTH)   // Authentification
	LOGME_F_CRON   LogFacility = LogFacility(syslog.LOG_CRON)   // Cron
	LOGME_F_DAEMON LogFacility = LogFacility(syslog.LOG_DAEMON) // Daemon sans classification
	LOGME_F_FTP    LogFacility = LogFacility(syslog.LOG_FTP)    // FTP
	LOGME_F_KERN   LogFacility = LogFacility(syslog.LOG_KERN)   // Kernel
	LOGME_F_LPR    LogFacility = LogFacility(syslog.LOG_LPR)    // Impression
	LOGME_F_MAIL   LogFacility = LogFacility(syslog.LOG_MAIL)   // Mail
	LOGME_F_NEWS   LogFacility = LogFacility(syslog.LOG_NEWS)   // Service Usenet, NNTP...
	LOGME_F_SYSLOG LogFacility = LogFacility(syslog.LOG_SYSLOG) // Syslog
	LOGME_F_USER   LogFacility = LogFacility(syslog.LOG_USER)   // Utilisateur générique
	LOGME_F_UUCP   LogFacility = LogFacility(syslog.LOG_UUCP)   // Système Unix to Unix Copy Program
	LOGME_F_LOCAL0 LogFacility = LogFacility(syslog.LOG_LOCAL0) // Utilisateurs Locaux
	LOGME_F_LOCAL1 LogFacility = LogFacility(syslog.LOG_LOCAL1)
	LOGME_F_LOCAL2 LogFacility = LogFacility(syslog.LOG_LOCAL2)
	LOGME_F_LOCAL3 LogFacility = LogFacility(syslog.LOG_LOCAL3)
	LOGME_F_LOCAL4 LogFacility = LogFacility(syslog.LOG_LOCAL4)
	LOGME_F_LOCAL5 LogFacility = LogFacility(syslog.LOG_LOCAL5)
	LOGME_F_LOCAL6 LogFacility = LogFacility(syslog.LOG_LOCAL6)
	LOGME_F_LOCAL7 LogFacility = LogFacility(syslog.LOG_LOCAL7)
)
