package main

type Order interface {
	Execute(nation *Nation)
	IsDone() bool
	Provence() *Provence
}

type BuildOrder struct {
	provence *Provence
	itemType ItemType
	count int
}

func (o *BuildOrder) Execute(nation *Nation) {
	if o.provence.Owner != nation {
		// Provence Lost. Cancel Order
		o.count = 0;
		return
	}
	numberToBuild := int(float64(o.provence.Population)/(ItemRate[o.itemType] * PopWorkerModifer))
	if numberToBuild > o.count {
		numberToBuild = o.count
	}
	cost := numberToBuild * ItemCost[o.itemType]
	if cost > nation.Metal {
		// Not enough Metal. Need a Player log to log this...
		// Keep the order though
		return
	}
	nation.Metal -= cost
	nation.Equipment[o.itemType] += numberToBuild
	o.count -= numberToBuild
}

func (o *BuildOrder) IsDone() bool {
	return o.count == 0
}

func (o *BuildOrder) Provence() *Provence {
	return o.provence
}

type TrainOrder struct {
	provence *Provence
	unitType UnitType
	count int
}

func (o *TrainOrder) Execute(nation *Nation) {
	if o.provence.Owner != nation {
		// Provence Lost. Cancel Order
		o.count = 0;
		return
	}
	numberToTrain := UnitRate[o.unitType]
	if numberToTrain > o.count {
		numberToTrain = o.count
	}
	cost := numberToTrain * UnitCost[o.unitType]
	if cost > nation.Food {
		// Not enough food (this might be an issue :p)
		// Keep the order though
		return
	}
	nation.Food -= cost
	o.count -= numberToTrain
	o.provence.Population -= numberToTrain
	//addToArmyInProvence(nation, o.provence, o.unitType, numberToTrain)
}

func (o *TrainOrder) IsDone() bool {
	return o.count == 0
}

func (o *TrainOrder) Provence() *Provence {
	return o.provence
}

type MoveOrder struct {
	army *Army
	target *Provence
}

func (o *MoveOrder) IsDone() bool {
	// A move order always counts as done, even if the move is interupted
	return true
}

func (o *MoveOrder) Provence() *Provence {
	return nil
}

func ProvincialOrders(OrderQueue []Order, provence *Provence) []Order {
	orders := []Order{}
	for _, order := range (OrderQueue) {
		if provence == order.Provence() {
			orders = append(orders, order)
		}
	}
	return orders
}
