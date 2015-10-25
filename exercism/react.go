package react

import "fmt"

//TestVersion testVersion
const TestVersion = 1

//MyCells

//MyInputCell my input cell
type MyInputCell struct {
	value         int
	computeCells1 []*MyComputeCell
	reactor       MyReactor
}

//MyComputeCell my input cell
type MyComputeCell struct {
	reactor         MyReactor
	callBacks       map[CallbackHandle]func(int)
	computeFunction func(int) int
	value           int
}

//MyComputeCell2 my input cell - Duplicated for now so I don't have to deal with it.
type MyComputeCell2 struct {
	callBacks       map[CallbackHandle]func(int)
	computeFunction func(int, int) int
	value           int
}

//MyReactor my reactor
type MyReactor struct {
	connections map[Cell][]MyComputeCell
}

//Implementation of the interfaces

//Value returns the value of my input cell
func (ci MyInputCell) Value() int {
	return ci.value
}

//SetValue sets a value for an input cell
func (ci *MyInputCell) SetValue(value int) {
	fmt.Println("Setting InputCell value too: ", value)
	fmt.Println("Current InputCell value is: ", ci.Value())
	if value == ci.Value() {
		return
	}
	ci.value = value
	ci.NotifyComputeCells()
}

//NotifyComputeCells notifies all listening cells
func (ci *MyInputCell) NotifyComputeCells() {
	//If value changed, notify all cells
	for _, computeCells := range ci.reactor.connections[ci] {
		computeCells.value = computeCells.computeFunction(ci.Value())
	}
}

//Value returns the value of my input cell
func (cc MyComputeCell) Value() int {
	return cc.value
}

//AddCallback adds a call back handle
func (cc *MyComputeCell) AddCallback(callBack func(int)) CallbackHandle {
	cc.callBacks[&callBack] = callBack
	return &callBack
}

//RemoveCallback adds a call back handle
func (cc *MyComputeCell) RemoveCallback(callBackHandle CallbackHandle) {
	delete(cc.callBacks, callBackHandle)
}

//Value returns the value of my input cell - Duplicated For Now
func (cc MyComputeCell2) Value() int {
	return cc.value
}

//AddCallback adds a call back handle  - Duplicated For Now
func (cc *MyComputeCell2) AddCallback(callBack func(int)) CallbackHandle {
	cc.callBacks[&callBack] = callBack
	return &callBack
}

//RemoveCallback adds a call back handle - Duplicated For Now
func (cc *MyComputeCell2) RemoveCallback(callBackHandle CallbackHandle) {
	delete(cc.callBacks, callBackHandle)
}

//CreateInput creates input for a reactor
func (r *MyReactor) CreateInput(value int) InputCell {
	myInputCell := MyInputCell{value: value, reactor: *r}
	return &myInputCell
}

//CreateCompute1 Creates a computed cell
func (r *MyReactor) CreateCompute1(cell Cell, callBack func(int) int) ComputeCell {
	myComputeCell := MyComputeCell{value: callBack(cell.Value()), computeFunction: callBack, callBacks: make(map[CallbackHandle]func(int))}
	cellConnections := r.connections[cell]
	cellConnections = append(cellConnections, myComputeCell)
	r.connections[cell] = cellConnections
	myComputeCell.reactor = *r
	return &myComputeCell
}

//CreateCompute2 Creates a computed cell  - Duplicated For Now
func (r *MyReactor) CreateCompute2(cell1 Cell, cell2 Cell, callBack func(int, int) int) ComputeCell {
	myComputeCell := MyComputeCell2{}
	myComputeCell.computeFunction = callBack
	myComputeCell.value = callBack(cell1.Value(), cell2.Value())
	myComputeCell.callBacks = make(map[CallbackHandle]func(int))
	return &myComputeCell
}

//New returns a new reactor
func New() Reactor {
	return &MyReactor{make(map[Cell][]MyComputeCell)}
}
