package main

func RunAI(n *Nation, tm *TurnManager) {
	// Put AI Code Here
	tm.SubmitTurn(n, n.OrderQueue)
}