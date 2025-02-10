package ui

import (
	"fyne.io/fyne/v2"
)

type Icons struct {
	MainIcon       fyne.Resource
	AddIcon        fyne.Resource
	BackIcon       fyne.Resource
	SaveIcon       fyne.Resource
	InfoIcon       fyne.Resource
	DeleteIcon     fyne.Resource
	ReadyIcon      fyne.Resource
	FirstLVLIcon   fyne.Resource
	SecondLVLIcon  fyne.Resource
	ThirdLVLICon   fyne.Resource
	FourthLVLICon  fyne.Resource
	FithLVLICon    fyne.Resource
	SixthLVLICon   fyne.Resource
	SeventhLVLICon fyne.Resource
	EighthLVLICon  fyne.Resource
	NinthLVLICon   fyne.Resource
	TenthLVLICon   fyne.Resource
	FireStreak     fyne.Resource
	FireLose       fyne.Resource
}

func LoadIcons() *Icons {
	// Использование встроенных ресурсов вместо загрузки с пути
	mainIcon := resourceMainiconPng
	addIcon := resourceAddPng
	backIcon := resourceBackPng
	saveIcon := resourceSavePng
	infoIcon := resourceInfoPng
	deleteIcon := resourceDeletePng
	readyIcon := resourceReadyPng
	firstLVLIcon := resourceFirstLVLPng
	secondLVLIcon := resourceSecondLVLPng
	thirdLVLICon := resourceThirdLVLPng
	fourthLVLICon := resourceFourthLVLPng
	fithLVLICon := resourceFifthLVLPng
	sixthLVLICon := resourceSixthLVLPng
	seventhLVLICon := resourceSeventhLVLPng
	eighthLVLICon := resourceEighthLVLPng
	ninthLVLICon := resourceNinthLVLPng
	tenthLVLICon := resourceTenthLVLPng
	firestreak := resourceFirestreakPng
	firelose := resourceFirelosePng

	return &Icons{
		MainIcon:       mainIcon,
		AddIcon:        addIcon,
		BackIcon:       backIcon,
		SaveIcon:       saveIcon,
		InfoIcon:       infoIcon,
		DeleteIcon:     deleteIcon,
		FirstLVLIcon:   firstLVLIcon,
		SecondLVLIcon:  secondLVLIcon,
		ThirdLVLICon:   thirdLVLICon,
		FourthLVLICon:  fourthLVLICon,
		FithLVLICon:    fithLVLICon,
		SixthLVLICon:   sixthLVLICon,
		SeventhLVLICon: seventhLVLICon,
		EighthLVLICon:  eighthLVLICon,
		NinthLVLICon:   ninthLVLICon,
		TenthLVLICon:   tenthLVLICon,
		ReadyIcon:      readyIcon,
		FireStreak:     firestreak,
		FireLose:       firelose,
	}
}
