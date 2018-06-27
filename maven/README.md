
## maven详解

### 官网文档
* [Maven in 5 Minutes](http://maven.apache.org/guides/getting-started/maven-in-five-minutes.html)
* [Maven Getting Started Guide](http://maven.apache.org/guides/getting-started/index.html)

### maven安装
```
export JAVA_HOME="/data/jdk"
export M2_HOME="/data/apache-maven-3.3.9"
export PATH=$PATH:"$M2_HOME/bin"
```

### 目录结构
```
src 
	main
		java
		resources
	test
		java
		resources
```

### setting文件
* 位置在 ~/.m2目录（推荐），或 M2_HOME/conf
```
~/.m2
	settings.xml
	repository
```

### maven公网地址
* [http://search.maven.org/](http://search.maven.org/)

### pom文件中标签详解
```
groupId和artifactId组成此工程坐标。

<groupId>com.kk</groupId>       定义组
<artifactId>utils</artifactId>    定义组内一个具体的条目
<version>1.0-SNAPSHOT</version>  版本
<packaging>jar</packaging>   jar/war，普通/web工程
<name>utils</name>       项目名称
```

### pom文件中，依赖管理
```
依赖在<dependencies></dependencies>标签中。

<dependency>
    <groupId>junit</groupId>
    <artifactId>junit</artifactId>
    <version>4.5</version>
    <scope>test</scope>
</dependency>

scope 默认compile，运行时级别。
		test 仅在单元测试时候生效
		provided  对于编译和测试有效，运行时无效。

排除间接依赖，
<dependencies>
	<exclusions></exclusions>
</dependencies>

如：	
<exclusions>
    <exclusion>
        <groupId>org.springframework</groupId>
        <artifactId>spring-beans</artifactId>
    </exclusion>
</exclusions>
```

### 模块管理，依赖父模块
* 参考：pom.parent
* 导入idea或eclipse时候，导入父模块时候可以导入全部子模块。 编译或打包的时候可以打包所有子模块。

### 创建工程
```
1.web
mvn archetype:generate -DgroupId=net.jianxi.tutorials.maven 
                       -DartifactId=webappdemo
                       -Dpackage=net.jianxi.tutorials.maven 
                       -DarchetypeArtifactId=maven-archetype-webapp 
                       -Dversion=1.0 -DinteractiveMode=No

可以把generate改成create

mvn archetype:create -DgroupId=com.renren.game.jisu  -DartifactId=renren-game-jisu -DarchetypeArtifactId=maven-archetype-webapp 


2.普通工程
mvn -B archetype:generate 
  -DarchetypeGroupId=org.apache.maven.archetypes 
  -DgroupId=com.mycompany.app 
  -DartifactId=my-app
```

### eclipse/idea 集成maven
```
mvn eclipse:eclipse
mvn eclipse:clean


mvn idea:idea
mvn idea:clean

idea自带集成maven。
```

### pom中采用变量
```
<properties>
    <spring_version>3.1.1.RELEASE</spring_version>
    <commons_lang_version>2.6</commons_lang_version>
</properties>

<dependency>
    <groupId>commons-lang</groupId>
    <artifactId>commons-lang</artifactId>
    <version>${commons-lang_version}</version>
</dependency>

```

### mvn执行单元测试
```
mvn test  执行所有单元测试

mvn test -Dtest=RandomGeneratorTest   以Random开头，Test结尾的测试类

mvn test -Dtest=Random*Test    用逗号分隔指定多个测试用例

mvn test -Dtest=ATest,BTest   指定即使没有任何测试用例也不要报错
```


### mvn跳过单元测试
```
1.mvn clean package  -Dmaven.test.skip=true

2.mvn package -DskipTests

3.plugin
<plugin>
    <groupId>org.apache.maven.plugins</groupId>
    <artifactId>maven-surefire-plugin</artifactId>
    <configuration>
        <skipTests>true</skipTests>
    </configuration>
</plugin>
```


### 执行main方法
```
1.mvn exec:java -Dexec.mainClass="com.kk.Main"

2.mvn clean compile exec:java -Dexec.mainClass="com.kk.Main" -Dexec.args="arg1 arg2"

3.打包时候带上所有jar包，指定target的jar包中的main方法，执行时候 java -jar kk.jar执行。

  <plugin>
    <artifactId>maven-assembly-plugin</artifactId>
    <configuration>
        <descriptorRefs>
            <descriptorRef>jar-with-dependencies</descriptorRef>
        </descriptorRefs>
        <archive>
            <manifest>
                <mainClass>com.kk.Main</mainClass>
            </manifest>
        </archive>
    </configuration>
    <executions>
        <execution>
            <id>make-assembly</id>
            <phase>package</phase>
            <goals>
                <goal>assembly</goal>
            </goals>
        </execution>
    </executions>
</plugin>
      
4.打包时候 把所有依赖的jar包copy到target/lib目录，执行时候 java -cp .:lib/*:classes kk.jar com.kk.Main

<plugin>
    <groupId>org.apache.maven.plugins</groupId>
    <artifactId>maven-dependency-plugin</artifactId>
    <executions>
        <execution>
            <id>copy-dependencies</id>
            <phase>package</phase>
            <goals>
                <goal>copy-dependencies</goal>
            </goals>
            <configuration>
                <outputDirectory>${project.build.directory}/lib</outputDirectory>
            </configuration>
        </execution>
    </executions>
</plugin>
                       
```

### 打包带上源代码
```
1.resources  会把源码打包到  jar中。

<resources>
    <resource>
        <directory>src/main/java</directory>
        <excludes>
            <exclude>.svn</exclude>
            <exclude>**.xml</exclude>
            <exclude>**.dtd</exclude>
            <exclude>**.conf</exclude>
        </excludes>
        <includes>
            <include>**/*.xml</include>
            <include>**/*.properties</include>
            <include>**/*.java</include>
        </includes>
    </resource>
    <resource>
        <directory>src/main/resources</directory>
        <includes>
            <include>**/*.xml</include>
            <include>**/*.dtd</include>
            <include>**/*.properties</include>
            <include>**/*.ftl</include>
            <include>**/*.txt</include>
            <include>**/*.online</include>
        </includes>
    </resource>
</resources>
<testResources>
    <testResource>
        <directory>src/test/java</directory>
        <includes>
            <include>**/*.xml</include>
            <include>**/*.properties</include>
            <include>**/*.java</include>
        </includes>
    </testResource>
    <testResource>
        <directory>src/test/resources</directory>
        <includes>
            <include>**/*.xml</include>
            <include>**/*.dtd</include>
            <include>**/*.properties</include>
            <include>**/*.ftl</include>
            <include>**/*.txt</include>
            <include>**/*.online</include>
        </includes>
    </testResource>
</testResources>
        
2.maven-source-plugin 打包源码到 source.jar中。

<plugin>
    <groupId>org.apache.maven.plugins</groupId>
    <artifactId>maven-source-plugin</artifactId>
    <version>2.1.2</version>
    <executions>
        <execution>
            <id>attach-sources</id>
            <goals>
                <goal>jar</goal>
            </goals>
        </execution>
    </executions>
</plugin>
        
```
### 指定java-language-level版本
```

在setting.xml中添加如下：
本地jdk版本为 1.8

<profiles>
	<profile>
 		<id>jdk-1.8</id>
 		<activation>
  			<activeByDefault>true</activeByDefault>
  			<jdk>1.8</jdk>
 		</activation>
	 	<properties>
	  		<maven.compiler.source>1.6</maven.compiler.source>
	  		<maven.compiler.target>1.6</maven.compiler.target>
	  		<maven.compiler.compilerVersion>1.6</maven.compiler.compilerVersion>
	 	</properties>
	</profile>
</profiles>

或者在项目中使用：

<plugin>
    <groupId>org.apache.maven.plugins</groupId>
    <artifactId>maven-compiler-plugin</artifactId>
    <configuration>
        <source>1.6</source>
        <target>1.6</target>
        <fork>true</fork>
        <verbose>true</verbose>
        <encoding>UTF-8</encoding>
        <compilerArguments>
            <sourcepath>${project.basedir}/src/main/java</sourcepath>
        </compilerArguments>
    </configuration>
 </plugin>
```

### 常见指令
```
Maven常用命令： 
1. 创建Maven的普通java项目： 
  mvn archetype:create 
  -DgroupId=packageName 
  -DartifactId=projectName  

可选  -DpackageName=packageName

2. 创建Maven的Web项目：   
   mvn archetype:create 
   -DgroupId=packageName    
   -DartifactId=webappName 
   -DarchetypeArtifactId=maven-archetype-webapp    
3. 编译源代码： mvn compile 
4. 编译测试代码：mvn test-compile    
5. 运行测试：mvn test   
6. 发布jar包：mvn deploy   
7. 打包：mvn package   
8. 在本地Repository中安装jar：mvn install 
9. 清除产生的项目：mvn clean   
10.mvn  dependency:tree  生成依赖树
11.mvn -version/-v  显示版本信息 
12.mvn jetty:run    运行项目于jetty上
13.mvn tomcat:run -Dmaven.tomcat.port=11116  maven 自带tomcat运行
```

### 发布jar包到 mvn私服中：
* mvn deploy:deploy-file -DgroupId=com.sun.media  -DartifactId=jai_codec -Dversion=1.1.3 -Dpackaging=jar -Dfile=jai_codec.jar  -DrepositoryId=releases  -Durl=http://mvn.kk.com/nexus/content/repositories/releases/



### 发布第三方Jar到本地库中： 
* mvn install:install-file -DgroupId=com -DartifactId=client -Dversion=0.1.0 -Dpackaging=jar -Dfile=d:\client-0.1.0.jar 

### 指定pom文件
* mvn -f pom-other.xml

### maven filter--管理不同环境的配置文件
```
1.pom.xml：

<build>
    <filters>
        <filter>env-filter.properties</filter>
    </filters>
    <resources>
        <resource>
            <filtering>true</filtering>
            <directory>src/main/resources</directory>
        </resource>
</resources>
</build>


2.env-filter.properties  位置和pom.xml 同级，文件内容如下：

#project.env=
project.env=-test


3.applicationContext-core.xml

mvn 打包后会把${project.env}替换为 env-filter.properties中的变量值。

<!-- 引入配置文件 -->
<bean id="propertyConfigurer"
      class="org.springframework.beans.factory.config.PropertyPlaceholderConfigurer">
    <property name="locations">
        <list>
            <value>classpath:config${project.env}.properties</value>
        </list>
    </property>
</bean>


4.config.properties， config-test.properties

app.name=kk    #config.properties  文件内容
app.name=kk_test   #config-test.properties  文件内容
```

# mvn 排包，检测jar包冲突
``
<plugin>
       <groupId>paibao</groupId>
      <artifactId>pai-maven-plugin</artifactId>
      <version>1.0.0-SNAPSHOT</version>
      <configuration>
         <versionCheckConfig>${basedir}/versionCheck.pb</versionCheckConfig>
     </configuration>
 </plugin>

命令：mvn pai:bao
```

5. mvn 打印依赖树指定插件版本
```
方案一：
mvn  org.apache.maven.plugins:maven-dependency-plugin:2.8:tree -X

方案二：
在pom.xml中添加：

<properties>
    <maven_dependency_plugin_version>3.1.0</maven_dependency_plugin_version>
</properties>
```
