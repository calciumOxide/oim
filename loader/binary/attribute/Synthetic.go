package attribute


/*
Synthetic 属性是定长属性，位于 ClassFile(§4.1)中的属性表。
如果一个类成员没 有在源文件中出现，则必须标记带有 Synthetic 属性，
或者设置 ACC_SYNTHETIC 标志。唯一 的例外是某些与人工实现无关的、由编译器自动产生的方法，
也就是说，Java 编程语言的默认的 实例初始化方法(无参数的实例初始化方法)、类初始化方法，
以及 Enum.values()和 Enum.valueOf()等方法是不用使用 Synthetic 属性或 ACC_SYNTHETIC 标记的。
Synthetic属性是在JDK 1.1中为了支持内部类或接口而引入的。
*/

type Synthetic struct {
}

func AllocSynthetic(b []byte) (*Synthetic, int) {
	return nil, 0
}