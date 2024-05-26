package usecase

import "github.com/JoshEvan/solomon/driver/usecase"

type Factory interface {
	NewProductUsecase() usecase.Usecase
}
