// Project: PA193_mnemonic_aller
// Mainteners UCO: 408788 497391 497577
// Description: Main file to test the API with command line

package main

import (
  "fmt"
  "./mnemonic"
  "encoding/hex"
)


func main() {
  fmt.Println("Testing...")

  mnemonic.EntropyToPhraseAndSeed()
  mnemonic.PhraseToEntropyAndSeed("phrase")
  mnemonic.VerifyPhraseAndSeed("phrase","phrase")

  tmp,_ := mnemonic.PhraseToSeed("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo wrong","TREZOR")
  fmt.Println(hex.EncodeToString(tmp))
}
