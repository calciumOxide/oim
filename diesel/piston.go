package diesel

import "../loader/clazz/attribute"
import "./instructions"
import "../loader/clazz"

func Nozzle(codes attribute.Codes) (int, error) {
	return 0, nil
}

func SteamCylinder() (int, error) {
	i := instructions.Instructions(0xb6)
	ctx := instructions.INSTRUCTION_MAP[i].Test(nil)
	ctx.Clazz = clazz.GetClass("com/oxide/A")
	i.Stroke(ctx)
	//i = instructions.Instructions(0xab)
	//ctx = instructions.INSTRUCTION_MAP[i].Test(ctx)
	//i.Stroke(ctx)
	return 0, nil
}