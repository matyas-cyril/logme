package logme

import (
	"crypto/md5"
	"fmt"
	"log"
	"log/syslog"
	"math/rand"
	"strings"
	"time"
)

// New "constructeur" du module logme
// args : arguments utiles à la déclaration
//
//	options :
//	   tag (string): OBLIGATOIRE
//	                 ajouter un nom identifiant le programme dans les logs
//
//	   length (int): définir le nombre de caractères du messageID [0;32]
//	                 défaut: 10
//
//	   logger (logPrint): façon d'afficher les logs (terminal, syslog, les deux, aucun)
//	                      défaut: LOGME_NO
//
//	   facility (logFacility): sous système applicatif dont les logs sont associés
//	                           défaut: LOGME_F_SYSLOG
//
// @return: OK: LogMe{}, nil
//
//	KO: nil, Error
func New(args map[string]interface{}) (*LogMe, error) {

	logger := LogMe{}

	// Valeurs par défaut
	logger.msgID.length = 10
	logger.opt.print = LOGME_NO
	logger.opt.facility = LOGME_F_SYSLOG

	// Parcourir les options
	for k, v := range args {

		if k == "length" {

			if !(isInstance(v, "int")) {
				return nil, fmt.Errorf("arg '%s' must be an integer", k)
			}

			var len int = v.(int)

			if len < 0 || len > 32 {
				return nil, fmt.Errorf("arg '%s' must be an integer in [0;32]", k)
			}

			logger.msgID.length = uint8(len)

		} else if k == "tag" {
			if !(isInstance(v, "string")) {
				return nil, fmt.Errorf("arg '%s' must be a string", k)
			}

			name := strings.TrimSpace(v.(string))
			if len(name) == 0 {
				return nil, fmt.Errorf("arg '%s' must be a string not null", k)
			}

			logger.opt.tag = name

		} else if k == "logger" {

			if !(isInstance(v, "LogPrint")) {
				return nil, fmt.Errorf("arg '%s' must be a LogPrint", k)
			}

			logger.opt.print = v.(LogPrint)

		} else if k == "facility" {
			if !(isInstance(v, "LogFacility")) {
				return nil, fmt.Errorf("arg '%s' must be a LogFacility", k)
			}
			logger.opt.facility = v.(LogFacility)

		} else {
			return nil, fmt.Errorf("arg '%s' not available", k)
		}

	}

	if len(logger.opt.tag) == 0 {
		return nil, fmt.Errorf("arg 'tag' must be defined")
	}

	// Prise en compte de syslog
	if logger.opt.print == LOGME_SYSLOG || logger.opt.print == LOGME_BOTH {

		logWriter, err := syslog.New(syslog.Priority(logger.opt.facility), logger.opt.tag)
		if err != nil {
			return nil, err
		}

		logger.writer = logWriter
		log.SetFlags(0)          // Laisse l'horodatage à Syslog
		log.SetOutput(logWriter) // On redirige dans Syslog
	}

	return &logger, nil
}

// MessageID génére un identifiant unique.
// Le nombre de caractères est défini à l'initialisation via l'argument 'len'
// @return msgID : identifiant généré
func (l *LogMe) MessageID() MsgID {

	newMsgID := ""
	if l.msgID.length > 0 {

		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		msgID := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v.%v", rnd.NormFloat64(), rnd.NormFloat64()))))
		newMsgID = strings.ToUpper(msgID[:l.msgID.length])

	}
	return MsgID(newMsgID)
}

// SetPrint modifier la sortie d'affichage des logs
func (l *LogMe) SetPrint(print LogPrint) {
	l.mu.Lock()
	l.opt.print = print
	l.mu.Unlock()
}

// GetPrint obtenir la sortie d'affichage active
func (l *LogMe) GetPrint() LogPrint {
	return l.opt.print
}

// DisableMessageID fonction désactivant le messageID dans les logs
func (l *LogMe) DisableMessageID() {
	l.mu.Lock()
	l.msgID.enable = false
	l.mu.Unlock()
}

// EnableMessageID fonction activant le messageID dans les logs
func (l *LogMe) EnableMessageID() {
	l.mu.Lock()
	l.msgID.enable = true
	l.mu.Unlock()
}

// IsEnabledMessageID fonction permettant de savoir si le MessageID est activé dans les logs
// @return bool
func (l *LogMe) IsEnabledMessageID() bool {
	return l.msgID.enable
}

// IsDisabledMessageID fonction permettant de savoir si le MessageID est désactivé dans les logs
// @return bool
func (l *LogMe) IsDisabledMessageID() bool {
	return !l.msgID.enable
}

// Close ferme la connexion au daemon Syslog
func (l *LogMe) Close() error {

	if l.writer == nil {
		return nil
	}

	return l.writer.Close()
}

// Debug priorité du log
func (l *LogMe) Debug(id MsgID, text ...string) {
	l.log(id, LOGME_P_DEBUG, text...)
}

// Info priorité du log
func (l *LogMe) Info(id MsgID, text ...string) {
	l.log(id, LOGME_P_INFO, text...)
}

// Notice priorité du log
func (l *LogMe) Notice(id MsgID, text ...string) {
	l.log(id, LOGME_P_NOTICE, text...)
}

// Warning priorité du log
func (l *LogMe) Warning(id MsgID, text ...string) {
	l.log(id, LOGME_P_WARNING, text...)
}

// Error priorité du log
func (l *LogMe) Error(id MsgID, text ...string) {
	l.log(id, LOGME_P_ERR, text...)
}

// Critical priorité du log
func (l *LogMe) Critical(id MsgID, text ...string) {
	l.log(id, LOGME_P_CRIT, text...)
}

// Alert priorité du log
func (l *LogMe) Alert(id MsgID, text ...string) {
	l.log(id, LOGME_P_ALERT, text...)
}

// Emergency priorité du log
func (l *LogMe) Emergency(id MsgID, text ...string) {
	l.log(id, LOGME_P_EMERG, text...)
}
