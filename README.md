# Lambda-Go

## Overview

Lambda-Go is a library designed to extend Go's capabilities with functional programming techniques inspired by Haskell. It aims to provide Go developers with tools to write cleaner, more expressive code while leveraging Go's strong typing and efficiency.

## Features

- **Core Functional Constructs**: Implements fundamental functional programming operations like `map`, `foldl`, `foldr`, and more.
- **Tuples Support**: Generic tuple data structure for handling paired data.
- **Utilities**: A collection of utilities such as `reverse`, `concat`, and `unique` to manipulate slices.
- **Predicate Functions**: Includes functions like `filter`, `any`, `all`, and `find` to work with predicates effectively.
- **Sorting and Grouping**: Functions for sorting slices and grouping elements based on specified criteria.

## Getting Started

### Prerequisites

- Go 1.18 or higher (requires support for generics)

### Installing

To use Lambda-Go in your project, start by cloning the repository:

```bash
git clone https://github.com/araujo88/lambda-go.git
```

Then, import the required subpackages in your Go application:

```go
import (
    "github.com/araujo88/lambda-go/pkg/core"
    "github.com/araujo88/lambda-go/pkg/utils"
)
```

### Usage Examples

Here are some simple examples to get you started:

**Using `Map` to double the values in a slice:**

```go
package main

import (
    "fmt"
    "github.com/araujo88/lambda-go/pkg/core"
)

func main() {
    slice := []int{1, 2, 3, 4, 5}
    doubled := core.Map(slice, func(x int) int { return x * 2 })
    fmt.Println(doubled)
}
```

**Filtering a slice of integers:**

```go
package main

import (
    "fmt"
    "github.com/araujo88/lambda-go/pkg/predicate"
)

func main() {
    slice := []int{1, 2, 3, 4, 5}
    filtered := predicate.Filter(slice, func(x int) bool { return x > 3 })
    fmt.Println(filtered)
}
```

## Contributing

We welcome contributions! If you have suggestions or want to improve the library, feel free to make a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
