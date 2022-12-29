package main

import (
	"testing"
	"os"
)


func TestNewdec(t *testing.T) {
 d:=newdeck()
 if len(d)!=16{
	t.Errorf("expected deck length of 16 but got %v", len(d))
 }
  if d[0]!="Aceofspades"{
	t.Errorf("it is supposed to be an ace of spades but it got  %v",d[0])
  }

  if d[15]!="fourofdiamonds"{
      t.Errorf("it is supposed to return a fourofdiamonds but it got %v",d[15])
  }
}

func testdeckfromfile( t *testing.T){
  os.Remove("_decktesting")
  deck := newdeck()
  deck.savetofile("_decktesting")
  loadedDeck:= createdeckfromfile("_decktesting")
  if len(loadedDeck) != 16{
	t.Errorf("length of deck is meant to be 16 but  is now %v" , len(loadedDeck))
  }
}