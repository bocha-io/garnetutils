import { mudConfig, resolveTableId  } from "@latticexyz/world/register";

export default mudConfig({
  systems: {
    TicSystem: {
      name: "tic",
      openAccess: true,
    },
  },
  tables: {
    Match: {
        dataStruct:false,
        schema:{
            created: "bool",
            gameType: "GameType",
        },
    },
    PlayerTwo: "bytes32",

    User: "bytes32",

    CurrentGameState: "GameState",
    CurrentHp: "uint32",
    Inmune: "uint32",
    Time: "uint32",
    Projectil: {
        dataStruct:false,
        schema:{
            spawned:"bool",
            x:"int32",
            y:"int32",
            shotDir: "int32",
        },
    },

    Position: {
      dataStruct: false,
      schema: {
        x: "int32",
        y: "int32",
      },
    },
  },
  enums: {
      GameState:["Playing", "Victory", "Defeat"],
      GameType:["Solo", "Online"]
  },
  modules: [
    {
      name: "KeysWithValueModule",
      root: true,
      args: [resolveTableId("User")],
    },
    {
      name: "KeysWithValueModule",
      root: true,
      args: [resolveTableId("PlayerTwo")],
    },

  ],
});
