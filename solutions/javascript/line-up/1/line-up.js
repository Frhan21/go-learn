//
// This is only a SKELETON file for the 'Line Up' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

const suffixes = new Map([
  ['one',   'st'],
  ['two',   'nd'],
  ['few',   'rd'],
  ['other', 'th'], // Catches zero, teen numbers, and 4-9
  ]);

export const format = (name, n) => {
  const pr = new Intl.PluralRules('en-US', { type: 'ordinal' });
  const rule = pr.select(n)
  const suffixed = suffixes.get(rule)
  return `${name}, you are the ${n}${suffixed} customer we serve today. Thank you!`
};


