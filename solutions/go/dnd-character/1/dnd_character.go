package dndcharacter
import ("math/rand"
        "slices"
       "math")

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	//panic("Please implement the Modifier() function")

    modifier := int(math.Floor(float64(score - 10) / 2.0))
    return modifier

}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	//panic("Please implement the Ability() function")
    
   // strength := 
    //dexterity :=
    //constitution :=
    //intelligence :=
    //wisdom :=
    //charisma :=
	ability := []int{}
    sum := 0
    
    for i := 0; i < 4; i++{
        ability = append(ability, 1 + rand.Intn(6))
    }
	minVal := slices.Min(ability)
    
    for _, value := range ability{
        sum += value
    }
    
    return sum - minVal
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	//panic("Please implement the GenerateCharacter() function")
		strength := Ability()
        dexterity := Ability()
        constitution := Ability()
        intelligence :=  Ability()
        wisdom := Ability()
        charisma := Ability()
    	
    hitpoints := 10 + Modifier(constitution)
    
   return Character{
       
        Strength: strength,
        Dexterity: dexterity,
        Constitution: constitution,
        Intelligence: intelligence,
        Wisdom: wisdom,
        Charisma: charisma,
        Hitpoints: hitpoints,
    }
}
