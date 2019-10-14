# UwUify Text

Use this library to UwUify your text! Can be used as a cli tool as well.

This was created as a random string manipulation exercise in Go.

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

cat tragedy.txt | uwuify
Did you evew heaw de twagedy of Dawd Pwagueis De Wise? I dought nowt. Iwt's
nowt a stowy de Jedi wouwd teww you. Iwt's a Siwth wegend. Dawd Pwagueis was a
Dawk Wowd of de Siwth, so powewfuw and so wise he couwd use de Fowce to
infwuence de midichwowians to cweawte wife... He had such a knowwedge of de dawk
side dawt he couwd even keep de ones he cawed abouwt fwom dying. De dawk side
of de Fowce is a pawthway to many abiwiwties some considew to be unnawtuwaw. He
became so powewfuw... de onwy ding he was afwaid of was wosing his powew, which
eventuawwy, of couwse, he did. Unfowtunawtewy, he taught his appwentice
evewyding he knew, den his appwentice kiwwed him in his sweep. Iwonic. He
couwd save owthews fwom deawth, buwt nowt himsewf.

```

## Tests

`go test ./...`
