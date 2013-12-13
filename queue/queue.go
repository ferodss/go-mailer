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

package queue

import (
    "fmt"
    "os"
    "regexp"
    "path/filepath"
    "log"

	"github.com/felipedjinn/go-mailer/conf"
	"github.com/felipedjinn/go-mailer/message"
	"github.com/felipedjinn/go-mailer/mailer"
)

// A Queue represents an active queue object.
type Queue struct {
    Files   []string
    Dir     string
}

// New creates a new Queue
// Return error when queue directory dont exists.
func New() *Queue {
    q := new(Queue)
    q.Dir = conf.QueueDir()

    //if _, err := os.Stat(q.Dir); os.IsNotExist(err) {
    //    return nil, fmt.Errorf("Queue dir \"%s\" not found", q.Dir)
    //}

    return q
}

// HasQueue checks if has files in queue directory and
// append this files in Queue.Files.
func (q *Queue) HasQueue() (hasQueue bool, err error) {
    dir, err := os.Open(q.Dir)
    if err != nil && os.IsNotExist(err) {
        return false, fmt.Errorf("Queue dir \"%s\" not found", q.Dir)
    }
    defer dir.Close()

    fileInfos, _ := dir.Readdir(-1)
    for _, fi := range fileInfos {
        if fi.IsDir() {
            continue
        }

        if ok, _ := regexp.MatchString("^[0-9]+\\.json$", fi.Name()); ok {
            hasQueue = true
            f, _ := filepath.Abs(q.Dir + string(os.PathSeparator) + fi.Name())
            q.Files = append(q.Files, f)
        }
    }

    return hasQueue, nil
}

// Process a file in argument.
func (q *Queue) Process(file string) {
    log.Printf("Processing file: %s\n", file)

    message, err := message.New(file);
    if err != nil {
        log.Printf(err.Error())
        return
    }

    // TODO: Remove file from queue when mail.Send return TRUE
    mailer.Send(message)
}
