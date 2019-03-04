package diesel

import "../loader/clazz/attribute"
import "./instructions"

func Nozzle(codes attribute.Codes) (int, error) {
	return 0, nil
}

func SteamCylinder() (int, error) {
	i := instructions.Instructions(0xb7)
	i.Stroke(nil)
	return 0, nil
}