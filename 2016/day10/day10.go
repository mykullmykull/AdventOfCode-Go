package day10

import (
	"fmt"
	"goAoc2016/utils"
	"strings"
)

type bot struct {
	id               int
	chipLow          int
	chipHigh         int
	giveLowToId      int
	giveHighToId     int
	giveLowToOutput  bool
	giveHighToOutput bool
}

func (b *bot) takeChip(chip int) bot {
	if b.chipLow != -1 {
		panic(fmt.Sprintf("bot %d already has a low chip %d", b.id, b.chipLow))
	}

	if b.chipHigh < chip {
		b.chipLow = b.chipHigh
		b.chipHigh = chip
		return *b
	}

	b.chipLow = chip
	return *b
}

func runB(input []string) int {
	bots := createBots(input)
	bots = fleshOutBots(bots, input)
	bots = giveChipsToInitialBots(bots, input)
	o1, o2, o3 := moveChipsImmediatelyWhenBotHasTwo(bots)
	return o1 * o2 * o3
}

func moveChipsImmediatelyWhenBotHasTwo(bots []bot) (int, int, int) {
	o1, o2, o3 := -1, -1, -1
	for {
		if o1 != -1 && o2 != -1 && o3 != -1 {
			return o1, o2, o3
		}

		for botId := 0; botId < len(bots); botId++ {
			bot := bots[botId]
			if bot.chipLow != -1 && bot.chipHigh != -1 {
				if bot.giveLowToOutput {
					switch bot.giveLowToId {
					case 0:
						o1 = bot.chipLow
					case 1:
						o2 = bot.chipLow
					case 2:
						o3 = bot.chipLow
					}
				} else {
					bots[bot.giveLowToId].takeChip(bot.chipLow)
				}
				bots[botId].chipLow = -1

				if bot.giveHighToOutput {
					switch bot.giveHighToId {
					case 0:
						o1 = bot.chipHigh
					case 1:
						o2 = bot.chipHigh
					case 2:
						o3 = bot.chipHigh
					}
				} else {
					bots[bot.giveHighToId].takeChip(bot.chipHigh)
				}
				bots[botId].chipHigh = -1
			}
		}
	}
}

func moveChipsToOutputs(bots []bot) (int, int, int) {
	o1, o2, o3 := -1, -1, -1
	for {
		printBots(bots)
		if o1 != -1 && o2 != -1 && o3 != -1 {
			return o1, o2, o3
		}

		botsToGive := []int{}
		for botId := 0; botId < len(bots); botId++ {
			bot := bots[botId]
			if bot.chipLow != -1 && bot.chipHigh != -1 {
				botsToGive = append(botsToGive, botId)
			}

		}

		for _, botId := range botsToGive {
			giver := bots[botId]
			if giver.giveHighToOutput && giver.giveHighToId == 0 {
				o1 = giver.chipHigh
			}

			if giver.giveHighToOutput && giver.giveHighToId == 1 {
				o2 = giver.chipHigh
			}

			if giver.giveHighToOutput && giver.giveHighToId == 2 {
				o3 = giver.chipHigh
			}

			if !giver.giveHighToOutput {
				bots[giver.giveHighToId].takeChip(giver.chipHigh)
			}
			bots[botId].chipHigh = -1

			if giver.giveHighToOutput {
				if giver.giveHighToId == 0 {
					o1 = giver.chipHigh
				}
				if giver.giveHighToId == 1 {
					o2 = giver.chipHigh
				}
				if giver.giveHighToId == 2 {
					o3 = giver.chipHigh
				}
			} else {
				bots[giver.giveLowToId].takeChip(giver.chipLow)
			}
			bots[botId].chipLow = -1
		}
	}
}

func runA(input []string, finalLow int, finalHigh int) int {
	bots := createBots(input)
	bots = fleshOutBots(bots, input)
	bots = giveChipsToInitialBots(bots, input)
	return moveChips(bots, finalLow, finalHigh)
}

func fleshOutBots(bots []bot, input []string) []bot {
	for _, row := range input {
		split := strings.Split(row, " ")
		if split[0] == "bot" {
			botId := utils.Atoi(split[1])
			giveLowToId := utils.Atoi(split[6])
			giveHighToId := utils.Atoi(split[11])
			bots[botId].giveLowToId = giveLowToId
			bots[botId].giveHighToId = giveHighToId
			bots[botId].giveLowToOutput = split[5] == "output"
			bots[botId].giveHighToOutput = split[10] == "output"
		}
	}
	return bots
}

func giveChipsToInitialBots(bots []bot, input []string) []bot {
	// give chips to initial bots
	for _, row := range input {
		split := strings.Split(row, " ")
		if split[0] == "value" {
			botId := utils.Atoi(split[5])
			chip := utils.Atoi(split[1])
			bots[botId] = bots[botId].takeChip(chip)
		}
	}
	return bots
}

func moveChips(bots []bot, finalLow int, finalHigh int) int {
	for {
		for botId := 0; botId < len(bots); botId++ {
			bot := bots[botId]
			if bot.chipLow == finalLow && bot.chipHigh == finalHigh {
				return botId
			}
		}

		botsToGive := []int{}
		for botId := 0; botId < len(bots); botId++ {
			bot := bots[botId]
			if bot.chipLow != -1 && bot.chipHigh != -1 {
				botsToGive = append(botsToGive, botId)
			}

		}

		for _, botId := range botsToGive {
			giver := bots[botId]
			bots[giver.giveHighToId].takeChip(giver.chipHigh)
			bots[botId].chipHigh = -1
			bots[giver.giveLowToId].takeChip(giver.chipLow)
			bots[botId].chipLow = -1
		}
	}
}

func createBots(input []string) []bot {
	maxBotId := -1
	for _, row := range input {
		split := strings.Split(row, " ")
		if split[0] == "bot" {
			botId := utils.Atoi(split[1])
			if botId > maxBotId {
				maxBotId = botId
			}
		}
		if split[0] == "value" {
			botId := utils.Atoi(split[5])
			if botId > maxBotId {
				maxBotId = botId
			}
		}
	}
	bots := []bot{}
	for i := 0; i <= maxBotId+1; i++ {
		bots = append(bots, bot{id: i, chipLow: -1, chipHigh: -1, giveLowToId: -1, giveHighToId: -1})
	}
	return bots
}

func printBots(bots []bot) {
	for _, bot := range bots {
		if bot.chipLow != -1 && bot.chipHigh != -1 {
			fmt.Printf("bot %d: low: %d high: %d lowToId: %d highToId: %d\n", bot.id, bot.chipLow, bot.chipHigh, bot.giveLowToId, bot.giveHighToId)
		}
	}
	fmt.Println()
}
