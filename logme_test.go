package logme_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/matyas-cyril/logme"
)

func TestLog(t *testing.T) {

	var wg sync.WaitGroup

	args := map[string]interface{}{
		"tag":      "monPrgmTest",
		"length":   20,                 // Définir un MessageID de 20 caractères
		"logger":   logme.LOGME_BOTH,   // Affichage Terminal et dans les logs
		"facility": logme.LOGME_F_MAIL, // Écriture dans /var/log/mail.log
	}

	l, e := logme.New(args)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer l.Close()

	l.EnableMessageID()

	for cpt := 1; cpt <= 5; cpt++ {

		go func(i int) {
			defer wg.Done()

			// Nbre aléatoire pour le test
			pause := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5) + 1

			msgId := l.MessageID()
			l.Info(msgId, fmt.Sprintf("Ceci est un test %d - pause: %d", i, pause))
			time.Sleep(time.Duration(pause) * time.Second)
			l.Info(msgId, fmt.Sprintf("Ceci est un test %d.%d", i, i))
		}(cpt)
		wg.Add(1)

	}

	wg.Wait()

}

/*
Résultat :
	2022-04-28 21:27:30 monPrgmTest[38826]: 0ABEB12F8B0BC1669635: Ceci est un test 5 - pause: 3
	2022-04-28 21:27:30 monPrgmTest[38826]: 9F921EE6457E072B8B45: Ceci est un test 2 - pause: 3
	2022-04-28 21:27:30 monPrgmTest[38826]: DE14FAFFA6BDE2FA1BFA: Ceci est un test 1 - pause: 4
	2022-04-28 21:27:30 monPrgmTest[38826]: 3DCF4B396E2EA52FB471: Ceci est un test 4 - pause: 2
	2022-04-28 21:27:30 monPrgmTest[38826]: 3F33A1EEDC39D5335A60: Ceci est un test 3 - pause: 3
	2022-04-28 21:27:32 monPrgmTest[38826]: 3DCF4B396E2EA52FB471: Ceci est un test 4.4
	2022-04-28 21:27:33 monPrgmTest[38826]: 0ABEB12F8B0BC1669635: Ceci est un test 5.5
	2022-04-28 21:27:33 monPrgmTest[38826]: 3F33A1EEDC39D5335A60: Ceci est un test 3.3
	2022-04-28 21:27:33 monPrgmTest[38826]: 9F921EE6457E072B8B45: Ceci est un test 2.2
	2022-04-28 21:27:34 monPrgmTest[38826]: DE14FAFFA6BDE2FA1BFA: Ceci est un test 1.1
*/
