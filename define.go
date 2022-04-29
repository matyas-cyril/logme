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
	facility logFacility
	print    logPrint // Afficher le log
}

type logPrint uint
type logFacility syslog.Priority
type logPriority syslog.Priority
type MsgID string

const (
	MSGID_EMPTY MsgID = ""

	LOGME_NO     logPrint = 0 // Log désactivé
	LOGME_TERM   logPrint = 1 // Log en sortie standard
	LOGME_SYSLOG logPrint = 2 // Log dans syslog
	LOGME_BOTH   logPrint = 3 // Log en sortie standard et syslog

	// Niveau du priorité du log (classement par ordre croissant)
	LOGME_P_DEBUG   logPriority = logPriority(syslog.LOG_DEBUG)
	LOGME_P_INFO    logPriority = logPriority(syslog.LOG_INFO)
	LOGME_P_NOTICE  logPriority = logPriority(syslog.LOG_NOTICE)
	LOGME_P_WARNING logPriority = logPriority(syslog.LOG_WARNING)
	LOGME_P_ERR     logPriority = logPriority(syslog.LOG_ERR)
	LOGME_P_CRIT    logPriority = logPriority(syslog.LOG_CRIT)
	LOGME_P_ALERT   logPriority = logPriority(syslog.LOG_ALERT)
	LOGME_P_EMERG   logPriority = logPriority(syslog.LOG_EMERG)

	// Sous système applicatif dont les logs sont associés
	LOGME_F_AUTH   logFacility = logFacility(syslog.LOG_AUTH)   // Authentification
	LOGME_F_CRON   logFacility = logFacility(syslog.LOG_CRON)   // Cron
	LOGME_F_DAEMON logFacility = logFacility(syslog.LOG_DAEMON) // Daemon sans classification
	LOGME_F_FTP    logFacility = logFacility(syslog.LOG_FTP)    // FTP
	LOGME_F_KERN   logFacility = logFacility(syslog.LOG_KERN)   // Kernel
	LOGME_F_LPR    logFacility = logFacility(syslog.LOG_LPR)    // Impression
	LOGME_F_MAIL   logFacility = logFacility(syslog.LOG_MAIL)   // Mail
	LOGME_F_NEWS   logFacility = logFacility(syslog.LOG_NEWS)   // Service Usenet, NNTP...
	LOGME_F_SYSLOG logFacility = logFacility(syslog.LOG_SYSLOG) // Syslog
	LOGME_F_USER   logFacility = logFacility(syslog.LOG_USER)   // Utilisateur générique
	LOGME_F_UUCP   logFacility = logFacility(syslog.LOG_UUCP)   // Système Unix to Unix Copy Program
	LOGME_F_LOCAL0 logFacility = logFacility(syslog.LOG_LOCAL0) // Utilisateurs Locaux
	LOGME_F_LOCAL1 logFacility = logFacility(syslog.LOG_LOCAL1)
	LOGME_F_LOCAL2 logFacility = logFacility(syslog.LOG_LOCAL2)
	LOGME_F_LOCAL3 logFacility = logFacility(syslog.LOG_LOCAL3)
	LOGME_F_LOCAL4 logFacility = logFacility(syslog.LOG_LOCAL4)
	LOGME_F_LOCAL5 logFacility = logFacility(syslog.LOG_LOCAL5)
	LOGME_F_LOCAL6 logFacility = logFacility(syslog.LOG_LOCAL6)
	LOGME_F_LOCAL7 logFacility = logFacility(syslog.LOG_LOCAL7)
)
