# -*- coding: utf-8 -*-

from mlog.logger import initlog


log = initlog(logfile="/data/logs/kk/slog.log")
log.info("发发火")
log.warn("测试 warn")