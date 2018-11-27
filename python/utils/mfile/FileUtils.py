# -*- coding: utf-8 -*-
import shutil
import os
import os.path

'''
os.path.dirname(path)  返回path的父目录

# 遍历文件夹下文件
rootdir = "/Users/kongzhihui/test"  # 指明被遍历的文件夹

    for parent, dirnames, filenames in os.walk(rootdir):  # 三个参数：分别返回1.父目录 2.所有文件夹名字（不含路径） 3.所有文件名字
        for dirname in dirnames:  # 输出文件夹信息
            print "parent is:" + parent
            print  "dirname is:" + dirname

        for filename in filenames:  # 输出文件信息
            print "parent is:"+ parent
            print "filename is:" + filename
            print "the full name of the file is:" + os.path.join(parent, filename)  # 输出文件路径信息
'''


def exists(path):
    return os.path.exists(path)


def isFile(path):
    return os.path.isfile(path)


def isDir(path):
    return os.path.isdir(path)


# 返回文件和目录，不包含路径，只有名字
def listdir(path):
    return os.listdir(path)


# 要求父目录必须存在
def touch(fname):
    fp = open(fname, 'a')
    fp.close()


# 要求src必须存在
def copyfile(src, dest):
    shutil.copy(src, dest)


# 要求src必须存在
def moveFile(src, dest):
    shutil.move(src, dest)


# 要求fname必须存在，fname必须是文件
def deleteFile(fname):
    os.remove(fname)


# 要求src必须存在，dest必须不存在； 目录下文件也会复制
def copyDirectory(src, dest):
    shutil.copytree(src, dest)


# 要求dname不存在，force=True会自动创建起父目录
def mkdir(dname, force=False):
    if force:
        os.makedirs(dname)
    else:
        os.mkdir(dname)


# 要求dname必须存在，force=False时候要求dname为空目录
def rmdir(dname, force=False):
    if force:
        shutil.rmtree(dname)
    else:
        os.rmdir(dname)


# 返回str，要求fname必须存在，utf-8编码
# 读取非UTF-8编码的文本文件，需要给open()函数传入encoding参数，例如，读取GBK编码的文件：f = open('/Users/michael/gbk.txt', 'r', encoding='gbk')
# 遇到有些编码不规范的文件，可能会遇到UnicodeDecodeError。遇到这种情况，open()函数还接收一个errors参数：f = open('/Users/michael/gbk.txt', 'r', encoding='gbk', errors='ignore')

# 对于二进制文件，如图片，需要制定使用 字节形式来读取，open(fname, "rb")，否则读取信息会丢失。

def readFileToString(fname):
    with open(fname) as fp:  # 或者fp = open("file.txt","r")
        cont = fp.read()
    # fp = None
    # cont = None
    # try:
    #     fp = open(fname)  # 或者fp = open("file.txt","r")
    #     cont = fp.read()
    # finally:
    #     if fp:
    #         fp.close()

    return cont


'''
for line in open(fname):
   print line.strip()
'''


# 返回list，每行删除回车，要求fname必须存在；  fp.readlines() 也返回list，但是每行会带上回车
def readLines(fname):
    ret = []
    with open(fname) as fp:
        for line in fp:
            ret.append(line.rstrip("\r\n"))
    return ret


# 会把原来的文件内容
def writeStringToFile(fname, s, append=False):
    mode = "w"
    if append:
        mode = "a"

    with open(fname, mode) as fp:
        fp.write(s)


# l为list，每个元素不用有回车，输出到文件的时候每个item会自动补上回车。
def writeLines(fname, l, append=False):
    mode = "w"
    if append:
        mode = "a"
    with open(fname, mode) as fp:
        for i in l:
            fp.write(str(i))
            fp.write("\n")
