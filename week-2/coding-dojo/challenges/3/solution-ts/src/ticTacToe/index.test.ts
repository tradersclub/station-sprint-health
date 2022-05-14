import {TicTacToeGame} from './index';

describe("who wins", () => {
  test('The O win', () => {
    const ticTacToeGame = new TicTacToeGame()

    ticTacToeGame.play(0,0)
    ticTacToeGame.play(1,1)
    ticTacToeGame.play(0,1)
    ticTacToeGame.play(2,2)
    ticTacToeGame.play(0,2)
    
    expect(ticTacToeGame.getWinner()).toBe("O");
  });
  
  test('Already end game', () => {
    const ticTacToeGame = new TicTacToeGame()

    ticTacToeGame.play(0,0)
    ticTacToeGame.play(1,1)
    ticTacToeGame.play(0,1)
    ticTacToeGame.play(2,2)
    ticTacToeGame.play(0,2)
    
    expect(() => ticTacToeGame.play(1,2)).toThrow("Game already end");
  });
  
  test('Position already used', () => {
    const ticTacToeGame = new TicTacToeGame()

    ticTacToeGame.play(0,0)
    
    expect(() => ticTacToeGame.play(0,0)).toThrow("This position is already used");
  });
  
})