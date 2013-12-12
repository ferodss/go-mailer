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
    "time"

	"github.com/felipedjinn/go-mailer/conf"
)

// A Queue represents an active queue object.
type Queue struct {
    Files   []string
    Dir     string
}

// New creates a new Queue
// Stop process when queue directory dont exists.
func New() *Queue {
    q := new(Queue)
    q.Dir = conf.QueueDir()

    if _, err := os.Stat(q.Dir); os.IsNotExist(err) {
        log.Fatal(fmt.Sprintf("Queue dir \"%s\" not found", q.Dir))
    }

    return q
}

// HasQueue checks if has files in queue directory and
// append this files in Queue.Files.
func (q *Queue) HasQueue() (hasQueue bool, err error) {
    markFn := func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if path[:1] == "." || info.IsDir() || filepath.Ext(path) != ".json" {
            return nil
        }

        if ok, _ := regexp.MatchString("^[0-9]+\\.json$", info.Name()); ok {
            hasQueue = true
            q.Files = append(q.Files, path)
        }

        return nil
    }

    return hasQueue, filepath.Walk(q.Dir, markFn)
}

// Process a file in argument.
func (q *Queue) Process(file string) {
    log.Printf("Processing file: %s\n", file)

    // TODO

    wait, _ := time.ParseDuration("12s")
    time.Sleep(wait)
    fmt.Println("Finished")
}

