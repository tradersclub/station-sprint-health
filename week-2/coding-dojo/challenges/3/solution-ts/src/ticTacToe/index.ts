import { TableTicTacToe, PlayerType, PositionValue } from "./types";

export class TicTacToeGame {
    private winner:PositionValue = undefined;
    private currentPlayer:PlayerType = "O";

    private table:TableTicTacToe = [
        [undefined,undefined,undefined],
        [undefined,undefined,undefined],
        [undefined,undefined,undefined],
    ]

    public play(line:number, column:number):PositionValue {
        if(this.winner) {
            throw new Error("Game already end");   
        }

        if(this.table[line][column]) {
            throw new Error("This position is already used");   
        }

        this.table[line][column] = this.currentPlayer;

        console.log(`Player "${this.currentPlayer}" plays in line ${line} and column ${column}`)
        
        this.calcWinner();
        this.getNextPlayer();
        
        return this.winner;
    }

    public getWinner():PositionValue {
        return this.winner;
    }
    
    public getCurrentPlayer():PlayerType {
        return this.currentPlayer;
    }

    private getNextPlayer() {
        if (this.currentPlayer === "O") {
            this.currentPlayer = "X";
        } else {
            this.currentPlayer = "O";
        }
    }

    private calcWinner() {
        if (this.transversalWinner()) return;

        for (let index = 0; index < 3; index++) {
            if (this.horizontalWinner(index)) return;
            if (this.verticalWinner(index)) return;
        }

    }

    private horizontalWinner(line:number):boolean {
        if(this.table[line][0] && this.table[line][0] === this.table[line][1] && this.table[line][1] === this.table[line][2]) {
            this.winner = this.table[line][0];
            console.log(`Player "${this.currentPlayer}" wins in horizontal`)
            return true;
        }
        return false;
    }
    
    private verticalWinner(column:number):boolean {
        if(this.table[0][column] && this.table[0][column] === this.table[1][column] && this.table[1][column] === this.table[2][column]) {
            this.winner = this.table[0][column];
            console.log(`Player "${this.winner}" wins in vertical`)
            return true;
        }
        return false;
    }
      
    private transversalWinner():boolean {
        if(this.table[0][0] && this.table[0][0] === this.table[1][1] && this.table[1][1] === this.table[2][2]) {
            this.winner = this.table[0][0];
            console.log(`Player "${this.winner}" wins in transversal`);
            return true;
        }
        if(this.table[2][0] && this.table[2][0] === this.table[1][1] && this.table[1][1] === this.table[0][2]) {
            this.winner = this.table[2][0]
            console.log(`Player "${this.winner}" wins in transversal`);
            return true;
        }
        return false;
    }
}