# -*- coding: utf-8 -*-
from mfile import FileUtils
import os
import os.path
import sys


if __name__ == "__main__":
    # print sys.argv[0]

    # FileUtils.touch("/Users/kongzhihui/test/t1/tfile")

    # print os.path.exists("/Users/kongzhihui/test/t1/s1/c4")
    # os.makedirs("/Users/kongzhihui/test/t1/s1/c4")
    # print os.path.dirname("./tfile")
    # print os.path.exists(os.path.dirname("tfile"))

    FileUtils.copyfile("/Users/kongzhihui/test/t1/tfile","/Users/kongzhihui/test/t1/tfile2")

    print FileUtils.exists("/Users/kongzhihui/test/t1/tfile2")
    # FileUtils.copyDirectory("/Users/kongzhihui/test/t1","/Users/kongzhihui/test/t3")

    # FileUtils.moveFile("/Users/kongzhihui/test/t2/tfile2", "/Users/kongzhihui/test/t2/tfile2_m")

    FileUtils.deleteFile("/Users/kongzhihui/test/t1/tfile2")

    print FileUtils.exists("/Users/kongzhihui/test/t1/tfile2")

    # print FileUtils.exists("/Users/kongzhihui/test/t5")

    # FileUtils.mkdir("/Users/kongzhihui/test/t5")
    # FileUtils.mkdir("/Users/kongzhihui/test/t5/s1/c1",True)


    # FileUtils.rmdir("/Users/kongzhihui/test/t5", True)

    print FileUtils.readFileToString("/Users/kongzhihui/test/t1/test.go")

    cont = FileUtils.readFileToString("/Users/kongzhihui/test/t1/gbk.go")
    print cont.decode("gbk").encode("utf-8")

    lines = FileUtils.readLines("/Users/kongzhihui/test/t1/test.go")
    for line in lines:
        print line


    FileUtils.writeStringToFile("/Users/kongzhihui/test/t1/wfile","呵呵，我来试试中文\n回车\n")

    FileUtils.writeStringToFile("/Users/kongzhihui/test/t1/wfile","呵呵，我来试试中文\n回车\n", True)

    FileUtils.writeLines("/Users/kongzhihui/test/t1/listfile", ["第一行","2","3"])

    l = FileUtils.readLines("/Users/kongzhihui/test/t1/listfile")
    print l


    print  FileUtils.isFile("/Users/kongzhihui/test/t1/listfile")
    print FileUtils.isDir("/Users/kongzhihui/test/t1")
    print FileUtils.listdir("/Users/kongzhihui/test")

    print os.path.dirname("/Users/kongzhihui/test/t1/listfile")

    rootdir = "/Users/kongzhihui/test"  # 指明被遍历的文件夹

    for parent, dirnames, filenames in os.walk(rootdir):  # 三个参数：分别返回1.父目录 2.所有文件夹名字（不含路径） 3.所有文件名字
        for dirname in dirnames:  # 输出文件夹信息
            print "parent is:" + parent
            print  "dirname is:" + dirname

        for filename in filenames:  # 输出文件信息
            print "parent is:"+ parent
            print "filename is:" + filename
            print "the full name of the file is:" + os.path.join(parent, filename)  # 输出文件路径信息