package common

import "log"

func PrintGreenOk(a ...any) {
	log.Println(COLOR_GREEN, a, COLOR_RESET)
}

func PrintBlueStatus(a ...any) {
	log.Println(COLOR_BLUE, a, COLOR_RESET)
}

func PrintCyanStatus(a ...any) {
	log.Println(COLOR_CYAN, a, COLOR_RESET)
}

func PrintYellowOperation(a ...any) {
	log.Println(COLOR_YELLOW, a, COLOR_RESET)
}

func PrintPurpleWarning(a ...any) {
	log.Println(COLOR_PURPLE, a, COLOR_RESET)
}

func PrintRedError(a ...any) {
	log.Println(COLOR_RED, a, COLOR_RESET)
}
