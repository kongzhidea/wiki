
```
def map = [:]     //空map
def map=['name':'john','age':14,'sex':'boy']
println map
println map.getClass() // java.util.LinkedHashMap

def mp = new HashMap<String,String>();
mp.put("a","1");
mp.put("b","2");
mp.c=3;
println mp

for( item in map){
    println item.key + ".." + item.value
}

map.each {key ,value->
    println(key + ".." + value)
}

获取元素:
map['name']
map.name

map操作： 
def map = [3:56, 'name':'Bruce'] 
def a = 'name' 
map.name    //Result: "Bruce" 
map['name'] //Result: "Bruce" 
map[a]      //Result: "Bruce" 
map[3]      //Result: 56 
以下访问是错误的，会抛出异常 
map[name] 
map.3 



键被解释成字符串： 
def x = 3 
def y = 5 
def map = [x:y, y:x] //Result: ["x":5, "y":3] 

如果要把值作为键，像下面这样： 
def city = 'shanghai' 
map."${city}" = 'china' 
map.shanghai //Result: "china" 


map方法： 
def map = ['name':'Bruce', 'age':27] 
map.containsKey('name')   //Result: true 
map.get('name')           //Result: "Bruce" 
map.get('weight', '60kg') //Result: "60kg"；会把key：value加进去 
map.getAt('age')          //Result: 27 
map.keySet()              //Result: [name, age, weight] 
map.put('height', '175')  //Result: ["name":"Bruce", "age":27, "weight":"60kg", "height":"175"] 
map.values().asList()     //Result: ["Bruce", 27, "60kg", "175"] 
map.size()                //Result: 4 

```