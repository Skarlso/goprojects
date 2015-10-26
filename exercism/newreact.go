package react

//TestVersion testVersion
const TestVersion = 1

//MyCells

//MyInputCell my input cell
type MyInputCell struct {
	value int
}

//MyComputeCell1 my input cell
type MyComputeCell1 struct {
	callBacks       map[CallbackHandle]func(int)
	computeFunction func(int) int
	value           int
	dependsOn       Cell
}

//MyComputeCell2 my input cell - Duplicated for now so I don't have to deal with it.
type MyComputeCell2 struct {
	callBacks       map[CallbackHandle]func(int)
	computeFunction func(int, int) int
	value           int
	dependsOn       []Cell
}

//MyReactor my reactor
type MyReactor struct {
	connections1 map[Cell][]*MyComputeCell1
	connections2 map[Cell][]*MyComputeCell2
}

/*
*
* REACTOR IMPLEMENTATIONS
*
 */

//CreateInput creates input cell
func (r MyReactor) CreateInput(value int) InputCell {
	return &MyInputCell{value: value}
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (r *MyReactor) CreateCompute1(cell Cell, computeF func(int) int) ComputeCell {
	myComputeCell1 := MyComputeCell1{}
	myComputeCell1.callBacks = make(map[CallbackHandle]func(int))
	myComputeCell1.computeFunction = computeF
	myComputeCell1.value = computeF(cell.Value())
	myComputeCell1.dependsOn = cell
	connections1 := append(r.connections1[cell], &myComputeCell1)
	r.connections1[cell] = connections1
	return &myComputeCell1
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (r *MyReactor) CreateCompute2(cell1 Cell, cell2 Cell, computeF func(int, int) int) ComputeCell {
	myComputeCell2 := MyComputeCell2{}
	myComputeCell2.callBacks = make(map[CallbackHandle]func(int))
	myComputeCell2.computeFunction = computeF
	myComputeCell2.value = computeF(cell1.Value(), cell2.Value())
	myComputeCell2.dependsOn = []Cell{cell1, cell2}
	return &myComputeCell2
}

/*
*
* INPUT CELL IMPLEMENTATIONS
*
 */

//SetValue sets the value of an input cell
func (ci *MyInputCell) SetValue(value int) {
	ci.value = value
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
func (cc *MyComputeCell1) AddCallback(callBack func(int)) CallbackHandle {
	cc.callBacks[&callBack] = callBack
	return &callBack
}

//RemoveCallback sets the value of an input cell
func (cc *MyComputeCell1) RemoveCallback(callBack CallbackHandle) {
	delete(cc.callBacks, callBack)
}

//Value returns value
func (cc MyComputeCell1) Value() int {
	return cc.value
}

/*
*
* COMPUTE2 CELL IMPLEMENTATIONS
*
 */

//AddCallback sets the value of an input cell
func (cc *MyComputeCell2) AddCallback(callBack func(int)) CallbackHandle {
	cc.callBacks[&callBack] = callBack
	return &callBack
}

//RemoveCallback sets the value of an input cell
func (cc *MyComputeCell2) RemoveCallback(callBack CallbackHandle) {
	delete(cc.callBacks, callBack)
}

//Value retruns value
func (cc MyComputeCell2) Value() int {
	return cc.value
}

/*
*
* REACTOR NEW METHOD
*
 */

//New returns a new reactor
func New() Reactor {
	return &MyReactor{connections1: make(map[Cell][]*MyComputeCell1), connections2: make(map[Cell][]*MyComputeCell2)}
}

