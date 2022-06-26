package domain

type tournMnger struct {
	type_   TourType
	players []Player
}

func (tm *tournMnger) Tree() *TourNode {
	switch tm.type_ {
	case 0:
		return tm.BuildFirstDefeat()
	case 1:
		return tm.BuildAllVsAll()
	}
	return nil
}

func (tm *tournMnger) BuildFirstDefeat() *TourNode {
	var tourNodes []*TourNode

	for i := 0; i < len(tm.players); i++ {
		tourNodes = append(tourNodes, &TourNode{Children: nil, Winner: tm.players[i]})
	}

	for len(tourNodes) > 1 {
		for j := 0; j < len(tourNodes); j += 2 {
			right := tourNodes[j]
			left := tourNodes[j+1]
			children := []*TourNode{left, right}
			var tourNode *TourNode = &TourNode{Children: children, Winner: Player{}}
			var tourNodes1 = append(tourNodes[0:j], tourNode)
			tourNodes = append(tourNodes1, tourNodes[j+2:]...)
		}
	}

	return tourNodes[0]
}

func (tm *tournMnger) BuildAllVsAll() *TourNode {
	var children []*TourNode
	var root *TourNode = &TourNode{Children: children, Winner: Player{}}
	for i := 0; i < len(tm.players); i++ {
		var child *TourNode = &TourNode{Children: nil, Winner: tm.players[i]}
		root.Children = append(root.Children, child)
	}
	return root
}
