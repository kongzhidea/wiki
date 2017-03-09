# -*- coding: utf-8 -*-

from hashlib import md5


def calMD5(str):
    m = md5()
    m.update(str)
    return m.hexdigest()

if __name__ == "__main__":
    print calMD5("123456")