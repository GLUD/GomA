
//ARBOL BINARIO

package main

import (
  "fmt"
  "math/rand"
)

type Arbol struct{
  Izquierda *Arbol
  Valor int
  Derecha *Arbol
}

func New(n, k int) *Arbol{
  var t *Arbol
  for _,v:=range rand.Perm(n){
    t = Insertar(t,(1+v)*k)
  }
  return t
  }

func Insertar(t *Arbol, v int) *Arbol{
  if t==nil{
    return &Arbol{nil, v, nil}
  }
  if v< t.Valor{
    t.Izquierda = Insertar(t.Izquierda,v)
    return t
  }
  t.Derecha = Insertar(t.Derecha,v)
  return t
}

func RecorrerInorden(t *Arbol){
  if t==nil{
    return
  }
  RecorrerInorden(t.Izquierda)
  fmt.Print(t.Valor)
  fmt.Print("-")
  RecorrerInorden(t.Derecha)
}

func RecorrerPreorden(t *Arbol){
  if t==nil{
    return
  }
  fmt.Print(t.Valor)
  fmt.Print("-")
  RecorrerPreorden(t.Izquierda)
  RecorrerPreorden(t.Derecha)
}

func RecorrerPostorden(t *Arbol){
  if t==nil{
    return
  }
  RecorrerPostorden(t.Izquierda)
  RecorrerPostorden(t.Derecha)
  fmt.Print(t.Valor)
  fmt.Print("-")
}

func main(){
  //t1:=New(10,3)
  t1 := &Arbol{ &Arbol{&Arbol{nil,3,nil}, 9, &Arbol{nil,6,nil} }, 12, &Arbol{&Arbol{nil,15,nil}, 18,&Arbol{nil,21,nil} } }
  RecorrerInorden(t1)
}

