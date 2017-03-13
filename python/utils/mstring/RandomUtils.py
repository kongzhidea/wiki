# -*- coding: utf-8 -*-
import random
import string

# l--list,　random.choice从序列中获取一个随机元素。其函数原型为：random.choice(sequence)
def choice(l):
    return random.choice(l)

# random.sample的函数原型为：random.sample(sequence, k)，从指定序列中随机获取指定长度的片断。sample函数不会修改原有序列。
def sample(l,k):
    return random.sample(l,k)

# random.randint()的函数原型为：random.randint(a, b)，用于生成一个指定范围内的整数。其中参数a是下限，参数b是上限，生成的随机数n: a <= n <= b

def randomInt(min,max):
    return random.randint(min,max)

'''
random.random()用于生成一个0到1的随机符点数: 0 <= n < 1.0

random.shuffle的函数原型为：random.shuffle(list)，用于将一个列表中的元素打乱
'''

# 随机k个 大小写字母
def randomAlpha(k):
    return random.sample(string.letters, k)

# 随机k个 数字，str类型，如 ['7', '6']， 如果要换成int类型：map(int,RandomUtils.randomNumeric(k))
def randomNumeric(k):
    return random.sample(string.digits, k)

# 随机k个 大小写字母和数字，
def randomAlphaNumeric(k):
    return random.sample(string.digits + string.letters, k)