# -*- coding: utf-8 -*-
import uuid

from mstring import UUIDGenerator

print uuid.uuid1()  # 带参的方法参见Python Doc

uid =UUIDGenerator.getUUID()
print len(uid),uid

