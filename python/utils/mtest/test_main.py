# -*- coding: utf-8 -*-

'''
获取main参数。
main参数中包含文件名。

例如：
python args.py 1 2 3

结果：4， ['args.py', '1', '2', '3']
'''


import sys

#print(sys.argc) //没有argc参数

print(sys.argv) # 参数个数

print(len(sys.argv))

for i in sys.argv:
   print i



### 解析main  长短参数
'''
import argparse

parser = argparse.ArgumentParser()
parser.add_argument("day")  ## 必须参数
parser.add_argument("-e", "--encoding")   ## 可选参数，  -和-- 表示可选参数。
args = parser.parse_args()


运行时候 可以指定key，如果不指定key则默认按照add顺序。
'''