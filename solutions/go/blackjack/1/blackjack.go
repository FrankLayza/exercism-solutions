package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch{
        case card == "ace":
        return 11
        case card == "two":
        return 2
        case card == "three":
        return 3
        case card == "four":
        return 4
        case card == "five":
        return 5
        case card == "six":
        return 6
        case card == "seven":
        return 7
        case card == "eight":
        return 8
        case card == "nine":
        return 9
        case card == "ten" || card == "jack" || card == "queen" || card == "king":
        return 10
        default: 
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    // Get the numeric value of the player's hand
    playerValue := ParseCard(card1) + ParseCard(card2)
    // Get the numeric value of the dealer's card
    dealerValue := ParseCard(dealerCard)

    switch {
    // RULE 1: Pair of aces? Always Split.
    // This must be checked first, as the value (22) isn't 21.
    case card1 == "ace" && card2 == "ace":
        return "P"

    // RULE 2: Blackjack (21) and dealer does NOT have a 10-value card or Ace?
    // This is an automatic Win.
    case (playerValue == 21 &&
        dealerCard != "ace" &&
        dealerCard != "jack" &&
        dealerCard != "king" &&
        dealerCard != "queen" &&
        dealerCard != "10"):
        return "W"

    // RULE 3: Blackjack (21) and dealer DOES have a 10-value card or Ace?
    // This is a Stand.
    // FIX: Note the ( ) around all the || conditions.
    case (playerValue == 21 &&
        (dealerCard == "ace" ||
        dealerCard == "jack" ||
        dealerCard == "king" ||
        dealerCard == "queen" ||
        dealerCard == "10")):
        return "S"

    // RULE 4: Value is 17-20? Always Stand.
    case playerValue >= 17 && playerValue <= 20:
        return "S"

    // RULE 5: Value is 12-16?
    // We must check the *specific* "Hit" rule FIRST.
    // FIX: Used dealerValue (number) instead of dealerCard (string).
    case (playerValue >= 12 && playerValue <= 16 && dealerValue >= 7):
        return "H"

    // RULE 6: Value is 12-16 and the previous rule was false?
    // This is the "Stand" rule.
    case playerValue >= 12 && playerValue <= 16:
        return "S"

    // RULE 7: Value is 11 or less? Always Hit.
    case playerValue <= 11:
        return "H"
    }
    
    // Default case: If no rules match (which shouldn't happen),
    // we return a safe bet, "S".
    return "S"
}