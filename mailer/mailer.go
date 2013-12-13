// Copyright © 2013 Felipe Rodrigues <lfrs.web@gmail.com>.
//
// Licensed under the Simple Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://opensource.org/licenses/Simple-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mailer

import (
    "encoding/base64"
    "fmt"
    "log"
    "net/mail"
    "net/smtp"
    "strconv"
    "strings"

    "github.com/felipedjinn/go-mailer/conf"
    "github.com/felipedjinn/go-mailer/message"
)

func Send(m *message.Message) bool {
    auth := smtpAuth()
    headers := mailHeader(m)

    body := ""
    for k, v := range headers {
        body += fmt.Sprintf("%s: %s\r\n", k, v)
    }

    // Recipients receives in Bcc
    body += fmt.Sprintf("Bcc: %s\r\n", strings.Join(m.To, ";"))
    body += "\r\n" + base64.StdEncoding.EncodeToString([]byte(m.Body))

    // TODO: Manualy create TSL connection
    err := smtp.SendMail(
        fmt.Sprintf("%s:%s", conf.Smtp().Host, strconv.Itoa(conf.Smtp().Port)),
        auth,
        conf.MailFrom(),
        []string{""},
        []byte(body),
    )
    if err != nil {
        log.Println("SendMail: " + err.Error())
        return false
    }

    return true
}

// Set up authentication information
func smtpAuth() smtp.Auth {
    return smtp.PlainAuth(
        "",
        conf.Smtp().User,
        conf.Smtp().Pass,
        conf.Smtp().Host,
    )
}

// mailHeader returns a map[string] with headers to use
// when send email.
func mailHeader(m *message.Message) map[string]string {
    header := make(map[string]string)

    from := mail.Address{conf.MailFromName(), conf.MailFrom()}
    header["From"] = from.String()
    header["To"] = ""
    header["Subject"] = m.Subject
    header["MIME-Version"] = "1.0"
    header["Content-Type"] = "text/html; charset=\"utf-8\""
    header["Content-Transfer-Encoding"] = "base64"

    return header
}
