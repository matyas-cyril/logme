package logme

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

// isInstance fonction permettant de tester le type de variable
// @param1 any : Objet a testé
// @param2 string : Type de référence
// @return bool
//
//	true  : le type est identique
//	false : le type est différent
//
// Exemple : isIntance(data,"string")
//
//	Permet de vérifier que data est un String
func isInstance(object any, typeRef string) bool {

	if object == nil {
		return false
	}

	var objectType string
	valueOf := reflect.ValueOf(object)
	if valueOf.Type().Kind() == reflect.Ptr {
		objectType = reflect.Indirect(valueOf).Type().Name()
	} else {
		objectType = valueOf.Type().Name()
	}

	return objectType == typeRef
}

func (l *LogMe) logTermPrintln(id MsgID, text ...string) {

	tag := ""
	msgId := ""

	if len(l.opt.tag) != 0 {
		tag = fmt.Sprintf(" %s[%d]:", l.opt.tag, os.Getpid())
	}

	if l.msgID.enable && len(id) != 0 {
		msgId = fmt.Sprintf(" %s:", string(id))
	}

	for _, t := range text {

		t = strings.TrimSpace(t)
		if len(t) == 0 {
			continue
		}
		now := time.Now()
		fmt.Printf("%s%s%s %s\n", now.Format("2006-01-02 15:04:05"), tag, msgId, t)

	}

}

func (l *LogMe) logSyslog(id MsgID, p LogPriority, text string) error {

	msgId := ""
	if len(id) != 0 {
		msgId = fmt.Sprintf("%s: ", string(id))
	}

	log := fmt.Sprintf("%s%s", msgId, text)

	switch p {

	case LOGME_P_INFO:
		return l.writer.Info(log)

	case LOGME_P_ERR:
		return l.writer.Err(log)

	case LOGME_P_WARNING:
		return l.writer.Warning(log)

	case LOGME_P_DEBUG:
		return l.writer.Debug(log)

	case LOGME_P_NOTICE:
		return l.writer.Notice(log)

	case LOGME_P_CRIT:
		return l.writer.Crit(log)

	case LOGME_P_ALERT:
		return l.writer.Alert(log)

	case LOGME_P_EMERG:
		return l.writer.Emerg(log)

	}

	return fmt.Errorf("facility not valid")
}

func (l *LogMe) log(id MsgID, p LogPriority, text ...string) {

	if l.opt.print == LOGME_NO {
		return
	}

	for _, t := range text {

		t = strings.TrimSpace(t)
		if len(t) == 0 {
			continue
		}

		switch l.opt.print {

		case LOGME_TERM:
			l.logTermPrintln(id, text...)

		case LOGME_SYSLOG:
			l.logSyslog(id, p, t)

		case LOGME_BOTH:

			l.logTermPrintln(id, text...)
			l.logSyslog(id, p, t)

		}

	}

}
