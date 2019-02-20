package attribute

/*
SourceDebugExtension 属性是可选属性，位于 ClassFile(§4.1)结构属性表。 ClassFile 结构的属性表最多只能包含一个 SourceDebugExtension 属性。
*/
type SourceDebugExtension struct {
	//common clazz.Attribute //attribute_length 项的值给出了当前属性的长度，不包括开始的 6 个字节。即 attribute_length 项的值是字节数组 debug_extension[]数组的长度
	//debugExtesion *DebugExtesionInfo //debug_extension[]数组用于保存扩展调试信息，扩展调试信息对于 Java 虚拟机来 说没有实际的语义。这个信息用改进版的 UTF-8 编码的字符串(§4.4.7)表示，这个 字符串不包含 byte 值为 0 的终止符。需要注意的是，debug_extension[]数组表示 的字符串可以比 Class 实例对应的字符串更长。
	DebugExtesion []uint8
}

/*type DebugExtesionInfo uint8

func AllocSourceDebugExtension(b []byte) (*SourceDebugExtension, int) {

}*/