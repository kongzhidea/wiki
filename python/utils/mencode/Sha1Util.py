# -*- coding: utf-8 -*-

from hashlib import sha1


def calcSHA1(str):
    m = sha1()
    m.update(str)
    return m.hexdigest()

if __name__ == "__main__":
    print calcSHA1("kk")