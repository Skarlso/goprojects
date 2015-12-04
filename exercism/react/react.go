package react

//TestVersion testVersion
const TestVersion = 1

//MyCells

//MyInputCell my input cell
type MyInputCell struct {
	value   int
	reactor MyReactor
}

//MyComputeCell my input cell
type MyComputeCell struct {
	callBacks       map[CallbackHandle]func(int)
	computeFunction interface{}
	value           int
	dependsOn       []Cell
	reactor         MyReactor
}

//MyReactor my reactor
type MyReactor struct {
	connections1 map[Cell][]*MyComputeCell
	connections2 map[Cell][]*MyComputeCell
}

/*
*
* REACTOR IMPLEMENTATIONS
*
 */

//CreateInput creates input cell
func (r MyReactor) CreateInput(value int) InputCell {
	return &MyInputCell{value: value, reactor: r}
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (r *MyReactor) CreateCompute1(cell Cell, computeF func(int) int) ComputeCell {
	return r.ProxyCompute([]Cell{cell}, computeF)
}

// CreateCompute2 is like CreateCompute1, but depending on two cells
// The compute function will only be called if the value of any of the
// passed cells changes.
func (r *MyReactor) CreateCompute2(cell1 Cell, cell2 Cell, computeF func(int, int) int) ComputeCell {
	return r.ProxyCompute([]Cell{cell1, cell2}, computeF)
}

//ProxyCompute a proxy to get the compute functions together
func (r *MyReactor) ProxyCompute(cell []Cell, computeF interface{}) ComputeCell {
	myComputeCell := MyComputeCell{}
	myComputeCell.callBacks = make(map[CallbackHandle]func(int))
	myComputeCell.computeFunction = computeF
	if len(cell) < 2 {
		myComputeCell.value = computeF.(func(int) int)(cell[0].Value())
		connections1 := append(r.connections1[cell[0]], &myComputeCell)
		r.connections1[cell[0]] = connections1
	} else {
		myComputeCell.value = computeF.(func(int, int) int)(cell[0].Value(), cell[1].Value())
		myComputeCell.dependsOn = []Cell{cell[0], cell[1]}
		connectionsCell1 := append(r.connections2[cell[0]], &myComputeCell)
		connectionsCell2 := append(r.connections2[cell[1]], &myComputeCell)
		r.connections2[cell[0]] = connectionsCell1
		r.connections2[cell[1]] = connectionsCell2
	}

	myComputeCell.reactor = *r
	return &myComputeCell

}

/*
*
* INPUT CELL IMPLEMENTATIONS
*
 */

//SetValue sets the value of an input cell
func (ci *MyInputCell) SetValue(value int) {
	if ci.value == value {
		return
	}
	ci.value = value
	for _, v := range ci.reactor.connections1[ci] {
		computedValue := v.computeFunction.(func(int) int)(ci.value)
		if v.value != computedValue {
			v.value = computedValue
			for _, callBack := range v.callBacks {
				callBack(v.value)
			}
		}
		if dependingComputeCells, ok := ci.reactor.connections2[v]; ok {
			for _, dcc := range dependingComputeCells {
				depValue := dcc.computeFunction.(func(int, int) int)(dcc.dependsOn[0].Value(), dcc.dependsOn[1].Value())
				if dcc.value != depValue {
					dcc.value = depValue
					for _, dcCallBack := range dcc.callBacks {
						dcCallBack(dcc.value)
					}
				}
			}
		}
	}
}

//Value return value
func (ci MyInputCell) Value() int {
	return ci.value
}

/*
*
* COMPUTE1 CELL IMPLEMENTATIONS
*
 */

//AddCallback sets the value of an input cell
func (cc *MyComputeCell) AddCallback(callBack func(int)) CallbackHandle {
	cc.callBacks[&callBack] = callBack
	return &callBack
}

//RemoveCallback sets the value of an input cell
func (cc *MyComputeCell) RemoveCallback(callBack CallbackHandle) {
	delete(cc.callBacks, callBack)
}

//Value returns value
func (cc MyComputeCell) Value() int {
	return cc.value
}

/*
*
* REACTOR NEW METHOD
*
 */

//New returns a new reactor
func New() Reactor {
	return &MyReactor{connections1: make(map[Cell][]*MyComputeCell), connections2: make(map[Cell][]*MyComputeCell)}
}
