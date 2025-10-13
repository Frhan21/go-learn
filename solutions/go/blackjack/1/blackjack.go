package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
        case card == "ace": 
        	return 11; 
        case card == "ten" || card == "jack" || card == "queen" || card == "king" :
        	return 10; 
        case card == "two": 
        	return 2; 
        case card == "three": 
        	return 3; 
        case card == "four": 
        	return 4; 
        case card == "five": 
        	return 5; 
        case card == "six": 
        	return 6; 
        case card == "seven": 
        	return 7;
        case card == "eight": 
        	return 8;
        case card == "nine":
        	return 9;
        default: 
    		return 0; 
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    playerCard := ParseCard(card1) + ParseCard(card2) 

    if card1 == "ace" && card2 == "ace" {
        return "P"
    }

    if playerCard == 21 {
        if dealerCard == "king" || dealerCard == "queen" || dealerCard == "jack" || dealerCard == "ten" || dealerCard == "ace"{
            return  "S"
        }
    	return "W"
    }

    if playerCard >= 17 && playerCard <= 20 {
        return "S"
    } 

    if playerCard >= 12 && playerCard <= 16 {
        if ParseCard(dealerCard) >= 7 {
            return "H"
        }
        return "S"
    }

    if playerCard <= 11 {
         return "H"
    }

    return "S"
}
