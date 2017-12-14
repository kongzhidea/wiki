### autoconfig ，结合maven使用，采用一套代码，多个antx.properties文件来区分线上和线下环境。
* [wiki](https://www.cnblogs.com/sooper/p/6148626.html)


### autoconfig命令用法
```
获取：autoconfig， 见 libs


基础用法：
1. 配置文件为 ~/antx.properties
2. 获取远程共享配置文件：autoconfig -s http://localhost/antx.properties
    a.推荐线上使用此方案，线上配置文件可以在管控台编辑
    b.测试环境建议直接把配置信息写到工程中，不执行autoconfig
    c.测试环境也可以使用antx配置文件，可以在本地，也可以在控制台编辑方式，但是都需要再执行一次autoconfig转换，没有策略b方便。
3. 使用方法：
    a.推荐：autoconfig -I project，也可以直接作用于war包： autoconfig -I project.war
        方式a位非交互式，自动化脚本中推荐使用。
    b.交互式：autoconfig project， 如果缺少配置文件会提示输入配置信息。
```

### maven工程中使用autoconfig
```
web工程中：在 webapp/META-INF/autoconf/ 目录，新建auto-config.xml文件和对应的模板文件（vm）。

vm文件采用 volacity写法，需要将'.' 替换为'_'。
vm 文件可适用于properties，也可以用于web.xml等。
```

### auto-config.xml样式如下：
```
<?xml version="1.0" encoding="UTF-8"?>
<config description="console config">

    <group name="common config">
        <!-- config -->
        <property name="jdbc.url" description="jdbc"/>
    </group>

    <script>
        <generate template="config.properties.vm" destfile="WEB-INF/classes/config.properties" charset="UTF-8"/>
    </script>

</config>
```
### config.properties.vm
```
#jdbc
jdbc.url=${jdbc_url}
```
