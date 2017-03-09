# -*- coding: utf-8 -*-
import base64


def encode(s):
    return base64.b64encode(s)


def decode(s):
    return base64.b64decode(s)


if __name__ == "__main__":
    print encode("abcd")
    print decode("YWJjZA==")
    print encode("孔明")
    print decode("5a2U5piO")
