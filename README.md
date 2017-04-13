# gentab-tool
使用 golang 编写的一个文本生成表格命令行工具, 有了这个小工具再也不用一个一个填表格啦!😁

```
Usage of gentab-tool:
  -c	centering control
  -d string
    	use -d to specify a separator
  -f string
    	use -f to specify a file
  -o string
    	use -o save result to file
  -t	use '\t' as separator
```

### example:

```
➜  gentab-tool git:(master) ✗ cat data #源文件内容 分隔符是\t
Filesystem	Size	Used
/dev/disk0s2	233Gi	111Gi
devfs	187Ki	187Ki
map-hosts	0Bi	0Bi
mapauto_home	0Bi	0Bi
➜  gentab-tool git:(master) ✗ ./tab -f data -t -o out.html #以 data 为输入文件, out.html 为输出文件 \t 为分隔符 转化
➜  gentab-tool git:(master) ✗ open out.html
```

通过网页或 markdown 预览preview:

| Filesystem   | Size  | Used  |
| ------------ | ----- | ----- |
| /dev/disk0s2 | 233Gi | 111Gi |
| devfs        | 187Ki | 187Ki |
| map-hosts    | 0Bi   | 0Bi   |
| mapauto_home | 0Bi   | 0Bi   |

### enjoy it!



