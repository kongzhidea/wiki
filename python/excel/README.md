
```
使用easy_install 安装:easy_install  xlrd  /  easy_install  xlwt
```

```
读excel:

# -*-coding:UTF-8-*-#

import xlrd

def showlist(l):
    print "[",
    for i in l:
        print i,",",
    print "]"

#打开excel
data = xlrd.open_workbook('../data/374.xls') #注意这里的workbook首字母是小写
#查看文件中包含sheet的名称
print data.sheet_names()
#得到第一个工作表，或者通过索引顺序 或 工作表名称  均从0开始
table = data.sheets()[0]
table = data.sheet_by_index(0)
table = data.sheet_by_name(u'OrderList')

#获取行数和列数  均从0开始
nrows = table.nrows
ncols = table.ncols
print "line,col number=",nrows,ncols

#获取整行和整列的值（数组）
row0 = table.row_values(0)
col0 = table.col_values(0)
showlist(row0)
showlist(col0)

#循环行,得到索引的列表
for rownum in range(table.nrows):
    showlist(table.row_values(rownum))

#单元格  均从0开始
cell_A1 = table.cell(0,0).value
cell_C3 = table.cell(2,2).value
print cell_A1,cell_C3

#分别使用行列索引
cell_A1 = table.row(0)[0].value
cell_A2 = table.col(1)[0].value
print cell_A1,cell_A2

```

```
写excel:



xlwt
http://pypi.python.org/pypi/xlrd


# -*-coding:UTF-8-*-#

import xlwt

def showlist(l):
    print "[",
    for i in l:
        print i,",",
    print "]"

#新建一个excel文件
file = xlwt.Workbook() #注意这里的Workbook首字母是大写，无语吧
#新建一个sheet
table = file.add_sheet(u'测试的')
#写入数据table.write(行,列,value)
table.write(0,0,'testd')
#如果对一个单元格重复操作，会引发
#returns error:
## Exception: Attempt to overwrite cell:
## sheetname=u'sheet 1' rowx=0 colx=0
#所以在打开时加cell_overwrite_ok=True 解决
#table = file.add_sheet('sheet name',cell_overwrite_ok=True )
#保存文件
file.save('../data/demo.xls')

```

```
另外，使用style
style = xlwt.XFStyle() # 初始化样式
font = xlwt.Font() #为样式创建字体
font.name = 'Times New Roman'
font.bold = True
style.font = font #为样式设置字体
table.write(0, 0, 'some bold Times text', style) # 使用样式
xlwt 允许单元格或者整行地设置格式。还可以添加链接以及公式。可以阅读源代码，那里有例子：
dates.py, 展示如何设置不同的数据格式
hyperlinks.py, 展示如何创建超链接 (hint: you need to use a formula)
merged.py, 展示如何合并格子
row_styles.py, 展示如何应用Style到整行格子中.
具体的例子可以看：
http://scienceoss.com/write-excel-files-with-python-using-xlwt/
google论坛：
http://groups.google.com/group/python-excel/
```