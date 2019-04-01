package main

//Wesley Monaris Rodrigues
import "container/list"
import "fmt"

type List []interface{}

type Book struct{
  Code string
  Name string
  Author string
  Company string
  Year string
  Edition string
  I int
}

type Client struct{
  Code string
  Name string
  CPF string
  Age string
  I int
}
type Address struct{
  Logradouro string
  Number string
  Neighborhood string
  City string
  CEP string
  UF string
  I int
}
type Employee struct{
  Code int
  Name string
  Age string
  SoldB int
  I int
}

type Sell struct{
  SoldB int

}

func main()  {
  BList := list.New();
  CList := list.New();
  AList := list.New();
  EList := list.New();
  AEList := list.New();
  SellList := list.New();
  choice := ""
  choose := ""
  var codeB , nameB , authorB , companyB , yearB , editionB  string
  iB := 0

  var codeC , nameC ,foneC , cpfC , ageC  string
  iC := 0

  var logradouroA, numberA, neighborhoodA, cityA, cepA, ufA string
  iA := 0
  iAE := 0

  var nameE, ageE, foneE string
  iE := 0
  var codeE int
  var soldB, sell int

  j:= 0

  var vF [100][10]interface{}
  var vE [100][10]interface{}

  for choose != "exit"{
    fmt.Println("Enter option")
    fmt.Println("\033[1;34mregis")
    fmt.Println("show\033[m")
    fmt.Scanln(&choose)

    if choose == "regis"{
      fmt.Println("Enter option")
      fmt.Println("\033[1;33mbook")
      fmt.Println("cli")
      fmt.Println("emplo\033[m")
      fmt.Scanln(&choice)
    //BOOK
      if choice == "book"{
        fmt.Println("enter a code")
        fmt.Scanln(&codeB)

        fmt.Println("enter a name")
        fmt.Scanln(&nameB)

        fmt.Println("enter a author")
        fmt.Scanln(&authorB)

        fmt.Println("enter a company")
        fmt.Scanln(&companyB)

        fmt.Println("enter a year")
        fmt.Scanln(&yearB)

        fmt.Println("enter a edition")
        fmt.Scanln(&editionB)

        iB += 1

        b := Book{
          Code: codeB,
          Name: nameB,
          Author: authorB,
          Company: companyB,
          Year: yearB,
          Edition: editionB,
          I: iB,
         }

         BList.PushBack(b)
       }

    //CLIENT
     if choice == "cli"{
       fmt.Println("enter a code")
       fmt.Scanln(&codeC)

       fmt.Println("enter a name")
       fmt.Scanln(&nameC)

       fmt.Println("\033[1;33mADDRESS")
       fmt.Println("enter a logradouro")
       fmt.Scanln(&logradouroA)

       fmt.Println("enter a number")
       fmt.Scanln(&numberA)

       fmt.Println("enter a neighborhood")
       fmt.Scanln(&neighborhoodA)

       fmt.Println("enter a city")
       fmt.Scanln(&cityA)

       fmt.Println("enter a cep")
       fmt.Scanln(&cepA)

       fmt.Println("enter a uf\033[m")
       fmt.Scanln(&ufA)

       fmt.Println("How many phones do you want to register?")
       fmt.Scanln(&j)

          for i := 0;i < j;i++ {
            fmt.Println("enter a fone")
            fmt.Scanln(&foneC)
            vF[iC][i]= foneC
          }

       fmt.Println("enter a cpf")
       fmt.Scanln(&cpfC)

       fmt.Println("enter a age")
       fmt.Scanln(&ageC)

       iC += 1
       iA += 1
       c := Client{
         Code: codeC,
         Name: nameC,
        // Fone: foneC,
         CPF: cpfC,
         Age: ageC,
         I: iC,
        }
        a := Address{
          Logradouro: logradouroA,
          Number: numberA,
          Neighborhood: neighborhoodA,
          City: cityA,
          CEP: cepA,
          UF: ufA,
          I: iA,
          }
        CList.PushBack(c)
        AList.PushBack(a)
      }
      //Employee

      if choice == "emplo" {
        fmt.Println("enter a code")
        fmt.Scanln(&codeE)

        fmt.Println("enter a name")
        fmt.Scanln(&nameE)

        fmt.Println("enter a age")
        fmt.Scanln(&ageE)

        fmt.Println("\033[1;33mADDRESS")
        fmt.Println("enter a logradouro")
        fmt.Scanln(&logradouroA)

        fmt.Println("enter a number")
        fmt.Scanln(&numberA)

        fmt.Println("enter a neighborhood")
        fmt.Scanln(&neighborhoodA)

        fmt.Println("enter a city")
        fmt.Scanln(&cityA)

        fmt.Println("enter a cep")
        fmt.Scanln(&cepA)

        fmt.Println("enter a uf\033[m")
        fmt.Scanln(&ufA)

        fmt.Println("How many phones do you want to register?")
        fmt.Scanln(&j)

           for i := 0;i < j;i++ {
             fmt.Println("enter a fone")
             fmt.Scanln(&foneE)
             vE[iE][i]= foneE
           }

        fmt.Println("enter a total sold books")
        fmt.Scanln(&soldB)

        iE += 1
        iAE += 1
        e := Employee{
          Code: codeE,
          Name: nameE,
          Age: ageE,
          SoldB: soldB,
          I: iE,
         }
         ae := Address{
           Logradouro: logradouroA,
           Number: numberA,
           Neighborhood: neighborhoodA,
           City: cityA,
           CEP: cepA,
           UF: ufA,
           I: iAE,
           }
          sell := Sell{
            SoldB: soldB,
          }
         EList.PushBack(e)
         AEList.PushBack(ae)
         SellList.PushBack(sell)
     }
   }

       if choose == "sell" {
         for pos := EList.Front(); pos != nil; pos = pos.Next(){
            var e Employee
            e = pos.Value.(Employee)

            fmt.Println("quem vendeu?")
            fmt.Scanln(&sell)
          if sell == e.Code {
            e.SoldB += 1
         }
       }
     }



    if choose == "show"{
      fmt.Println("--Enter option--")
      fmt.Println("\033[1;35m--book--")
      fmt.Println("--cli--")
      fmt.Println("--emplo--\033[1;35m")
      fmt.Scanln(&choice)

       if choice == "book"{
          for pos := BList.Front(); pos != nil; pos = pos.Next(){
           var b Book
           b = pos.Value.(Book)
           fmt.Println("\033[1;34mThis is the book\033[m" , b.I)
           fmt.Println("The code is" , b.Code)
           fmt.Println("The name is" , b.Name)
           fmt.Println("The author is" , b.Author)
           fmt.Println("The company is" , b.Company)
           fmt.Println("The year is" , b.Year)
           fmt.Println("The edition is" , b.Edition)
           fmt.Println("\033[1;32mnext BOOK\033[m")
           }
         }
       if choice == "cli"{
          for pos := CList.Front(); pos != nil; pos = pos.Next(){
             var c Client

             c = pos.Value.(Client)
             fmt.Println("\033[1;34mThis is the client\033[m" , c.I)
             fmt.Println("The code is" , c.Code)
             fmt.Println("The name is" , c.Name)
             //fmt.Println("The fone is" , c.Fone[c.I-1])
             fmt.Println("The cpf is" , c.CPF)
             fmt.Println("The age is" , c.Age)
             fmt.Println("\033[1;32mnext Adress\033[m")

             for pos := AList.Front(); pos != nil; pos = pos.Next(){
               var a Address
               a = pos.Value.(Address)
               if c.I == a.I {
                 fmt.Println("\033[1;34mAddress the client\033[m" , a.I)
                 fmt.Println("The logradouro is" , a.Logradouro)
                 fmt.Println("The number is" , a.Number)
                 fmt.Println("The neighborhood is" , a.Neighborhood)
                 fmt.Println("The city is" , a.City)
                 fmt.Println("The cep is" , a.CEP)
                 fmt.Println("The uf is" , a.UF)
                 fmt.Println("\033[1;33mnext FONE\033[m")

                  fmt.Println("The FONE the client" , a.I)
                  for i := 0;  i< j; i++ {
                    if vF[c.I-1][i] != nil{
                      fmt.Println(vF[c.I-1][i])
                    }
                  }
                  fmt.Println("\033[1;32mnext CLIENT\033[m")
              }
            }
          }
        }

        if choice == "emplo"{
           for pos := EList.Front(); pos != nil; pos = pos.Next(){
              var e Employee

              e = pos.Value.(Employee)
              fmt.Println("\033[1;34mThis is the Employee\033[m" , e.I)
              fmt.Println("The code is" , e.Code)
              fmt.Println("The name is" , e.Name)
              fmt.Println("The age is" , e.Age)
              fmt.Println("The  is Sold Book" , e.SoldB)

              fmt.Println("\033[1;32mnext Adress\033[m")

              for pos := AEList.Front(); pos != nil; pos = pos.Next(){
                var ae Address
                ae = pos.Value.(Address)
                if e.I == ae.I {
                  fmt.Println("\033[1;34mAddress the employee\033[m" , ae.I)
                  fmt.Println("The logradouro is" , ae.Logradouro)
                  fmt.Println("The number is" , ae.Number)
                  fmt.Println("The neighborhood is" , ae.Neighborhood)
                  fmt.Println("The city is" , ae.City)
                  fmt.Println("The cep is" , ae.CEP)
                  fmt.Println("The uf is" , ae.UF)
                  fmt.Println("\033[1;33mnext FONE\033[m")

                   fmt.Println("The FONE the Employee" , ae.I)
                   for i := 0;  i< j; i++ {
                     if vE[e.I-1][i] != nil{
                       fmt.Println(vE[e.I-1][i])
                     }
                   }
             }
           }
         }
       }
     }
   }
 }
