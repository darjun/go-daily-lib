package main

// CheckCollision 两个物体之间是否碰撞
func CheckCollision(entityA, entityB Entity) bool {
	top, left := entityA.Y(), entityA.X()
	bottom, right := entityA.Y()+float64(entityA.Height()), entityA.X()+float64(entityA.Width())
	// 左上角
	x, y := entityB.X(), entityB.Y()
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	// 右上角
	x, y = entityB.X()+float64(entityB.Width()), entityB.Y()
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	// 左下角
	x, y = entityB.X(), entityB.Y()+float64(entityB.Height())
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	// 右下角
	x, y = entityB.X()+float64(entityB.Width()), entityB.Y()+float64(entityB.Height())
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	return false
}
