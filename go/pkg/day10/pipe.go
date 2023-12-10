package day10

type Pipe interface {
	NextPosition(pos Position, dir Direction) (Position, Direction)
	CanConnect(tile Tile, dir Direction) bool
}

type VerticalPipe struct{}

func (p *VerticalPipe) NextPosition(pos Position, dir Direction) (Position, Direction) {
	var next Position
	if dir == South {
		next = Position{pos.X, pos.Y + 1}
	} else {
		next = Position{pos.X, pos.Y - 1}
	}
	return next, dir
}

func (p *VerticalPipe) CanConnect(tile Tile, dir Direction) bool {
	if tile.Pipe == nil || !dir.IsVertical() {
		return false
	}

	if dir == South {
		return tile.Value == '|' || tile.Value == 'L' || tile.Value == 'J'
	}

	return tile.Value == '|' || tile.Value == '7' || tile.Value == 'F'
}

type HorizontalPipe struct{}

func (p *HorizontalPipe) NextPosition(pos Position, dir Direction) (Position, Direction) {
	var next Position
	if dir == East {
		next = Position{pos.X + 1, pos.Y}
	} else {
		next = Position{pos.X - 1, pos.Y}
	}
	return next, dir
}

func (p *HorizontalPipe) CanConnect(tile Tile, dir Direction) bool {
	if tile.Pipe == nil || dir.IsVertical() {
		return false
	}

	if dir == East {
		return tile.Value == '-' || tile.Value == 'J' || tile.Value == '7'
	}

	return tile.Value == '-' || tile.Value == 'L' || tile.Value == 'F'
}

type NorthEastBendPipe struct{}

func (p *NorthEastBendPipe) NextPosition(pos Position, dir Direction) (Position, Direction) {
	var nextPos Position
	var nextDir Direction
	if dir == South {
		nextPos = Position{pos.X + 1, pos.Y}
		nextDir = East
	} else {
		nextPos = Position{pos.X, pos.Y - 1}
		nextDir = North
	}
	return nextPos, nextDir
}

func (p *NorthEastBendPipe) CanConnect(tile Tile, dir Direction) bool {
	if tile.Pipe == nil {
		return false
	}

	if dir == West {
		return tile.Value == '|' || tile.Value == '7' || tile.Value == 'F'
	}

	if dir == South {
		return tile.Value == '-' || tile.Value == '7' || tile.Value == 'J'
	}

	return false
}

type NorthWestBendPipe struct{}

func (p *NorthWestBendPipe) NextPosition(pos Position, dir Direction) (Position, Direction) {
	var nextPos Position
	var nextDir Direction
	if dir == South {
		nextPos = Position{pos.X - 1, pos.Y}
		nextDir = West
	} else {
		nextPos = Position{pos.X, pos.Y - 1}
		nextDir = North
	}
	return nextPos, nextDir
}

func (p *NorthWestBendPipe) CanConnect(tile Tile, dir Direction) bool {
	if tile.Pipe == nil {
		return false
	}

	if dir == East {
		return tile.Value == '|' || tile.Value == '7' || tile.Value == 'F'
	}

	if dir == South {
		return tile.Value == '-' || tile.Value == 'F' || tile.Value == 'L'
	}

	return false
}

type SouthWestBendPipe struct{}

func (p *SouthWestBendPipe) NextPosition(pos Position, dir Direction) (Position, Direction) {
	var nextPos Position
	var nextDir Direction
	if dir == North {
		nextPos = Position{pos.X - 1, pos.Y}
		nextDir = West
	} else {
		nextPos = Position{pos.X, pos.Y + 1}
		nextDir = South
	}
	return nextPos, nextDir
}

func (p *SouthWestBendPipe) CanConnect(tile Tile, dir Direction) bool {
	if tile.Pipe == nil {
		return false
	}

	if dir == East {
		return tile.Value == '|' || tile.Value == 'J' || tile.Value == 'L'
	}

	if dir == North {
		return tile.Value == '-' || tile.Value == 'F' || tile.Value == 'L'
	}

	return false
}

type SouthEastBendPipe struct{}

func (p *SouthEastBendPipe) NextPosition(pos Position, dir Direction) (Position, Direction) {
	var nextPos Position
	var nextDir Direction
	if dir == North {
		nextPos = Position{pos.X + 1, pos.Y}
		nextDir = East
	} else {
		nextPos = Position{pos.X, pos.Y + 1}
		nextDir = South
	}
	return nextPos, nextDir
}

func (p *SouthEastBendPipe) CanConnect(tile Tile, dir Direction) bool {
	if tile.Pipe == nil {
		return false
	}

	if dir == West {
		return tile.Value == '|' || tile.Value == 'J' || tile.Value == 'L'
	}

	if dir == North {
		return tile.Value == '-' || tile.Value == 'J' || tile.Value == '7'
	}

	return false
}
