# semaphoreci

semaphoreci is Go wrapper for [Semaphore](https://semaphoreapp.com/) [official API](http://docs.semaphoreapp.com/api).
You can find Go package documentation [here](http://godoc.org/github.com/mytrile/semaphoreci).

* [User's projects](#users-projects)
* [Project’s branches](#projects-branches)
* [Branch Status](#branch-status)
* [Branch History](#branch-history)
* [Build Information](#build-information)
* [Build Log](#build-log)

## Installation

      $ go install github.com/mytrile/semaphoreci

## Usage

``` go
package main

import (
  "fmt"
  "github.com/mytrile/wunderlist"
)

func main() {
  client  := semaphoreci.NewClient("auth_token")
  project := client.Project(hash_id)
  // ...
}
```

### [User's projects](http://docs.semaphoreapp.com/api#projects)

``` go
projects, err := client.Projects()
```

### [Project’s branches](http://docs.semaphoreapp.com/api#branches)

``` go
branches, err := project.Branches()
```

### [Branch Status](http://docs.semaphoreapp.com/api#branch_status)

``` go
branches, err := project.BranchStatus(branch_id)
```

### [Branch History](http://docs.semaphoreapp.com/api#branch_history)

``` go
branches, err := project.BranchHistory(branch_id)
```

### [Build Information](http://docs.semaphoreapp.com/api#build_information)

``` go
branches, err := project.BuildInfo(branch_id, build_num)
```

### [Build Log](http://docs.semaphoreapp.com/api#build_log)

``` go
branches, err := project.BuildLog(branch_id, build_num)
```

### Meta

* Author  - Dimitar Kostov
* Email   - mitko.kostov@gmail.com
* Blog    - <http://mytrile.github.io>
* Twitter - [mytrile](https://twitter.com/mytrile)

## License

(The MIT License)

Copyright (c) 2013 Dimitar Kostov <mitko.kostov@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
