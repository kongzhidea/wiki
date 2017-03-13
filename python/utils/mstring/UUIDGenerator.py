# -*- coding: utf-8 -*-
import uuid

def getUUID():
    uid = uuid.uuid1()
    return str(uid).replace("-","")