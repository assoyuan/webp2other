## 起因：

设计师们：`从网上保存的图片是 webp 格式的，在 PS 里用不了，怎么保存 jpg 格式呢？`

我：`那就写个工具，解决你们的问题吧`

## 示例：

main.go
```golang
package main

import w2o webp2other

func main(){
    w2o.Exec()
}
```
> 把`main.go`打包成可执行文件，双击运行