package logme

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

// isInstance fonction permettant de tester le type de variable
// @param1 interface{} : Objet a testé
// @param2 string : Type de référence
// @return bool
//
//	true  : le type est identique
//	false : le type est différent
//
// Exemple : isIntance(data,"string")
//
//	Permet de vérifier que data est un String
func isInstance(object interface{}, typeRef string) bool {

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
		fmt.Println(fmt.Sprintf("%s%s%s %s", now.Format("2006-01-02 15:04:05"), tag, msgId, t))

	}

}

func (l *LogMe) logSyslog(id MsgID, p LogPriority, text string) error {

	msgId := ""
	if len(id) != 0 {
		msgId = fmt.Sprintf("%s: ", string(id))
	}

	log := fmt.Sprintf("%s%s", msgId, text)

	var err error = nil

	if p == LOGME_P_INFO {
		err = l.writer.Info(log)

	} else if p == LOGME_P_ERR {
		err = l.writer.Err(log)

	} else if p == LOGME_P_WARNING {
		err = l.writer.Warning(log)

	} else if p == LOGME_P_DEBUG {
		err = l.writer.Debug(log)

	} else if p == LOGME_P_NOTICE {
		err = l.writer.Notice(log)

	} else if p == LOGME_P_CRIT {
		err = l.writer.Crit(log)

	} else if p == LOGME_P_ALERT {
		err = l.writer.Alert(log)

	} else if p == LOGME_P_EMERG {
		err = l.writer.Emerg(log)

	} else {
		err = fmt.Errorf("facility not valid")
	}

	return err
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

		if l.opt.print == LOGME_TERM {
			l.logTermPrintln(id, text...)

		} else if l.opt.print == LOGME_SYSLOG {
			l.logSyslog(id, p, t)

		} else if l.opt.print == LOGME_BOTH {

			l.logTermPrintln(id, text...)
			l.logSyslog(id, p, t)

		}

	}

}
