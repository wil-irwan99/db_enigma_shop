package main

import (
	"fmt"
	"gorm/config"
	"gorm/repository"
	"gorm/usecase"
	"gorm/utils"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	cfg := config.NewConfig()
	//cfg.DbConn()
	db := cfg.DbConn()
	defer cfg.DbClose()

	//repo := repository.NewCustomerRepository(db)
	//uscs := usecase.NewMemberActivitionUsecase(repo)

	//db.AutoMigrate(&model.Customer{}) //mengenerate/membuat table dari struct

	//Insert
	// customer := model.Customer{
	// 	Id:      utils.GenerateId(),
	// 	Name:    "Bulan Sutisna",
	// 	Address: "Bali",
	// 	Balance: 100000,
	// 	UserCredential: model.UserCredential{
	// 		UserName: "bulanbulanan",
	// 		Password: "password",
	// 	},
	// 	Email: "bulan.s@gmail.com",
	// 	Phone: "20202020",
	// }

	// err := repo.Create(&customer)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//Update
	// customerExisting := model.Customer{
	// 	Id: "e9f4862f-7d4f-405f-ad9e-ac96eed9db09",
	// }

	// err := repo.Update(&customerExisting, map[string]interface{}{
	// 	"address":   "",
	// 	"balance":   15000,
	// 	"is_status": 0,
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//Delete by ID
	// err := repo.Delete(&customerExisting)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//Find by ID
	// customerExisting, err := repo.FindById(customerExisting.Id)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//fmt.Println(customerExisting)

	//FindAllBy
	// customers := []model.Customer{}
	// customers, err := repo.FindAllBy(map[string]interface{}{
	// 	"address": "Denmark",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customers)

	//FindFirstBy
	// customers := model.Customer{}
	// customers, err := repo.FindFirstBy(map[string]interface{}{
	// 	"address": "Denmark",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customers)

	//FindBy
	// customers := []model.Customer{}
	// customers, err := repo.FindBy("name like ? and is_status = ?", "%a%", 1)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customers)

	//Count
	// var TotalCustomerStatus []struct {
	// 	Name     string
	// 	IsStatus int
	// 	Total    int64
	// }
	// err := repo.Count(&TotalCustomerStatus, "is_status")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(TotalCustomerStatus)
	//Count jika Group by kosong
	// var total int64
	// err := repo.Count(&total, "")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(total)

	//Group By
	// var Result []struct {
	// 	Address string
	// 	Total   int64
	// }
	// err := repo.GroupBy(&Result, "address, count(address) as total", nil, "address")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(Result)

	//Paging
	// customerPaging, err := repo.Paging(1, 2)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println("Result fro customerPaging")
	// fmt.Println(customerPaging)

	//Update With Preload
	// customer, err := repo.FindByIdWithPreload(map[string]interface{}{"id": "73686850-4ac3-4133-9946-f6c3d1f353c0"}, "UserCredential")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println("Before", customer)
	// c := customer.(model.Customer)
	// c.UserCredential.Password = "rahasia"
	// err = repo.UpdateBy(&c)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// customer, err = repo.FindByIdWithPreload(map[string]interface{}{"id": "73686850-4ac3-4133-9946-f6c3d1f353c0"}, "UserCredential")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println("After", customer)

	//Insert data with multiple address
	// pass, _ := HashPassword("password")
	// customer := model.Customer{
	// 	Id:             utils.GenerateId(),
	// 	Name:           "Fadli Zona Nyaman",
	// 	Phone:          "14045",
	// 	Email:          "orderaja@gmail.com",
	// 	Balance:        5000,
	// 	UserCredential: model.UserCredential{UserName: "fadlizone", Password: pass},
	// 	Address: []model.Address{
	// 		{
	// 			StreetName: "Jl Nin Aja",
	// 			City:       "Bogor",
	// 			PostalCode: "123",
	// 		},
	// 		{
	// 			StreetName: "Jl Terus",
	// 			City:       "Bandung",
	// 			PostalCode: "456",
	// 		},
	// 	},
	// }
	// err := repo.Create(&customer)
	// utils.IsError(err)

	// customer, err := repo.FindByIdWithPreload(
	// 	map[string]interface{}{"id": "376b2dc0-1c2b-4319-a9ef-80288f4318b0"},
	// 	"Address",
	// )
	// utils.IsError(err)
	// fmt.Println(customer.ToString())

	//tugas 1 login
	//custRepo := repository.NewCustomerRepository(db)
	//userCredRepo := repository.NewUserCredsRepository(db)

	//loginUscs := usecase.NewAuthenticationUsecase(custRepo, userCredRepo)
	//loginUscs.Authentication("fadlizone", "pass")

	//tugas 2
	// memberActUscs := usecase.NewMemberActivitionUsecase(custRepo)
	// memberActUscs.MemberActivation("2e245d7e-4f1c-46ef-abe2-f8ae6a477b01")

	//tugas 3
	//saldoUscs := usecase.NewCustomerBalanceUsecase(custRepo)
	//saldoUscs.BalanceProcessing("2e245d7e-4f1c-46ef-abe2-f8ae6a477b01", "+", 10000)

	//create Product
	// product01 := model.Product{
	// 	ProductName: "Kacang Manis",
	// }

	// productRepo := repository.NewProductRepository(db)
	// err := productRepo.Create(&product01)
	// utils.IsError(err)

	// passHash, _ := HashPassword("pass")
	// customer01 := model.Customer{
	// 	Id:   utils.GenerateId(),
	// 	Name: "Bulan Matahari",
	// 	Address: []model.Address{
	// 		{
	// 			StreetName: "Jl Nin Aja",
	// 			City:       "Bogor",
	// 			PostalCode: "123",
	// 		},
	// 	},
	// 	Phone:   "1872642",
	// 	Email:   "bulan.matahari@gmail.com",
	// 	Balance: 10000,
	// 	UserCredential: model.UserCredential{
	// 		UserName: "fadlizone",
	// 		Password: passHash},
	// }

	// customerRepo := repository.NewCustomerRepository(db)
	// err = customerRepo.Create(&customer01)
	// utils.IsError(err)

	//case 1 -> Insert Cust baru sekalian Product
	// passHash, _ := HashPassword("pass")
	// customer01 := model.Customer{
	// 	Id:   utils.GenerateId(),
	// 	Name: "Orang Miskin",
	// 	Address: []model.Address{
	// 		{
	// 			StreetName: "Jl Suka Suka",
	// 			City:       "Sukabumi",
	// 			PostalCode: "456721",
	// 		},
	// 	},
	// 	Phone:   "6666312312",
	// 	Email:   "perlu.duitg@gmail.com",
	// 	Balance: 0,
	// 	UserCredential: model.UserCredential{
	// 		UserName: "oragnmiskin",
	// 		Password: passHash},
	// Products: []*model.Product{
	// 	{
	// 		ProductName: "Caca Marica",
	// 	},
	// 	{
	// 		ProductName: "Beng Beng",
	// 	},
	// },
	// }
	// customerRepo := repository.NewCustomerRepository(db)
	// err := customerRepo.Create(&customer01)
	// utils.IsError(err)

	//case 2 ->  menambahkan produk baru ke cust yg sudah terdaftar
	// customerRepo := repository.NewCustomerRepository(db)
	// cust, err := customerRepo.FindById("a6f27fd9-9b8d-4f93-aca0-e7cb86dd8f4a")
	// utils.IsError(err)
	// cust.Products = []*model.Product{
	// 	{
	// 		ProductName: "Yakult",
	// 	},
	// 	// {
	// 	// 	ProductName: "Ale-ale",
	// 	// },
	// 	// {
	// 	// 	ProductName: "Chocolatos",
	// 	// },
	// }
	// err = customerRepo.UpdateByModel(&cust)
	// utils.IsError(err)

	//case 3 -> Cust baru dengan item yg sudah exist
	// prdRepo := repository.NewProductRepository(db)
	// prod, err := prdRepo.FindById(7)
	// utils.IsError(err)
	// fmt.Println(prod)
	// pass, _ := HashPassword("password")
	// customer01 := model.Customer{
	// 	Id:   utils.GenerateId(),
	// 	Name: "Tes Joni",
	// 	Address: []model.Address{
	// 		{
	// 			StreetName: "Jl Dulu Aja",
	// 			City:       "Bandung",
	// 			PostalCode: "15422",
	// 		},
	// 	},
	// 	Phone:   "12329976",
	// 	Email:   "jonites@gmail.com",
	// 	Balance: 15000,
	// 	UserCredential: model.UserCredential{
	// 		UserName: "jontes",
	// 		Password: pass,
	// 	},
	// 	Products: []*model.Product{
	// 		&prod,
	// 	},
	// }
	// cstRepo := repository.NewCustomerRepository(db)
	// err = cstRepo.Create(&customer01)
	// utils.IsError(err)

	//case 4 -> cust exist mengambil data prod exist
	// //prod exist
	// prdRepo := repository.NewProductRepository(db)
	// prod, err := prdRepo.FindById(8)
	// utils.IsError(err)
	// //cust exist
	// cstRepo := repository.NewCustomerRepository(db)
	// cust, err := cstRepo.FindById("0e305546-5efc-4aac-be98-4cba4dd9580b")
	// utils.IsError(err)
	// //append then update
	// cust.Products = append(cust.Products, &prod)
	// err = cstRepo.UpdateByModel(&cust)
	// utils.IsError(err)

	//case exist customer ingin menghapus relasinya, tapi tidak menghapus product di mst_product
	// prdRepo := repository.NewProductRepository(db)
	// cstRepo := repository.NewCustomerRepository(db)
	// cust, err := cstRepo.FindById("0e305546-5efc-4aac-be98-4cba4dd9580b")
	// utils.IsError(err)
	// prod, err := prdRepo.FindById(8)
	// utils.IsError(err)
	// err = cstRepo.DeleteAssociation(&cust, "Products", &prod)
	// utils.IsError(err)

	//case exist customer ingin mengupdate product relasinya
	// prdRepo := repository.NewProductRepository(db)
	// cstRepo := repository.NewCustomerRepository(db)
	// cust, err := cstRepo.FindByIdWithPreload(map[string]interface{}{"id": "0e305546-5efc-4aac-be98-4cba4dd9580b"}, "Products")
	// utils.IsError(err)
	// newProd, err := prdRepo.FindById(8)
	// utils.IsError(err)
	// var oldProduct uint = 7
	// var newProductSlice []model.Product
	// for _, prod := range cust.Products {
	// 	if (prod).ID != oldProduct {
	// 		newProductSlice = append(newProductSlice, *prod)
	// 	}
	// }
	// newProductSlice = append(newProductSlice, newProd)
	// err = cstRepo.UpdateAssociation(&cust, "Products", newProductSlice)
	// utils.IsError(err)

	//Latihan II tugas 1
	// cstRepo := repository.NewCustomerRepository(db)
	// totalProdPerCust := usecase.NewProductPerCustomer(cstRepo)
	// result, err := totalProdPerCust.CountProductPerCustomer("0e305546-5efc-4aac-be98-4cba4dd9580b")
	// utils.IsError(err)
	// fmt.Println("Total Produk : ", result)

	//Latihan II tugas 2
	// cstRepo := repository.NewCustomerRepository(db)
	// uscsProdlesCust := usecase.NewProdlessCust(cstRepo)
	// result, err := uscsProdlesCust.FindProdlesCust()
	// utils.IsError(err)
	// fmt.Println("Productless Customers: ")
	// for _, v := range result {
	// 	fmt.Println(v.Id, " ", v.Name)
	// }

	//Latihan II tugas 3
	cstRepo := repository.NewCustomerRepository(db)
	uscsProdBuyerList := usecase.NewProductBuyerList(cstRepo)
	result, err := uscsProdBuyerList.FindProductBuyerList()
	utils.IsError(err)
	fmt.Println(result)

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
