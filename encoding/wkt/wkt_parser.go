package wkt

import (
	"fmt"
	"strconv"

	"github.com/spatial-go/geoos/space"
)

// Parser ...
type Parser struct {
	*Lexer
}

// Parse ...
func (p *Parser) Parse() (space.Geometry, error) {
	t, err := p.scanToken()
	if err != nil {
		return nil, err
	}
	var srid = 0
	var geom space.Geometry
	switch t.ttype {
	case PointEnum:
		geom, err = p.parsePoint()
	case Linestring:
		geom, err = p.parseLineString()
	case PolygonEnum:
		geom, err = p.parsePolygon()
	case Multipoint:
		line, err := p.parseLineString()
		if err != nil {
			return nil, err
		}
		geom = space.MultiPoint(line.ToPointArray())
	case MultilineString:
		poly, err := p.parsePolygon()
		if err != nil {
			return nil, err
		}
		multiline := make(space.MultiLineString, 0, len(poly))
		for _, ring := range poly {
			multiline = append(multiline, space.LineString(ring))
		}
		geom = multiline
	case MultiPolygonEnum:
		geom, err = p.parseMultiPolygon()
	case GeometryCollection:
		geom, err = p.parseGeometryCollection()

	case Srid:
		if srid == 0 {
			srid, _ = p.parseSrid()
		}
		geom, err = p.Parse()
	default:
		return nil, fmt.Errorf("Parse unexpected token %s on pos %d expected geometry type", t.lexeme, t.pos)
	}
	if err != nil {
		return nil, err
	}
	if srid != 0 {
		return space.CreateElementValidWithCoordSys(geom, srid)
	}
	return geom, nil
}

func (p *Parser) parseSrid() (srid int, err error) {
	t, err := p.scanToken()
	if err != nil {
		return 0, err
	}
	switch t.ttype {
	case EqualSign:
		s, err := p.scanToken()
		if err != nil {
			return 0, err
		}
		srid, _ = strconv.Atoi(s.lexeme)
		if s, err := p.scanToken(); err != nil || s.ttype != Semicolon {
			return srid, fmt.Errorf("parse srid unexpected token %s on pos %d", s.lexeme, s.pos)
		}
	}
	return
}

func (p *Parser) parsePoint() (point space.Point, err error) {
	t, err := p.scanToken()
	if err != nil {
		return point, err
	}
	switch t.ttype {
	case Empty:
		point = space.Point{0, 0}
	case Z, M, ZM:
		t1, err := p.scanToken()
		if err != nil {
			return point, err
		}
		if t1.ttype == Empty {
			point = space.Point{0, 0}
			break
		}
		if t1.ttype != LeftParen {
			return point, fmt.Errorf("parse point unexpected token %s on pos %d", t.lexeme, t.pos)
		}
		fallthrough
	case LeftParen:
		switch t.ttype { // reswitch on the type because of the fallthrough
		case Z, M:
			point, err = p.parseCoordDrop1()
		case ZM:
			point, err = p.parseCoordDrop2()
		default:
			point, err = p.parseCoord()
		}
		if err != nil {
			return point, err
		}

		t, err = p.scanToken()
		if err != nil {
			return point, err
		}

		if t.ttype != RightParen {
			return point, fmt.Errorf("parse point unexpected token %s on pos %d expected )", t.lexeme, t.pos)
		}
	default:
		return point, fmt.Errorf("parse point unexpected token %s on pos %d", t.lexeme, t.pos)
	}

	return point, nil
}

func (p *Parser) parseLineString() (line space.LineString, err error) {
	line = make([][]float64, 0)
	t, err := p.scanToken()
	if err != nil {
		return nil, err
	}
	switch t.ttype {
	case Empty:
	case Z, M, ZM:
		t1, err := p.scanToken()
		if err != nil {
			return line, err
		}
		if t1.ttype == Empty {
			break
		}
		if t1.ttype != LeftParen {
			return line, fmt.Errorf("unexpected token %s on pos %d expected '('", t.lexeme, t.pos)
		}
		fallthrough
	case LeftParen:
		line, err = p.parseLineStringText(t.ttype)
		if err != nil {
			return line, err
		}
	default:
		return line, fmt.Errorf("unexpected token %s on pos %d", t.lexeme, t.pos)
	}

	return line, nil
}

func (p *Parser) parseLineStringText(ttype tokenType) (line space.LineString, err error) {
	line = make([][]float64, 0)
	for {
		var point space.Point
		switch ttype {
		case Z, M:
			point, err = p.parseCoordDrop1()
		case ZM:
			point, err = p.parseCoordDrop2()
		default:
			point, err = p.parseCoord()
		}
		if err != nil {
			return line, err
		}
		line = append(line, point)
		t, err := p.scanToken()
		if err != nil {
			return line, err
		}
		if t.ttype == RightParen {
			break
		} else if t.ttype != Comma {
			return line, fmt.Errorf("unexpected token %s on pos %d expected ','", t.lexeme, t.pos)
		}
	}
	return line, nil
}

func (p *Parser) parsePolygon() (poly space.Polygon, err error) {
	poly = make([][][]float64, 0)
	t, err := p.scanToken()
	if err != nil {
		return poly, err
	}
	switch t.ttype {
	case Empty:
	case Z, M, ZM:
		t1, err := p.scanToken()
		if err != nil {
			return poly, err
		}
		if t1.ttype == Empty {
			break
		}
		if t1.ttype != LeftParen {
			return poly, fmt.Errorf("unexpected token %s on pos %d expected '('", t.lexeme, t.pos)
		}
		fallthrough
	case LeftParen:
		poly, err = p.parsePolygonText(t.ttype)
		if err != nil {
			return poly, err
		}
	default:
		return poly, fmt.Errorf("unexpected token %s on pos %d", t.lexeme, t.pos)
	}

	return poly, nil
}

func (p *Parser) parsePolygonText(ttype tokenType) (poly space.Polygon, err error) {
	poly = make([][][]float64, 0)
	for {
		var line space.LineString
		t, err := p.scanToken()
		if err != nil {
			return poly, err
		}
		if t.ttype != LeftParen {
			return poly, fmt.Errorf("unexpected token %s on pos %d expected '('", t.lexeme, t.pos)
		}
		line, err = p.parseLineStringText(ttype)
		if err != nil {
			return poly, err
		}
		poly = append(poly, space.Ring(line))
		t, err = p.scanToken()
		if err != nil {
			return poly, err
		}
		if t.ttype == RightParen {
			break
		} else if t.ttype != Comma {
			return poly, fmt.Errorf("unexpected token %s on pos %d expected ','", t.lexeme, t.pos)
		}
	}
	return poly, nil
}

func (p *Parser) parseMultiPolygon() (multi space.MultiPolygon, err error) {
	multi = make([]space.Polygon, 0)
	t, err := p.scanToken()
	if err != nil {
		return multi, err
	}
	switch t.ttype {
	case Empty:
	case Z, M, ZM:
		t1, err := p.scanToken()
		if err != nil {
			return multi, err
		}
		if t1.ttype == Empty {
			break
		}
		if t1.ttype != LeftParen {
			return multi, fmt.Errorf("unexpected token %s on pos %d expected '('", t.lexeme, t.pos)
		}
		fallthrough
	case LeftParen:
		multi, err = p.parseMultiPolygonText(t.ttype)
		if err != nil {
			return multi, err
		}
	default:
		return multi, fmt.Errorf("unexpected token %s on pos %d", t.lexeme, t.pos)
	}

	return multi, nil
}

func (p *Parser) parseMultiPolygonText(ttype tokenType) (multi space.MultiPolygon, err error) {
	multi = make([]space.Polygon, 0)
	for {
		var poly space.Polygon
		t, err := p.scanToken()
		if err != nil {
			return multi, err
		}
		if t.ttype != LeftParen {
			return multi, fmt.Errorf("unexpected token %s on pos %d expected '('", t.lexeme, t.pos)
		}
		poly, err = p.parsePolygonText(ttype)
		if err != nil {
			return multi, err
		}
		multi = append(multi, poly)
		t, err = p.scanToken()
		if err != nil {
			return multi, err
		}
		if t.ttype == RightParen {
			break
		} else if t.ttype != Comma {
			return multi, fmt.Errorf("unexpected token %s on pos %d expected ','", t.lexeme, t.pos)
		}
	}
	return multi, nil
}
func (p *Parser) parseGeometryCollection() (coll space.Collection, err error) {
	coll = make(space.Collection, 0)
	t, err := p.scanToken()
	if err != nil {
		return coll, err
	}
	switch t.ttype {
	case Empty:
	case Z, M, ZM:
		t1, err := p.scanToken()
		if err != nil {
			return coll, err
		}
		if t1.ttype == Empty {
			break
		}
		if t1.ttype != LeftParen {
			return coll, fmt.Errorf("unexpected token %s on pos %d expected '('", t.lexeme, t.pos)
		}
		fallthrough
	case LeftParen:
		coll, err = p.parseGeometryCollectionText(t.ttype)
		if err != nil {
			return coll, err
		}
	default:
		return coll, fmt.Errorf("unexpected token %s on pos %d", t.lexeme, t.pos)
	}

	return coll, nil
}

func (p *Parser) parseGeometryCollectionText(ttype tokenType) (coll space.Collection, err error) {
	coll = make(space.Collection, 0)
	for {
		geom, err := p.Parse()
		if err != nil {
			return coll, err
		}
		coll = append(coll, geom)

		t, err := p.scanToken()
		if err != nil {
			return coll, err
		}
		if t.ttype == RightParen {
			break
		} else if t.ttype != Comma {
			return coll, fmt.Errorf("unexpected token %s on pos %d expected ','", t.lexeme, t.pos)
		}
	}
	return coll, nil
}

func (p *Parser) parseCoord() (point space.Point, err error) {
	t1, err := p.scanToken()
	if err != nil {
		return point, err
	}
	if t1.ttype != Float {
		return point, fmt.Errorf("parse coordinates unexpected token %s on pos %d", t1.lexeme, t1.pos)
	}
	t2, err := p.scanToken()
	if err != nil {
		return point, err
	}
	if t2.ttype != Float {
		return point, fmt.Errorf("parse coordinates unexpected token %s on pos %d", t1.lexeme, t2.pos)
	}

	c1, err := strconv.ParseFloat(t1.lexeme, 64)
	if err != nil {
		return point, fmt.Errorf("invalid lexeme %s for token on pos %d", t1.lexeme, t1.pos)
	}
	c2, err := strconv.ParseFloat(t2.lexeme, 64)
	if err != nil {
		return point, fmt.Errorf("invalid lexeme %s for token on pos %d", t2.lexeme, t2.pos)
	}

	return space.Point{c1, c2}, nil
}

func (p *Parser) parseCoordDrop1() (point space.Point, err error) {
	point, err = p.parseCoord()
	if err != nil {
		return point, err
	}

	// drop the last value Z or M coordinates are not really supported
	t, err := p.scanToken()
	if err != nil {
		return point, err
	}
	if t.ttype != Float {
		return point, fmt.Errorf("parseCoordDrop1 unexpected token %s on pos %d expected Float", t.lexeme, t.pos)
	}

	return point, nil
}

func (p *Parser) parseCoordDrop2() (point space.Point, err error) {
	point, err = p.parseCoord()
	if err != nil {
		return point, err
	}

	// drop the last value M values
	// and Z coordinates are not really supported
	for i := 0; i < 2; i++ {
		t, err := p.scanToken()
		if err != nil {
			return point, err
		}
		if t.ttype != Float {
			return point, fmt.Errorf("parseCoordDrop2 unexpected token %s on pos %d expected Float", t.lexeme, t.pos)
		}
	}

	return point, nil
}
