# Introduction

将FBX转换成tileset的工具

使用需要用到fbxsdk(libfbxsdk.so),可从官网下载 https://www.autodesk.com/developer-network/platform-technologies/fbx-sdk-2019-2.

# Usage
```shell
usage: fbx-to-tileset --input=INPUT [<flags>]

Flags:
      --help                  Show context-sensitive help (also try --help-long and --help-man).
  -i, --input=INPUT           输入的FBX模型文件
  -o, --output="./out/"       输出切片文件的目录.
  -t, --texture=TEXTURE       贴图文件目录，如果设置了则会进行尺寸优化。如果没有设置，则依据FBX2glTF的规则进行贴图搜寻，且不对贴图做任何处理，搜寻规则为FBX文件目录/.fbm目录/当前工作目录
  -c, --clear                 清理临时目录
      --lng=39.90691          生成切片的经度
      --lat=116.39123         生成切片的纬度
      --height=0              生成切片的高度
  -v, --verticeslimit=500000  单个b3dm文件的最大顶点数
  -m, --minsize=2             单个模型文件的最小尺寸，小于该尺寸的会合并成cmpt文件
```