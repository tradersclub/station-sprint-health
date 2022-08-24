import {isDictionary} from "../utils";

describe("run test", () => {
  test('should EGG and ADD is dictionary', () => {
    expect(isDictionary('egg', 'add').isDictionary).toBeTruthy();
    expect(isDictionary('add', 'egg').isDictionary).toBeTruthy();
  });

  test('should FOO and BAR is dictionary', () => {
    expect(isDictionary('foo', 'bar').isDictionary).toBeFalsy();
    expect(isDictionary('bar', 'foo').isDictionary).toBeFalsy();
  });

  test('should PAPER and TITLE is dictionary', () => {
    expect(isDictionary('paper', 'title').isDictionary).toBeTruthy();
    expect(isDictionary('title', 'paper').isDictionary).toBeTruthy();
  });

  test('should PPAPER and TTITLE is dictionary', () => {
    expect(isDictionary('ppaper', 'ttitle').isDictionary).toBeTruthy();
    expect(isDictionary('ttitle', 'ppaper').isDictionary).toBeTruthy();
  });

  test('should ARTHUR and LEO is dictionary', () => {
    expect(isDictionary('arthur', 'leo').isDictionary).toBeFalsy();
  });
})
