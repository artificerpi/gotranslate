## gotranslate (forked)


## Usage
* import the gotranslate package
``` go
import (
	"github.com/artificerpi/gotranslate"
)
```
* Example
``` go
// auto detecting from language type
gotranslate.QuickTranslate("Hello", language.Chinese)
// output: 你好

gotranslate.QuickTranslation("Hello", language.English, language.Chinese)
// output: 你好
```

## Original Repo
[eefret/gotranslate](https://github.com/eefret/gotranslate)