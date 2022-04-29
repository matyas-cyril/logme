# logme

Surcharge du module syslog permettant un rendu des logs dans le terminal et / ou syslog.<br>De plus une personnalisation type "postfix/cyrus" est possible via l'intégration d'un "MessageID".<br>Le MessageID est une chaîne de caractères générée aléatoirement permettant de suivre plus finement une transaction dans les logs.

# Fonctions

### **New(args map[string]interface{}) (\*LogMe, error)** : instancier logme

### **(l \*LogMe) MessageID() msgID** : générer un MessageID.

### **(l \*LogMe) EnableMessageID()** : activer l'utilisation du MessageID. Par défaut, l'utilisation du MessageID est désactivée dans les logs.

### **(l \*LogMe) DisableMessageID()** : desactiver l'utilisation du MessageID

### **(l \*LogMe) IsEnabledMessageID() bool** : le MessageID est-il activé

### **(l \*LogMe) IsDisabledMessageID() bool** : le MessageID est-il desactivé

### **(l \*LogMe) Close() error** : clôturer la connexion au daemon Syslog
<br>

### **Priorités des logs par ordre croissant :**

### **(l \*LogMe) Debug(id msgID, text ...string)**

### **(l \*LogMe) Info(id msgID, text ...string)**

### **(l \*LogMe) Notice(id msgID, text ...string)**

### **(l \*LogMe) Warning(id msgID, text ...string)**

### **(l \*LogMe) Error(id msgID, text ...string)**

### **(l \*LogMe) Critical(id msgID, text ...string)**

### **(l \*LogMe) Alert(id msgID, text ...string)**

### **(l \*LogMe) Emergency(id msgID, text ...string)**
<br>

# Arguments

| CLEF | TYPE | DÉFAUT | DESCRIPTION |
|:----:|:----:|:-------:|:----------:|
| tag | string | | Nom identifiant le prgm dans les logs |
| length | int | 10 | Définie le nombre de caractères du messageID<br>Min:1<br>Max: 32 |
| logger | logPrint | LOGME_NO | Affichage des logs |
| facility | logFacility | LOGME_F_SYSLOG | sous système applicatif dont les logs sont associés |

## logPrint

```go
LOGME_NO      // Log désactivé
LOGME_TERM    // Log en sortie standard sur le terminal
LOGME_SYSLOG  // Log dans syslog
LOGME_BOTH    // Log en sortie terminal et syslog
```

## logFacility

```go
LOGME_F_AUTH    // Authentification
LOGME_F_CRON    // Cron
LOGME_F_DAEMON  // Daemon sans classification
LOGME_F_FTP     // FTP
LOGME_F_KERN    // Kernel
LOGME_F_LPR     // Impression
LOGME_F_MAIL    // Mail
LOGME_F_NEWS    // Service Usenet, NNTP...
LOGME_F_SYSLOG  // Syslog
LOGME_F_USER    // Utilisateur générique
LOGME_F_UUCP    // Système Unix to Unix Copy Program
LOGME_F_LOCAL0  // Utilisateurs Locaux
LOGME_F_LOCAL1
LOGME_F_LOCAL2
LOGME_F_LOCAL3
LOGME_F_LOCAL4
LOGME_F_LOCAL5
LOGME_F_LOCAL6
LOGME_F_LOCAL7
```

```go
MSGID_EMPTY  // Bouchon pour un MessageID vide. A pour effet de désactiver le rendu sans avoir à utiliser la fonction DisableMessageID()
```

# Exemple :

```go
import (
	"github.com/matyas-cyril/logme"
)

// Déclaration avec l'argument obligatoire
args := map[string]interface{}{
    "tag":      "monPrgm",
}

l, err := logme.New(args)
if err != nil {
    return err
}
defer l.Close()

// Activation du rendu du messageID dans les logs
l.EnableMessageID()

// Génération d'un messageID
msgId := l.MessageID()
l.Info(msgId, "Ceci est un test")

Résultat :
   2022-04-28 19:51:53 monPrgm[28930]: FB0C3A364B: Ceci est un test

```
