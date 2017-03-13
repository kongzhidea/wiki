# -*- coding: utf-8 -*-

print "%s %d"%("kong",123)  ## 参数格式要完全对应

id=1
print "{0}" .format(id)  ## 参数数量不能少

print "{name} {sex}" .format(name="kong",sex="F")  ## key不能少

print '{name} {age}'.format(age=22, name='管理员')

print '{0} {1} {2}'.format('a', 'b', 'c')


print '{array[2]}'.format(array=range(10))
print '{attr.__class__}'.format(attr=0)

## unicode字符，需要加!r; 或者引入  reload(sys); sys.setdefaultencoding('utf8')
print '{name!r}'.format(name=u'汉字')
