#### [面试题 01.03. URL化](https://leetcode-cn.com/problems/string-to-url-lcci/)

难度简单8

URL化。编写一种方法，将字符串中的空格全部替换为`%20`。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用`Java`实现的话，请使用字符数组实现，以便直接在数组上操作。）

**示例1:**

```
 输入："Mr John Smith    ", 13
 输出："Mr%20John%20Smith"
```

**示例2:**

```
 输入："               ", 5
 输出："%20%20%20%20%20"
```

**提示：**

1. 字符串长度在[0, 500000]范围内。



------

##### 思路：使用标准库api

将空格部分之间变成%20

```go
func replaceSpaces(S string, length int) string {
    return strings.ReplaceAll(S[:length], " ", "%20")
}
```

------

##### 思路：

获取空格数，创建一个byte数组，按要求往里面添加数据

```go
/*
执行用时 :16 ms, 在所有 Go 提交中击败了94.76%的用户
内存消耗 :8 MB, 在所有 Go 提交中击败了100.00%的用户
*/

func replaceSpaces(S string, length int) string {
    num := 0
    //获取0-length之间空格的数量
    for i:=0;i <length; i++ {
        if S[i] == ' ' {
            num++
        }
    }
    //创建一个byte数组
    result := make([]byte, 3*num + (length-num))
    k := 0
    //遍历S，遇到空格就在byte数组里添加%20
    for i:=0;i <length; i++ {
        if S[i] == ' ' {
            result[k] = '%'
            result[k+1] = '2'
            result[k+2] = '0'
            k += 3
        } else {
            result[k] = S[i]
            k++
        }
    }
    //建byte数组转换成string并返回
    return string(result)
}
```

