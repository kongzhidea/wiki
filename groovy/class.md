
```

1.类格式：
class Person{
    def id
    def name
}

Groovy类和java类一样，你完全可以用标准java bean的语法定义一个groovy 类。
我们也可以使用更groovy的方式定义和使用类，这样的好处是，你可以少写一半以上的javabean代码：
(1)    不需要public修饰符
groovy的默认访问修饰符就是public，如果你的groovy类成员需要public修饰，则你根本不用写它。
(2)    不需要类型说明
groovy也不关心变量和方法参数的具体类型。
(3)    不需要getter/setter方法
在groovy中，则彻底不需要getter/setter方法——所有类成员（如果是默认的public）根本不用通过getter/setter方法引用它们。
当然，如果你一定要通过get/set方法访问成员属性，groovy也提供了它们。
(4)    不需要构造函数
不在需要程序员声明任何构造函数，因为groovy自动提供了足够你使用的构造函数。
不用担心构造函数不够多，因为实际上只需要两个构造函数
（1个不带参数的默认构造函数，1个只带一个map参数的构造函数—由于是map类型，通过这个参数你可以在构造对象时任意初始化它的成员变量）。
(5) 不需要return
Groovy中，方法不需要return来返回值吗？这个似乎很难理解。看后面的代码吧。
因此，groovy风格的类是这样的：
(6)    不需要()号
Groovy中方法调用可以省略()号（构造函数除外），也就是说下面两句是等同的：
 
7)get，set 方法可以不填，groovy会默认生成。
person1.setName 'kk'
person1.setName('kk')
 
下面看一个完整类定义的例子：
class Person {
 def name
 def age
 String toString(){//注意方法的类型String，因为我们要覆盖的方法为String类型
     "$name,$age"
 }
 
我们可以使用默认构造方法实例化Person类：
def person1=new Person()
person1.name='kk'
person1.age=20
println person1


也可以用groovy的风格做同样的事：
def person2=new Person(['name':'gg','age':22]) //[]号可以省略
println person2


00..    动态性
Groovy所有的对象都有一个元类metaClass，我们可以通过metaClass属性访问该元类。通过元类，可以为这个对象增加方法（在java中不可想象）！见下面的代码，msg是一个String,通过元类，我们为msg增加了一个String 类中所没有的方法up：

def msg = "Hello!"
println msg.metaClass
String.metaClass.up = {  delegate.toUpperCase() }
println msg.up()

通过元类，我们还可以检索对象所拥有的方法和属性（就象反射）：

msg.metaClass.methods.each { println it.name }
msg.metaClass.properties.each { println it.name }

甚至我们可以看到我们刚才添加的up方法。
我们可以通过元类判断有没有一个叫up的方法，然后再调用它：

if (msg.metaClass.respondsTo(msg, 'up')) {
    println msg.toUpperCase()
}

当然，也可以推断它有没有一个叫bytes的属性：

if (msg.metaClass.hasProperty(msg, 'bytes')) {
    println msg.bytes.encodeBase64()
}




2.安全占位符
这个很有用，可以避免很多NullPointerException，但是也不能滥用了

def result = obj?.property  

代码里，obj是个对象，property是对象的一个熟悉，这行代码的意思，如果obj不为null，则会返回property属性的值，如果obj为null，这会直接返回null。语句可以一直串下去

def result = a?.b?.c?.d...  


4.字段操作
按照Groovy Bean的标准，默认的时候类里面的所有字段，Groovy都会帮忙生成一个get方法。在类的外部，即便你直接用了属性名而不用get方法去取值，拿到的也是通过get方法拿到的值。如果想直接拿属性值怎么办呢？ 通过字段运算符：
class A {
   String b
}

A a = new A()
a.b //通过get方法拿值
a.getB() //通过get方法拿值
a.@b //直接拿值

如果在A中定义  def getAge()，那么外面可以直接通过a.age来获取age的值。

5.构造函数
groovy会默认加上 不带参数的构造函数。

也可以指定构造函数：
Person(id) {
    this.id = id
}

如果没有构造函数，可以这样初始化类：
p = new Person(id:10,name:"孔")， 也可以穿入map，如 p = new Person([id:10,name:"孔"])
会对id和name分别赋值。

5.Java中的方法返回写为return; 或者return obj; 在Groovy的方法中return是可选的

Java:
public String sayHello() {
  return "Hello, 山风小子";
}

Groovy:
public String sayHello() {
  return "Hello, 山风小子"
}

//或者
public String sayHello() {
  "Hello, 山风小子"
}

//或者
String sayHello() {
  "Hello, 山风小子"
}

//或者
public sayHello() {
  "Hello, 山风小子"
}

// 或者
def sayHello() {
  "Hello, 山风小子"
}
```