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

package message

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/felipedjinn/go-mailer/conf"
)

// A Message representes a mail to send
type Message struct {
    To              []string
    ReplayTo        string      `json:replay-to`
    Subject         string
    Body            string
}

// New creates a new Message loading your
// data from path argument JSON file.
func New(path string) (m *Message, err error) {
    m = new(Message)

    // Defaults
    m.ReplayTo = conf.MailFrom()

    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(m)
    if err != nil {
        return nil, err
    }

    return m, nil
}

// Converts a Message object to string
func (m *Message) String() string {
    s := "Message:"
    s += fmt.Sprintf("To       : %v\n", m.To)
    s += fmt.Sprintf("ReplayTo : %s\n", m.ReplayTo)
    s += fmt.Sprintf("Subject  : %s\n", m.Subject)
    s += fmt.Sprintf("Body     : %s\n", m.Body)

    return s
}
