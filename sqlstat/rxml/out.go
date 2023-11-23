package rxml

import (
	"fmt"
	"strings"
)

func GenOutput(s SqlStat) (string, error) {
	var sqlLines []string
	content := fmt.Sprintf("### %s\n", s.Path)
	content = content + fmt.Sprintf("> sql统计:复杂%d个，简单%d个\n>\n", len(s.ComplexSql), len(s.SimpleSql))
	content = content + "> 使用的关键字：" + strings.Join(s.Fs, ",") + "\n"
	for _, cs := range s.ComplexSql {
		contentItem := fmt.Sprintf("#### %s\n", cs.Id)
		contentItem = contentItem + "```\n" + cs.XMLValue + "\n```"
		sqlLines = append(sqlLines, contentItem)
	}
	return content + strings.Join(sqlLines, "\n"), nil
}

/**
# 复杂sql统计
## cloud-cmp-core
### mapper文件地址

#### 使用的关键字

#### selectAll

```
sql内容
```


*/
