// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package graphdb

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

var GoUnusedProtection__ int

type Int int32

type GraphVertex struct {
	Name        Int     `thrift:"name,1"`
	Color       Int     `thrift:"color,2"`
	Weight      float64 `thrift:"weight,3"`
	Description string  `thrift:"description,4"`
}

func NewGraphVertex() *GraphVertex {
	return &GraphVertex{}
}

func (p *GraphVertex) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *GraphVertex) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Name = Int(v)
	}
	return nil
}

func (p *GraphVertex) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Color = Int(v)
	}
	return nil
}

func (p *GraphVertex) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return fmt.Errorf("error reading field 3: %s")
	} else {
		p.Weight = v
	}
	return nil
}

func (p *GraphVertex) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 4: %s")
	} else {
		p.Description = v
	}
	return nil
}

func (p *GraphVertex) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GraphVertex"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *GraphVertex) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:name: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Name)); err != nil {
		return fmt.Errorf("%T.name (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:name: %s", p, err)
	}
	return err
}

func (p *GraphVertex) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("color", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:color: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Color)); err != nil {
		return fmt.Errorf("%T.color (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:color: %s", p, err)
	}
	return err
}

func (p *GraphVertex) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("weight", thrift.DOUBLE, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:weight: %s", p, err)
	}
	if err := oprot.WriteDouble(float64(p.Weight)); err != nil {
		return fmt.Errorf("%T.weight (3) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:weight: %s", p, err)
	}
	return err
}

func (p *GraphVertex) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("description", thrift.STRING, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:description: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Description)); err != nil {
		return fmt.Errorf("%T.description (4) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:description: %s", p, err)
	}
	return err
}

func (p *GraphVertex) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GraphVertex(%+v)", *p)
}

type GraphEdge struct {
	FirstVertex     *GraphVertex `thrift:"firstVertex,1"`
	SecondVertex    *GraphVertex `thrift:"secondVertex,2"`
	Weight          float64      `thrift:"weight,3"`
	IsBidirectional bool         `thrift:"isBidirectional,4"`
	Description     string       `thrift:"description,5"`
}

func NewGraphEdge() *GraphEdge {
	return &GraphEdge{}
}

func (p *GraphEdge) Equals(e *GraphEdge) bool {
	if p.FirstVertex.Name == e.FirstVertex.Name && p.SecondVertex.Name == e.SecondVertex.Name {
		return true
	} else if  p.IsBidirectional && p.FirstVertex.Name == e.SecondVertex.Name && p.SecondVertex.Name == e.FirstVertex.Name {
		return true
	}
	return false
}

func (p *GraphEdge) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *GraphEdge) readField1(iprot thrift.TProtocol) error {
	p.FirstVertex = NewGraphVertex()
	if err := p.FirstVertex.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.FirstVertex)
	}
	return nil
}

func (p *GraphEdge) readField2(iprot thrift.TProtocol) error {
	p.SecondVertex = NewGraphVertex()
	if err := p.SecondVertex.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.SecondVertex)
	}
	return nil
}

func (p *GraphEdge) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return fmt.Errorf("error reading field 3: %s")
	} else {
		p.Weight = v
	}
	return nil
}

func (p *GraphEdge) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return fmt.Errorf("error reading field 4: %s")
	} else {
		p.IsBidirectional = v
	}
	return nil
}

func (p *GraphEdge) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 5: %s")
	} else {
		p.Description = v
	}
	return nil
}

func (p *GraphEdge) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GraphEdge"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *GraphEdge) writeField1(oprot thrift.TProtocol) (err error) {
	if p.FirstVertex != nil {
		if err := oprot.WriteFieldBegin("firstVertex", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:firstVertex: %s", p, err)
		}
		if err := p.FirstVertex.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.FirstVertex)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:firstVertex: %s", p, err)
		}
	}
	return err
}

func (p *GraphEdge) writeField2(oprot thrift.TProtocol) (err error) {
	if p.SecondVertex != nil {
		if err := oprot.WriteFieldBegin("secondVertex", thrift.STRUCT, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:secondVertex: %s", p, err)
		}
		if err := p.SecondVertex.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.SecondVertex)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:secondVertex: %s", p, err)
		}
	}
	return err
}

func (p *GraphEdge) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("weight", thrift.DOUBLE, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:weight: %s", p, err)
	}
	if err := oprot.WriteDouble(float64(p.Weight)); err != nil {
		return fmt.Errorf("%T.weight (3) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:weight: %s", p, err)
	}
	return err
}

func (p *GraphEdge) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("isBidirectional", thrift.BOOL, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:isBidirectional: %s", p, err)
	}
	if err := oprot.WriteBool(bool(p.IsBidirectional)); err != nil {
		return fmt.Errorf("%T.isBidirectional (4) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:isBidirectional: %s", p, err)
	}
	return err
}

func (p *GraphEdge) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("description", thrift.STRING, 5); err != nil {
		return fmt.Errorf("%T write field begin error 5:description: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Description)); err != nil {
		return fmt.Errorf("%T.description (5) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 5:description: %s", p, err)
	}
	return err
}

func (p *GraphEdge) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GraphEdge(%+v)", *p)
}

type Graph struct {
	Vertices []*GraphVertex `thrift:"vertices,1"`
	Edges    []*GraphEdge   `thrift:"edges,2"`
}

func NewGraph() *Graph {
	return &Graph{}
}

func (p *Graph) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Graph) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Vertices = make([]*GraphVertex, 0, size)
	for i := 0; i < size; i++ {
		_elem0 := NewGraphVertex()
		if err := _elem0.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem0)
		}
		p.Vertices = append(p.Vertices, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *Graph) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Edges = make([]*GraphEdge, 0, size)
	for i := 0; i < size; i++ {
		_elem1 := NewGraphEdge()
		if err := _elem1.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem1)
		}
		p.Edges = append(p.Edges, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *Graph) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Graph"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *Graph) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Vertices != nil {
		if err := oprot.WriteFieldBegin("vertices", thrift.LIST, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:vertices: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Vertices)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Vertices {
			if err := v.Write(oprot); err != nil {
				return fmt.Errorf("%T error writing struct: %s", v)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:vertices: %s", p, err)
		}
	}
	return err
}

func (p *Graph) writeField2(oprot thrift.TProtocol) (err error) {
	if p.Edges != nil {
		if err := oprot.WriteFieldBegin("edges", thrift.LIST, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:edges: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Edges)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Edges {
			if err := v.Write(oprot); err != nil {
				return fmt.Errorf("%T error writing struct: %s", v)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:edges: %s", p, err)
		}
	}
	return err
}

func (p *Graph) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Graph(%+v)", *p)
}
