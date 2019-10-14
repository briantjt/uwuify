# UwUify Text

Use this library to UwUify your text! Can be used as a cli tool as well. This was created as a random string manipulation exercise for myself.

## Usage

### Import

```go
import "github.com/briantjt/uwuify/uwu"

func main() {
    fmt.Println(uwu.UwUify("Hello there"))
}
```

### CLI

`go install github.com/briantjt/uwuify/`

CLI reads from stdin so `echo` or `cat` will work

```bash
echo "Hello there" | uwuify
Hewwo dewe
```
