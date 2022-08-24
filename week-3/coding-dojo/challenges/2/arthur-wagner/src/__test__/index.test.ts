import {isDictionary} from "../utils";
import {isIsomorphic, isIsomorphic2} from "../utils_leo";

describe("run test", () => {
  test('should EGG and ADD is dictionary', () => {
    expect(isDictionary('egg', 'add').isDictionary).toBeTruthy();
    expect(isDictionary('add', 'egg').isDictionary).toBeTruthy();
    expect(isIsomorphic('egg', 'add')).toBeTruthy();
    expect(isIsomorphic('add', 'egg')).toBeTruthy();
    expect(isIsomorphic2('egg', 'add')).toBeTruthy();
    expect(isIsomorphic2('add', 'egg')).toBeTruthy();
  });

  test('should FOO and BAR is dictionary', () => {
    expect(isDictionary('foo', 'bar').isDictionary).toBeFalsy();
    expect(isDictionary('bar', 'foo').isDictionary).toBeFalsy();
    expect(isIsomorphic('foo', 'bar')).toBeFalsy();
    expect(isIsomorphic('bar', 'foo')).toBeFalsy();
    expect(isIsomorphic2('foo', 'bar')).toBeFalsy();
    expect(isIsomorphic2('bar', 'foo')).toBeFalsy();
  });

  test('should PAPER and TITLE is dictionary', () => {
    expect(isDictionary('paper', 'title').isDictionary).toBeTruthy();
    expect(isDictionary('title', 'paper').isDictionary).toBeTruthy();
    expect(isIsomorphic('paper', 'title')).toBeTruthy();
    expect(isIsomorphic('title', 'paper')).toBeTruthy();
    expect(isIsomorphic2('paper', 'title')).toBeTruthy();
    expect(isIsomorphic2('title', 'paper')).toBeTruthy();
  });

  test('should PPAPER and TTITLE is dictionary', () => {
    expect(isDictionary('ppaper', 'ttitle').isDictionary).toBeTruthy();
    expect(isDictionary('ttitle', 'ppaper').isDictionary).toBeTruthy();
    expect(isIsomorphic('ppaper', 'ttitle')).toBeTruthy();
    expect(isIsomorphic('ttitle', 'ppaper')).toBeTruthy();
    expect(isIsomorphic2('ppaper', 'ttitle')).toBeTruthy();
    expect(isIsomorphic2('ttitle', 'ppaper')).toBeTruthy();
  });

  test('should ARTHUR and LEO is dictionary', () => {
    expect(isDictionary('arthur', 'leo').isDictionary).toBeFalsy();
    expect(isIsomorphic('arthur', 'leo')).toBeFalsy();
    expect(isIsomorphic2('arthur', 'leo')).toBeFalsy();
  });
})
