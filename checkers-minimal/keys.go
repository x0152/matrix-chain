package checkers

import "cosmossdk.io/collections"

const ModuleName = "checkers"
const MaxIndexLenght = 256

var (
    ParamsKey  = collections.NewPrefix("Params")
    StoredGamesKey = collections.NewPrefix("StoredGames/value/")
)
