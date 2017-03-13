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

