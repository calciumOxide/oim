package binary

type AccessFlags uint16

const (
	ACC_EMPTY        AccessFlags = 0x0000 //空
	ACC_PUBLIC       AccessFlags = 0x0001 //可以被包的类外访问
	ACC_PRIVATE      AccessFlags = 0x0002 //表示字段仅能该类自身调用
	ACC_PROTECTED    AccessFlags = 0x0004 //表示字段可以被子类调用
	ACC_STATIC       AccessFlags = 0x0008 //表示静态字段
	ACC_FINAL        AccessFlags = 0x0010 //不允许有子类 || final，方法不能被重写(覆盖)
	ACC_SUPER        AccessFlags = 0x0020 //当用到 invokespecial 指令时，需要特殊处理的父类方法 || synchronized，方法由管程同步
	ACC_SYNCHRONIZED AccessFlags = 0x0020 //当用到 invokespecial 指令时，需要特殊处理的父类方法 || synchronized，方法由管程同步
	ACC_VOLATILE     AccessFlags = 0x0040 //表示字段是易变的 || bridge，方法由编译器产生
	ACC_BRIDGE       AccessFlags = 0x0040 //表示字段是易变的 || bridge，方法由编译器产生
	ACC_TRANSIENT    AccessFlags = 0x0080 // transient，表示字段不会被序列化 || 表示方法带有变长参数
	ACC_VARARGS      AccessFlags = 0x0080 // transient，表示字段不会被序列化 || 表示方法带有变长参数
	ACC_NATIVE       AccessFlags = 0x0100 //native，方法引用非 java 语言的本地方法
	ACC_INTERFACE    AccessFlags = 0x0200 //标识定义的是接口而不是类
	ACC_ABSTRACT     AccessFlags = 0x0400 //不能被实例化 || abstract，方法没有具体实现
	ACC_STRICT       AccessFlags = 0x0800 //strictfp，方法使用 FP-strict 浮点格式
	ACC_SYNTHETIC    AccessFlags = 0x1000 //表示字段由编译器自动产生, 标识并非 Java 源码生成的代码
	ACC_ENUM         AccessFlags = 0x4000 //enum，表示字段为枚举类型
)
