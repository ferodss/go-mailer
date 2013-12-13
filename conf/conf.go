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

package conf

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "path"
    "path/filepath"
)

// Constants
const (
    // Path for the configuration file
    Path    = "./config.conf"
)

// A smtpConfig represents a configuration to connect into SMTP server.
type smtpConfig struct {
    Host    string
    Port    int
    User    string
    Pass    string
}

// A mailConfig represents a configuration for send emails.
type mailConfig struct {
    From        string
    FromName    string
}

// A Config represents an active configuration object.
type Config struct {
    Smtp        smtpConfig      // SMTP options
    Mail        mailConfig      // Mail options
    QueueDir    string          // Queue dir
    WaitFor     string          // Waiting time between check queue
}

// New creates a new Config
// Stop process when config file can't be read or
// not is a valid JSON format.
func New() *Config {
    c := new(Config)

    // Defaults
    d, _ := os.Getwd()
    c.QueueDir = filepath.Clean(d + "/tmp")

    file, err := os.Open(Path)
    if err != nil {
        log.Printf("Error: could not read config file %s\n", Path)
        log.Fatal(err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(c)
    if err != nil {
        log.Printf("Error decoding file %s\n", Path)
        log.Fatal(err)
    }

    if false == path.IsAbs(c.QueueDir) {
        log.Fatal("QueueDir must be an absolute path")
    }

    return c
}

var conf = New()


// Converts a Config object to string
func (c *Config) String() string {
    s := "Config:"

    s += fmt.Sprint("\n  Smtp:\n")
    s += fmt.Sprintf("    Host: %s\n", c.Smtp.Host)
    s += fmt.Sprintf("    Port: %d\n", c.Smtp.Port)
    s += fmt.Sprintf("    User: %s\n", c.Smtp.User)
    s += fmt.Sprintf("    Pass: %s\n", c.Smtp.Pass)
    s += fmt.Sprint("\n")

    s += fmt.Sprint("  Mail:\n")
    s += fmt.Sprintf(  "    From: %s\n", c.Mail.From)
    s += fmt.Sprintf(  "    FromName: %s\n", c.Mail.FromName)
    s += fmt.Sprint("\n")

    s += fmt.Sprintf(  "    QueueDir: %s\n", c.QueueDir)
    s += fmt.Sprintf(  "    WaitFor: %s\n", c.WaitFor)

    return s
}

// Access to SMTP configuration
func Smtp() smtpConfig {
    return conf.Smtp
}

// Access to Config.waitFor
func WaitFor() string {
    return conf.WaitFor
}

// Access to Config.QueueDir
func QueueDir() string {
    return conf.QueueDir
}

// Access to Config.Mail.FromName
func MailFromName() string {
    return conf.Mail.FromName
}

// Access to Config.Mail.From
func MailFrom() string {
    return conf.Mail.From
}

// Converts a Config object to string
func String() string {
    return conf.String()
}
