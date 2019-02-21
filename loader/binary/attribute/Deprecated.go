package attribute

/*
Deprecated属性是可选定长属性，位于ClassFile(§4.1), field_info(§4.5) 或 method_info(§4.6)结构的属性表中。
类、接口、方法或字段都可以带有为 Deprecated 属性，如果类、接口、方法或字段标记了此属性，则说明它将会在后续某个版本中被取代。
在运行 时解释器或工具(譬如编译器)读取 Class 文件格式时，可以用 Deprecated 属性来告诉使用者 避免使用这些类、接口、方法或字段，
选择其他更好的方式。Deprecated 属性的出现不会修改 类或接口的语义。
*/
type Deprecated struct {
}
