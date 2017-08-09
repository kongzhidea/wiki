# -*- coding: utf-8 -*-

import csv
import sys


GBK = 'gbk'
UTF8 = 'utf-8'


'''
csv文件为gbk编码，读取的时候需要转成utf-8编码，写会到csv文件时候需要再转成gbk编码.
'''

# 将编码为f的str转成编码t
def encodeTransfer(str, f, t):
    return str.decode(f).encode(t)


def test_readcsv():
    csv_file = "/Users/kongzhihui/test/csv/kk.csv"

    csv_in_file = open(csv_file, 'rb')

    # csv文件的读取，默认csv分隔符为 逗号。
    #reader = csv.reader(csv_in_file, delimiter=",")
    reader = csv.reader(csv_in_file)

    # line 为 list类型。
    for line in reader:
        id = line[0]
        type = encodeTransfer(line[1], GBK, UTF8)
        cate = encodeTransfer(line[2], GBK, UTF8)

        print id, type, cate

def test_writecsv():
    data = [
        ("And Now For Something Completely Different", 1971, "姓名"),
        ("Monty Python And The Holy Grail", 1975, "赵"),
        ("Monty Python's Life Of Brian", 1979, "钱"),
        ("Monty Python Live At The Hollywood Bowl", 1982, "孙"),
        ("Monty Python's The Meaning Of Life", 1983, "李")
    ]

    writer = csv.writer(sys.stdout)

    for item in data:
        writer.writerow(item)


def test_writecsv2():
    csv_file = "/Users/kongzhihui/test/csv/out.csv"

    csv_out_file = open(csv_file, 'wb')

    # 对应写csv的对象
    writer = csv.writer(csv_out_file)

    data = [
        ("And Now For Something Completely Different", 1971, "姓名"),
        ("Monty Python And The Holy Grail", 1975, "赵"),
        ("Monty Python's Life Of Brian", 1979, "钱"),
        ("Monty Python Live At The Hollywood Bowl", 1982, "孙"),
        ("Monty Python's The Meaning Of Life", 1983, "李")
    ]

    for item in data:
        tup = []
        for col in item:
            if(isinstance(col,str)):
                tup.append(encodeTransfer(col,UTF8,GBK))
            else:
                tup.append(col)

        writer.writerow(tup)

if __name__ == "__main__":
    test_readcsv()

    test_writecsv()

    test_writecsv2()