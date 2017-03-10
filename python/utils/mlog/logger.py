# -*- coding: utf-8 -*-
import logging
import sys

def initlog(name=__file__, logfile = None):
    # 生成一个日志对象
    logger = logging.getLogger(name)

    # 成一个格式器，用于规范日志的输出格式。如果没有这行代码，那么缺省的
    # 格式就是："%(message)s"。也就是写日志时，信息是什么日志中就是什么，
    # 没有日期，没有信息级别等信息。logging支持许多种替换值，详细请看
    # Formatter的文档说明。这里有三项：时间，信息级别，日志信息

    # initlog()中的name属性对应formatter中的$(name)s
    formatter = logging.Formatter('%(asctime)s %(levelname)s %(filename)s:%(lineno)s - %(message)s')

    # 生成一个Handler。logging支持许多Handler，
    # 象FileHandler, SocketHandler, SMTPHandler等，我由于要写
    # 输出到文件 文件就使用了FileHandler。
    # logfile是一个全局变量，它就是一个文件名，如：'crawl.log'
    if logfile:
        hdlr = logging.FileHandler(logfile)

        # 将格式器设置到处理器上
        hdlr.setFormatter(formatter)

        # 将处理器加到日志对象上
        logger.addHandler(hdlr)

    #输出到控制台
    # 创建一个控制台的handler
    c_handler = logging.StreamHandler(sys.stdout);
    #c_handler.setLevel(logging.WARNING);
    c_handler.setFormatter(formatter)


    logger.addHandler(c_handler)
    # 设置日志信息输出的级别。logging提供多种级别的日志信息，如：NOTSET,
    # DEBUG, INFO, WARNING, ERROR, CRITICAL等。每个级别都对应一个数值。
    # 如果不执行此句，缺省为30(WARNING)。可以执行：logging.getLevelName
    # (logger.getEffectiveLevel())来查看缺省的日志级别。日志对象对于不同
    # 的级别信息提供不同的函数进行输出，如：info(), error(), debug()等。当
    # 写入日志时，小于指定级别的信息将被忽略。因此为了输出想要的日志级别一定
    # 要设置好此参数。这里我设为NOTSET（值为0），也就是想输出所有信息
    logger.setLevel(logging.INFO)   ## info级别
    return logger

if __name__ == "__main__":
    #定义logger
    # logger=initlog(logfile="/data/logs/kk/slog.log")
    logger = initlog()

    logger.debug("debug")
    logger.info('info,呵呵')
    logger.warn('warn')
    logger.error("error")
