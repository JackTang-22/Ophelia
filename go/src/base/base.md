TODO: 2019.10.14-2019.10.20
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
3. 变量
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

  