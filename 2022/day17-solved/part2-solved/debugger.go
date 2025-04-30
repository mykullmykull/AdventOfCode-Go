package day

func (s simulator) printTower() {
	max := s.state.towerCols.getMaxSpace()
	for x := 0; x < max; x++ {
		print("|")
		print(string(s.state.towerCols.col1[x]))
		print(string(s.state.towerCols.col2[x]))
		print(string(s.state.towerCols.col3[x]))
		print(string(s.state.towerCols.col4[x]))
		print(string(s.state.towerCols.col5[x]))
		print(string(s.state.towerCols.col6[x]))
		print(string(s.state.towerCols.col7[x]))
		println("|")
	}
	println()
}

func (s simulator) printTowerWithRock() {
	towerMap := make([]string, s.state.towerCols.getMaxSpace())
	for x := 0; x < len(towerMap); x++ {
		row := "  |"
		row += string(s.state.towerCols.col1[x])
		row += string(s.state.towerCols.col2[x])
		row += string(s.state.towerCols.col3[x])
		row += string(s.state.towerCols.col4[x])
		row += string(s.state.towerCols.col5[x])
		row += string(s.state.towerCols.col6[x])
		row += string(s.state.towerCols.col7[x])
		row += "|"
		towerMap[x] = row
	}

	for _, point := range s.rock.points {
		row := point[0]
		col := point[1] + 2
		towerMap[row] = towerMap[row][:col] + "0" + towerMap[row][col+1:]
	}
	for _, row := range towerMap {
		println(row)
	}
}
