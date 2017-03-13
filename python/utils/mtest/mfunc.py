# -*- coding: utf-8 -*-

'''
函数可变参数, type(args)=tuple
'''
def calcsum(*args):
    # print args #(1, 2, 3, 4, 5)
    return sum(args)


print calcsum(1,2,3,4,5)


'''
函数可变参数, type(args)=dict
'''
def calcValueSum(**args):
    # print args  #{'a': 1, 'b': 2}
    t = 0
    for key,value in args.items():
        t += value
    return t

print calcValueSum(a=1,b=2)