# Algorithm
**常用算法/数据结构实现**

## Yami64

> 一个类Base64的编码实现，最后两位特殊字符与base64有所不同，
> 方便不经过转义用于http协议。
> 采用无填充方式，每3字节源二进制为一单位，内部分割出每6bit为一编码单元
> 8bit x 3 = 24bit = 6bit x 4。

Base64字符编码集:
```
[   
'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
'U', 'V', 'W', 'X', 'Y', 'Z',
'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
'u', 'v', 'w', 'x', 'y', 'z',
'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
'+', '/'
]
```
Yami64字符编码集:
```
[   
'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
'U', 'V', 'W', 'X', 'Y', 'Z',
'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
'u', 'v', 'w', 'x', 'y', 'z',
'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
'-', '_'
]
```


## 排序

- [x] 排序算法

- [x]  冒泡排序

- [x]  选择排序

- [x]  直接插入排序

- [x]  快速排序

- [x]  归并排序

- [x]  堆排序

## 树

- [x] 二叉堆（完全二叉树）
