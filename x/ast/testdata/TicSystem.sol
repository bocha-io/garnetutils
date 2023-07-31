// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {System} from "@latticexyz/world/src/System.sol";
import {User} from "../codegen/tables/User.sol";
import {Match} from "../codegen/tables/Match.sol";
import {PlayerTwo} from "../codegen/tables/PlayerTwo.sol";
import {Position} from "../codegen/tables/Position.sol";
import {Projectil} from "../codegen/tables/Projectil.sol";
import {CurrentHp} from "../codegen/tables/CurrentHp.sol";
import {Inmune} from "../codegen/tables/Inmune.sol";
import {Time} from "../codegen/tables/Time.sol";
import {GameState, GameType} from "../codegen/Types.sol";
import {CurrentGameState} from "../codegen/tables/CurrentGameState.sol";
import {addressToEntityKey} from "../addressToEntityKey.sol";
import {KeysLib} from "../libs/KeysLib.sol";

contract TicSystem is System {
    function checkCollision(int32 x, int32 y, int32 size, int32 targetX, int32 targetY, int32 targetSize)
        private
        pure
        returns (bool)
    {
        bool collisionX = x + size >= targetX && targetX + targetSize >= x;
        bool collisionY = y + size >= targetY && targetY + targetSize >= y;
        // collision only if on both axes
        return collisionX && collisionY;
    }

    function updatePlayerMovement(bytes32 senderKey, int32 hDir, int32 vDir) private returns (int32, int32) {
        (int32 x, int32 y) = Position.get(senderKey);
        int32 newX = x + int32(10) * hDir;
        if (newX < 16) {
            newX = 16;
        }
        if (newX > 388 - 20) {
            newX = 388 - 20;
        }
        int32 newY = y + int32(10) * vDir;
        if (newY < 0) {
            newY = 0;
        }
        if (newY > 244 - 20 - 36) {
            newY = 244 - 20 - 36;
        }
        Position.set(senderKey, newX, newY);
        return (newX, newY);
    }

    function updateEnemyPosition(bytes32 enemyID, int32 playerX, int32 playerY) private returns (int32, int32) {
        // Enemy position will be 0,0 if not set
        (int32 enemyX, int32 enemyY) = Position.get(enemyID);
        // Enemy will move 5 units at the time, trying to get to the player
        if (enemyX > playerX) {
            enemyX -= 5;
        } else if (enemyX < playerX) {
            enemyX += 5;
        }

        if (enemyY > playerY) {
            enemyY -= 5;
        } else if (enemyY < playerY) {
            enemyY += 5;
        }
        Position.set(enemyID, enemyX, enemyY);
        return (enemyX, enemyY);
    }

    function updateProjectil(bytes32 projectilID, int32 playerX, int32 playerY, int32 playerShootDir)
        private
        returns (int32, int32)
    {
        (bool spawned, int32 x, int32 y, int32 shootDir) = Projectil.get(projectilID);
        if (!spawned && playerShootDir != 0) {
            x = playerX - 8;
            y = playerY + 14;
            Projectil.set(projectilID, true, x, y, playerShootDir);
            return (x, y);
        }

        if (spawned) {
            if (shootDir == 1) {
                x = x + 20;
            } else if (shootDir == 2) {
                x = x - 20;
            } else if (shootDir == 3) {
                y = y + 20;
            } else if (shootDir == 4) {
                y = y - 20;
            }

            if (x <= 0) {
                Projectil.set(projectilID, false, 0, 0, 0);
                return (1000, 1000);
            }
            if (x >= 388) {
                Projectil.set(projectilID, false, 0, 0, 0);
                return (1000, 1000);
            }
            if (y <= 0) {
                Projectil.set(projectilID, false, 0, 0, 0);
                return (1000, 1000);
            }
            if (y >= 244) {
                Projectil.set(projectilID, false, 0, 0, 0);
                return (1000, 1000);
            }
            Projectil.set(projectilID, true, x, y, shootDir);
            return (x, y);
        }

        return (1000, 1000);
    }

    function updateInmune(bytes32 key) private returns (uint32) {
        uint32 inmune = Inmune.get(key);
        if (inmune != 0) {
            inmune--;
            Inmune.set(key, inmune);
        }
        return inmune;
    }

    function updateTime(bytes32 key) private returns (uint32) {
        uint32 time = Time.get(key);
        if (time != 0) {
            time--;
            Time.set(key, time);
        }
        return time;
    }

    function reduceHp(bytes32 key) private returns (uint32) {
        uint32 hp = CurrentHp.get(key);
        if (hp == 1) {
            CurrentHp.set(key, 0);
            return 0;
        }

        CurrentHp.set(key, hp - 1);
        return hp - 1;
    }

    function tic(int32 hDir, int32 vDir, int32 shootDir, int32 hDir2, int32 vDir2, int32 shootDir2) public {
        bytes32 senderKey = addressToEntityKey(_msgSender());

        // Get the match
        (bool created, GameType gameType) = Match.get(senderKey);
        require(created == true, "the game was not created yet");
        bytes32 playerTwo = 0;
        if (gameType == GameType.Online) {
            playerTwo = PlayerTwo.get(senderKey);
            require(playerTwo != 0, "player two is missing");
        }
        require(CurrentGameState.get(senderKey) == GameState.Playing, "the game has not started yet");

        if (updateTime(senderKey) == 0) {
            CurrentGameState.set(senderKey, GameState.Defeat);
            return;
        }

        (int32 playerX, int32 playerY) = updatePlayerMovement(senderKey, hDir, vDir);
        (int32 pX, int32 pY) = updateProjectil(KeysLib.projectilP1(senderKey), playerX, playerY, shootDir);

        // Update the enemy pos
        (int32 enemyX, int32 enemyY) = updateEnemyPosition(KeysLib.enemyKey(senderKey), playerX, playerY);

        // Projectil1 collision
        if (checkCollision(pX, pY, 5, enemyX - 30, enemyY + 8, 24)) {
            Projectil.set(KeysLib.projectilP1(senderKey), false, 0, 0, 0);
            if (reduceHp(KeysLib.enemyKey(senderKey)) == 0) {
                CurrentGameState.set(senderKey, GameState.Victory);
            }
        }

        uint32 inmune = updateInmune(senderKey);
        // Enemy player collision
        if (checkCollision(playerX - 8, playerY + 14, 5, enemyX - 30, enemyY + 8, 24)) {
            if (inmune == 0) {
                if (reduceHp(senderKey) == 0) {
                    CurrentGameState.set(senderKey, GameState.Defeat);
                }
                Inmune.set(senderKey, 30);
            }
        }

        // Player two logic
        if (playerTwo != 0) {
            (playerX, playerY) = updatePlayerMovement(playerTwo, hDir2, vDir2);
            (pX, pY) = updateProjectil(KeysLib.projectilP2(senderKey), playerX, playerY, shootDir2);

            // Projectil2 collision
            if (checkCollision(pX, pY, 13, enemyX - 30, enemyY + 8, 24)) {
                Projectil.set(KeysLib.projectilP2(senderKey), false, 0, 0, 0);
                if (reduceHp(KeysLib.enemyKey(senderKey)) == 0) {
                    CurrentGameState.set(senderKey, GameState.Victory);
                }
            }
        }
    }
}
