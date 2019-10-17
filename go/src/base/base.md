TODO: 2019.10.14-2019.10.31
1. golang命令
2. 变量、常量、类型、函数、包、标识符
3. 数组、切片、结构体、指针、map
4. 方法、接口
5. goroutine、select、defer、mutex、panic、recover等

20191015
1. golang命令
  -
    build: 编译包和依赖
    clean: 移除对象文件
    doc: 显示包或者符号的文档
    env: 打印go的环境信息
    bug: 启动错误报告
    fix: 运行go tool fix
    fmt: 运行gofmt进行格式化
    generate: 从processing source生成go文件
    get: 下载并安装包和依赖
    install: 编译并安装包和依赖
    list: 列出包
    run: 编译并运行go程序
    test: 运行测试
    tool: 运行go提供的工具
    version: 显示go的版本
    vet: 运行go tool vet
20191016
2. 关键字
  - 
    if      for     func    case        struct      import               
    go      type    chan    defer       default     package
    map     const   else    break       select      interface
    var     goto    range   return      switch      continue     fallthrough  
  - 
    内建常量：  
            true        false       iota        nil
    内建类型：  
            int         int8        int16       int32       int64
            uint        uint8       uint16      uint32      uint64      uintptr
            float32     float64 
            complex128  complex64
    bool：      
            byte        rune        string 	    error
    内建函数：   
            make        delete      complex     panic       append      copy    
            close       len         cap	        real        imag        new   	recover
  -
3. 变量 & 常量
  - var variable type
  - var variable = value
  - variable := value 只能在函数内部使用 不能声明全局变量 
  - 已经声明的变量再次声明golang编译报错
  - 推荐使用驼峰命名法
  - var var1, var2 type
  - var var1, var2 type = value1, value2
  - var1, var2 := value1, value2  至少有一个变量声明的，如果全部是赋值 编译报错
  - var (
            var1 type1
            var2 type2
    )
  - variables exchange
    var1, var2 = var2, var1
  - _ 用于丢弃某个不需要的值
  - iota 声明初始化值为0
  - 常量：在编译阶段就确定下来的值，程序运行时无法改变。
    - 常量赋值是一个编译期行为，右边的值不能出现在运行时才能得到结果的值。
    - wrong: const HOME = os.GetEnv("HOME") 
3. 数据类型
  - 值类型
    整型    int8,uint               # 基础类型之数字类型
    浮点型  float32，float64         # 基础类型之数字类型
    复数                            # 基础类型之数字类型
    布尔型  bool                    # 基础类型，只能存true/false，占据1个字节，不能转换为整型，0和1也不能转换为布尔
    字符串  string                  # 基础类型
    数组                            # 复合类型 
    结构体  struct                  # 复合类型
  - 引用类型
    指针    *
    切片    slice
    字典    map
    函数    func
    管道    chan
    接口    interface
  - golang没有字符类型 可使用byte保存单个字符
  - 值类型默认的0值
    int     0
    int8    0
    int32   0
    int64   0
    uint    0x0
    rune    0           //rune的实际类型是 int32
    byte    0x0         // byte的实际类型是 uint8
    float32 0           //长度为 4 byte
    float64 0           //长度为 8 byte
    bool    false
    string  ""
  - 输出格式化
     %%	%字面量
     %b	二进制整数值，基数为2，或者是一个科学记数法表示的指数为2的浮点数
     %c	字符型
     %d	十进制数值，基数为10
     %e	科学记数法e表示的浮点或者复数
     %E	科学记数法E表示的浮点或者附属
     %f	标准计数法表示的浮点或者附属
     %o	8进制度
     %p	十六进制表示的一个地址值
     %s	字符串
     %T	输出值的类型，注意int32和int是两种不同的类型，编译器不会自动转换，需要类型转换。
  - 数据类型转换 type()  转换可能会溢出， 溢出不会报错 只是结果会是未知的 
  - 基本数据类型转字符串：fmt.Sprintf() 该函数会返回转换后的字符串
  - 字符串转基本数据类型：使用包strconv 如果转换成数字的字符串不合法  会返回0
  - alias golang1.9 之后  type aliasName = type
4. 流程控制
  - if true{}  if err := value; err != nil {}
  - switch 每一个case后面默认添加了break Go提供fallthrough，代表不跳出switch，后面的语句无条件执行
  - for ; ; {}   for {} for v1, v2 := Xx {}
  - 循环关键字
    break用于函数内跳出当前for、switch、select语句的执行
    continue用于跳出for循环的本次迭代。
    goto可以退出多层循环
5. 运算符
  - 
    算术运算符：	+	-	*	/	%	++	--	
    关系运算符：	==	!=	<=	>=	<	>	
    逻辑运算符：	!	&&	||
    位运算：		&（按位与）	|（按位或）	^（按位取反）	<<（左移）	>>（右移）
    赋值运算符：	=	+=	-=	*=	/=	%=	<<=	>>=	&=	^=	|=
    其他运算符：	&（取地址）	*（取指针值） <-（Go Channel相关运算符）
  - Go中只有后--和后++，且自增自减不能用于表达式中，只能独立使用
  - 位运算
    &     按位与，参与运算的两个数二进制位相与：同时为1，结果为1，否则为0
    |     按位或，参与运算的两个数二进制位相或：有一个为1，结果为1，否则为0
    ^     按位异或：二进位不同，结果为1，否则为0
    <<    按位左移：二进位左移若干位，高位丢弃，低位补0，左移n位其实就是乘以2的n次方
    >>    按位右移：二进位右移若干位，右移n位其实就是除以2的n次方
20191017
6. 数值类型
  - 整数
    1. 有符号
      int     32位系统等价于int32 4字节    64位系统等价于int64 8字节 
      int8    占据1字节   范围 -128 ~ 127
      int16   占据2字节   范围 -2(15次方) ~ 2（15次方）-1
      int32   占据4字节   范围 -2(31次方) ~ 2（31次方）-1
      int64   占据8字节   范围 -2(63次方) ~ 2（63次方）-1
      rune	int32的别称
    2. 无符号
      uint	32位系统占4字节（与uint32范围一样），64位系统占8字节（与uint64范围一样）     
      uint8   占据1字节   范围 0 ~ 255
      uint16  占据2字节   范围 0 ~ 2（16次方）-1
      uint32  占据4字节   范围 0 ~ 2（32次方）-1
      uint64  占据8字节   范围 0 ~ 2（64次方）-1
      byte	uint8的别称
    3. golang中不同类型的变量不能相互赋值  例如 int  int8
    4. golang默认整数类型是int
    5. unsafe.Sizeof() 查看字节数
  - 浮点型
    - float32 单精度  占据4字节   范围 -3.403E38 ~ 3.403E38    (math.MaxFloat32)
    - float64 双精度  占据8字节   范围 -1.798E208 ~ 1.798E308  (math.MaxFloat64)
    - 浮点数 = 符号位 + 指数位 + 尾数位
    - 浮点数的范围和操作系统无关
    - 可能会出现精度损失的情况
    - 不能用 == 判断两个浮点数是否相同
  - 复数
    - 复数默认类型是complex128（64位实数+64位虚数）
    - complex64(32位实数+32位虚数)
  - NAN
    - 不是一个数字
    - 函数math.IsNaN用于测试一个数是否是非数NaN
    - 函数math.NaN则返回非数对应的值
    - NaN和任何数都是不相等的
7. 字符
  - 字符
    1. golang中没有char类型，单个字符用byte存储 rune
    2. 直接输出字符类型为ASCII码值 %c 输出字符
    3. 字符可以和整数进行运算
  - 字符串
    1. golang字符串由单个字节连接起来，Go字符串是一串固定长度的字符连接起来的字符序列。
    2. golang中字符串是基本数据类型，初始化后内容不能被修改
    3. UTF-8编码 "" 或者 ``定义 ``没有字符转义功能 可以写在多行
    4. 字符串的常见操作
      - len() 长度
      - string() 其他类型转换为string类型
      - + 字符串拼接
      - strings 包
      - strconv 包
8. 数组
  - 数组是一段固定长度的连续内存区域。数组的长度定义后不可更改，长度使用 len() 获取
  - var a [len]type
  - a [len]type := [len]type{...}
  - a := [len]type{...}
  - a := [...]type{...} 自动计算长度
  - 数组长度不可变
  - 长度是数组类型的一部分 [1]int 和 [2]int 是不同的类型
  - 数组之间是值赋值，传入函数作为参数的时候，传入的是副本，而不是指针
9. 结构体
  - struct_test.go
  - 结构体中的数据类型可以是任意类型，存储空间连续，字段按照声明时的顺序存放
  - 结构体本身和成员都可用于比较，== !=  不支持> <
  - 结构体成员首字母小写 对包外不可见
  - p := &Person{} === new(Person) 指针类型
  - 匿名结构体嵌套 两层都有相同的字段，优先访问上层
10. 引用类型
  - 切片
    1. 切片(slice)解决了数组长度不能扩展，以及基本类型数组传递时产生副本的问题
    2. slice_test.go
    3. slice的默认开始位置是0，ar[:n]等价于ar[0:n]
    4. slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
    5. 如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]
    6. 切片的遍历可以使用for循环，也可以使用range函数
    7. 切片通过内部的指针和相关属性引用数组片段，实现了变长方案，Slice并不是真正意义上的动态数组
  - map
    1. map是一个无序键值对集合
    2. map_test.go
    3. golang中map的key通常为int string，
       但也可以是其他类型如：bool、数字、string、指针、channel，
       还可以是只包含前面几个类型的接口、结构体、数组。
       slice、map、function由于不能使用 == 来判断，不能作为map的key
    4. map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
    5. map的长度是不固定的，也就是和slice一样，也是一种引用类型
    6. 内置的len函数同样适用于map，返回map拥有的key的数量
    7. go没有提供清空元素的方法，可以重新make一个新的map，不用担心垃圾回收的效率，
       因为go中并行垃圾回收效率比写一个清空函数高效很多
    8. map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
    9. Go内置的map只有读是线程安全的，读写是线程不安全的
    10. 需要并发读写时，一般都是加锁，但是这样做性能不高，在go1.9版本中提供了更高效并发安全的sync.Map
    11. sync.Map
      - 无须初始化，直接声明即可
      - sync.Map不能使用map的方式进行取值和设值操作，而是使用sync.Map的方法进行调用。Store表示存储，Load表示获取，Delete表示删除。
      - 使用Range配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，需要继续迭代时，返回true，终止迭代返回false。
      - map没有提供获取map数量的方法，可以在遍历时手动计算。sync.Map为了并发安全。损失了一定的性能
  - 指针
    1. TODO